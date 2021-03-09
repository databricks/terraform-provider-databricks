package api

import (
	"encoding/json"
)

// Query ...
type Query struct {
	ID             string            `json:"id,omitempty"`
	DataSourceID   string            `json:"data_source_id"`
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	Query          string            `json:"query"`
	Schedule       *QuerySchedule    `json:"schedule,omitempty"`
	Options        *QueryOptions     `json:"options,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Visualizations []json.RawMessage `json:"visualizations,omitempty"`
}

// QuerySchedule ...
type QuerySchedule struct {
	Interval int `json:"interval"`
}

// QueryOptions ...
type QueryOptions struct {
	Parameters    []interface{}     `json:"-"`
	RawParameters []json.RawMessage `json:"parameters,omitempty"`
}

// MarshalJSON ...
func (o *QueryOptions) MarshalJSON() ([]byte, error) {
	if o.Parameters != nil {
		o.RawParameters = []json.RawMessage{}
		for _, p := range o.Parameters {
			b, err := json.Marshal(p)
			if err != nil {
				return nil, err
			}
			o.RawParameters = append(o.RawParameters, b)
		}
	}

	type localQueryOptions QueryOptions
	return json.Marshal((*localQueryOptions)(o))
}

// UnmarshalJSON ...
func (o *QueryOptions) UnmarshalJSON(b []byte) error {
	type localQueryOptions QueryOptions
	err := json.Unmarshal(b, (*localQueryOptions)(o))
	if err != nil {
		return err
	}

	o.Parameters = []interface{}{}
	for _, rp := range o.RawParameters {
		var qp QueryParameter

		// Unmarshal into base parameter type to figure out the right type.
		err = json.Unmarshal(rp, &qp)
		if err != nil {
			return err
		}

		// Acquire pointer to the correct parameter type.
		var i interface{}
		switch qp.Type {
		case queryParameterTextTypeName:
			i = &QueryParameterText{}
		case queryParameterNumberTypeName:
			i = &QueryParameterNumber{}
		case queryParameterEnumTypeName:
			i = &QueryParameterEnum{}
		case queryParameterQueryTypeName:
			i = &QueryParameterQuery{}
		case queryParameterDateTypeName:
			i = &QueryParameterDate{}
		case queryParameterDateTimeTypeName:
			i = &QueryParameterDateTime{}
		case queryParameterDateTimeSecTypeName:
			i = &QueryParameterDateTimeSec{}
		default:
			panic("don't know what to do...")
		}

		// Unmarshal into correct parameter type.
		err = json.Unmarshal(rp, &i)
		if err != nil {
			return err
		}

		o.Parameters = append(o.Parameters, i)
	}

	return nil
}

// QueryParameter ...
type QueryParameter struct {
	Name  string `json:"name"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type"`
}

// Valid type values.
const (
	queryParameterTextTypeName        = "text"
	queryParameterNumberTypeName      = "number"
	queryParameterEnumTypeName        = "enum"
	queryParameterQueryTypeName       = "query"
	queryParameterDateTypeName        = "date"
	queryParameterDateTimeTypeName    = "datetime-local"
	queryParameterDateTimeSecTypeName = "datetime-with-seconds"
)

// QueryParameterText ...
type QueryParameterText struct {
	QueryParameter

	Value string `json:"value"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterText) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterTextTypeName
	type localQueryParameter QueryParameterText
	return json.Marshal((localQueryParameter)(p))
}

// QueryParameterNumber ...
type QueryParameterNumber struct {
	QueryParameter

	Value float64 `json:"value"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterNumber) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterNumberTypeName
	type localQueryParameter QueryParameterNumber
	return json.Marshal((localQueryParameter)(p))
}

// QueryParameterMultipleValuesOptions ...
type QueryParameterMultipleValuesOptions struct {
	Prefix    string `json:"prefix"`
	Suffix    string `json:"suffix"`
	Separator string `json:"separator"`
}

// QueryParameterEnum ...
type QueryParameterEnum struct {
	QueryParameter

	Value   string                               `json:"value"`
	Options string                               `json:"enumOptions"`
	Multi   *QueryParameterMultipleValuesOptions `json:"multiValuesOptions,omitempty"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterEnum) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterEnumTypeName
	type localQueryParameter QueryParameterEnum
	return json.Marshal((localQueryParameter)(p))
}

// QueryParameterQuery ...
type QueryParameterQuery struct {
	QueryParameter

	Value   string                               `json:"value"`
	QueryID string                               `json:"queryId"`
	Multi   *QueryParameterMultipleValuesOptions `json:"multiValuesOptions,omitempty"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterQuery) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterQueryTypeName
	type localQueryParameter QueryParameterQuery
	return json.Marshal((localQueryParameter)(p))
}

// QueryParameterDate ...
type QueryParameterDate struct {
	QueryParameter

	Value string `json:"value"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDate) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterDateTypeName
	type localQueryParameter QueryParameterDate
	return json.Marshal((localQueryParameter)(p))
}

// QueryParameterDateTime ...
type QueryParameterDateTime struct {
	QueryParameter

	Value string `json:"value"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDateTime) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterDateTimeTypeName
	type localQueryParameter QueryParameterDateTime
	return json.Marshal((localQueryParameter)(p))
}

// QueryParameterDateTimeSec ...
type QueryParameterDateTimeSec struct {
	QueryParameter

	Value string `json:"value"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDateTimeSec) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterDateTimeSecTypeName
	type localQueryParameter QueryParameterDateTimeSec
	return json.Marshal((localQueryParameter)(p))
}
