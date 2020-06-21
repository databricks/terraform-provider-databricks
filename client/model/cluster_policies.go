package model

// ClusterPolicy defines cluster policy
type ClusterPolicy struct {
	PolicyID           string `json:"policy_id,omitempty"`
	Name               string `json:"name"`
	Definition         string `json:"definition"`
	CreatedAtTimeStamp int64  `json:"created_at_timestamp"`
}

// ClusterPolicyCreate is the endity used for request
type ClusterPolicyCreate struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}
