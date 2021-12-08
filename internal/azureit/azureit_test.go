package main

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       200,
			Response:     `{}`,
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		responseWriter := httptest.NewRecorder()
		azure.PublicCloud.ResourceManagerEndpoint = client.Host
		os.Setenv("MSI_ENDPOINT", client.Host)
		os.Setenv("MSI_SECRET", "secret")
		os.Setenv("ACI_CONTAINER_GROUP", "")
		triggerStart(responseWriter, nil)
		assert.Equal(t, "400 Bad Request", responseWriter.Result().Status)

		responseWriter = httptest.NewRecorder()
		os.Setenv("ACI_CONTAINER_GROUP", "/abc")
		triggerStart(responseWriter, nil)
		assert.Equal(t, "200 OK", responseWriter.Result().Status)

		// test that app properly fails
		os.Setenv("FUNCTIONS_CUSTOMHANDLER_PORT", "abc")
		defer func() {
			err := recover()
			require.NotNil(t, err)
		}()
		main()
	})
}
