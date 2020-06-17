package service

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GenString(times int) []byte {
	var buf bytes.Buffer
	for i := 0; i < times; i++ {
		buf.WriteString("Hello world how are you doing?\n")
	}
	return buf.Bytes()
}

func TestCreateFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	dir := "/client-test/"
	dir2 := "/client-test/dir2/"
	path := "/client-test/randomfile"
	path2 := "/client-test/dir2/randomfile"
	path3 := "/client-test/dir2/randomfile2"
	copyPath := "/client-test/dir2/randomfilecopy2"

	randomStr := GenString(500)
	t.Log(len(randomStr))
	t.Log(len(base64.StdEncoding.EncodeToString(randomStr)))

	client := GetIntegrationDBAPIClient()

	err := client.DBFS().Mkdirs(dir)
	assert.NoError(t, err, err)

	err = client.DBFS().Mkdirs(dir2)
	assert.NoError(t, err, err)

	inputData := base64.StdEncoding.EncodeToString(randomStr)
	err = client.DBFS().Create(path, true, inputData)
	assert.NoError(t, err, err)

	err = client.DBFS().Create(path2, true, inputData)
	assert.NoError(t, err, err)

	err = client.DBFS().Create(path3, true, inputData)
	assert.NoError(t, err, err)

	defer func() {
		err := client.DBFS().Delete(dir, true)
		assert.NoError(t, err, err)
	}()

	base64Resp, err := client.DBFS().Read(path)
	assert.NoError(t, err, err)
	assert.True(t, inputData == base64Resp)

	items, err := client.DBFS().List(dir, false)
	assert.NoError(t, err, err)
	assert.True(t, len(items) == 2)

	items, err = client.DBFS().List(dir, true)
	assert.NoError(t, err, err)
	assert.True(t, len(items) == 3)

	err = client.DBFS().Copy(path, copyPath, client, true)
	assert.NoError(t, err, err)

	base64Resp, err = client.DBFS().Read(copyPath)
	assert.NoError(t, err, err)
	assert.True(t, inputData == base64Resp)

	items, err = client.DBFS().List(dir, true)
	assert.NoError(t, err, err)
	assert.True(t, len(items) == 4)
}

////15500000
////15500000
//func TestReadFile(t *testing.T) {
//	if testing.Short() {
//		t.Skip("skipping integration test in short mode.")
//	}
//
//	path := "/sri/randomfile"
//	client := GetIntegrationDBAPIClient()
//
//	data, err := client.DBFS().Status(path)
//	assert.NoError(t, err, err)
//	t.Log(data)
//
//}
//
//func TestListRecursive(t *testing.T) {
//	if testing.Short() {
//		t.Skip("skipping integration test in short mode.")
//	}
//
//	path := "/andre_mesarovic/mlflow"
//	client := GetIntegrationDBAPIClient()
//	data, err := client.DBFS().List(path, true)
//	assert.NoError(t, err, err)
//
//	t.Log(data)
//}
