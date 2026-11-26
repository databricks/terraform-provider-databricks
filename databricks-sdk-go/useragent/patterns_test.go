package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchSemVer(t *testing.T) {
	assert.NoError(t, matchSemVer("1.2.3"))
	assert.NoError(t, matchSemVer("0.0.0-dev+2e014739024a"))
	assert.Error(t, matchSemVer("1.2.3.4"))
	assert.Error(t, matchSemVer("1.2"))
}

func TestMatchAlphanum(t *testing.T) {
	for _, v := range []string{
		"foo",
		"FOO",
		"FOO123",
		"foo_bar",
		"foo-bar",
		"foo.bar",
	} {
		assert.NoError(t, matchAlphanum(v))
	}

	for _, v := range []string{
		"foo bar",
		"foo/bar",
	} {
		assert.Error(t, matchAlphanum(v))
	}
}

func TestMatchAlphanumOrSemVer(t *testing.T) {
	for _, v := range []string{
		"foo",
		"1.2.3",
		"0.0.0-dev+2e014739024a",
		"client.0",
	} {
		assert.NoError(t, matchAlphanumOrSemVer(v))
	}
	for _, v := range []string{
		"foo/bar",
		"1/2/3",
	} {
		assert.Error(t, matchAlphanumOrSemVer(v))
	}
}

func TestSanitize(t *testing.T) {
	for _, v := range []string{
		"foo",
		"FOO",
		"FOO123",
		"foo_bar",
		"foo-bar",
		"foo+bar",
		"foo.bar",
		"1.2.3",
		"client.0",
	} {
		assert.Equal(t, v, Sanitize(v))
	}

	sanitizeMap := map[string]string{
		"1@2#3?4,5/6!7 8 ": "1-2-3-4-5-6-7-8-",
		"foo bar":          "foo-bar",
		"foo/bar":          "foo-bar",
		"foo:)bar":         "foo--bar",
		"foo😊bar":          "foo-bar",
	}
	for k, v := range sanitizeMap {
		assert.Equal(t, v, Sanitize(k))
	}
}
