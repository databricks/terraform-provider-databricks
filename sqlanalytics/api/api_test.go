package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringOrIntBasics(t *testing.T) {
	s := NewStringOrInt("1234")
	assert.Equal(t, "1234", s.String())
}

func TestStringOrIntMarshal(t *testing.T) {
	{
		s := stringOrInt("1234")
		b, err := json.Marshal(s)
		assert.NoError(t, err)
		assert.Equal(t, "1234", string(b))
	}
	{
		s := stringOrInt("6de5037b-5afb-4473-9d0d-50046543bdc3")
		b, err := json.Marshal(s)
		assert.NoError(t, err)
		assert.Equal(t, "\"6de5037b-5afb-4473-9d0d-50046543bdc3\"", string(b))
	}
}

func TestStringOrIntUnmarshal(t *testing.T) {
	{
		var s stringOrInt
		err := json.Unmarshal([]byte("1234"), &s)
		assert.NoError(t, err)
		assert.Equal(t, "1234", string(s))
	}
	{
		var s stringOrInt
		err := json.Unmarshal([]byte("\"1235\""), &s)
		assert.NoError(t, err)
		assert.Equal(t, "1235", string(s))
	}
	{
		var s stringOrInt
		err := json.Unmarshal([]byte("[1]"), &s)
		assert.Error(t, err)
	}
	{
		var s stringOrInt
		err := json.Unmarshal([]byte("this is invalid json"), &s)
		assert.Error(t, err)
	}
}
