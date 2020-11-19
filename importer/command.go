package importer

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
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

// Run import according to flags
func Run(args ...string) error {
	log.SetOutput(&logLevel)
	c := common.NewClientFromEnvironment()
	ic := newImportContext(c)

	flags := flag.NewFlagSet("importer", flag.ExitOnError)
	flags.StringVar(&ic.Module, "module", "",
		"Terraform module name, that changes are imported. "+
			"Defaults to empty string. Makes effect on generated "+
			"import.sh file")

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	flags.StringVar(&ic.Directory, "directory", cwd,
		"Directory to generate sources in. Defaults to current directory.")
	flags.Int64Var(&ic.lastActiveDays, "last-active-days", 3650,
		"Items with older than activity specified won't be imported.")
	flags.BoolVar(&ic.debug, "debug", false, "Print extra debug information")

	listing := ""
	services := ""
	for _, ir := range ic.Importables {
		if strings.Contains(services, ir.Service) {
			continue
		}
		if len(services) > 0 {
			services += ","
		}
		services += ir.Service
		if ir.List != nil {
			if len(listing) > 0 {
				listing += ","
			}
			listing += ir.Service
		}
	}
	flags.StringVar(&ic.services, "services", services,
		"Coma-separated list of services to import. By default all services are imported.")
	flags.StringVar(&ic.listing, "listing", listing,
		"Coma-separated list of services to be listed and further passed on for importing. "+
			"`-services` parameter controls which transitive dependencies will be processed. "+
			"We recommend limiting services with `-listing` more often, than `-services`.")
	flags.StringVar(&ic.match, "match", "", "Match resource names during listing operation. "+
		"This filter applies to all resources that are getting listed, so if you want to import "+
		"all dependencies of just one cluster, specify -listing=compute")
	err = flags.Parse(args)
	if err != nil {
		return err
	}
	if ic.debug {
		logLevel = append(logLevel, "[DEBUG]")
	}
	return ic.Run()
}
