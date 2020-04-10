package db

import (
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAWSS3Mount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSS3Create,
		Read:   resourceAWSS3Read,
		Delete: resourceAWSS3Delete,

		Schema: map[string]*schema.Schema{
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"s3_bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mount_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAWSS3Create(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	clusterId := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterId, client)
	if err != nil {
		return err
	}
	s3BucketName := d.Get("s3_bucket_name").(string)
	mountName := d.Get("mount_name").(string)

	s3BucketMount := service.NewAWSIamMount(s3BucketName, mountName)

	err = s3BucketMount.Create(client, clusterId)
	if err != nil {
		return err
	}

	d.SetId(mountName)

	err = d.Set("cluster_id", clusterId)
	if err != nil {
		return err
	}
	err = d.Set("mount_name", mountName)
	if err != nil {
		return err
	}

	return resourceAWSS3Read(d, m)
}

func resourceAWSS3Read(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	clusterId := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterId, client)
	if err != nil {
		return err
	}
	s3BucketName := d.Get("s3_bucket_name").(string)
	mountName := d.Get("mount_name").(string)

	s3BucketMount := service.NewAWSIamMount(s3BucketName, mountName)

	s3BucketNameMounted, err := s3BucketMount.Read(client, clusterId)
	if err != nil {
		return err
	}

	err = d.Set("s3_bucket_name", s3BucketNameMounted)
	return err
}

func resourceAWSS3Delete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	clusterId := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterId, client)
	if err != nil {
		return err
	}
	s3BucketName := d.Get("s3_bucket_name").(string)
	mountName := d.Get("mount_name").(string)
	s3BucketMount := service.NewAWSIamMount(s3BucketName, mountName)
	return s3BucketMount.Delete(client, clusterId)
}
