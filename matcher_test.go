package pagepath

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatcher_Compare(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		matcher := NewMatcher()
		page, err := New("https://pkg.go.dev/net/url?xxx=bbb#Parse")
		require.NoError(t, err)

		comparisonPage, err := New("https://pkg.go.dev/net/url?xxx=bbb")
		require.NoError(t, err)

		assert.True(t, matcher.Compare(page, comparisonPage, ElementScheme))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementHost))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementPath))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementQuery))
		assert.False(t, matcher.Compare(page, comparisonPage, ElementFragment))
	})
	t.Run("ignore trailing slash", func(t *testing.T) {
		matcher := NewMatcher(
			WithIgnoreTrailingSlash(),
		)
		page, err := New("https://pkg.go.dev/net/url?xxx=bbb#Parse")
		require.NoError(t, err)

		comparisonPage, err := New("https://pkg.go.dev/net/url/?xxx=bbb")
		require.NoError(t, err)

		assert.True(t, matcher.Compare(page, comparisonPage, ElementScheme))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementHost))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementPath))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementQuery))
		assert.False(t, matcher.Compare(page, comparisonPage, ElementFragment))
	})
	t.Run("ignore subdomain www", func(t *testing.T) {
		matcher := NewMatcher(
			WithPrepareFuncs(func(p *Page) {
				p.Host = strings.TrimPrefix(p.Host, "www.")
			}),
		)
		page, err := New("https://example.com/net/url?xxx=bbb#Parse")
		require.NoError(t, err)

		comparisonPage, err := New("https://www.example.com/net/url?xxx=bbb")
		require.NoError(t, err)

		assert.True(t, matcher.Compare(page, comparisonPage, ElementScheme))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementHost))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementPath))
		assert.True(t, matcher.Compare(page, comparisonPage, ElementQuery))
		assert.False(t, matcher.Compare(page, comparisonPage, ElementFragment))
	})
}
