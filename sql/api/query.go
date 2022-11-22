package api

import (
	"encoding/json"
	"fmt"
)

// Query ...
type Query struct {
	ID             string            `json:"id,omitempty"`
	DataSourceID   string            `json:"data_source_id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Query          string            `json:"query"`
	Schedule       *QuerySchedule    `json:"schedule"`
	Options        *QueryOptions     `json:"options,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Visualizations []json.RawMessage `json:"visualizations,omitempty"`
}

// QuerySchedule ...
type QuerySchedule struct {
	// Interval in seconds.
	//
	// For daily schedules, this MUST be a multiple of 86400.
	// For weekly schedules, this MUST be a multiple of 604800.
	//
	Interval int `json:"interval"`

	// Time of day, for daily and weekly schedules.
	Time *string `json:"time"`

	// Day of week, for weekly schedules.
	DayOfWeek *string `json:"day_of_week"`

	// Schedule should be active until this date.
	Until *string `json:"until"`
}

// QueryOptions ...
type QueryOptions struct {
	Parameters    []any             `json:"-"`
	RawParameters []json.RawMessage `json:"parameters,omitempty"`

	RunAsRole string `json:"run_as_role,omitempty"`
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

	o.Parameters = []any{}
	for _, rp := range o.RawParameters {
		var qp QueryParameter

		// Unmarshal into base parameter type to figure out the right type.
		err = json.Unmarshal(rp, &qp)
		if err != nil {
			return err
		}

		// Acquire pointer to the correct parameter type.
		var i any
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
		case queryParameterDateRangeTypeName:
			i = &QueryParameterDateRange{}
		case queryParameterDateTimeRangeTypeName:
			i = &QueryParameterDateTimeRange{}
		case queryParameterDateTimeSecRangeTypeName:
			i = &QueryParameterDateTimeSecRange{}
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
	queryParameterTextTypeName             = "text"
	queryParameterNumberTypeName           = "number"
	queryParameterEnumTypeName             = "enum"
	queryParameterQueryTypeName            = "query"
	queryParameterDateTypeName             = "date"
	queryParameterDateTimeTypeName         = "datetime-local"
	queryParameterDateTimeSecTypeName      = "datetime-with-seconds"
	queryParameterDateRangeTypeName        = "date-range"
	queryParameterDateTimeRangeTypeName    = "datetime-range"
	queryParameterDateTimeSecRangeTypeName = "datetime-range-with-seconds"
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

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterText) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterText
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.Type = ""
	return nil
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

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterNumber) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterNumber
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.Type = ""
	return nil
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

	Values []string `json:"-"`

	Value   json.RawMessage                      `json:"value"`
	Options string                               `json:"enumOptions"`
	Multi   *QueryParameterMultipleValuesOptions `json:"multiValuesOptions,omitempty"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterEnum) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterEnumTypeName

	// Set `Value` depending on multiple options being allowed or not.
	var err error
	if p.Multi == nil {
		// Set as single string.
		p.Value, err = json.Marshal(p.Values[0])
		if err != nil {
			return nil, err
		}
	} else {
		// Set as array of strings.
		p.Value, err = json.Marshal(p.Values)
		if err != nil {
			return nil, err
		}
	}

	type localQueryParameter QueryParameterEnum
	return json.Marshal((localQueryParameter)(p))
}

// UnmarshalJSON deals with polymorphism of the `value` field.
func (p *QueryParameterEnum) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterEnum
	err := json.Unmarshal(b, (*localQueryParameter)(p))
	if err != nil {
		return err
	}

	// If multiple options aren't configured, assume `value` is a string.
	// Otherwise, it's an array of strings.
	if p.Multi == nil {
		var v string
		err = json.Unmarshal(p.Value, &v)
		if err != nil {
			return nil
		}
		p.Values = []string{v}
	} else {
		var vs []string
		err = json.Unmarshal(p.Value, &vs)
		if err != nil {
			return nil
		}
		p.Values = vs
	}

	p.Type = ""
	p.Value = nil
	return nil
}

// QueryParameterQuery ...
type QueryParameterQuery struct {
	QueryParameter

	Values []string `json:"-"`

	Value   json.RawMessage                      `json:"value"`
	QueryID string                               `json:"queryId"`
	Multi   *QueryParameterMultipleValuesOptions `json:"multiValuesOptions,omitempty"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterQuery) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterQueryTypeName

	// Set `Value` depending on multiple options being allowed or not.
	var err error
	if p.Multi == nil {
		// Set as single string.
		p.Value, err = json.Marshal(p.Values[0])
		if err != nil {
			return nil, err
		}
	} else {
		// Set as array of strings.
		p.Value, err = json.Marshal(p.Values)
		if err != nil {
			return nil, err
		}
	}

	type localQueryParameter QueryParameterQuery
	return json.Marshal((localQueryParameter)(p))
}

