package logger

import (
	"context"
	"fmt"
	"strconv"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TfLogger struct {
	Name string
}

// This function is always enabled because we extend the Logger from Go SDK and there we check
// the logging is enabled based on level (which default to Info). This however isn't required in terraform
// since we use the tflog package which automatically reads from TF_LOG environment variable.
func (tfLogger *TfLogger) Enabled(_ context.Context, _ logger.Level) bool {
	return true
}

func (tfLogger *TfLogger) Tracef(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Trace(ctx, fmt.Sprintf(format, v...), convertToMap(v))
	} else {
		tflog.SubsystemTrace(ctx, tfLogger.Name, fmt.Sprintf(format, v...), convertToMap(v))
	}
}

func (tfLogger *TfLogger) Debugf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Debug(ctx, fmt.Sprintf(format, v...), convertToMap(v))
	} else {
		tflog.SubsystemDebug(ctx, tfLogger.Name, fmt.Sprintf(format, v...), convertToMap(v))
	}
}

func (tfLogger *TfLogger) Infof(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Info(ctx, fmt.Sprintf(format, v...), convertToMap(v))
	} else {
		tflog.SubsystemInfo(ctx, tfLogger.Name, fmt.Sprintf(format, v...), convertToMap(v))
	}
}

func (tfLogger *TfLogger) Warnf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Warn(ctx, fmt.Sprintf(format, v...), convertToMap(v))
	} else {
		tflog.SubsystemWarn(ctx, tfLogger.Name, fmt.Sprintf(format, v...), convertToMap(v))
	}
}

func (tfLogger *TfLogger) Errorf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Error(ctx, fmt.Sprintf(format, v...), convertToMap(v))
	} else {
		tflog.SubsystemError(ctx, tfLogger.Name, fmt.Sprintf(format, v...), convertToMap(v))
	}
}

func convertToMap(args ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i, arg := range args {
		key := strconv.Itoa(i)
		m[key] = arg
	}
	return m
}

func SetLogger() {
	logger.DefaultLogger = &TfLogger{}
}
