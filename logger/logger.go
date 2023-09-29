package logger

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TfLogger struct {
	Name  string
	Level logger.Level
}

func (tfLogger *TfLogger) Enabled(_ context.Context, level logger.Level) bool {
	return true
}

func (tfLogger *TfLogger) Tracef(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Trace(ctx, format, convertToMap(v))
	} else {
		tflog.SubsystemTrace(ctx, tfLogger.Name, format, convertToMap(v))
	}
}

func (tfLogger *TfLogger) Debugf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Debug(ctx, format, convertToMap(v))
	} else {
		tflog.SubsystemDebug(ctx, tfLogger.Name, format, convertToMap(v))
	}
}

func (tfLogger *TfLogger) Infof(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Info(ctx, format, convertToMap(v))
	} else {
		tflog.SubsystemInfo(ctx, tfLogger.Name, format, convertToMap(v))
	}
}

func (tfLogger *TfLogger) Warnf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Warn(ctx, format, convertToMap(v))
	} else {
		tflog.SubsystemWarn(ctx, tfLogger.Name, format, convertToMap(v))
	}
}

func (tfLogger *TfLogger) Errorf(ctx context.Context, format string, v ...any) {
	if tfLogger == nil {
		tflog.Error(ctx, format, convertToMap(v))
	} else {
		tflog.SubsystemError(ctx, tfLogger.Name, format, convertToMap(v))
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
