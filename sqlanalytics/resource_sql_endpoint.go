package sqlanalytics

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// SQLEndpoint ...
type SQLEndpoint struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name"`
	ClusterSize       string `json:"cluster_size"`
	AutoStopMinutes   int    `json:"auto_stop_mins,omitempty"`
	MinNumClusters    int    `json:"min_num_clusters,omitempty"`
	MaxNumClusters    int    `json:"max_num_clusters,omitempty"`
	NumClusters       int    `json:"num_clusters,omitempty"`
	NumActiveSessions int    `json:"num_active_sessions,omitempty"`
	State             string `json:"state,omitempty" tf:"computed"`
	JdbcURL           string `json:"jdbc_url,omitempty" tf:"computed"`
}

// EndpointList ...
type EndpointList struct {
	Endpoints []SQLEndpoint `json:"endpoints"`
}

// NewSQLEndpointsAPI ...
func NewSQLEndpointsAPI(ctx context.Context, m interface{}) SQLEndpointsAPI {
	return SQLEndpointsAPI{m.(*common.DatabricksClient), ctx}
}

// SQLEndpointsAPI ...
type SQLEndpointsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// List all SQL endpoints
func (a SQLEndpointsAPI) List() (l EndpointList, err error) {
	err = a.client.Get(a.context, "/sql/endpoints", nil, &l)
	return
}

// Start ..
func (a SQLEndpointsAPI) Start(endpointID string, timeout time.Duration) error {
	err := a.client.Post(a.context, fmt.Sprintf("/sql/endpoints/%s/start", endpointID), nil, nil)
	if err != nil {
		return err
	}
	return a.waitForRunning(endpointID, timeout)
}

// Stop ...
func (a SQLEndpointsAPI) Stop(endpointID string) error {
	return a.client.Post(a.context, fmt.Sprintf("/sql/endpoints/%s/stop", endpointID), nil, nil)
}

// Get ...
func (a SQLEndpointsAPI) Get(endpointID string) (se SQLEndpoint, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/sql/endpoints/%s", endpointID), nil, &se)
	return
}

// Create ...
func (a SQLEndpointsAPI) Create(se *SQLEndpoint, timeout time.Duration) error {
	// maybe response should be something else...
	err := a.client.Post(a.context, "/sql/endpoints", se, se)
	if err != nil {
		return err
	}
	return a.waitForRunning(se.ID, timeout)
}

func (a SQLEndpointsAPI) waitForRunning(id string, timeout time.Duration) error {
	return resource.RetryContext(a.context, timeout, func() *resource.RetryError {
		endpoint, err := a.Get(id)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		switch endpoint.State {
		case "RUNNING":
			return nil
		case "DELETED":
			return resource.NonRetryableError(
				fmt.Errorf("Endpoint got deleted during creation"))
		default:
			msg := fmt.Errorf("Endpoint %s is %s", id, endpoint.State)
			log.Printf("[INFO] %s", msg.Error())
			return resource.RetryableError(msg)
		}
	})
}

// Edit ...
func (a SQLEndpointsAPI) Edit(se SQLEndpoint, timeout time.Duration) error {
	err := a.client.Post(a.context, fmt.Sprintf("/sql/endpoints/%s/edit", se.ID), se, nil)
	if err != nil {
		return err
	}
	r, err := a.Get(se.ID)
	if err != nil {
		return err
	}
	if r.State != "RUNNING" {
		return a.Start(se.ID, timeout)
	}
	return a.waitForRunning(se.ID, timeout)
}

// Delete ...
func (a SQLEndpointsAPI) Delete(endpointID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/sql/endpoints/%s", endpointID), nil)
}

// ResourceSQLEndpoint ...
func ResourceSQLEndpoint() *schema.Resource {
	s := internal.StructToSchema(SQLEndpoint{}, func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	})
	return util.CommonResource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var se SQLEndpoint
			if err := internal.DataToStructPointer(d, s, &se); err != nil {
				return err
			}
			if err := NewSQLEndpointsAPI(ctx, c).Create(&se, d.Timeout(schema.TimeoutCreate)); err != nil {
				return err
			}
			d.SetId(se.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			endpointsAPI := NewSQLEndpointsAPI(ctx, c)
			se, err := endpointsAPI.Get(d.Id())
			if err != nil {
				return err
			}
			if err = internal.StructToData(se, s, d); err != nil {
				return err
			}
			return endpointsAPI.waitForRunning(d.Id(), d.Timeout(schema.TimeoutRead))
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var se SQLEndpoint
			if err := internal.DataToStructPointer(d, s, &se); err != nil {
				return err
			}
			return NewSQLEndpointsAPI(ctx, c).Edit(se, d.Timeout(schema.TimeoutUpdate))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSQLEndpointsAPI(ctx, c).Delete(d.Id())
		},
		Schema: s,
	}.ToResource()
}
