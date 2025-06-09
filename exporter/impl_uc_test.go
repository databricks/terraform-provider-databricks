package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmitUserSpOrGroup(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("users,groups")
	emitUserSpOrGroup(ic, "user@example.com")
	assert.Equal(t, 1, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_user[<unknown>] (user_name: user@example.com)")

	emitUserSpOrGroup(ic, "users")
	assert.Equal(t, 2, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_group[<unknown>] (display_name: users)")

	emitUserSpOrGroup(ic, "abcd1234-ab12-cd34-ef56-abcdef123456")
	assert.Equal(t, 3, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_service_principal[<unknown>] (application_id: abcd1234-ab12-cd34-ef56-abcdef123456)")

	emitUserSpOrGroup(ic, "users @ test.com")
	assert.Equal(t, 4, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_group[<unknown>] (display_name: users @ test.com)")

}
