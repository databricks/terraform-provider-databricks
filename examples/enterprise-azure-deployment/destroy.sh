#!/bin/bash
terraform destroy -var-file='secret.tfvars' -var-file='dbws.tfvars'
