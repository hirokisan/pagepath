package pagepath

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPage_New(t *testing.T) {
	pagePath := "https://pkg.go.dev/net/url?xxx=bbb#Parse"
	page, err := New(pagePath)
	require.NoError(t, err)

	assert.Equal(t, "https", page.Scheme)
	assert.Equal(t, "pkg.go.dev", page.Host)
	assert.Equal(t, "/net/url", page.Path)
	assert.Equal(t, "xxx=bbb", page.Query)
	assert.Equal(t, "Parse", page.Fragment)

	assert.Equal(t, pagePath, page.String())
}

func TestPage_Extract(t *testing.T) {
	page, err := New("https://pkg.go.dev/net/url?xxx=bbb#Parse")
	require.NoError(t, err)

	assert.Equal(t, "https", page.Extract(ElementScheme))
	assert.Equal(t, "pkg.go.dev", page.Extract(ElementHost))
	assert.Equal(t, "/net/url", page.Extract(ElementPath))
	assert.Equal(t, "xxx=bbb", page.Extract(ElementQuery))
	assert.Equal(t, "Parse", page.Extract(ElementFragment))
}

func TestPage_ComparePageWithElements(t *testing.T) {
	page, err := New("https://pkg.go.dev/net/url?xxx=bbb#Parse")
	require.NoError(t, err)

	comparisonPage, err := New("https://pkg.go.dev/net/url?xxx=bbb")
	require.NoError(t, err)

	assert.True(t, page.ComparePageWithElements(comparisonPage, ElementScheme))
	assert.True(t, page.ComparePageWithElements(comparisonPage, ElementHost))
	assert.True(t, page.ComparePageWithElements(comparisonPage, ElementPath))
	assert.True(t, page.ComparePageWithElements(comparisonPage, ElementQuery))
	assert.False(t, page.ComparePageWithElements(comparisonPage, ElementFragment))
}
