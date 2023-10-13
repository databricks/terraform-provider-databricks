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

// This function is always enabled because TfLogger implements the Logger interface from Go SDK and there we check
// if the logging is enabled based on level (which default to Info).
// This however isn't possible here since tflog isn't enabled / disabled based on log level.
// Omitting is done internally through the `ShouldOmit` method that filters based on logger configurations.
func (tfLogger *TfLogger) Enabled(_ context.Context, _ logger.Level) bool {
	return true
}

func (tfLogger *TfLogger) Tracef(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Trace(ctx, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemTrace(ctx, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Debugf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Debug(ctx, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemDebug(ctx, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Infof(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Info(ctx, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemInfo(ctx, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Warnf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Warn(ctx, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemWarn(ctx, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func (tfLogger *TfLogger) Errorf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Error(ctx, fmt.Sprintf(format, v...), nil)
	} else {
		tflog.SubsystemError(ctx, tfLogger.Name, fmt.Sprintf(format, v...), nil)
	}
}

func SetLogger() {
	logger.DefaultLogger = &TfLogger{}
}
