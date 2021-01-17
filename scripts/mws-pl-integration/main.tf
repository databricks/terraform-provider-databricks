
provider "aws" {
  // region                  = "us-west-2"
  shared_credentials_file = "/Users/navid/.aws/credentials"
  profile                 = "aws-field-eng_databricks-power-user"
}

// get any env var to tf
data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

provider "databricks" {
  alias = "mws"
  host  = data.external.env.result.DATABRICKS_HOST
  username = data.external.env.result.DATABRICKS_USERNAME
  password = data.external.env.result.DATABRICKS_PASSWORD
}

resource "databricks_mws_vpc_endpoint" "vpce_relay"{
    account_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
    aws_vpc_endpoint_id =data.external.env.result.DATABRICKS_AWS_VPCE_RELAY_ID
    vpc_endpoint_name=data.external.env.result.DATABRICKS_AWS_VPCE_RELAY_NAME
    region= data.external.env.result.TEST_REGION
}

resource "databricks_mws_vpc_endpoint" "vpce_rest"{
    account_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
    aws_vpc_endpoint_id =data.external.env.result.DATABRICKS_AWS_VPCE_REST_ID
    vpc_endpoint_name=data.external.env.result.DATABRICKS_AWS_VPCE_REST_NAME
    region= data.external.env.result.TEST_REGION
}


resource "databricks_mws_private_access_settings" "pas"{
    account_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
    private_access_settings_name=data.external.env.result.DATABRICKS_PAS_NAME
    region= data.external.env.result.TEST_REGION

}

output "test_databricks_mws_vpc_relay_endpoint" {
    value = databricks_mws_vpc_endpoint.vpce_relay.vpc_endpoint_id
}

output "test_databricks_mws_vpc_rest_endpoint" {
    value = databricks_mws_vpc_endpoint.vpce_rest.vpc_endpoint_id
}

output "test_databricks_mws_private_access_settings"{
  value = databricks_mws_private_access_settings.pas.private_access_settings_id
}



