package tap

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jenkins-x-plugins/jx-tap/pkg/assets"
	"github.com/jenkins-x-plugins/jx-tap/pkg/common"
	"github.com/jenkins-x/go-scm/scm"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras/helper"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras/templates"
	"github.com/jenkins-x/jx-helpers/v3/pkg/files"
	"github.com/jenkins-x/jx-helpers/v3/pkg/options"
	"github.com/jenkins-x/jx-helpers/v3/pkg/scmhelpers"
	"github.com/jenkins-x/jx-helpers/v3/pkg/stringhelpers"
	"github.com/jenkins-x/jx-helpers/v3/pkg/termcolor"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/mpontillo/tap13"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	info = termcolor.ColorInfo

	cmdLong = templates.LongDesc(`
		Promotes a version of an application to an Environment
`)

	cmdExample = templates.Examples(`
		# promotes your current app to the staging environment
		%s 
	`)
)

// Options the options for this command
type Options struct {
	options.BaseOptions
	scmhelpers.PullRequestOptions

	Dir                     string
	Namespace               string
	GenerateHTML            bool
	WriteLogToBucketTimeout time.Duration
	Template                *template.Template
	PassedLanguages         []string
}

// NewCmdTap creates a command object for the command
func NewCmdTap() (*cobra.Command, *Options) {
	options := &Options{}

	cmd := &cobra.Command{
		Use:     "tap",
		Short:   "Processes the .tap files generated by a pipeline",
		Long:    cmdLong,
		Example: fmt.Sprintf(cmdExample, common.BinaryName),
		Run: func(cmd *cobra.Command, args []string) {
			err := options.Run()
			helper.CheckErr(err)
		},
	}

	cmd.Flags().StringVarP(&options.Dir, "dir", "d", ".", "The directory to scan for *.tap files")
	cmd.Flags().StringVarP(&options.Namespace, "namespace", "n", "", "The namespace. Defaults to the current namespace")
	cmd.Flags().BoolVarP(&options.GenerateHTML, "html", "", false, "generates HTML rather than commenting on the Pull Request")
	cmd.Flags().DurationVarP(&options.WriteLogToBucketTimeout, "write-log-timeout", "", time.Minute*30, "The timeout for writing pipeline logs to the bucket")

	options.BaseOptions.AddBaseFlags(cmd)
	return cmd, options
}

func (o *Options) Run() error {
	err := filepath.Walk(o.Dir, func(path string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() {
			return nil
		}
		name := f.Name()
		if !strings.HasSuffix(name, ".tap") {
			return nil
		}
		return o.processTapFile(path)
	})
	if err != nil {
		return errors.Wrapf(err, "failed to process tap files in dir %s", o.Dir)
	}

	if o.GenerateHTML || len(o.PassedLanguages) == 0 {
		return nil
	}

	comment := "Valid linters:\n* " + strings.Join(o.PassedLanguages, "\n* ")
	return o.commentOnPullRequest(comment)
}

func (o *Options) processTapFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "failed to read file %s", path)
	}
	if len(data) == 0 {
		return nil
	}
	lines := strings.Split(string(data), "\n")
	results := tap13.Parse(lines)
	if results == nil {
		return errors.Errorf("nil returned from parsing %d lines of text from file %s", len(lines), path)
	}

	log.Logger().Infof("processing tap %s", info(path))

	var tests []Test
	for i := range results.Tests {
		t := &results.Tests[i]
		state := "passed"
		style := "info"
		if t.Failed {
			state = "failed"
			style = "danger"
		}

		message := strings.ReplaceAll(string(t.YamlBytes), "\\n", "\n")
		message = strings.TrimSpace(message)
		message = strings.TrimPrefix(message, "message:")
		message = strings.TrimSpace(message)

		t.Description = strings.TrimSpace(t.Description)
		t.Description = strings.TrimPrefix(t.Description, "-")
		t.Description = strings.TrimSpace(t.Description)

		errorDetails, err := ParseErrors(message, t.Description)
		if err != nil {
			return errors.Wrapf(err, "failed to parse errors for file %s, message %s", path, message)
		}

		tests = append(tests, Test{
			Test:        *t,
			Status:      state,
			StatusStyle: style,
			Message:     message,
			Errors:      errorDetails,
		})
	}

	return o.processTapFileResults(path, tests)
}

