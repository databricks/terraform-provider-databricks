+++
title = "Installation"
date = 2020-04-20T23:34:03-04:00
weight = 6
chapter = false
+++

## Installation Options
{{< tabs groupId="installationOptions" >}}
{{% tab name="Install Quickly using cURL" %}}
* To quickly install the binary please execute the following curl command in your shell.

```bash
$ curl https://raw.githubusercontent.com/databrickslabs/terraform-provider-databricks/master/godownloader-databricks-provider.sh | bash -s -- -b $HOME/.terraform.d/plugins
```

* The command should have moved the binary into your `~/.terraform.d/plugins` folder.

* You can `ls` the previous directory to verify.

{{% /tab %}}

{{% tab name="Install by building source" %}}
* After installing `golang`, `terraform`, and `make` you will now build the artifact.

```bash
$ go get -v -u github.com/databrickslabs/terraform-provider-databricks && cd $GOPATH/src/github.com/databrickslabs/terraform-provider-databricks 
```

* When you are in the root directory of the repository please run:

```bash
$ make build
```

* Locate your [terraform plugins directory](https://www.terraform.io/docs/extend/how-terraform-works.html#plugin-locations) 
    or the root folder of your terraform code

* Copy the `terraform-provider-databricks` artifact to that terraform plugins locations

```bash
$ mkdir -p ~/.terraform.d/plugins/ && cp terraform-provider-databricks ~/.terraform.d/plugins/terraform-provider-databricks
``` 

* Now your plugin for the Databricks Terraform provider is installed correctly. You can actually use the provider.

{{% /tab %}}

{{% tab name="Run via Docker Container" %}}

* To install and build the code if you dont want to install golang, terraform, etc. All you need is docker and git.

* First make sure you clone the repository and you are in the directory.

* Then build the docker image with this command (this command will trigger a multi-stage docker build):

```bash
$ docker build -t terraform-provider-databricks . 
```

* Then run the execute the terraform binary via the following command and volume mount. Make sure that you are in the directory
 with the terraform code. The following command you can execute the following commands and additional ones as part of 
 the terraform binary.
 
```bash
$ docker run -it -v $(pwd):/workpace -w /workpace terraform-provider-databricks init
$ docker run -it -v $(pwd):/workpace -w /workpace terraform-provider-databricks plan
$ docker run -it -v $(pwd):/workpace -w /workpace terraform-provider-databricks apply
```

{{% /tab %}}

{{< /tabs >}}
