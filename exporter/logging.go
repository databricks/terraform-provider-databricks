package exporter

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	sdklogger "github.com/databricks/databricks-sdk-go/logger"
)

type levelWriter struct {
	prefixes []string
	output   io.Writer
}

type exporterLogLevel struct {
	name     string
	sdkLevel sdklogger.Level
	prefixes []string
}

func (lw *levelWriter) Write(p []byte) (n int, err error) {
	a := string(p)
	msg := strings.TrimSuffix(a, "\n")
	for _, l := range lw.prefixes {
		if strings.HasPrefix(a, l) {
			timeStr := time.Now().Local().Format(time.RFC3339Nano)
			logStr := msg[0:len(l)] + " " + timeStr + ": " + strings.TrimLeft(msg[len(l):], " ") + "\n"
			output := lw.output
			if output == nil {
				output = os.Stdout
			}
			_, err := output.Write([]byte(logStr))
			return len(p), err
		}
	}
	return len(p), nil
}

func configureExporterLogging(debug, trace bool) exporterLogLevel {
	level, warning := determineExporterLogLevel(debug, trace, os.Getenv("TF_LOG"))
	log.SetOutput(&levelWriter{
		prefixes: level.prefixes,
		output:   os.Stdout,
	})
	sdklogger.DefaultLogger = &sdklogger.SimpleLogger{Level: level.sdkLevel}
	if warning != "" {
		log.Printf("[WARN] %s", warning)
	}
	return level
}

func determineExporterLogLevel(debug, trace bool, tfLog string) (exporterLogLevel, string) {
	if trace {
		return traceExporterLogLevel(), ""
	}
	if debug {
		return debugExporterLogLevel(), ""
	}

	switch strings.ToUpper(strings.TrimSpace(tfLog)) {
	case "":
		return infoExporterLogLevel(), ""
	case "ERROR":
		return exporterLogLevel{
			name:     "ERROR",
			sdkLevel: sdklogger.LevelError,
			prefixes: []string{"[ERROR]"},
		}, ""
	case "WARN":
		return exporterLogLevel{
			name:     "WARN",
			sdkLevel: sdklogger.LevelWarn,
			prefixes: []string{"[WARN]", "[ERROR]"},
		}, ""
	case "INFO":
		return infoExporterLogLevel(), ""
	case "DEBUG":
		return debugExporterLogLevel(), ""
	case "TRACE":
		return traceExporterLogLevel(), ""
	default:
		return infoExporterLogLevel(), "Invalid TF_LOG value " + strconv.Quote(tfLog) + "; defaulting exporter logging to INFO. Valid values are: TRACE, DEBUG, INFO, WARN, ERROR"
	}
}

func infoExporterLogLevel() exporterLogLevel {
	return exporterLogLevel{
		name:     "INFO",
		sdkLevel: sdklogger.LevelInfo,
		prefixes: []string{"[INFO]", "[WARN]", "[ERROR]"},
	}
}

func debugExporterLogLevel() exporterLogLevel {
	return exporterLogLevel{
		name:     "DEBUG",
		sdkLevel: sdklogger.LevelDebug,
		prefixes: []string{"[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
	}
}

func traceExporterLogLevel() exporterLogLevel {
	return exporterLogLevel{
		name:     "TRACE",
		sdkLevel: sdklogger.LevelTrace,
		prefixes: []string{"[TRACE]", "[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
	}
}
