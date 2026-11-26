package poll_test

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	waiter := sql.WaitGetWarehouseRunning[int]{
		Poll: poll.Simple(sql.GetWarehouseResponse{
			Id: "test",
		}),
	}
	res, err := waiter.Get()
	assert.NoError(t, err)
	assert.Equal(t, "test", res.Id)
}

func TestSimpleError(t *testing.T) {
	waiter := sql.WaitGetWarehouseRunning[int]{
		Poll: poll.SimpleError[sql.GetWarehouseResponse](errors.New("test")),
	}
	_, err := waiter.Get()
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("test"), err)
	}
}