func (o *Options) processTapFileResults(path string, tests []Test) error {
	if o.GenerateHTML {
		return o.generateTapResultsHTML(path, tests)
	}

	err := o.PullRequestOptions.Validate()
	if err != nil {
		return errors.Wrapf(err, "failed to validate pull request options")
	}

	pr, err := o.DiscoverPullRequest()
	if err != nil {
		return errors.Wrapf(err, "failed to discover the pull request")
	}
	if pr == nil {
		return errors.Errorf("no Pull Request could be found for %d in repository %s", o.Number, o.Repository)
	}

	commentMarkdown, lang, err := o.generatePullRequestComment(pr, path, tests)
	if err != nil {
		return errors.Wrapf(err, "failed to generate pull request comment")
	}

	if commentMarkdown == "" {
		o.PassedLanguages = append(o.PassedLanguages, lang)
		return nil
	}
	return o.commentOnPullRequest(commentMarkdown)
}

func (o *Options) commentOnPullRequest(commentMarkdown string) error {
	ctx := context.Background()
	comment := &scm.CommentInput{Body: commentMarkdown}
	_, _, err := o.ScmClient.PullRequests.CreateComment(ctx, o.FullRepositoryName, o.Number, comment)
	prName := "#" + strconv.Itoa(o.Number)
	if err != nil {
		return errors.Wrapf(err, "failed to comment on pull request %s on repository %s", prName, o.FullRepositoryName)
	}
	log.Logger().Infof("commented on pull request %s on repository %s", prName, o.FullRepositoryName)
	return nil
}

func (o *Options) generateTapResultsHTML(path string, tests []Test) error {
	if o.Template == nil {
		templateName := "templates/report.html"
		reportTemplate, err := assets.Asset(templateName)
		if err != nil {
			return errors.Wrapf(err, "failed to load template %s", templateName)
		}

		funcMap := template.FuncMap{
			"htmlSafe": func(html string) template.HTML {
				return template.HTML(html)
			},
		}

		o.Template, err = template.New("name").Funcs(funcMap).Parse(string(reportTemplate))
		if err != nil {
			return errors.Wrapf(err, "failed to parse template")
		}
	}
	buf := &bytes.Buffer{}
	templateData := map[string]interface{}{
		"Tests": tests,
	}
	err := o.Template.Execute(buf, templateData)
	if err != nil {
		return errors.Wrapf(err, "failed to execute template for %s", path)
	}

	outFile := strings.TrimSuffix(path, ".tap") + ".html"
	err = ioutil.WriteFile(outFile, buf.Bytes(), files.DefaultFileWritePermissions)
	if err != nil {
		return errors.Wrapf(err, "failed to save file %s", outFile)
	}

	log.Logger().Infof("saved tap file %s", info(outFile))
	return nil
}

func (o *Options) generatePullRequestComment(pr *scm.PullRequest, path string, tests []Test) (string, string, error) {
	buf := strings.Builder{}

	sourcePrefix := pr.Base.Repo.Link
	branch := o.PullRequestOptions.Branch
	if sourcePrefix != "" && branch != "" {
		sourcePrefix = stringhelpers.UrlJoin(sourcePrefix, "blob", branch)
	}
	lang := ""
	if strings.HasSuffix(path, ".tap") {
		lang = strings.TrimSuffix(path, ".tap")
		idx := strings.LastIndex(lang, "-")
		if idx >= 0 {
			lang = lang[idx+1:]
		}
	}

	// lets check if all the tests are green...
	allPassed := true
	for _, t := range tests {
		if t.Failed {
			allPassed = false
			break
		}
		if len(t.Errors) > 0 {
			allPassed = false
			break
		}
	}
	if allPassed || len(tests) == 0 {
		return "", lang, nil
	}

	if lang != "" {
		buf.WriteString(lang + " Linter\n")
	}
	for _, t := range tests {
		if t.Passed {
			continue
		}
		for _, e := range t.Errors {
			fileLink := "* "
			file := e.File
			if file != "" {
				lineSuffix := ""
				if e.Line != "" {
					lineSuffix = "#L" + e.Line
				}
				fileLink += "[" + file + "](" + sourceLink(sourcePrefix, file) + lineSuffix + ") : "
			}
			buf.WriteString(fileLink + e.Heading + "\n")

			message := e.Message
			if message != "" {
				if !strings.HasSuffix(message, "\n") {
					message += "\n"
				}
				buf.WriteString("\n")
				if strings.TrimSpace(message) != "" {
					buf.WriteString("```" + toMarkdownLang(lang) + "\n")
					buf.WriteString(message)
					buf.WriteString("```\n")
				}
			}
		}
	}
	return buf.String(), lang, nil
}

func sourceLink(sourcePrefix string, file string) string {
	if sourcePrefix == "" {
		return file
	}
	return stringhelpers.UrlJoin(sourcePrefix, file)
}

func toMarkdownLang(lang string) string {
	lower := strings.ToLower(lang)
	switch lower {
	case "kubernetes_kubeval":
		return "bash"
	default:
		return lower
	}
}

type Test struct {
	tap13.Test

	Status      string
	StatusStyle string
	Message     string
	Errors      []*Error
}
