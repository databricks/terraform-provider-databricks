package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResults_Error(t *testing.T) {
	cr := Results{}
	assert.NoError(t, cr.Err())
	cr.ResultType = "error"
	assert.EqualError(t, cr.Err(), "")

	cr.Summary = "<b>NotFoundException: Things are going wrong; nested exception is: with something</b>"
	assert.Equal(t, "Things are going wrong with something", cr.Error())

	cr.Summary = ""
	cr.Cause = "ExecutionError: \nStatusCode=400\nStatusDescription=ABC\nSomething else"
	assert.Equal(t, "\nStatusCode=400\nStatusDescription=ABC", cr.Error())

	cr.Cause = "ErrorMessage=Error was here\n"
	assert.Equal(t, "Error was here", cr.Error())

	assert.False(t, cr.Scan())
}

func TestResults_Scan(t *testing.T) {
	cr := Results{
		ResultType: "table",
		Data: []interface{}{
			[]interface{}{"foo", 1, true},
			[]interface{}{"bar", 2, false},
		},
	}
	a := ""
	b := 0
	c := false
	assert.True(t, cr.Scan(&a, &b, &c))
	assert.Equal(t, "foo", a)
	assert.Equal(t, 1, b)
	assert.Equal(t, true, c)

	assert.True(t, cr.Scan(&a, &b, &c))
	assert.Equal(t, "bar", a)
	assert.Equal(t, 2, b)
	assert.Equal(t, false, c)

	assert.False(t, cr.Scan(&a, &b, &c))
}
