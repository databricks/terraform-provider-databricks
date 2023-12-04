package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringIsUUID(t *testing.T) {
	assert.True(t, StringIsUUID("3f670caf-9a4b-4479-8143-1a0878da8f57"))
	assert.False(t, StringIsUUID("abc"))
}

func TestContains(t *testing.T) {
	assert.True(t, Contains[string]([]string{"a", "b", "c"}, "a"))
	assert.False(t, Contains[string]([]string{"a", "b", "c"}, "d"))
	assert.True(t, Contains[int]([]int{1, 2, 3}, 1))
	assert.False(t, Contains[int]([]int{1, 2, 3}, 4))
}
