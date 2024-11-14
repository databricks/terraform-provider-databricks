package aws

var AwsConfig = map[string]map[string]string{
	"aws": {
		"accountId":            "414351767826",
		"logDeliveryIamArn":    "arn:aws:iam::414351767826:role/SaasUsageDeliveryRole-prod-IAMRole-3PLHICCRR1TK",
		"unityCatalogueIamArn": "arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL",
	},
	"aws-us-gov": {
		"accountId":            "044793339203",
		"logDeliveryIamArn":    "arn:aws-us-gov:iam::044793339203:role/SaasUsageDeliveryRole-prod-aws-gov-IAMRole-L4QM0RCHYQ1G",
		"unityCatalogueIamArn": "arn:aws-us-gov:iam::044793339203:role/unity-catalog-prod-UCMasterRole-1QRFA8SGY15OJ",
	},
}

var AwsPartitions = []string{"aws", "aws-us-gov"}
var AwsPartitionsValidationError = "aws_partition must be either 'aws' or 'aws-us-gov'"
