package exporter

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
)

type levelWriter []string

var logLevel = levelWriter{"[INFO]", "[ERROR]", "[WARN]"}

func (lw *levelWriter) Write(p []byte) (n int, err error) {
	a := string(p)
	for _, l := range *lw {
		if strings.Contains(a, l) {
			return os.Stdout.Write(p)
		}
	}
	return
}

func (ic *importContext) allServicesAndListing() (string, string) {
	services := ""
	listing := ""
	for _, ir := range ic.Importables {
		if !strings.Contains(services, ir.Service) {
			if len(services) > 0 {
				services += ","
			}
			services += ir.Service
		}
		if ir.List != nil && !strings.Contains(listing, ir.Service) {
			if len(listing) > 0 {
				listing += ","
			}
			listing += ir.Service
		}
	}
	return services, listing
}

func (ic *importContext) interactivePrompts() {
	req, _ := http.NewRequest("GET", "/", nil)
	for ic.Client.DatabricksClient.Config.Authenticate(req) != nil {
		ic.Client.DatabricksClient.Config.Host = askFor("🔑 Databricks Workspace URL:")
		ic.Client.DatabricksClient.Config.Token = askFor("🔑 Databricks Workspace PAT:")
	}
	ic.match = askFor("🔍 Match entity names (optional):")
	listing := ""
	for r, ir := range ic.Importables {
		if ir.List == nil {
			continue
		}
		if !askFlag(fmt.Sprintf("✅ Generate `%s` and related resources?", r)) {
			continue
		}
		if len(listing) > 0 {
			listing += ","
		}
		listing += ir.Service
		if ir.Service == "mounts" {
			ic.mounts = true
		}
	}
	ic.listing = listing
}

// Run import according to flags
func Run(args ...string) error {
	log.SetOutput(&logLevel)
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
	var skipInteractive bool
	flags.BoolVar(&skipInteractive, "skip-interactive", false, "Skip interactive mode")
	flags.BoolVar(&ic.includeUserDomains, "includeUserDomains", false, "Include domain portion in `databricks_user` resource name")
	flags.BoolVar(&ic.importAllUsers, "importAllUsers", false,
		"Import all users and service principals, even if they aren't referenced in any resource")
	flags.StringVar(&ic.Directory, "directory", cwd,
		"Directory to generate sources in. Defaults to current directory.")
	flags.Int64Var(&ic.lastActiveDays, "last-active-days", 3650,
		"Items with older than activity specified won't be imported.")
	flags.BoolVar(&ic.debug, "debug", false, "Print extra debug information.")
	flags.BoolVar(&ic.mounts, "mounts", false, "List DBFS mount points.")
	flags.BoolVar(&ic.generateDeclaration, "generateProviderDeclaration", true,
		"Generate Databricks provider declaration.")
	services, listing := ic.allServicesAndListing()
	flags.StringVar(&ic.services, "services", services,
		"Comma-separated list of services to import. By default all services are imported.")
	flags.StringVar(&ic.listing, "listing", listing,
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
	if ic.debug {
		logLevel = append(logLevel, "[DEBUG]")
	}
	return ic.Run()
}
