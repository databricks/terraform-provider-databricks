#!/bin/bash
terraform plan -var-file='secret.tfvars' -var-file='dbws.tfvars'
