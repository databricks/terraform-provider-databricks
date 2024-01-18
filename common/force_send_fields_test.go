package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func expectPanic(t *testing.T, message string) func() {
	return func() {
		r := recover()
		if r == nil {
			t.Errorf("expected panic: %s", message)
		}
		if str, ok := r.(string); ok && str != message {
			t.Errorf("expected panic: %s, got: %v", message, r)
		}
		if err, ok := r.(error); ok && err.Error() != message {
			t.Errorf("expected panic: %s, got: %v", message, r)
		}
	}
}

func TestSetForceSendFields_RequestMustBePointer(t *testing.T) {
	defer expectPanic(t, "request argument to setForceSendFields must be a pointer")()

	SetForceSendFields(1, nil, nil)
}

type noForceSendFields struct {
	A string `json:"a"`
}

func TestSetForceSendFields_RequestMustHaveForceSendFields(t *testing.T) {
	defer expectPanic(t, "request argument to setForceSendFields must have ForceSendFields field")()

	SetForceSendFields(&noForceSendFields{}, nil, nil)
}

type incorrectForceSendFields struct {
	A               string `json:"a"`
	ForceSendFields int    `json:"force_send_fields"`
}

func TestSetForceSendFields_RequestMustHaveForceSendFieldsWithCorrectType(t *testing.T) {
	defer expectPanic(t, "request argument to setForceSendFields must have ForceSendFields field of type []string (got int)")()

	SetForceSendFields(&incorrectForceSendFields{}, nil, nil)
}

type withForceSendFields struct {
	A               string   `json:"a"`
	B               string   `json:"-"`
	ForceSendFields []string `json:"force_send_fields"`
}

func TestSetForceSendFields_NoFields(t *testing.T) {
	req := &withForceSendFields{}
	SetForceSendFields(req, nil, nil)
	assert.Len(t, req.ForceSendFields, 0)
}

func TestSetForceSendFields_ForceAWhenPresent(t *testing.T) {
	req := &withForceSendFields{}
	SetForceSendFields(req, data{"a": ""}, []string{"a"})
	assert.Len(t, req.ForceSendFields, 1)
	assert.Equal(t, "A", req.ForceSendFields[0])
}

func TestSetForceSendFields_DoNotForceAWhenAbsent(t *testing.T) {
	req := &withForceSendFields{}
	SetForceSendFields(req, data{}, []string{"a"})
	assert.Len(t, req.ForceSendFields, 0)
}

func TestSetForceSendFields_DoNotDuplicate(t *testing.T) {
	req := &withForceSendFields{ForceSendFields: []string{"A"}}
	SetForceSendFields(req, data{"a": ""}, []string{"a"})
	assert.Len(t, req.ForceSendFields, 1)
	assert.Equal(t, "A", req.ForceSendFields[0])
}

func TestSetForceSendFields_CannotForceNonJsonFields(t *testing.T) {
	defer expectPanic(t, "unexpected field b not found in request structure, expected one of: a")()
	req := &withForceSendFields{}
	SetForceSendFields(req, data{"b": ""}, []string{"b"})
}

func TestSetForceSendFields_CannotForceUnknownFields(t *testing.T) {
	defer expectPanic(t, "unexpected field c not found in request structure, expected one of: a")()
	req := &withForceSendFields{}
	SetForceSendFields(req, data{"b": ""}, []string{"c"})
}
