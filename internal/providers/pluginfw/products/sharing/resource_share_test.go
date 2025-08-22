package sharing

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestShareSyncEffectiveFields(t *testing.T) {
	shareName := "test-share-name"
	shares := ShareResource{}
	plan := ShareInfoExtended{
		ShareInfo_SdkV2: sharing_tf.ShareInfo_SdkV2{
			Name: types.StringValue(shareName),
		},
	}
	state := ShareInfoExtended{
		ShareInfo_SdkV2: sharing_tf.ShareInfo_SdkV2{
			Name: types.StringValue(shareName),
		},
	}
	_, diagnostics := shares.syncEffectiveFields(context.Background(), plan, state, effectiveFieldsActionRead{})
	assert.False(t, diagnostics.HasError())
}
