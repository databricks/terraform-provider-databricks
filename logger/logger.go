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

func (tfLogger *TfLogger) Tracef(_ context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Trace(loggerContext, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemTrace(loggerContext, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Debugf(_ context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Debug(loggerContext, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemDebug(loggerContext, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Infof(_ context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Info(loggerContext, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemInfo(loggerContext, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Warnf(_ context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Warn(loggerContext, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemWarn(loggerContext, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Errorf(_ context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Error(loggerContext, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemError(loggerContext, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func SetLogger(ctx context.Context) {
	var tfLogger *TfLogger
	logger.DefaultLogger = tfLogger
	loggerContext = ctx
}
