package service

import (
	"bytes"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	dat, _ := ioutil.ReadFile("/Users/Sri.Tikkireddy/Downloads/DatabricksConcepts&BestPractices.pdf")
	t.Log(len(dat))
}

func TestSplit(t *testing.T) {
	dat, _ := ioutil.ReadFile("/Users/Sri.Tikkireddy/Downloads/DatabricksConcepts&BestPractices.pdf")
	t.Log(len(dat))
	t.Log(len(split(dat, 1e6)))
}

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

	err = client.DBFS().Delete(path, false)
	assert.NoError(t, err, err)
}

func TestReadFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	path := "/sri/randomfile"
	client := GetIntegrationDBAPIClient()
	data, err := client.DBFS().Read(path)
	assert.NoError(t, err, err)

	byteArr, err := base64.StdEncoding.DecodeString(data)
	assert.NoError(t, err, err)

	t.Log(len(strings.Split(string(byteArr), "\n")))
	t.Log(strings.Split(string(byteArr), "\n")[0])
	t.Log(strings.Split(string(byteArr), "\n")[1])
	t.Log(strings.Split(string(byteArr), "\n")[2])
	t.Log(strings.Split(string(byteArr), "\n")[499999])
	t.Log(strings.Split(string(byteArr), "\n")[500000])
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
