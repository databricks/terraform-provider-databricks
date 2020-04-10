package service

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestInstanceProfiles(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	arn := "arn:aws:iam::997819012307:instance-profile/helloworldsritestingterraform"
	client := GetIntegrationDBAPIClient()

	defer func() {
		err := client.InstanceProfiles().Delete(arn)
		assert.NoError(t, err, err)
	}()
	err := client.InstanceProfiles().Create(arn, true)
	assert.NoError(t, err, err)

	arnSearch, err := client.InstanceProfiles().Read(arn)
	assert.NoError(t, err, err)
	assert.True(t, len(arnSearch) > 0)
	log.Println(arnSearch)
}
