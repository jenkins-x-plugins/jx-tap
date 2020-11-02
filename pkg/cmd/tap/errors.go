package tap

import (
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Error struct {
	Status      string
	StatusStyle string
	File        string
	Line        string
	Col         string
	Heading     string
	Message     string
}

var (
	fileLineColRegex = regexp.MustCompile(`([^<>:;,?\"*|/]+)\s+(\d+)\s+(\d+)\s+(.*)`)
	fileLineRegex    = regexp.MustCompile(`([^<>:;,?\"*|/]+)\s+(\d+)\s+(.*)`)
)

// HeadingHTML returns the answer as html
func (e *Error) HeadingHTML() string {
	if e.Heading == "" {
		return ""
	}
	answer := e.Heading
	answer = string(markdown.ToHTML([]byte(answer), nil, nil))
	answer = strings.TrimSpace(answer)
	if strings.HasPrefix(answer, "<p>") && strings.HasSuffix(answer, "</p>") {
		answer = strings.TrimPrefix(answer, "<p>")
		answer = strings.TrimSuffix(answer, "</p>")
	}
	return answer
}

// ParseErrors parsers the errors from a message
func ParseErrors(text, defaultFile string) ([]*Error, error) {
	var answer []*Error
	lines := strings.Split(text, "\n")
	buf := strings.Builder{}

	var lastError *Error
	for _, line := range lines {
		var e *Error
		matches := fileLineColRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			e = &Error{
				File:    matches[1],
				Line:    matches[2],
				Col:     matches[3],
				Heading: matches[4],
			}
		} else {
			matches := fileLineRegex.FindStringSubmatch(line)
			if len(matches) > 0 {
				e = &Error{
					File:    matches[1],
					Line:    matches[2],
					Heading: matches[3],
				}
			}
		}
		if e == nil {
			buf.WriteString(line)
			buf.WriteString("\n")
			continue
		}
		// avoid file name of "myscript.sh line 5" ending up as "myscrip.sh line"
		e.File = strings.TrimSuffix(e.File, " line")

		if buf.Len() > 0 {
			if lastError == nil {
				lastError = &Error{File: defaultFile}
				answer = append(answer, lastError)
			}

			// lets keep the old buffer
			lastError.Message = buf.String()
			buf.Reset()
		}
		answer = append(answer, e)
		lastError = e
	}
	if lastError == nil {
		lastError = &Error{File: defaultFile}
		answer = append(answer, lastError)
	}
	if buf.Len() > 0 && lastError != nil {
		lastError.Message = buf.String()
	}
	return answer, nil
}
