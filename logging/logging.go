package logging

import (
	"os"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)

// Interface to mock os.Getenv() operations
type envOperations interface {
	Getenv(key string) string
}

type realEnvOperations struct{}

func (r *realEnvOperations) Getenv(key string) string {
	return os.Getenv(key)
}

// Unfortunately LogLevel() in terraform plugin sdk v2 is not test friendly so we have to mock it since it calls os.GetEnv() underneath
type logLevelGetter interface {
	GetLogLevel() string
}

type realLogLevelGetter struct{}

func (r *realLogLevelGetter) GetLogLevel() string {
	return logging.LogLevel()
}

func SetLogger() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: terraformToSDKLogLevel(&realEnvOperations{}, &realLogLevelGetter{}),
	}
}

// Golang doesn't have mutable maps, having global mutable variable is less safe as compared to a function returning map
func terraformToSDKLogLevelMapping() map[string]logger.Level {
	return map[string]logger.Level{
		"TRACE": logger.LevelTrace,
		"DEBUG": logger.LevelDebug,
		"INFO":  logger.LevelInfo,
		"WARN":  logger.LevelWarn,
		"ERROR": logger.LevelError,
	}
}

func terraformToSDKLogLevel(envOps envOperations, logGetter logLevelGetter) logger.Level {
	// Exception for unknown log levels are handled in underlying LogLevel() method
	// Defaulting to TRACE in such cases
	terraformLogLevel := logGetter.GetLogLevel()

	// If TF_LOG environment variable is not set
	if terraformLogLevel == "" {
		// If TF_ACC_LOG_PATH is set, we use TRACE log level to be consistent with LogOutput() method
		if envOps.Getenv(logging.EnvAccLogFile) != "" {
			return logger.Level(logger.LevelTrace)
		} else {
			// In all other cases we use INFO log level to be consistent with Go SDK
			return logger.Level(logger.LevelInfo)
		}
	}

	// In most cases environment variable won't be passed,
	// We should only compute the mapping function when it is passed explicitly.
	mapping := terraformToSDKLogLevelMapping()
	return mapping[terraformLogLevel]
}
