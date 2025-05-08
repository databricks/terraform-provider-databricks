package aws

var AwsConfig = map[string]map[string]string{
	"aws": {
		"accountId":            "414351767826",
		"awsNamespace":         "aws",
		"logDeliveryIamArn":    "arn:aws:iam::414351767826:role/SaasUsageDeliveryRole-prod-IAMRole-3PLHICCRR1TK",
		"unityCatalogueIamArn": "arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL",
	},
	"aws-us-gov": {
		"accountId":            "044793339203",
		"awsNamespace":         "aws-us-gov",
		"logDeliveryIamArn":    "arn:aws-us-gov:iam::044793339203:role/SaasUsageDeliveryRole-prod-aws-gov-IAMRole-L4QM0RCHYQ1G",
		"unityCatalogueIamArn": "arn:aws-us-gov:iam::044793339203:role/unity-catalog-prod-UCMasterRole-1QRFA8SGY15OJ",
	},
	"aws-us-gov-dod": {
		"accountId":            "170661010020",
		"awsNamespace":         "aws-us-gov",
		"logDeliveryIamArn":    "arn:aws-us-gov:iam::170661010020:role/SaasUsageDeliveryRole-prod-aws-gov-dod-IAMRole-1DMEHBYR8VC5P",
		"unityCatalogueIamArn": "arn:aws-us-gov:iam::170661010020:role/unity-catalog-prod-UCMasterRole-1DI6DL6ZP26AS",
	},
}

var AwsPartitions = []string{"aws", "aws-us-gov", "aws-us-gov-dod"}
var AwsPartitionsValidationError = "aws_partition must be either 'aws' or 'aws-us-gov' or 'aws-us-gov-dod'"
