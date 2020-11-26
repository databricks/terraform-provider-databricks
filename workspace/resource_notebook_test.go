package workspace

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
)

func notebookToB64(filePath string) (string, error) {
	notebookBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to find notebook to convert to base64; %w", err)
	}
	return base64.StdEncoding.EncodeToString(notebookBytes), nil
}

func TestValidateNotebookPath(t *testing.T) {
	testCases := []struct {
		name         string
		notebookPath string
		errorCount   int
	}{
		{"empty_path",
			"",
			2},
		{"correct_path",
			"/directory",
			0},
		{"path_starts_with_no_slash",
			"directory",
			1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, errs := ValidateNotebookPath(tc.notebookPath, "key")

			assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.notebookPath)
		})
	}
}

func TestResourceNotebookCreate_DirDoesNotExists(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("acceptance/testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	path := "/test/path.py"
	content := pythonNotebookDataB64
	objectId := 12345

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "not found",
				},
				Status: 404,
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/test",
				},
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: ImportRequest{
					Content:   content,
					Path:      path,
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: NotebookContent{
					Content: pythonNotebookDataB64,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   int64(objectId),
					ObjectType: Notebook,
					Path:       path,
					Language:   "PYTHON",
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"path":     path,
			"content":  content,
			"language": "PYTHON",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, "PYTHON", d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestResourceNotebookRead(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("acceptance/testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	testId := "/test/path.py"
	objectId := 12345
	assert.NoError(t, err, err)
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: NotebookContent{
					Content: pythonNotebookDataB64,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   int64(objectId),
					ObjectType: Notebook,
					Path:       testId,
					Language:   "PYTHON",
				},
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		New:      true,
		ID:       testId,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, testId, d.Get("path"))
	assert.Equal(t, "PYTHON", d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestResourceNotebookDelete(t *testing.T) {
	testId := "/test/path.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: NotebookDeleteRequest{Path: testId, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       testId,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
}

func TestResourceNotebookDelete_TooManyRequests(t *testing.T) {
	testId := "/test/path.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/delete",
				Status:   http.StatusTooManyRequests,
			},
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: NotebookDeleteRequest{Path: testId, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       testId,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
}

func TestResourceNotebookRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		ID:       "/test/path.py",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceNotebookRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		ID:       "/test/path.py",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "/test/path.py", d.Id(), "Id should not be empty for error reads")
}

func TestResourceNotebookCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: ImportRequest{
					Content:   "YWJjCg==",
					Path:      "/path.py",
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Fpath.py",
				Response: NotebookContent{
					Content: "YWJjCg==",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: "NOTEBOOK",
					Path:       "/path.py",
					Language:   "PYTHON",
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"content":  "YWJjCg==",
			"language": "PYTHON",
			"path":     "/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/path.py", d.Id())
}

func TestResourceNotebookCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"content":   "YWJjCg==",
			"language":  "PYTHON",
			"path":      "/path.py",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceNotebookDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestUri(t *testing.T) {
	uri := "https://sri-e2-test-workspace-3.cloud.databricks.com/api/2.0/workspace/export?format=DBC\u0026path=/demo-notebook-rbc"
	t.Log(url.PathUnescape(uri))
}

// nolint this should be refactored to support dbc files
func convertZipBytesToCRC(b64 []byte) (string, error) {
	r, err := zip.NewReader(bytes.NewReader(b64), int64(len(b64)))
	if err != nil {
		return "0", err
	}
	var totalSum int64
	for _, f := range r.File {
		if f.FileInfo().IsDir() == false {
			file, err := f.Open()
			if err != nil {
				return "", err
			}
			crc, err := getDBCCheckSumForCommands(file)
			if err != nil {
				return "", err
			}
			totalSum += int64(crc)
		}
	}
	return strconv.Itoa(int(totalSum)), nil
}

// nolint this should be refactored to support dbc files
func getDBCCheckSumForCommands(fileIO io.Reader) (int, error) {
	var stringBuff bytes.Buffer
	scanner := bufio.NewScanner(fileIO)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		stringBuff.WriteString(scanner.Text())
	}
	jsonString := stringBuff.Bytes()
	var notebook map[string]interface{}
	err := json.Unmarshal(jsonString, &notebook)
	if err != nil {
		return 0, err
	}
	var commandsBuffer bytes.Buffer
	commandsMap := map[int]string{}
	commands := notebook["commands"].([]interface{})
	for _, command := range commands {
		commandsMap[int(command.(map[string]interface{})["position"].(float64))] = command.(map[string]interface{})["command"].(string)
	}
	keys := make([]int, 0, len(commandsMap))
	for k := range commandsMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		commandsBuffer.WriteString(commandsMap[k])
	}
	return int(crc32.ChecksumIEEE(commandsBuffer.Bytes())), nil
}
