package fieldmask

import (
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	paths := []string{"name", "age", "email"}
	fieldMask := New(paths)

	if diff := cmp.Diff(fieldMask.Paths, paths); diff != "" {
		t.Errorf("New() paths mismatch (-want +got):\n%s", diff)
	}
}

func TestFieldMask_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		fieldMask FieldMask
		expected  string
		wantErr   bool
	}{
		{
			name:      "empty fields",
			fieldMask: FieldMask{Paths: []string{}},
			expected:  `""`,
		},
		{
			name:      "single field",
			fieldMask: FieldMask{Paths: []string{"name"}},
			expected:  `"name"`,
		},
		{
			name:      "multiple fields",
			fieldMask: FieldMask{Paths: []string{"name", "age", "email"}},
			expected:  `"name,age,email"`,
		},
		{
			name:      "fields with spaces",
			fieldMask: FieldMask{Paths: []string{"first name", "last name"}},
			expected:  `"first name,last name"`,
		},
		{
			name:      "fields with special characters",
			fieldMask: FieldMask{Paths: []string{"user_id", "created_at", "is_active"}},
			expected:  `"user_id,created_at,is_active"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.fieldMask.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldMask.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != tt.expected {
				t.Errorf("FieldMask.MarshalJSON() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestFieldMask_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    FieldMask
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   `""`,
			want:    FieldMask{Paths: []string{}},
			wantErr: false,
		},
		{
			name:    "single field",
			input:   `"name"`,
			want:    FieldMask{Paths: []string{"name"}},
			wantErr: false,
		},
		{
			name:    "multiple fields",
			input:   `"name,age,email"`,
			want:    FieldMask{Paths: []string{"name", "age", "email"}},
			wantErr: false,
		},
		{
			name:    "fields with spaces",
			input:   `"first name,last name"`,
			want:    FieldMask{Paths: []string{"first name", "last name"}},
			wantErr: false,
		},
		{
			name:    "fields with special characters",
			input:   `"user_id,created_at,is_active"`,
			want:    FieldMask{Paths: []string{"user_id", "created_at", "is_active"}},
			wantErr: false,
		},
		{
			name:    "trailing comma",
			input:   `"name,age,"`,
			want:    FieldMask{Paths: []string{"name", "age", ""}},
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			input:   `invalid`,
			want:    FieldMask{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FieldMask
			err := result.UnmarshalJSON([]byte(tt.input))

			if (err != nil) != tt.wantErr {
				t.Errorf("FieldMask.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(result, tt.want); diff != "" {
				t.Errorf("FieldMask.UnmarshalJSON() = %v, want %v, diff %v", result, tt.want, diff)
			}
		})
	}
}

func TestFieldMask_UnmarshalJSON_NilPointer(t *testing.T) {
	var fieldMask *FieldMask
	err := fieldMask.UnmarshalJSON([]byte(`"name,age"`))

	if err == nil {
		t.Error("FieldMask.UnmarshalJSON() on nil pointer should return error")
	}

	expectedErr := "FieldMask.UnmarshalJSON on nil pointer"
	if err.Error() != expectedErr {
		t.Errorf("FieldMask.UnmarshalJSON() error = %v, want %v", err.Error(), expectedErr)
	}
}

func TestFieldMask_EncodeValues(t *testing.T) {
	tests := []struct {
		name      string
		fieldMask FieldMask
		key       string
		want      string
		wantErr   bool
	}{
		{
			name:      "empty fields",
			fieldMask: FieldMask{Paths: []string{}},
			key:       "fields",
			want:      "",
			wantErr:   false,
		},
		{
			name:      "single field",
			fieldMask: FieldMask{Paths: []string{"name"}},
			key:       "fields",
			want:      "name",
			wantErr:   false,
		},
		{
			name:      "multiple fields",
			fieldMask: FieldMask{Paths: []string{"name", "age", "email"}},
			key:       "fields",
			want:      "name,age,email",
			wantErr:   false,
		},
		{
			name:      "fields with spaces",
			fieldMask: FieldMask{Paths: []string{"first name", "last name"}},
			key:       "fields",
			want:      "first name,last name",
			wantErr:   false,
		},
		{
			name:      "custom key",
			fieldMask: FieldMask{Paths: []string{"id", "status"}},
			key:       "select",
			want:      "id,status",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			err := tt.fieldMask.EncodeValues(tt.key, &values)

			if (err != nil) != tt.wantErr {
				t.Errorf("FieldMask.EncodeValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result := values.Get(tt.key)
			if result != tt.want {
				t.Errorf("FieldMask.EncodeValues() = %v, want %v", result, tt.want)
			}
		})
	}
}
