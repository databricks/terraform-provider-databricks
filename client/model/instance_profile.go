package model

// InstanceProfileInfo contains the ARN for aws instance profiles
type InstanceProfileInfo struct {
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
}

// InstanceProfileList ...
type InstanceProfileList struct {
	InstanceProfiles []InstanceProfileInfo `json:"instance_profiles,omitempty"`
}
