package regex

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlPattern(t *testing.T) {
	t.Run("valid: valid url", func(t *testing.T) {
		url := "https://www.google.com/"

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.True(t, matched)
		assert.NoError(t, err)
	})

	t.Run("valid: valid url", func(t *testing.T) {
		url := "https://xkcd.com/2293/"

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.True(t, matched)
		assert.NoError(t, err)
	})

	t.Run("valid: valid url", func(t *testing.T) {
		url := "github.com/orgs/4rchr4y"

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.True(t, matched)
		assert.NoError(t, err)
	})

	t.Run("valid: valid url", func(t *testing.T) {
		url := "https://github.com/orgs/4rchr4y/repositories?type=source"

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.True(t, matched)
		assert.NoError(t, err)
	})

	t.Run("invalid: empty string as a value", func(t *testing.T) {
		url := ""

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.False(t, matched)
		assert.NoError(t, err)
	})

	t.Run("invalid: random letters as a value", func(t *testing.T) {
		url := "abcdef"

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.False(t, matched)
		assert.NoError(t, err)
	})

	t.Run("invalid: invalid url", func(t *testing.T) {
		url := "https://this-shouldn't.match@example.com"

		matched, err := regexp.MatchString(UrlPatternString, url)

		assert.False(t, matched)
		assert.NoError(t, err)
	})
}
