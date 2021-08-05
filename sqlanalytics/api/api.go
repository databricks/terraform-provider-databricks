// Package api contains API structures for use with the Databricks SQL API.
//
// These are not intended to be used directly.
//
// They are kept in a separate package to avoid confusion between the
// Terraform resources and the API objects. While their structure is
// very similar, there are a handful of nuanced differences to improve
// the UX of the Terraform resource.
//
package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// stringOrInt is a type wrapper for a JSON value that can either be encoded
// as a Javascript number (integer) or a Javascript string (UUID).
// Also see `Widget` and `Visualization`.
type stringOrInt string

func NewStringOrInt(s string) stringOrInt {
	return stringOrInt(s)
}

func (s stringOrInt) String() string {
	return string(s)
}

func (s stringOrInt) MarshalJSON() ([]byte, error) {
	i, err := strconv.Atoi(string(s))
	if err == nil {
		return json.Marshal(i)
	}

	return json.Marshal(string(s))
}

func (s *stringOrInt) UnmarshalJSON(b []byte) error {
	var tmp interface{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	switch v := tmp.(type) {
	case float64:
		*s = stringOrInt(strconv.Itoa(int(v)))
	case string:
		*s = stringOrInt(v)
	default:
		return fmt.Errorf("json: expected to unmarshal a string or an int, got %T", v)
	}

	return nil
}
