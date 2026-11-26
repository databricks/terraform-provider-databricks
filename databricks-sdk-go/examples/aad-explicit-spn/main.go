package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:              askFor("Host:"),
		AzureResourceID:   askFor("Azure Resource ID:"),
		AzureTenantID:     askFor("AAD Tenant ID:"),
		AzureClientID:     askFor("AAD Client ID:"),
		AzureClientSecret: askFor("AAD Client Secret:"),
		Credentials:       config.AzureClientSecretCredentials{},
	}))
	all, err := w.Clusters.ListAll(context.Background(), compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

func askFor(prompt string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, prompt+" ")
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s != "" {
			break
		}
	}
	return s
}
