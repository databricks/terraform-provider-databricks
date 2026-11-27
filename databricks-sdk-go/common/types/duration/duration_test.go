package duration

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"
)

func TestAsDuration(t *testing.T) {
	d := time.Second * 5
	dur := New(d)
	result := dur.AsDuration()
	if result != d {
		t.Errorf("AsDuration() = %v, want %v", result, d)
	}
}

func TestDuration_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		expected string
		wantErr  bool
	}{
		{
			name:     "zero duration",
			duration: *New(0),
			expected: "0s",
		},
		{
			name:     "positive duration",
			duration: *New(5 * time.Second),
			expected: "5s",
		},
		{
			name:     "negative duration",
			duration: *New(-2 * time.Minute),
			expected: "-120s",
		},
		{
			name:     "negative duration with fractional seconds",
			duration: *New(-2*time.Minute + 100*time.Millisecond),
			expected: "-119.900s",
		},
		{
			name:     "fractional seconds",
			duration: *New(1500 * time.Millisecond),
			expected: "1.500s",
		},
		{
			name:     "large duration",
			duration: *New(9223372036*time.Second + 854775000*time.Nanosecond),
			expected: "9223372036.854775s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.duration.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != `"`+tt.expected+`"` {
				t.Errorf("Duration.MarshalJSON() = %v, want %v", string(result), `"`+tt.expected+`"`)
			}
		})
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Duration
		wantErr bool
	}{
		{
			name:    "zero duration",
			input:   `"0s"`,
			want:    *New(0),
			wantErr: false,
		},
		{
			name:    "positive duration",
			input:   `"5s"`,
			want:    *New(5 * time.Second),
			wantErr: false,
		},
		{
			name:    "negative duration",
			input:   `"-2s"`,
			want:    *New(-2 * time.Second),
			wantErr: false,
		},
		{
			name:    "negative duration with fractional seconds",
			input:   `"-2.1s"`,
			want:    *New(-2*time.Second - 100*time.Millisecond),
			wantErr: false,
		},
		{
			name:    "fractional seconds",
			input:   `"1.5s"`,
			want:    *New(1500 * time.Millisecond),
			wantErr: false,
		},
		{
			name:    "large duration",
			input:   `"9223372036.854775000s"`,
			want:    *New(9223372036*time.Second + 854775000*time.Nanosecond),
			wantErr: false,
		},
		{
			name:    "invalid duration format",
			input:   `"invalid"`,
			want:    *New(0),
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   `""`,
			want:    *New(0),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Duration
			err := d.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// We cannot compare Proto messages directly, so we compare the underlying time.Duration
				if d.AsDuration() != tt.want.AsDuration() {
					t.Errorf("Duration.UnmarshalJSON() = %v, want %v", d.AsDuration(), tt.want.AsDuration())
				}
			}

		})
	}
}

func TestDuration_EncodeValues(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		key      string
		expected string
	}{
		{
			name:     "zero duration",
			duration: *New(0),
			key:      "duration",
			expected: "0s",
		},
		{
			name:     "positive duration",
			duration: *New(5 * time.Second),
			key:      "timeout",
			expected: "5s",
		},
		{
			name:     "negative duration",
			duration: *New(-2 * time.Minute),
			key:      "delay",
			expected: "-120s",
		},
		{
			name:     "fractional seconds",
			duration: *New(1500 * time.Millisecond),
			key:      "interval",
			expected: "1.500s",
		},
		{
			name:     "large duration",
			duration: *New(24 * time.Hour),
			key:      "period",
			expected: "86400s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			err := tt.duration.EncodeValues(tt.key, &values)
			if err != nil {
				t.Errorf("Duration.EncodeValues() error = %v", err)
				return
			}
			result := values.Get(tt.key)
			if result != tt.expected {
				t.Errorf("Duration.EncodeValues() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDuration_JSONRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
	}{
		{
			name:     "zero duration",
			duration: *New(0),
		},
		{
			name:     "positive duration",
			duration: *New(5 * time.Second),
		},
		{
			name:     "negative duration",
			duration: *New(-2 * time.Minute),
		},
		{
			name:     "fractional seconds",
			duration: *New(1500 * time.Millisecond),
		},
		{
			name:     "complex duration",
			duration: *New(1*time.Hour + 2*time.Minute + 3*time.Second),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonData, err := json.Marshal(tt.duration)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}

			// Unmarshal from JSON
			var result Duration
			err = json.Unmarshal(jsonData, &result)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			// Check that the round trip preserved the value
			if result.AsDuration() != tt.duration.AsDuration() {
				t.Errorf("Duration.UnmarshalJSON() = %v, want %v", result.AsDuration(), tt.duration.AsDuration())
			}

		})
	}
}
