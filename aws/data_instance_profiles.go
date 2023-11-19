package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

func instanceProfileMatchesFilter(ip *instanceProfileData, filter *instanceProfileFilter) bool {
	var ipMap map[string]interface{}
	m, _ := json.Marshal(ip)
	_ = json.Unmarshal(m, &ipMap)
	val := ipMap[filter.Name]
	stringVal := fmt.Sprint(val)
	re := regexp.MustCompile(filter.Pattern)
	return re.Match([]byte(regexp.QuoteMeta(stringVal)))
}

type instanceProfileData struct {
	Name    string `json:"name"`
	Arn     string `json:"arn"`
	RoleArn string `json:"role_arn"`
	IsMeta  bool   `json:"is_meta"`
}

type instanceProfileFilter struct {
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}

func DataSourceInstanceProfiles() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		InstanceProfiles []instanceProfileData `json:"instance_profiles,omitempty" tf:"computed"`
		Filter           instanceProfileFilter `json:"filter,omitempty" tf:"optional"`
	}, w *databricks.WorkspaceClient) error {

		if data.Filter != (instanceProfileFilter{}) {
			if data.Filter.Pattern == "" {
				return fmt.Errorf("field `pattern` cannot be empty")
			}
			var fieldNames []string
			val := reflect.ValueOf(instanceProfileData{})
			for i := 0; i < val.Type().NumField(); i++ {
				fieldNames = append(fieldNames, val.Type().Field(i).Tag.Get("json"))
			}
			if !slices.Contains(fieldNames, data.Filter.Name) {
				if data.Filter.Name == "" {
					return fmt.Errorf("field `name` cannot be empty")
				}
				return fmt.Errorf("`%s` is not a valid value for the name field. Must be one of [%s]", data.Filter.Name, strings.Join(fieldNames, ", "))
			}
		}

		instanceProfiles, err := w.InstanceProfiles.ListAll(ctx)
		if err != nil {
			return err
		}

		for _, v := range instanceProfiles {
			arnSlices := strings.Split(v.InstanceProfileArn, "/")
			name := arnSlices[len(arnSlices)-1]
			ipData := instanceProfileData{
				Name:    name,
				Arn:     v.InstanceProfileArn,
				RoleArn: v.IamRoleArn,
				IsMeta:  v.IsMetaInstanceProfile,
			}
			if data.Filter != (instanceProfileFilter{}) && !instanceProfileMatchesFilter(&ipData, &data.Filter) {
				continue
			}
			data.InstanceProfiles = append(data.InstanceProfiles, ipData)
		}
		return nil
	})
}
