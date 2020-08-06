package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParentDirPath(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		expectedDirPath string
		expectedError   error
	}{
		{
			name:            "basic_path",
			path:            "/test/abc/file.py",
			expectedDirPath: "/test/abc",
			expectedError:   nil,
		},
		{
			name:            "root_path",
			path:            "/file.py",
			expectedDirPath: "",
			expectedError:   DirPathRootDirError,
		},
		{
			name:            "empty_path",
			path:            "",
			expectedDirPath: "",
			expectedError:   PathEmptyError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dirPath, err := GetParentDirPath(tt.path)
			assert.Equal(t, tt.expectedDirPath, dirPath, "dirPath values should match")
			assert.Equal(t, tt.expectedError, err, "err values should match")
		})
	}
}
