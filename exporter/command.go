package exporter

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"golang.org/x/exp/maps"
)

type levelWriter []string

var logLevel = levelWriter{"[INFO]", "[ERROR]", "[WARN]"}

func (lw *levelWriter) Write(p []byte) (n int, err error) {
	a := string(p)
	for _, l := range *lw {
		if strings.HasPrefix(a, l) {
			timeStr := time.Now().Local().Format(time.RFC3339Nano)
			logStr := a[0:len(l)] + " " + timeStr + ": " + strings.TrimLeft(a[len(l):len(a)-1], " ") + "\n"
			return os.Stdout.WriteString(logStr)
		}
	}
	return
}

func (ic *importContext) allServicesAndListing() (string, string) {
	services := map[string]struct{}{}
	listing := map[string]struct{}{}
	for _, ir := range ic.Importables {
		services[ir.Service] = struct{}{}
		if ir.List != nil {
			listing[ir.Service] = struct{}{}
		}
	}
	return strings.Join(maps.Keys(services), ","), strings.Join(maps.Keys(listing), ",")
}

func (ic *importContext) interactivePrompts() {
	req, _ := http.NewRequest("GET", "/", nil)
	for ic.Client.DatabricksClient.Config.Authenticate(req) != nil {
		ic.Client.DatabricksClient.Config.Host = askFor("🔑 Databricks Workspace URL:")
		ic.Client.DatabricksClient.Config.Token = askFor("🔑 Databricks Workspace PAT:")
	}
	ic.match = askFor("🔍 Match entity names (optional):")

	services := map[string][]string{}
	for r, ir := range ic.Importables {
		if ir.List == nil {
			continue
		}
		service := ir.Service
		v, exists := services[service]
		if exists {
			services[service] = append(v, r)
		} else {
			services[service] = []string{r}
		}
	}

	ic.listing = map[string]struct{}{}
	keys := maps.Keys(services)
	slices.Sort(keys)
	for _, service := range keys {
		resources := services[service]
		if !askFlag(fmt.Sprintf("✅ Generate for service `%s` (%s) and related resources?",
			service, strings.Join(resources, ","))) {
			continue
		}
		ic.listing[service] = struct{}{}
		if service == "mounts" {
			ic.mounts = true
		}
	}
}

// Run import according to flags
func Run(args ...string) error {
	log.SetOutput(&logLevel)
	log.Printf("[WARN] This tooling is experimental and provided as is. It has an evolving interface, which may change or be removed in future versions of the provider.")
	client, err := client.New(&config.Config{})
	if err != nil {
		return err
	}
	ic := newImportContext(&common.DatabricksClient{
		DatabricksClient: client,
	})

	flags := flag.NewFlagSet("exporter", flag.ExitOnError)
	flags.StringVar(&ic.Module, "module", "",
		"Terraform module name, that changes are imported. "+
			"Defaults to empty string. Makes effect on generated "+
			"import.sh file")

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	var skipInteractive, trace, debug bool
	flags.BoolVar(&skipInteractive, "skip-interactive", false, "Skip interactive mode")
	flags.BoolVar(&ic.includeUserDomains, "includeUserDomains", false, "Include domain portion in `databricks_user` resource name")
	flags.BoolVar(&ic.importAllUsers, "importAllUsers", false,
		"Import all users and service principals, even if they aren't referenced in any resource")
	flags.BoolVar(&ic.exportDeletedUsersAssets, "exportDeletedUsersAssets", false,
		"Export assets (notebooks, etc.) of deleted users & service principals")
	flags.StringVar(&ic.Directory, "directory", cwd,
		"Directory to generate sources in. Defaults to current directory.")
	flags.Int64Var(&ic.lastActiveDays, "last-active-days", 3650,
		"Items with older than activity specified won't be imported.")
	flags.BoolVar(&ic.incremental, "incremental", false, "Incremental export of the data. Requires -updated-since parameter")
	flags.BoolVar(&ic.exportSecrets, "export-secrets", false, "Generate terraform.tfvars with secrets")
	flags.BoolVar(&ic.noFormat, "noformat", false, "Don't run `terraform fmt` on exported files")
	flags.BoolVar(&ic.nativeImportSupported, "native-import", false, "Generate native import blocks (requires Terraform 1.5+)")
	flags.StringVar(&ic.updatedSinceStr, "updated-since", "",
		"Include only resources updated since a given timestamp (in ISO8601 format, i.e. 2023-07-01T00:00:00Z)")
	flags.BoolVar(&debug, "debug", false, "Print extra debug information.")
	flags.BoolVar(&trace, "trace", false, "Print full debug information.")
	flags.BoolVar(&ic.mounts, "mounts", false, "List DBFS mount points.")
	flags.BoolVar(&ic.generateDeclaration, "generateProviderDeclaration", true,
		"Generate Databricks provider declaration.")
	flags.StringVar(&ic.notebooksFormat, "notebooksFormat", "SOURCE",
		"Format to export notebooks: SOURCE, DBC, JUPYTER. Default: SOURCE")
	services, listing := ic.allServicesAndListing()
	var configuredServices string
	flags.StringVar(&configuredServices, "services", services,
		"Comma-separated list of services to import. By default all services are imported.")
	var configuredListing string
	flags.StringVar(&configuredListing, "listing", listing,
		"Comma-separated list of services to be listed and further passed on for importing. "+
			"`-services` parameter controls which transitive dependencies will be processed. "+
			"We recommend limiting services with `-listing` more often, than `-services`.")
	flags.StringVar(&ic.match, "match", "", "Match resource names during listing operation. "+
		"This filter applies to all resources that are getting listed, so if you want to import "+
		"all dependencies of just one cluster, specify -listing=compute")
	prefix := ""
	flags.StringVar(&prefix, "prefix", "", "Prefix that will be added to the name of all exported resources")
	newArgs := args
	if len(args) > 1 && args[1] == "exporter" {
		newArgs = args[2:]
	}
	err = flags.Parse(newArgs)
	if err != nil {
		return err
	}
	if !skipInteractive {
		ic.interactivePrompts()
	}
	if len(prefix) > 0 {
		ic.prefix = prefix + "_"
	}
	if trace {
		logLevel = append(logLevel, "[DEBUG]", "[TRACE]")
	} else if debug {
		logLevel = append(logLevel, "[DEBUG]")
	}
	ic.enableServices(configuredServices)
	ic.enableListing(configuredListing)
	return ic.Run()
}
