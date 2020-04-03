package db

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"strings"
)

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Delete: resourceSecretDelete,

		Schema: map[string]*schema.Schema{
			"string_value": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

type SecretId map[string]string

// go binary encoder
func toGOB64(m SecretId) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// go binary decoder
func fromGOB64(str string) (SecretId, error) {
	m := SecretId{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return m, err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func getSecretId(scope string, key string) (string, error) {
	return scope + "|||" + key, nil
}

func getScopeAndKeyFromSecretId(secretIdString string) (string, string, error) {
	return strings.Split(secretIdString, "|||")[0], strings.Split(secretIdString, "|||")[1], nil
}

func resourceSecretCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	scopeName := d.Get("scope").(string)
	key := d.Get("key").(string)
	secretValue := d.Get("string_value").(string)
	err := client.Secrets().Create(secretValue, scopeName, key)
	if err != nil {
		return err
	}
	id, err := getSecretId(scopeName, key)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceSecretRead(d, m)
}

func resourceSecretRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretId(id)
	if err != nil {
		return err
	}
	err = d.Set("scope", scope)
	if err != nil {
		return err
	}
	err = d.Set("key", key)
	if err != nil {
		return err
	}
	d.SetId(id)
	return nil
}

func resourceSecretDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretId(id)
	if err != nil {
		return err
	}
	err = client.Secrets().Delete(scope, key)
	return err
}
