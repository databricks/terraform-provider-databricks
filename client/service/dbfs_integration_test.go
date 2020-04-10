package service

import (
	"bytes"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
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

	path := "/sri/randomfile"

	randomStr := GenString(500000)
	t.Log(len(randomStr))
	t.Log(len(base64.StdEncoding.EncodeToString(randomStr)))

	client := GetIntegrationDBAPIClient()
	err := client.DBFS().Create(path, true, base64.StdEncoding.EncodeToString(randomStr))
	assert.NoError(t, err, err)

	//err = client.DBFS().Delete(path, false)
	//assert.NoError(t, err, err)
}

//15500000
//15500000
func TestReadFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	path := "/sri/randomfile"
	client := GetIntegrationDBAPIClient()

	data, err := client.DBFS().Status(path)
	assert.NoError(t, err, err)
	t.Log(data)

}

func TestListRecursive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	path := "/andre_mesarovic/mlflow"
	client := GetIntegrationDBAPIClient()
	data, err := client.DBFS().List(path, true)
	assert.NoError(t, err, err)

	t.Log(data)
}
