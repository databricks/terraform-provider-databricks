package storage

import (
	"bytes"
	"context"
	"crypto/md5"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func GenString(times int) []byte {
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
	dir2 := dir + "/dir2/"
	path := dir + "/randomfile"
	path2 := dir + "/dir2/randomfile"
	path3 := dir + "/dir2/randomfile2"

	randomStr := GenString(500)
	client := common.NewClientFromEnvironment()

	dbfsAPI := NewDbfsAPI(context.Background(), client)
	err := dbfsAPI.Mkdirs(dir)
	assert.NoError(t, err, err)

	err = dbfsAPI.Mkdirs(dir2)
	assert.NoError(t, err, err)

	err = dbfsAPI.Create(path, randomStr, true)
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
