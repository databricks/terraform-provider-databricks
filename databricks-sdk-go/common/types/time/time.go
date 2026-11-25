package time

import (
	"net/url"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Time is a wrapper for timestamppb.Timestamp to provide custom marshaling
// for JSON and URL query strings.
//
// It embeds timestamppb.Timestamp and exposes the .AsTime() method to
// easily convert to time.Time.
//
// Example:
//
//	customTime := time.New(stdtime.Now())
//	goTime := customTime.AsTime()
type Time struct {
	internal *timestamppb.Timestamp
}

// New creates a custom Time from a standard time.Time.
func New(t time.Time) *Time {
	return &Time{internal: timestamppb.New(t)}
}

// AsTime returns the underlying time.Time value.
func (x *Time) AsTime() time.Time {
	if x == nil {
		return time.Time{}
	}
	return x.internal.AsTime()
}

// MarshalJSON implements the [json.Marshaler] interface
// by marshalling the time as a protobuf Timestamp.
func (t Time) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(t.internal)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface
// by unmarshalling the time from a protobuf Timestamp.
func (t *Time) UnmarshalJSON(b []byte) error {
	var pb timestamppb.Timestamp
	if err := protojson.Unmarshal(b, &pb); err != nil {
		return err
	}
	*t = *New(pb.AsTime())
	return nil
}

// EncodeValues implements the [query.Encoder] interface by encoding the
// time as a protobuf Timestamp.
func (t Time) EncodeValues(key string, v *url.Values) error {
	res, err := protojson.Marshal(t.internal)
	if err != nil {
		return err
	}
	// remove the quotes from the string.
	queryValue := strings.Trim(string(res), "\"")
	v.Set(key, queryValue)
	return nil
}
