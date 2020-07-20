#!/bin/bash

CURRENT_DIR=`pwd`
make install
cd examples/enterprise-azure-deployment/enterprise-workspace-deploy
terraform init
./apply.sh
cd $CURRENT_DIR