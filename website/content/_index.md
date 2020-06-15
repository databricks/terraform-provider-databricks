+++
title = "Databricks Terraform Provider"
date = 2020-04-20T23:34:03-04:00
weight = 1
chapter = false
pre = ""
+++

# Lets lay some bricks! 

## Quick install

To quickly install the binary please execute the following curl command in your shell.

```bash
$ curl https://raw.githubusercontent.com/databrickslabs/terraform-provider-databricks/master/godownloader-databricks-provider.sh | bash -s -- -b $HOME/.terraform.d/plugins
```

The command should have moved the binary into your `~/.terraform.d/plugins` folder.

You can `ls` the previous directory to verify.

## Feedback

Please provide feedback in github issues. There is a template for this:

{{% button href="https://github.com/databrickslabs/terraform-provider-databricks/issues/new?assignees=stikkireddy&labels=question&template=feedback.md&title=%5BFEEDBACK%2FQUESTION%5D+Short+Description+of+feedback" icon="fas fa-comment-dots" %}}Please provide feedback!{{% /button %}}
