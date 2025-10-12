package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []catalog.ArtifactMatcher `json:"artifact_matchers" tf:"slice_set,alias:artifact_matcher"`
	// The artifact type of the allowlist.
	ArtifactType catalog.ArtifactType `json:"artifact_type" tf:"force_new"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty" tf:"computed"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `json:"created_by,omitempty" tf:"computed"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty" tf:"computed"`
}

func ResourceArtifactAllowlist() common.Resource {
	allowlistSchema := common.StructToSchema(ArtifactAllowlistInfo{}, common.NoCustomize)
	p := common.NewPairID("metastore_id", "artifact_type")

	createOrUpdate := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
		if err != nil {
			return err
		}
		err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
		if err != nil {
			return err
		}

		var createAllowlist ArtifactAllowlistInfo
		common.DataToStructPointer(d, allowlistSchema, &createAllowlist)

		al, err := w.ArtifactAllowlists.Update(ctx, catalog.SetArtifactAllowlist{
			ArtifactMatchers: createAllowlist.ArtifactMatchers,
			ArtifactType:     createAllowlist.ArtifactType,
		})

		if err != nil {
			return err
		}

		d.Set("metastore_id", al.MetastoreId)
		p.Pack(d)
		return nil
	}
	return common.Resource{
		Schema: allowlistSchema,
		Create: createOrUpdate,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			_, artifactType, err := p.Unpack(d)
			if err != nil {
				return err
			}

			al, err := w.ArtifactAllowlists.GetByArtifactType(ctx, catalog.ArtifactType(artifactType))
			if err != nil {
				return err
			}

			allowlist := ArtifactAllowlistInfo{
				ArtifactMatchers: al.ArtifactMatchers,
				CreatedAt:        al.CreatedAt,
				CreatedBy:        al.CreatedBy,
				MetastoreId:      al.MetastoreId,
				ArtifactType:     catalog.ArtifactType(artifactType),
			}

			return common.StructToData(allowlist, allowlistSchema, d)
		},
		Update: createOrUpdate,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}

			_, artifactType, err := p.Unpack(d)
			if err != nil {
				return err
			}

			_, err = w.ArtifactAllowlists.Update(ctx, catalog.SetArtifactAllowlist{
				ArtifactType:     catalog.ArtifactType(artifactType),
				ArtifactMatchers: []catalog.ArtifactMatcher{},
			})
			if err != nil {
				return err
			}

			return nil
		},
	}
}
