package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessAzureWasbAbfssUrisCorrectlySplitsURI(t *testing.T) {
	testCases := []struct {
		URI                string
		ExpectedContainer  string
		ExpectedStorageAcc string
		ExpectedDirectory  string
	}{
		{
			URI:                "abfss://wibble@mystorage.dfs.core.windows.net/wobble",
			ExpectedContainer:  "wibble",
			ExpectedStorageAcc: "mystorage",
			ExpectedDirectory:  "/wobble",
		},
		{
			URI:                "abfss://wibble@mystorage.dfs.core.windows.net",
			ExpectedContainer:  "wibble",
			ExpectedStorageAcc: "mystorage",
			ExpectedDirectory:  "",
		},
	}

	for _, tc := range testCases {
		container, storageAcc, dir, err := ProcessAzureWasbAbfssUris(tc.URI)
		assert.Equal(t, tc.ExpectedContainer, container)
		assert.Equal(t, tc.ExpectedStorageAcc, storageAcc)
		assert.Equal(t, tc.ExpectedDirectory, dir)
		assert.Nil(t, err)
	}
}
