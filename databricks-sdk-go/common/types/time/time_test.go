package time

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"
)

func TestAsTime(t *testing.T) {
	now := time.Now()
	timeVal := New(now)
	result := timeVal.AsTime()
	if !result.Equal(now) {
		t.Errorf("AsTime() = %v, want %v", result, now)
	}
}

func TestAsTime_NilPointer(t *testing.T) {
	var timeVal *Time
	result := timeVal.AsTime()
	expected := time.Time{}
	if !result.Equal(expected) {
		t.Errorf("AsTime() on nil = %v, want %v", result, expected)
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		expected string
		wantErr  bool
	}{
		{
			name:     "zero time",
			time:     *New(time.Time{}),
			expected: `"0001-01-01T00:00:00Z"`,
		},
		{
			name:     "specific time",
			time:     *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)),
			expected: `"2023-12-25T10:30:00Z"`,
		},
		{
			name:     "time with nanoseconds",
			time:     *New(time.Date(2023, 12, 25, 10, 30, 0, 123456789, time.UTC)),
			expected: `"2023-12-25T10:30:00.123456789Z"`,
		},
		{
			name:     "time with timezone",
			time:     *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.FixedZone("EST", -5*3600))),
			expected: `"2023-12-25T15:30:00Z"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.time.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != tt.expected {
				t.Errorf("Time.MarshalJSON() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Time
		wantErr bool
	}{
		{
			name:    "zero time",
			input:   `"0001-01-01T00:00:00Z"`,
			want:    *New(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "specific time",
			input:   `"2023-12-25T10:30:00Z"`,
			want:    *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "time with nanoseconds",
			input:   `"2023-12-25T10:30:00.123456789Z"`,
			want:    *New(time.Date(2023, 12, 25, 10, 30, 0, 123456789, time.UTC)),
			wantErr: false,
		},
		{
			name:    "time with timezone",
			input:   `"2023-12-25T10:30:00-05:00"`,
			want:    *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.FixedZone("", -5*3600))),
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   `""`,
			wantErr: true,
		},
		{
			name:    "invalid time format",
			input:   `"invalid-time"`,
			want:    *New(time.Time{}),
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   `"invalid-json"`,
			want:    *New(time.Time{}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var timeVal Time
			err := timeVal.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !timeVal.AsTime().Equal(tt.want.AsTime()) {
				t.Errorf("Time.UnmarshalJSON() = %v, want %v", timeVal, tt.want)
			}
		})
	}
}

func TestTime_EncodeValues(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		key      string
		expected string
	}{
		{
			name:     "zero time",
			time:     *New(time.Time{}),
			key:      "created_at",
			expected: "0001-01-01T00:00:00Z",
		},
		{
			name:     "specific time",
			time:     *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)),
			key:      "updated_at",
			expected: "2023-12-25T10:30:00Z",
		},
		{
			name:     "time with nanoseconds",
			time:     *New(time.Date(2023, 12, 25, 10, 30, 0, 123456789, time.UTC)),
			key:      "timestamp",
			expected: "2023-12-25T10:30:00.123456789Z",
		},
		{
			name:     "time with timezone",
			time:     *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.FixedZone("EST", -5*3600))),
			key:      "local_time",
			expected: "2023-12-25T15:30:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			err := tt.time.EncodeValues(tt.key, &values)
			if err != nil {
				t.Errorf("Time.EncodeValues() error = %v", err)
				return
			}
			result := values.Get(tt.key)
			if result != tt.expected {
				t.Errorf("Time.EncodeValues() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTime_JSONRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		time Time
	}{
		{
			name: "zero time",
			time: *New(time.Time{}),
		},
		{
			name: "specific time",
			time: *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)),
		},
		{
			name: "time with nanoseconds",
			time: *New(time.Date(2023, 12, 25, 10, 30, 0, 123456789, time.UTC)),
		},
		{
			name: "time with timezone",
			time: *New(time.Date(2023, 12, 25, 10, 30, 0, 0, time.FixedZone("EST", -5*3600))),
		},
		{
			name: "current time",
			time: *New(time.Now()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonData, err := json.Marshal(tt.time)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}

			// Unmarshal from JSON
			var result Time
			err = json.Unmarshal(jsonData, &result)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			// Check that the round trip preserved the value
			if !result.AsTime().Equal(tt.time.AsTime()) {
				t.Errorf("JSON round trip failed: original = %v, result = %v", tt.time, result)
			}
		})
	}
}

func TestTime_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
	}{
		{
			name: "very old time",
			time: time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "very future time",
			time: time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			name: "leap year time",
			time: time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeVal := *New(tt.time)

			jsonData, err := json.Marshal(timeVal)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}

			var result Time
			err = json.Unmarshal(jsonData, &result)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			if !result.AsTime().Equal(tt.time) {
				t.Errorf("Time round trip failed: original = %v, result = %v", tt.time, result)
			}
		})
	}
}
