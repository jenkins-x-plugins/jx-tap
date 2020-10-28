package tap_test

import (
	"testing"

	"github.com/jenkins-x-plugins/jx-tap/pkg/cmd/tap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	sample = "main.go 30 6  `Something` is unused (deadcode)\nfunc Something() {\n     ^\nmain.go 27 21  Error return value of `http.ListenAndServe` is not checked (errcheck)\n\thttp.ListenAndServe(\" 8080\", nil)\n\t                   ^\nmain.go 23  File is not `goimports`-ed (goimports)"
)

func TestParseErrors(t *testing.T) {
	results, err := tap.ParseErrors(sample)
	require.NoError(t, err, "failed to parse %s", sample)
	require.Len(t, results, 3, "size of results")

	e0 := results[0]
	assert.Equal(t, "main.go", e0.File, "e0.File")
	assert.Equal(t, "30", e0.Line, "e0Line")
	assert.Equal(t, "6", e0.Col, "e0Col")
	assert.Equal(t, "`Something` is unused (deadcode)", e0.Heading, "e0Heading")

	e1 := results[1]
	assert.Equal(t, "main.go", e1.File, "e1.File")
	assert.Equal(t, "27", e1.Line, "e1.Line")
	assert.Equal(t, "21", e1.Col, "e1.Col")
	assert.Equal(t, "Error return value of `http.ListenAndServe` is not checked (errcheck)", e1.Heading, "e1.Heading")

	e2 := results[2]
	assert.Equal(t, "main.go", e2.File, "e2.File")
	assert.Equal(t, "23", e2.Line, "e2.Line")
	assert.Equal(t, "", e2.Col, "e2.Col")
	assert.Equal(t, "File is not `goimports`-ed (goimports)", e2.Heading, "e2.Heading")
}
