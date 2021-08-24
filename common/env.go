package common

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/mitchellh/go-homedir"
)

var (
	envMutex     sync.Mutex
	onceClient   sync.Once
	commonClient *DatabricksClient
)

// NewClientFromEnvironment makes very good client for testing purposes
func NewClientFromEnvironment() *DatabricksClient {
	client := DatabricksClient{}
	for _, attr := range ClientAttributes() {
		found := false
		var value interface{}
		for _, envName := range attr.EnvVars {
			v := os.Getenv(envName)
			if v == "" {
				continue
			}
			switch attr.Kind {
			case reflect.String:
				value = v
				found = true
			case reflect.Bool:
				if vv, err := strconv.ParseBool(v); err == nil {
					value = vv
					found = true
				}
			case reflect.Int:
				if vv, err := strconv.Atoi(v); err == nil {
					value = vv
					found = true
				}
			default:
				continue
			}
		}
		if found {
			attr.Set(&client, value)
		}
	}
	err := client.Configure()
	if err != nil {
		panic(err)
	}
	return &client
}

// ResetCommonEnvironmentClient resets test dummy
func ResetCommonEnvironmentClient() {
	commonClient = nil
	onceClient = sync.Once{}
}

// CommonEnvironmentClient configured once per run of application
func CommonEnvironmentClient() *DatabricksClient {
	if commonClient != nil {
		return commonClient
	}
	onceClient.Do(func() {
		commonClient = NewClientFromEnvironment()
	})
	return commonClient
}

// CleanupEnvironment backs up environment - use as `defer CleanupEnvironment()()`
// clears it and restores it in the end. It's meant strictly for "unit" tests
// as last resort, because it slows down parallel execution with mutex.
func CleanupEnvironment() func() {
	// make a backed-up pristine environment
	envMutex.Lock()
	prevEnv := os.Environ()
	oldPath := os.Getenv("PATH")
	pwd := os.Getenv("PWD")
	os.Clearenv()
	err := os.Setenv("PATH", oldPath)
	if err != nil {
		log.Printf("[WARN] Cannot bring back PATH: %v", err)
	}
	err = os.Setenv("HOME", pwd)
	if err != nil {
		log.Printf("[WARN] Cannot set HOME to old PWD: %v", err)
	}
	homedir.DisableCache = true
	// and return restore function
	return func() {
		for _, kv := range prevEnv {
			kvs := strings.SplitN(kv, "=", 2)
			os.Setenv(kvs[0], kvs[1])
		}
		envMutex.Unlock()
	}
}
