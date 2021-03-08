package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/pkg/errors"
)

// Wrapper is an API wrapper for the SQL Analytics queries, visualizations, and dashboards API.
type Wrapper struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewWrapper creates and returns an API wrapper.
func NewWrapper(ctx context.Context, m interface{}) Wrapper {
	return Wrapper{m.(*common.DatabricksClient), ctx}
}

const sqlaBasePath string = "/preview/sql"

// CreateVisualization ...
func (a Wrapper) CreateVisualization(v *Visualization) (*Visualization, error) {
	var vp Visualization
	err := a.client.Post(a.context, fmt.Sprintf("%s/visualizations", sqlaBasePath), v, &vp)
	if err != nil {
		return nil, err
	}

	// Set query ID on returned object.
	// It's not included in the POST response.
	vp.QueryID = v.QueryID

	return &vp, err
}

// ReadVisualization ...
func (a Wrapper) ReadVisualization(v *Visualization) (*Visualization, error) {
	if v.QueryID == "" {
		return nil, errors.Errorf("Cannot read visualization without query ID")
	}

	var q Query
	err := a.client.Get(a.context, fmt.Sprintf("%s/queries/%s", sqlaBasePath, v.QueryID), nil, &q)
	if err != nil {
		return nil, err
	}

	// Look for matching visualization ID.
	for _, vp := range q.Visualizations {
		var vnew Visualization
		err = json.Unmarshal(vp, &vnew)
		if err != nil {
			return nil, err
		}

		if vnew.ID == v.ID {
			// Include query ID in returned object.
			// It's not part of the API response.
			vnew.QueryID = v.QueryID
			return &vnew, nil
		}
	}

	return nil, errors.Errorf("Cannot find visualization %d attached to query %s", v.ID, v.QueryID)
}

// UpdateVisualization ...
func (a Wrapper) UpdateVisualization(v *Visualization) (*Visualization, error) {
	var vp Visualization
	err := a.client.Post(a.context, fmt.Sprintf("%s/visualizations/%d", sqlaBasePath, v.ID), v, &vp)
	if err != nil {
		return nil, err
	}

	// Set query ID on returned object.
	// It's not included in the POST response.
	vp.QueryID = v.QueryID

	return &vp, nil
}

// DeleteVisualization ...
func (a Wrapper) DeleteVisualization(v *Visualization) error {
	return a.client.Delete(a.context, fmt.Sprintf("%s/visualizations/%d", sqlaBasePath, v.ID), nil)
}

// CreateQuery ...
func (a Wrapper) CreateQuery(q *Query) (*Query, error) {
	var qp Query
	err := a.client.Post(a.context, fmt.Sprintf("%s/queries", sqlaBasePath), q, &qp)
	if err != nil {
		return nil, err
	}

	return &qp, err
}

// ReadQuery ...
func (a Wrapper) ReadQuery(q *Query) (*Query, error) {
	var qp Query
	err := a.client.Get(a.context, fmt.Sprintf("%s/queries/%s", sqlaBasePath, q.ID), nil, &qp)
	if err != nil {
		return nil, err
	}

	return &qp, nil
}

// UpdateQuery ...
func (a Wrapper) UpdateQuery(q *Query) (*Query, error) {
	var qp Query
	err := a.client.Post(a.context, fmt.Sprintf("%s/queries/%s", sqlaBasePath, q.ID), q, &qp)
	if err != nil {
		return nil, err
	}

	return &qp, nil
}

// DeleteQuery ...
func (a Wrapper) DeleteQuery(q *Query) error {
	return a.client.Delete(a.context, fmt.Sprintf("%s/queries/%s", sqlaBasePath, q.ID), nil)
}
