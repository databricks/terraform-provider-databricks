#!/bin/bash
terraform apply -var-file='secret.tfvars' -var-file='dbws.tfvars'
