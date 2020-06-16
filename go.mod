module github.com/databrickslabs/databricks-terraform

go 1.13

require (
	github.com/Azure/go-autorest/autorest v0.10.2
	github.com/Azure/go-autorest/autorest/adal v0.8.3
	github.com/aws/aws-sdk-go v1.32.2
	github.com/fatih/color v1.9.0 // indirect
	github.com/google/go-querystring v1.0.0
	github.com/hashicorp/go-retryablehttp v0.6.6
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.13.1
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/r3labs/diff v0.0.0-20191120142937-b4ed99a31f5a
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	gopkg.in/ini.v1 v1.57.0
)

replace github.com/Azure/go-autorest => github.com/tombuildsstuff/go-autorest v14.0.1-0.20200317095413-f2d2d0252c3c+incompatible

replace github.com/Azure/go-autorest/autorest => github.com/tombuildsstuff/go-autorest/autorest v0.10.1-0.20200317095413-f2d2d0252c3c

replace github.com/Azure/go-autorest/autorest/azure/auth => github.com/tombuildsstuff/go-autorest/autorest/azure/auth v0.4.3-0.20200317095413-f2d2d0252c3c
