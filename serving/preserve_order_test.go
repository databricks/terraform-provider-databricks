package serving

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/stretchr/testify/assert"
)

func TestReorderByName(t *testing.T) {
	t.Run("reorders items to match config order", func(t *testing.T) {
		configNames := []string{"prod", "staging", "dev"}
		apiItems := []serving.ServedModelOutput{
			{Name: "dev", ModelName: "model1"},
			{Name: "prod", ModelName: "model2"},
			{Name: "staging", ModelName: "model3"},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Equal(t, "prod", result[0].Name)
		assert.Equal(t, "staging", result[1].Name)
		assert.Equal(t, "dev", result[2].Name)
	})

	t.Run("handles alphabetical API ordering", func(t *testing.T) {
		// Simulates the actual bug: API returns alphabetically, config has different order
		configNames := []string{"prod_model", "candidate_model"}
		apiItems := []serving.ServedModelOutput{
			{Name: "candidate_model", ModelVersion: "2"},
			{Name: "prod_model", ModelVersion: "1"},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Equal(t, "prod_model", result[0].Name)
		assert.Equal(t, "candidate_model", result[1].Name)
	})

	t.Run("appends extra API items not in config", func(t *testing.T) {
		configNames := []string{"model1", "model2"}
		apiItems := []serving.ServedModelOutput{
			{Name: "model2", ModelName: "m2"},
			{Name: "model3", ModelName: "m3"}, // Not in config
			{Name: "model1", ModelName: "m1"},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Len(t, result, 3)
		assert.Equal(t, "model1", result[0].Name)
		assert.Equal(t, "model2", result[1].Name)
		assert.Equal(t, "model3", result[2].Name) // Appended at end
	})

	t.Run("handles empty config names", func(t *testing.T) {
		configNames := []string{}
		apiItems := []serving.ServedModelOutput{
			{Name: "model1"},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Equal(t, apiItems, result) // Returns original unchanged
	})

	t.Run("handles empty API items", func(t *testing.T) {
		configNames := []string{"model1"}
		apiItems := []serving.ServedModelOutput{}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Empty(t, result)
	})

	t.Run("handles nil slices", func(t *testing.T) {
		var apiItems []serving.ServedModelOutput
		result := reorderByName([]string{"model1"}, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Nil(t, result)
	})

	t.Run("maintains order when already correct", func(t *testing.T) {
		configNames := []string{"model1", "model2", "model3"}
		apiItems := []serving.ServedModelOutput{
			{Name: "model1"},
			{Name: "model2"},
			{Name: "model3"},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Equal(t, "model1", result[0].Name)
		assert.Equal(t, "model2", result[1].Name)
		assert.Equal(t, "model3", result[2].Name)
	})

	t.Run("works with served entities", func(t *testing.T) {
		configNames := []string{"entity_b", "entity_a"}
		apiItems := []serving.ServedEntityOutput{
			{Name: "entity_a", EntityName: "cat.sch.entity_a"},
			{Name: "entity_b", EntityName: "cat.sch.entity_b"},
		}

		result := reorderByName(configNames, apiItems,
			func(e serving.ServedEntityOutput) string { return e.Name })

		assert.Equal(t, "entity_b", result[0].Name)
		assert.Equal(t, "entity_a", result[1].Name)
	})

	t.Run("handles partial matches", func(t *testing.T) {
		// Some items in config don't exist in API response (shouldn't happen, but be safe)
		configNames := []string{"model1", "model2", "model3"}
		apiItems := []serving.ServedModelOutput{
			{Name: "model2"},
			{Name: "model1"},
			// model3 doesn't exist in API
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Len(t, result, 2)
		assert.Equal(t, "model1", result[0].Name)
		assert.Equal(t, "model2", result[1].Name)
	})

	t.Run("handles single item", func(t *testing.T) {
		configNames := []string{"only_model"}
		apiItems := []serving.ServedModelOutput{
			{Name: "only_model", ModelName: "m1"},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		assert.Len(t, result, 1)
		assert.Equal(t, "only_model", result[0].Name)
	})

	t.Run("preserves all item properties", func(t *testing.T) {
		configNames := []string{"model2", "model1"}
		apiItems := []serving.ServedModelOutput{
			{
				Name:               "model1",
				ModelName:          "my_model",
				ModelVersion:       "1",
				ScaleToZeroEnabled: true,
			},
			{
				Name:               "model2",
				ModelName:          "another_model",
				ModelVersion:       "3",
				ScaleToZeroEnabled: false,
			},
		}

		result := reorderByName(configNames, apiItems,
			func(m serving.ServedModelOutput) string { return m.Name })

		// Check first item (model2)
		assert.Equal(t, "model2", result[0].Name)
		assert.Equal(t, "another_model", result[0].ModelName)
		assert.Equal(t, "3", result[0].ModelVersion)
		assert.False(t, result[0].ScaleToZeroEnabled)

		// Check second item (model1)
		assert.Equal(t, "model1", result[1].Name)
		assert.Equal(t, "my_model", result[1].ModelName)
		assert.Equal(t, "1", result[1].ModelVersion)
		assert.True(t, result[1].ScaleToZeroEnabled)
	})
}
