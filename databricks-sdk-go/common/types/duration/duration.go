package duration

import (
	"net/url"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Duration is a wrapper for durationpb.Duration to provide custom marshaling
// for JSON and URL query strings.
//
// It embeds durationpb.Duration and exposes the .AsDuration() method to
// easily convert to time.Duration.
//
// Example:
//
//	customDur := duration.New(30 * time.Second)
//	goDur := customDur.AsDuration()
type Duration struct {
	internal *durationpb.Duration
}

// New creates a custom Duration from a standard time.Duration.
func New(d time.Duration) *Duration {
	return &Duration{internal: durationpb.New(d)}
}

// AsDuration returns the underlying time.Duration value.
func (x *Duration) AsDuration() time.Duration {
	if x == nil {
		return 0
	}
	return x.internal.AsDuration()
}

// MarshalJSON implements the [json.Marshaler] interface
// by marshalling the duration as a protobuf Duration.
func (d Duration) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(d.internal)
}

// EncodeValues implements the [query.Encoder] interface by encoding the
// duration as a string, like "3.3s".
func (d Duration) EncodeValues(key string, v *url.Values) error {
	res, err := protojson.Marshal(d.internal)
	if err != nil {
		return err
	}
	// remove the quotes from the string
	queryValue := strings.Trim(string(res), "\"")
	v.Set(key, queryValue)
	return nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface. It can parse a
// duration from the protobuf Duration.
func (d *Duration) UnmarshalJSON(b []byte) error {
	var pb durationpb.Duration
	if err := protojson.Unmarshal(b, &pb); err != nil {
		return err
	}
	*d = *New(pb.AsDuration())
	return nil
}
