package tap_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/jenkins-x-plugins/jx-tap/pkg/cmd/tap"
	"github.com/jenkins-x/go-scm/scm"
	fakescm "github.com/jenkins-x/go-scm/scm/driver/fake"
	"github.com/jenkins-x/jx-helpers/v3/pkg/files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	// generateTestOutput enable to regenerate the expected output
	generateTestOutput = true
)

func TestTapPullRequestComments(t *testing.T) {
	owner := "myorg"
	repo := "myrepo"
	fullName := scm.Join(owner, repo)

	prNumber := 123
	prBranch := "my-pr-branch-name"
	commentPrefix := fmt.Sprintf("%s#%d:", fullName, prNumber)

	sourceDir := "test_data"

	tmpDir, err := ioutil.TempDir("", "")
	require.NoError(t, err, "failed to create temp dir")

	fs, err := ioutil.ReadDir(sourceDir)
	require.NoError(t, err, "failed to iterate over source %s", sourceDir)

	for _, f := range fs {
		if f == nil || f.IsDir() {
			continue
		}
		name := f.Name()
		if !strings.HasSuffix(name, ".tap") {
			continue
		}

		path := filepath.Join(sourceDir, name)

		to := filepath.Join(tmpDir, name)
		err := files.CopyFile(path, to)
		require.NoError(t, err, "failed to copy file from %s to %s", path, to)

		scmClient, fakeData := fakescm.NewDefault()
		_, o := tap.NewCmdTap()

		o.Dir = tmpDir
		o.Number = prNumber
		o.Branch = prBranch
		o.Owner = owner
		o.Repository = repo
		o.ScmClient = scmClient
		o.SourceURL = "https://github.com/" + fullName

		pr := &scm.PullRequest{
			Number: prNumber,
			Title:  "my awesome pull request",
			Body:   "some text",
			Source: prBranch,
			Base: scm.PullRequestBranch{
				Repo: scm.Repository{
					Namespace: owner,
					Name:      repo,
					FullName:  fullName,
				},
			},
			Link: o.SourceURL + "/pull/" + strconv.Itoa(prNumber),
		}
		fakeData.PullRequests[prNumber] = pr

		err = o.Run()
		require.NoError(t, err, "failed to run command")

		require.Len(t, fakeData.PullRequestCommentsAdded, 1, "should have added one comment for %s", name)
		comment := strings.TrimPrefix(fakeData.PullRequestCommentsAdded[0], commentPrefix)
		expectedFile := filepath.Join(sourceDir, name+".expected.md")
		if generateTestOutput {
			err = ioutil.WriteFile(expectedFile, []byte(comment), files.DefaultFileWritePermissions)
			require.NoError(t, err, "failed to save file %s", expectedFile)
		} else {
			data, err := ioutil.ReadFile(expectedFile)
			require.NoError(t, err, "failed to load file %s", expectedFile)
			assert.Equal(t, string(data), comment, "comment for %s", name)
		}

		err = os.Remove(to)
		require.NoError(t, err, "failed to remove file %s", to)
	}

}
