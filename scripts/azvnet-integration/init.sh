#!/bin/bash

# Docs: https://docs.databricks.com/clusters/init-scripts.html

mkdir /dbfs/mnt/sample/$DB_CLUSTER_ID
echo "DB_CLUSTER_ID=${DB_CLUSTER_ID}" >> /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt
echo "DB_CONTAINER_IP=${DB_CONTAINER_IP}" >> /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt
echo "DB_IS_DRIVER=${DB_IS_DRIVER}" >> /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt
echo "DB_DRIVER_IP=${DB_DRIVER_IP}" >> /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt
echo "DB_INSTANCE_TYPE=${DB_INSTANCE_TYPE}" >> /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt
echo "DB_CLUSTER_NAME=${DB_CLUSTER_NAME}" >> /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt

echo "Init script finished. Look at /dbfs/mnt/sample/$DB_CLUSTER_ID/vars-$DB_CONTAINER_IP.txt"
