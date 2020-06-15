package databricks

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceNotebook() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebookCreate,
		Read:   resourceNotebookRead,
		Delete: resourceNotebookDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(i interface{}) string {
					base64String := i.(string)
					base64, err := convertBase64ToCheckSum(base64String)
					if err != nil {
						return ""
					}
					return base64
				},
				ValidateFunc: validation.StringIsBase64,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(model.Scala),
					string(model.Python),
					string(model.R),
					string(model.SQL),
				}, false),
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"mkdirs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(model.Source),
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(model.DBC),
					string(model.Jupyter),
					string(model.Source),
					string(model.HTML),
				}, false),
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceNotebookCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	path := d.Get("path").(string)
	content := d.Get("content").(string)
	language := d.Get("language").(string)
	format := d.Get("format").(string)
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs {
		parentDir := filepath.Dir(path)
		err := client.Notebooks().Mkdirs(parentDir)
		if err != nil {
			return err
		}
	}

	err := client.Notebooks().Create(path, content, model.Language(language), model.ExportFormat(format), overwrite)
	if err != nil {
		return err
	}
	d.SetId(path)
	//err = d.Set("content", content)
	//if err != nil {
	//	return err
	//}
	err = d.Set("format", format)
	if err != nil {
		return err
	}
	err = d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	return resourceNotebookRead(d, m)
}

func resourceNotebookRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	format := d.Get("format").(string)
	notebookData, err := client.Notebooks().Export(id, model.ExportFormat(format))
	if err != nil {
		if isNotebookMissing(err.Error(), id) {
			log.Printf("Missing notebook with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}
	notebookInfo, err := client.Notebooks().Read(id)
	if err != nil {
		return err
	}
	err = d.Set("path", id)
	if err != nil {
		return err
	}

	crc, err := convertBase64ToCheckSum(notebookData)
	if err != nil {
		return err
	}
	err = d.Set("content", crc)
	if err != nil {
		return err
	}
	err = d.Set("language", string(notebookInfo.Language))
	if err != nil {
		return err
	}
	err = d.Set("object_id", int(notebookInfo.ObjectID))
	if err != nil {
		return err
	}
	err = d.Set("object_type", string(notebookInfo.ObjectType))

	return err
}

func resourceNotebookDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	err := client.Notebooks().Delete(id, true)
	return err
}

func convertBase64ToCheckSum(b64 string) (string, error) {
	dataArr, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Printf("Error while trying to decode base64 content: %v\n", err)
		return "error", err
	}
	checksum, err := convertZipBytesToCRC(dataArr)
	if err != nil {
		return strconv.Itoa(int(crc32.ChecksumIEEE(dataArr))), nil
	}
	return checksum, nil
}

func convertZipBytesToCRC(b64 []byte) (string, error) {
	r, err := zip.NewReader(bytes.NewReader(b64), int64(len(b64)))
	if err != nil {
		return "0", err
	}
	var totalSum int64
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
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

func isNotebookMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST") &&
		strings.Contains(errorMsg, fmt.Sprintf("Path (%s) doesn't exist.", resourceID))
}
