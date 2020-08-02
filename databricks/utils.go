package databricks

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
)

func convertListInterfaceToString(m []interface{}) []string {
	response := []string{}
	for _, v := range m {
		if v != nil {
			response = append(response, v.(string))
		}
	}
	return response
}

func getMapFromOneItemList(input interface{}) map[string]interface{} {
	inputList := input.([]interface{})
	if len(inputList) >= 1 {
		return inputList[0].(map[string]interface{})
	}
	return nil
}

// PackagedMWSIds is a struct that contains both the MWS acct id and the ResourceId (resources are networks, creds, etc.)
type PackagedMWSIds struct {
	MwsAcctID  string
	ResourceID string
}

// Helps package up MWSAccountId with another id such as credentials id or network id
// uses format mwsAcctID/otherId
func packMWSAccountID(idsToPackage PackagedMWSIds) string {
	return fmt.Sprintf("%s/%s", idsToPackage.MwsAcctID, idsToPackage.ResourceID)
}

// Helps unpackage MWSAccountId from another id such as credentials id or network id
func unpackMWSAccountID(combined string) (PackagedMWSIds, error) {
	var packagedMWSIds PackagedMWSIds
	parts := strings.Split(combined, "/")
	if len(parts) != 2 {
		// TODO: set id to "" if invalid format
		return packagedMWSIds, fmt.Errorf("unpacked account has more than or less than two parts, combined id: %s", combined)
	}
	packagedMWSIds.MwsAcctID = parts[0]
	packagedMWSIds.ResourceID = parts[1]
	return packagedMWSIds, nil
}

// ValidateInstanceProfileARN is a ValidateFunc that ensures the role id is a valid aws iam instance profile arn
func ValidateInstanceProfileARN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)

	if v == "" {
		return nil, []error{fmt.Errorf("%s is empty got: %s, must be an aws instance profile arn", key, v)}
	}

	// Parse and verify instance profiles
	instanceProfileArn, err := arn.Parse(v)
	if err != nil {
		return nil, []error{fmt.Errorf("%s is invalid got: %s received error: %w", key, v, err)}
	}
	// Verify instance profile resource type, Resource gets parsed as instance-profile/<profile-name>
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return nil, []error{fmt.Errorf("%s must be an instance profile resource, got: %s in %s",
			key, instanceProfileArn.Resource, v)}
	}
	return nil, nil
}

var PathEmptyError error = errors.New("provided path is empty")

// we would never want to handle root directories in regards to creating them
var DirPathRootDirError error = errors.New("dir path is root directory")

// Os libraries behave bizarely on windows as they will replace slashes with other values.
// This causes issues & errors when submitting the request
func GetParentDirPath(filePath string) (string, error) {
	if filePath == "" {
		return "", PathEmptyError
	}

	pathParts := strings.Split(filePath, "/")

	// if length of pathParts is just two items then the parent should be the root directory
	if len(pathParts) == 2 {
		return "", DirPathRootDirError
	}

	dirPath := strings.Join(pathParts[0:len(pathParts)-1], "/")

	return dirPath, nil
}
