package logger

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TfLogger struct {
	Name string
}

var loggerContext context.Context

// This function is always enabled because TfLogger implements the Logger interface from Go SDK and there we check
// if the logging is enabled based on level (which default to Info).
// This however isn't possible here since tflog isn't enabled / disabled based on log level.
// Omitting is done internally through the `ShouldOmit` method that filters based on logger configurations.
func (tfLogger *TfLogger) Enabled(_ context.Context, _ logger.Level) bool {
	return true
}

func (tfLogger *TfLogger) Tracef(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Trace(getLoggerValidContext(ctx), fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemTrace(getLoggerValidContext(ctx), tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Debugf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Debug(getLoggerValidContext(ctx), fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemDebug(getLoggerValidContext(ctx), tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Infof(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Info(getLoggerValidContext(ctx), fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemInfo(getLoggerValidContext(ctx), tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Warnf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Warn(getLoggerValidContext(ctx), fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemWarn(getLoggerValidContext(ctx), tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Errorf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Error(getLoggerValidContext(ctx), fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemError(getLoggerValidContext(ctx), tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

// Tflogger is called through Databricks provider and also Go SDK
// We need a way to check if the context is correctly set for logging, if not then we use the context we define in configureContextFunc
// We can't use the logging.GetProviderRootLogger(ctx) to check as it is internal
func getLoggerValidContext(ctx context.Context) context.Context {
	logger := ctx.Value("provider")
	if logger == nil {
		return loggerContext
	}
	return ctx
}

func SetLogger(ctx context.Context) {
	var tfLogger *TfLogger
	logger.DefaultLogger = tfLogger
	loggerContext = ctx
}
