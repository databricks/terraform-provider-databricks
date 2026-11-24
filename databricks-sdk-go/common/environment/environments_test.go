package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultEnvironmentIsReturned(t *testing.T) {
	assert.Equal(t, ".cloud.databricks.com", GetEnvironmentForHostname("").DnsZone)
}
