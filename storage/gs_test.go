package storage

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrValidateClusterForGoogleStorage_Failures(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceMount().TestResourceData()
		err := createOrValidateClusterForGoogleStorage(ctx, client, d, "a", "")
		assert.EqualError(t, err, "cannot re-create mounting cluster: nope")

		err = createOrValidateClusterForGoogleStorage(ctx, client, d, "", "b")
		assert.EqualError(t, err, "cannot create mounting cluster: nope")
	})
}
