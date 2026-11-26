package main

import (
	"context"
	"flag"
	"os"

	"github.com/databricks/databricks-sdk-go"
	sdk "github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"golang.org/x/exp/slog"
)

var json bool

func init() {
	flag.BoolVar(&json, "json", false, "log json messages")
}

func main() {
	flag.Parse()

	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	var handler slog.Handler
	if json {
		handler = opts.NewJSONHandler(os.Stderr)
	} else {
		handler = opts.NewTextHandler(os.Stderr)
	}

	// Define global logger for the SDK.
	// This instance is used when no context is present (e.g. upon construction),
	// or when a context is present but doesn't hold a logger to use.
	sdk.DefaultLogger = &slogAdapter{
		slog.New(handler).With("global", true),
	}

	// Define logger on context for the SDK to use.
	ctx := sdk.NewContext(context.Background(), &slogAdapter{
		slog.New(handler).With("global", false),
	})

	// Construct client and make a request.
	// Observe a mix of global=true and global=false in the log output.
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
	if err != nil {
		panic(err)
	}
}
