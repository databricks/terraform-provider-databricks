package access

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceTableACLRead(t *testing.T) {
	t.Skip()
	d, err := qa.ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "table",
				Data: []interface{}{
					[]interface{}{"DENIED_SELECT", 1, false},
					[]interface{}{"DENIED_MODIFY", 2, true},
				},
			}
		},
		Fixtures: []qa.HTTPFixture{
			//..
		},
		Resource: ResourceTableACL(),
		Read:     true,
		ID:       "DATABASE/reporting/somebody",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "DATABASE/reporting/somebody", d.Id())
	assert.Equal(t, "reporting", d.Get("database"))
	assert.Equal(t, "..", d.Get("principal"))
}