package main

import (
	"context"
	"flag"
	"os"

	"github.com/databricks/databricks-sdk-go"
	sdk "github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var json bool

func init() {
	flag.BoolVar(&json, "json", false, "log json messages")
}

func main() {
	flag.Parse()

	var l zerolog.Logger
	if json {
		l = log.Output(os.Stderr)
	} else {
		l = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Define global logger for the SDK.
	// This instance is used when no context is present (e.g. upon construction),
	// or when a context is present but doesn't hold a logger to use.
	sdk.DefaultLogger = &zerologAdapter{l.With().Bool("global", true).Logger()}

	// Define logger on context for the SDK to use.
	ctx := sdk.NewContext(context.Background(), &zerologAdapter{
		l.With().Bool("global", false).Logger(),
	})

	// Construct client and make a request.
	// Observe a mix of global=true and global=false in the log output.
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
	if err != nil {
		panic(err)
	}
}
