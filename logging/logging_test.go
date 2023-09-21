package logging

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/stretchr/testify/assert"
)

type mockEnvOperations struct {
	values map[string]string
}

func (m *mockEnvOperations) Getenv(key string) string {
	return m.values[key]
}

type mockLogLevelGetter struct {
	level string
}

func (m *mockLogLevelGetter) GetLogLevel() string {
	if m.level == "" {
		return m.level
	}
	mapping := terraformToSDKLogLevelMapping()
	if _, ok := mapping[m.level]; ok {
		return m.level
	}
	return "TRACE"
}

func TestTerraformToSDKLogLevel_NoEnvSet(t *testing.T) {
	mock := &mockEnvOperations{
		values: map[string]string{},
	}
	logMock := &mockLogLevelGetter{
		level: "",
	}
	logLevel := terraformToSDKLogLevel(mock, logMock)
	expectedLogLevel := logger.Level(logger.LevelInfo)
	actualLogLevel := logLevel
	assert.Equal(t, expectedLogLevel, actualLogLevel)
}

func TestTerraformToSDKLogLevel_TFLOGSet_FileEnvNotSet(t *testing.T) {
	TF_LOG := "DEBUG"
	mock := &mockEnvOperations{
		values: map[string]string{
			logging.EnvLog: TF_LOG,
		},
	}
	logMock := &mockLogLevelGetter{
		level: TF_LOG,
	}
	logLevel := terraformToSDKLogLevel(mock, logMock)
	expectedLogLevel := logger.Level(logger.LevelDebug)
	actualLogLevel := logLevel
	assert.Equal(t, expectedLogLevel, actualLogLevel)
}

func TestTerraformToSDKLogLevel_TFLOGSet_FileEnvSet(t *testing.T) {
	TF_LOG := "DEBUG"
	mock := &mockEnvOperations{
		values: map[string]string{
			logging.EnvLog:        TF_LOG,
			logging.EnvAccLogFile: "/path/to/file",
		},
	}
	logMock := &mockLogLevelGetter{
		level: TF_LOG,
	}
	logLevel := terraformToSDKLogLevel(mock, logMock)
	expectedLogLevel := logger.Level(logger.LevelDebug)
	actualLogLevel := logLevel
	assert.Equal(t, expectedLogLevel, actualLogLevel)
}

func TestTerraformToSDKLogLevel_TFLOGNotSet_FileEnvSet(t *testing.T) {
	TF_LOG := ""
	mock := &mockEnvOperations{
		values: map[string]string{
			logging.EnvLog:        TF_LOG,
			logging.EnvAccLogFile: "/path/to/file",
		},
	}
	logMock := &mockLogLevelGetter{
		level: TF_LOG,
	}
	logLevel := terraformToSDKLogLevel(mock, logMock)
	expectedLogLevel := logger.Level(logger.LevelTrace)
	actualLogLevel := logLevel
	assert.Equal(t, expectedLogLevel, actualLogLevel)
}

func TestTerraformToSDKLogLevel_TFLOGInvalid(t *testing.T) {
	TF_LOG := "NonValid"
	mock := &mockEnvOperations{
		values: map[string]string{
			logging.EnvLog: TF_LOG,
		},
	}
	logMock := &mockLogLevelGetter{
		level: TF_LOG,
	}
	logLevel := terraformToSDKLogLevel(mock, logMock)
	expectedLogLevel := logger.Level(logger.LevelTrace)
	actualLogLevel := logLevel
	assert.Equal(t, expectedLogLevel, actualLogLevel)
}
