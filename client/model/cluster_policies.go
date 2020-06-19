package model

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// AttributePolicy defines JSON mapping for attribute policy
type AttributePolicy struct {
	Path         *string       `json:"-"`
	Type         string        `json:"type"`
	Value        interface{}   `json:"value,omitempty"`
	DefaultValue interface{}   `json:"defaultValue,omitempty"`
	Values       []interface{} `json:"values,omitempty"`
	Hidden       bool          `json:"hidden,omitempty"`
	IsOptional   bool          `json:"isOptional,omitempty"`
	Pattern      string        `json:"pattern,omitempty"`
	MinValue     int           `json:"minValue,omitempty"`
	MaxValue     int           `json:"maxValue,omitempty"`
}

// ClusterPolicy defines cluster policy
type ClusterPolicy struct {
	PolicyID           string `json:"policy_id,omitempty"`
	Name               string `json:"name"`
	Definition         string `json:"definition"`
	CreatedAtTimeStamp int64  `json:"created_at_timestamp"`

	Attributes map[string]*AttributePolicy
}

// ParseDefinition parses policy definition
func (clusterPolicy *ClusterPolicy) ParseDefinition(definition string) error {
	clusterPolicy.Definition = definition

	err := json.Unmarshal([]byte(definition), &clusterPolicy.Attributes)
	if err != nil {
		return err
	}

	return nil
}

// ClusterPolicyCreate is the endity used for request
type ClusterPolicyCreate struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

// // Prepare sets definition from attributes
// func (clusterPolicy *ClusterPolicy) Prepare() ([]byte, error) {
// 	policyJSONBytes, err := json.Marshal(clusterPolicy.attributes)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "Problem serializing %s policy", clusterPolicy.Name)
// 	}
// 	clusterPolicy.Definition = string(policyJSONBytes)
// }

// MarshalJSON is called when json.Marshal is invoked
func (clusterPolicy *ClusterPolicy) MarshalJSON() ([]byte, error) {
	policyJSONBytes, err := json.Marshal(clusterPolicy.Attributes)
	if err != nil {
		return nil, errors.Wrapf(err, "Problem serializing %s policy", clusterPolicy.Name)
	}
	clusterPolicy.Definition = string(policyJSONBytes)
	return json.Marshal(&struct {
		PolicyID   string `json:"policy_id,omitempty"`
		Name       string `json:"name"`
		Definition string `json:"definition"`
	}{
		clusterPolicy.PolicyID, clusterPolicy.Name, clusterPolicy.Definition,
	})
}

// ToString returns debug JSON, ignoring errors
func (clusterPolicy *ClusterPolicy) ToString() string {
	j, err := clusterPolicy.MarshalJSON()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", j)
}
