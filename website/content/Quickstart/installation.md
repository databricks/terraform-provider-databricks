+++
title = "Installation"
date = 2020-04-20T23:34:03-04:00
weight = 6
chapter = false
+++

A Terraform provider for Databricks workspace components

## Installing databricks-terraform with Go

* Install Go 1.13. For previous versions, you may have to set your `$GOPATH` manually, if you haven't done it yet visit [here](https://golang.org/doc/install).
* Install Terraform 0.12.x [from here](https://www.terraform.io/downloads.html) and save it into `/usr/local/bin/terraform` folder (create it if it doesn't exists). This provider DOES NOT SUPPORT Terraform 0.12 or above.
* Download the code by issuing a `go get` command.

```bash
# Download the source code for databricks-terraform
# and build the needed binary, by saving it inside $GOPATH/bin
$ go get -u github.com/databrickslabs/databricks-terraform


# After fetching the code base we will switch into the directory for the code base.
$ cd $GOPATH/src/github.com/databrickslabs/databricks-terraform 

# Once in the directory you will run the build using the make command provided by the make file
$ make build

# Once the file is made we will then move the file to where terraform can pick it up
$ mkdir -p ~/.terraform.d/plugins/ && cp terraform-provider-db ~/.terraform.d/plugins/terraform-provider-db
```

If you wish to uninstall the binary simply remove the file from the directory.

```bash
$ rm /usr/local/bin/terraform-provider-db
```

## Using databricks-terraform with Docker (TBD!)



