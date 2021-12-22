package storage

import (
	"bytes"
	"context"
	"crypto/md5"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestCreateFileFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: createHandle{
				Path:      "/create-fails",
				Overwrite: true,
			},
			Status:   404,
			Response: common.NotFound("fails"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/close",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		err := a.Create("/create-fails", []byte("abc"), true)
		assert.EqualError(t, err, "cannot create handle: fails")
	})
}

func TestCreateFile_AddBlockFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: createHandle{
				Path:      "/add-fails",
				Overwrite: true,
			},
			Response: handleResponse{123},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/add-block",
			ExpectedRequest: addBlock{
				Data:   "YWJj",
				Handle: 123,
			},
			Status:   404,
			Response: common.NotFound("fails"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/close",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		err := a.Create("/add-fails", []byte("abc"), true)
		assert.EqualError(t, err, "cannot add block: fails")
	})
}

func TestCreateFile_CloseFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/create",
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/add-block",
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/dbfs/close",
			Status:   404,
			Response: common.NotFound("fails"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		err := a.Create("/close-fails", []byte("abc"), true)
		assert.EqualError(t, err, "cannot close handle: fails")
	})
}

func TestDbfsListRecursiveFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/list?path=abc",
			Status:   404,
			Response: common.NotFound("fails"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/list?path=sub",
			Response: FileList{
				Files: []FileInfo{
					{
						Path:  "bcd",
						IsDir: true,
					},
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/dbfs/list?path=bcd",
			Status:       404,
			Response:     common.NotFound("fails"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		_, err := a.List("abc", true)
		assert.EqualError(t, err, "cannot list abc: fails")
		_, err = a.List("sub", true)
		assert.EqualError(t, err, "cannot list subfolder: cannot list bcd: fails")
		_, err = a.List("bcd", false)
		assert.EqualError(t, err, "cannot list bcd: fails")
	})
}

func TestDbfsReadFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny: true,
			Status:   404,
			Response: common.NotFound("fails"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		a := NewDbfsAPI(ctx, client)
		_, err := a.Read("abc")
		assert.EqualError(t, err, "cannot read abc: fails")
	})
}

func genString(times int) []byte {
	var buf bytes.Buffer
	for i := 0; i < times; i++ {
		buf.WriteString("Hello world how are you doing?\n")
	}
	return buf.Bytes()
}

func TestAccCreateFile(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	dir := "/client-test/" + randomName
	path := dir + "/randomfile"
	path2 := dir + "/dir2/randomfile"
	path3 := dir + "/dir2/randomfile2"

	randomStr := genString(10)
	client := common.NewClientFromEnvironment()

	dbfsAPI := NewDbfsAPI(context.Background(), client)

	err := dbfsAPI.Create(path, randomStr, true)
	assert.NoError(t, err, err)

	err = dbfsAPI.Create(path2, randomStr, true)
	assert.NoError(t, err, err)

	err = dbfsAPI.Create(path3, randomStr, true)
	assert.NoError(t, err, err)

	defer func() {
		err := dbfsAPI.Delete(dir, true)
		assert.NoError(t, err, err)
	}()

	resp, err := dbfsAPI.Read(path)
	assert.NoError(t, err, err)
	assert.True(t, md5.Sum(randomStr) == md5.Sum(resp))

	items, err := dbfsAPI.List(dir, false)
	assert.NoError(t, err, err)
	assert.Len(t, items, 2)

	items, err = dbfsAPI.List(dir, true)
	assert.NoError(t, err, err)
	assert.Len(t, items, 3)
}
