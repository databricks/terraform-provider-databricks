package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	lifeTimeSeconds := int32(30)
	comment := "Hello world"

	token, err := client.Tokens().Create(lifeTimeSeconds, comment)
	assert.NoError(t, err, err)
	assert.True(t, len(token.TokenValue) > 0, "Token value is empty")

	defer func() {
		err := client.Tokens().Delete(token.TokenInfo.TokenID)
		assert.NoError(t, err, err)
	}()

	_, err = client.Tokens().Read(token.TokenInfo.TokenID)
	assert.NoError(t, err, err)

	tokenList, err := client.Tokens().List()
	assert.NoError(t, err, err)
	assert.True(t, len(tokenList) > 0, "Token list is empty")
}
