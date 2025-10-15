package aws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataAwsUnityCatalogAssumeRolePolicy() common.Resource {
	type AwsUcAssumeRolePolicy struct {
		common.Namespace
		RoleName           string `json:"role_name"`
		UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty" tf:"computed"`
		ExternalId         string `json:"external_id"`
		AwsAccountId       string `json:"aws_account_id"`
		AwsPartition       string `json:"aws_partition,omitempty" tf:"default:aws"`
		JSON               string `json:"json" tf:"computed"`
		Id                 string `json:"id" tf:"computed"`
	}
	return common.NoClientData(func(ctx context.Context, data *AwsUcAssumeRolePolicy) error {
		if !slices.Contains(AwsPartitions, data.AwsPartition) {
			return errors.New(AwsPartitionsValidationError)
		}
		awsNamespace := AwsConfig[data.AwsPartition]["awsNamespace"]
		if data.UnityCatalogIamArn == "" {
			data.UnityCatalogIamArn = AwsConfig[data.AwsPartition]["unityCatalogueIamArn"]
		}

		policy := awsIamPolicy{
			Version: "2012-10-17",
			Statements: []*awsIamPolicyStatement{
				{
					Sid:     "UnityCatalogAssumeRole",
					Effect:  "Allow",
					Actions: "sts:AssumeRole",
					Condition: map[string]map[string]string{
						"StringEquals": {
							"sts:ExternalId": data.ExternalId,
						},
					},
					Principal: map[string]string{
						"AWS": data.UnityCatalogIamArn,
					},
				},
				{
					Sid:     "ExplicitSelfRoleAssumption",
					Effect:  "Allow",
					Actions: "sts:AssumeRole",
					Condition: map[string]map[string]string{
						"ArnLike": {
							"aws:PrincipalArn": fmt.Sprintf("arn:%s:iam::%s:role/%s", awsNamespace, data.AwsAccountId, data.RoleName),
						},
						"StringEquals": {
							"sts:ExternalId": data.ExternalId,
						},
					},
					Principal: map[string]string{
						"AWS": fmt.Sprintf("arn:%s:iam::%s:root", awsNamespace, data.AwsAccountId),
					},
				},
			},
		}
		policyJSON, err := json.MarshalIndent(policy, "", "  ")
		if err != nil {
			return err
		}
		data.Id = fmt.Sprintf("%s-%s-%s", data.AwsAccountId, data.RoleName, data.ExternalId)
		data.JSON = string(policyJSON)
		return nil
	})
}
