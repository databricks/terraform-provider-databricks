package sharing

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildAddUpdates(n int) []sharing.SharedDataObjectUpdate {
	updates := make([]sharing.SharedDataObjectUpdate, n)
	for i := range updates {
		updates[i] = sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &sharing.SharedDataObject{Name: fmt.Sprintf("obj-%d", i)},
		}
	}
	return updates
}

func TestUpdateShareInBatches(t *testing.T) {
	ctx := context.Background()
	for _, tc := range []struct {
		name       string
		numObjects int
		wantCalls  []int // expected number of object updates per call
	}{
		{"empty", 0, []int{0}},
		{"below limit", 50, []int{50}},
		{"exactly at limit", 100, []int{100}},
		{"just above limit", 101, []int{100, 1}},
		{"uneven multiple", 250, []int{100, 100, 50}},
		{"exact multiple", 200, []int{100, 100}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var gotCalls []int
			var gotNames []string
			updateFn := func(_ context.Context, u sharing.UpdateShare) (*sharing.ShareInfo, error) {
				// Every batch must carry the non-update fields.
				assert.Equal(t, "s", u.Name)
				assert.Equal(t, "o", u.Owner)
				gotCalls = append(gotCalls, len(u.Updates))
				for _, up := range u.Updates {
					gotNames = append(gotNames, up.DataObject.Name)
				}
				return &sharing.ShareInfo{Name: u.Name}, nil
			}
			update := sharing.UpdateShare{Name: "s", Owner: "o", Updates: buildAddUpdates(tc.numObjects)}

			info, err := updateShareInBatches(ctx, update, updateFn)

			require.NoError(t, err)
			require.NotNil(t, info)
			assert.Equal(t, "s", info.Name)
			assert.Equal(t, tc.wantCalls, gotCalls)

			// Order must be preserved across batches, without dropping any object.
			var wantNames []string
			for _, u := range update.Updates {
				wantNames = append(wantNames, u.DataObject.Name)
			}
			assert.Equal(t, wantNames, gotNames)
		})
	}
}

func TestUpdateShareInBatchesStopsOnError(t *testing.T) {
	ctx := context.Background()
	wantErr := fmt.Errorf("boom")
	calls := 0
	updateFn := func(_ context.Context, u sharing.UpdateShare) (*sharing.ShareInfo, error) {
		calls++
		if calls == 2 {
			return nil, wantErr
		}
		return &sharing.ShareInfo{Name: u.Name}, nil
	}

	_, err := updateShareInBatches(ctx, sharing.UpdateShare{Name: "s", Updates: buildAddUpdates(250)}, updateFn)

	assert.ErrorIs(t, err, wantErr)
	assert.Equal(t, 2, calls) // stops after the failing batch, does not issue the third
}
