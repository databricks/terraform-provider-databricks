# Integration Testing on AWS with MWS

The `run.sh` script will setup the required resources in AWS defined in `prereqs.tf` and then pass these as environment variables to the `golang` integration tests. 

Once the tests have finished the `run.sh` will attempt to cleanup resources unless an environment variable of `SKIP_CLEANUP` is set.

> Note: Recreating the resouces each run is a good practive when running the tests as it ensures that past runs haven't made changes which effect future tests. For a quicker loop when debugging see the `Debugging` section.

*Requirements*
- `.env` file at root of project is set with a SP which has ability to assign roles (easiest to set OWNER on sub)
- `terraform` installed
- AWS_ACCESS_KEY_ID environment variable set
- AWS_SECRET_ACCESS_KEY environment variable set
- DATABRICKS_MWS_AWS_REGION environment variable set
- DATABRICKS_MWS_ACCT_ID environment variable set
- DATABRICKS_MWS_USERNAME environment variable set
- DATABRICKS_MWS_PASSWORD environment variable set
- DATABRICKS_MWS_HOST environment variable set

## Debugging

If you want to run in a tighter loop without waiting on resource creation each time you invoke the tests you can use the `SKIP_CLEANUP` env like so:

```
export SKIP_CLEANUP=true
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
export DATABRICKS_MWS_AWS_REGION=...
export DATABRICKS_MWS_ACCT_ID=...
export DATABRICKS_MWS_USERNAME=...
export DATABRICKS_MWS_PASSWORD=...
export DATABRICKS_MWS_HOST=...
integration-environment-azure/run.sh
```

In this case the same workspace, storage and other pre-reqs will be used each run.
                                                                                                                                                                                                 -var 'databricks_mws_acct_id=$DATABRICKS_MWS_ACCT_ID'