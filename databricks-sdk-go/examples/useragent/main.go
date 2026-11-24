package main

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// Run this example by running `go run ./examples/useragent` from the root of this repository.
//
// To run this example, ensure that you have a DEFAULT profile configured in your `.databrickscfg`.
func main() {
	useragent.WithProduct("databricks-sdk-example", "0.0.1")
	useragent.WithPartner("databricks-sdk-example-abc")
	useragent.WithPartner("databricks-sdk-example-def")
	useragent.WithUserAgentExtra("test-key", "test-value")
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(context.Background(), iam.MeRequest{})
	if err != nil {
		panic(err)
	}
}
