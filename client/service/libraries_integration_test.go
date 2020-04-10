package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLibraryCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	libraries := []model.Library{
		model.Library{
			Pypi: &model.PyPi{
				Package: "networkx",
			},
		},
	}

	err := client.Libraries().Create("0406-055906-bunch228", libraries)
	assert.NoError(t, err, err)

	err = client.Libraries().Delete("0406-055906-bunch228", libraries)
	assert.NoError(t, err, err)

	libraryStatusList, err := client.Libraries().List("0406-055906-bunch228")
	assert.NoError(t, err, err)

	t.Log(libraryStatusList)
	t.Log(libraryStatusList[0].Library.Pypi)

}
