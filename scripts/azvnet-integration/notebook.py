# Databricks notebook source
# MAGIC %fs ls /mnt/sample

# COMMAND ----------
# MAGIC %sh ls /dbfs/mnt/sample

# COMMAND ----------
dbutils.widgets.text("department", "")

print(f'Department is {dbutils.widgets.get("department")}')
print(f'But secret is redacted: {dbutils.secrets.get("terraform", "blob_storage_key")}') # lgtm [py/clear-text-logging-sensitive-data]
