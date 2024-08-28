package mws

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// NewVPCEndpointAPI creates VPCEndpointAPI instance from provider meta
func NewVPCEndpointAPI(ctx context.Context, m any) VPCEndpointAPI {
	return VPCEndpointAPI{m.(*common.DatabricksClient), ctx}
}

// VPCEndpointAPI exposes the mws VPC endpoint API
type VPCEndpointAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates the VPC endpoint registeration process
func (a VPCEndpointAPI) Create(vpcEndpoint *VPCEndpoint) error {
	vpcEndpointAPIPath := fmt.Sprintf("/accounts/%s/vpc-endpoints", vpcEndpoint.AccountID)
	err := a.client.Post(a.context, vpcEndpointAPIPath, vpcEndpoint, &vpcEndpoint)
	if err != nil {
		return err
	}
	return resource.RetryContext(a.context, 15*time.Minute, func() *resource.RetryError {
		ve, err := a.Read(vpcEndpoint.AccountID, vpcEndpoint.VPCEndpointID)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if ve.GcpVpcEndpointInfo != nil {
			return nil
		}
		switch state := strings.ToLower(ve.State); state {
		case "available":
			return nil
		case "pending", "pendingacceptance":
			return resource.RetryableError(
				fmt.Errorf("endpoint %s is still %s",
					ve.AwsVPCEndpointID, ve.State))
		default:
			return resource.NonRetryableError(
				fmt.Errorf("cannot register %s: %s",
					ve.AwsVPCEndpointID, ve.State))
		}
	})
}

// Read returns the VPCEndpoint object along with metadata and any additional errors when attaching to workspace
func (a VPCEndpointAPI) Read(mwsAcctID, vpcEndpointID string) (ve VPCEndpoint, err error) {
	vpcEndpointAPIPath := fmt.Sprintf("/accounts/%s/vpc-endpoints/%s", mwsAcctID, vpcEndpointID)
	err = a.client.Get(a.context, vpcEndpointAPIPath, nil, &ve)
	return
}

// Delete deletes the VPCEndpoint object given a VPCEndpoint id
func (a VPCEndpointAPI) Delete(mwsAcctID, vpcEndpointID string) error {
	vpcEndpointAPIPath := fmt.Sprintf("/accounts/%s/vpc-endpoints/%s", mwsAcctID, vpcEndpointID)
	return a.client.Delete(a.context, vpcEndpointAPIPath, nil)
}

// List lists all the available network objects in the mws account
func (a VPCEndpointAPI) List(mwsAcctID string) ([]VPCEndpoint, error) {
	var mwsVPCEndpointList []VPCEndpoint
	vpcEndpointAPIPath := fmt.Sprintf("/accounts/%s/vpc-endpoints", mwsAcctID)
	err := a.client.Get(a.context, vpcEndpointAPIPath, nil, &mwsVPCEndpointList)
	return mwsVPCEndpointList, err
}

func ResourceMwsVpcEndpoint() common.Resource {
	s := common.StructToSchema(VPCEndpoint{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["aws_vpc_endpoint_id"].ExactlyOneOf = []string{"aws_vpc_endpoint_id", "gcp_vpc_endpoint_info"}
		s["region"].ExactlyOneOf = []string{"region", "gcp_vpc_endpoint_info"}
		s["aws_endpoint_service_id"].ConflictsWith = []string{"gcp_vpc_endpoint_info"}
		s["aws_account_id"].ConflictsWith = []string{"gcp_vpc_endpoint_info"}
		s["gcp_vpc_endpoint_info"].ConflictsWith = []string{"aws_vpc_endpoint_id", "region", "aws_endpoint_service_id", "aws_account_id"}
		return s
	})
	p := common.NewPairSeparatedID("account_id", "vpc_endpoint_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var vpcEndpoint VPCEndpoint
			common.DataToStructPointer(d, s, &vpcEndpoint)
			if err := NewVPCEndpointAPI(ctx, c).Create(&vpcEndpoint); err != nil {
				return err
			}
			d.Set("vpc_endpoint_id", vpcEndpoint.VPCEndpointID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, vpcEndpointID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			vpcEndpoint, err := NewVPCEndpointAPI(ctx, c).Read(accountID, vpcEndpointID)
			if err != nil {
				return err
			}
			return common.StructToData(vpcEndpoint, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, vpcEndpointID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewVPCEndpointAPI(ctx, c).Delete(accountID, vpcEndpointID)
		},
	}
}
