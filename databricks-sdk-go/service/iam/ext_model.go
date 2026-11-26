package iam

type PartialUpdate struct {
	// Unique ID in the Databricks workspace.
	Id string `json:"-" url:"-"`

	Operations []Patch `json:"Operations,omitempty"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas []PatchSchema `json:"schemas,omitempty"`
}
