package storage

import (
	"bytes"
	"encoding/base64"
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
	t.Log(len(randomStr))
	t.Log(len(base64.StdEncoding.EncodeToString(randomStr)))

	client := common.NewClientFromEnvironment()

	err := NewDBFSAPI(client).Mkdirs(dir)
	assert.NoError(t, err, err)

	err = NewDBFSAPI(client).Mkdirs(dir2)
	assert.NoError(t, err, err)

	inputData := base64.StdEncoding.EncodeToString(randomStr)
	err = NewDBFSAPI(client).Create(path, true, inputData)
	assert.NoError(t, err, err)

	err = NewDBFSAPI(client).Create(path2, true, inputData)
	assert.NoError(t, err, err)

	err = NewDBFSAPI(client).Create(path3, true, inputData)
	assert.NoError(t, err, err)

	defer func() {
		err := NewDBFSAPI(client).Delete(dir, true)
		assert.NoError(t, err, err)
	}()

	base64Resp, err := NewDBFSAPI(client).Read(path)
	assert.NoError(t, err, err)
	assert.True(t, inputData == base64Resp)

	items, err := NewDBFSAPI(client).List(dir, false)
	assert.NoError(t, err, err)
	assert.Len(t, items, 2)

	items, err = NewDBFSAPI(client).List(dir, true)
	assert.NoError(t, err, err)
	assert.Len(t, items, 3)
}