// UnmarshalJSON deals with polymorphism of the `value` field.
func (p *QueryParameterQuery) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterQuery
	err := json.Unmarshal(b, (*localQueryParameter)(p))
	if err != nil {
		return err
	}

	// If multiple options aren't configured, assume `value` is a string.
	// Otherwise, it's an array of strings.
	if p.Multi == nil {
		var v string
		err = json.Unmarshal(p.Value, &v)
		if err != nil {
			return nil
		}
		p.Values = []string{v}
	} else {
		var vs []string
		err = json.Unmarshal(p.Value, &vs)
		if err != nil {
			return nil
		}
		p.Values = vs
	}

	p.Type = ""
	p.Value = nil
	return nil
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

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterDate) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterDate
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.Type = ""
	return nil
}

// QueryParameterDateTime ...
type QueryParameterDateTime struct {
	QueryParameter

	Value       any    `json:"value"`
	StringValue string `json:"-"`
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDateTime) MarshalJSON() ([]byte, error) {
	p.QueryParameter.Type = queryParameterDateTimeTypeName
	type localQueryParameter QueryParameterDateTime
	return json.Marshal((localQueryParameter)(p))
}

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterDateTime) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterDateTime
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.Type = ""
	return nil
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

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterDateTimeSec) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterDateTimeSec
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.Type = ""
	return nil
}

type DateTimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type QueryParameterRangeBase struct {
	QueryParameter

	Value       any            `json:"value"`
	StringValue string         `json:"-"`
	RangeValue  *DateTimeRange `json:"-"`
}

func (p *QueryParameterRangeBase) toParameterObject() {
	if p.RangeValue != nil {
		p.Value = p.RangeValue
	} else {
		p.Value = p.StringValue
	}
}

func (p *QueryParameterRangeBase) decodeQueryParameter() {
	if v, ok := p.Value.(map[string]any); ok {
		p.RangeValue = &DateTimeRange{Start: v["start"].(string), End: v["end"].(string)}
	} else {
		p.StringValue = fmt.Sprintf("%v", p.Value)
	}
	p.Value = nil
}

// QueryParameterDateRange ...
type QueryParameterDateRange struct {
	QueryParameterRangeBase
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDateRange) MarshalJSON() ([]byte, error) {
	type localQueryParameter QueryParameterDateRange
	p.QueryParameter.Type = queryParameterDateRangeTypeName
	p.toParameterObject()
	return json.Marshal((localQueryParameter)(p))
}

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterDateRange) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterDateRange
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.decodeQueryParameter()
	p.Type = ""
	return nil
}

// QueryParameterDateTimeRange ...
type QueryParameterDateTimeRange struct {
	QueryParameterRangeBase
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDateTimeRange) MarshalJSON() ([]byte, error) {
	type localQueryParameter QueryParameterDateTimeRange
	p.QueryParameter.Type = queryParameterDateTimeRangeTypeName
	p.toParameterObject()
	return json.Marshal((localQueryParameter)(p))
}

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterDateTimeRange) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterDateTimeRange
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.decodeQueryParameter()
	p.Type = ""
	return nil
}

// QueryParameterDateTimeSecRange ...
type QueryParameterDateTimeSecRange struct {
	QueryParameterRangeBase
}

// MarshalJSON sets the type before marshaling.
func (p QueryParameterDateTimeSecRange) MarshalJSON() ([]byte, error) {
	type localQueryParameter QueryParameterDateTimeSecRange
	p.QueryParameter.Type = queryParameterDateTimeSecRangeTypeName
	p.toParameterObject()
	return json.Marshal((localQueryParameter)(p))
}

// UnmarshalJSON clears the type after marshaling.
func (p *QueryParameterDateTimeSecRange) UnmarshalJSON(b []byte) error {
	type localQueryParameter QueryParameterDateTimeSecRange
	if err := json.Unmarshal(b, (*localQueryParameter)(p)); err != nil {
		return err
	}
	p.decodeQueryParameter()
	p.Type = ""
	return nil
}
