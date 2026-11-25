package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

type CustomCredentials struct{}

func (c *CustomCredentials) Name() string {
	return "custom"
}

func (c *CustomCredentials) Configure(
	ctx context.Context, cfg *config.Config,
) (credentials.CredentialsProvider, error) {
	visitor := func(r *http.Request) error {
		token := askFor("Token:")
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
	return credentials.CredentialsProviderFn(visitor), nil
}

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:        askFor("Host:"),
		Credentials: &CustomCredentials{},
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
