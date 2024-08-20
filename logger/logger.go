package logger

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TfLogger struct {
	ctx context.Context
}

func NewTfLogger(ctx context.Context) *TfLogger {
	return &TfLogger{ctx: ctx}
}

func SetTfLogger(ctx context.Context) {
	logger.DefaultLogger = NewTfLogger(ctx)
}

// This function is always enabled because TfLogger implements the Logger interface from Go SDK and there we check
// if the logging is enabled based on level (which default to Info).
// This however isn't possible here since tflog isn't enabled / disabled based on log level.
// Omitting is done internally through the `ShouldOmit` method that filters based on logger configurations.
func (tfLogger *TfLogger) Enabled(_ context.Context, _ logger.Level) bool {
	return true
}

func (tfLogger *TfLogger) Tracef(_ context.Context, format string, v ...any) {
	tflog.Trace(tfLogger.ctx, fmt.Sprintf(format, v...), nil)
}

func (tfLogger *TfLogger) Debugf(_ context.Context, format string, v ...any) {
	tflog.Debug(tfLogger.ctx, fmt.Sprintf(format, v...), nil)
}

func (tfLogger *TfLogger) Infof(_ context.Context, format string, v ...any) {
	tflog.Info(tfLogger.ctx, fmt.Sprintf(format, v...), nil)
}

func (tfLogger *TfLogger) Warnf(_ context.Context, format string, v ...any) {
	tflog.Warn(tfLogger.ctx, fmt.Sprintf(format, v...), nil)
}

func (tfLogger *TfLogger) Errorf(_ context.Context, format string, v ...any) {
	tflog.Error(tfLogger.ctx, fmt.Sprintf(format, v...), nil)
}
