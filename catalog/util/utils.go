package util

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func SliceToSet(in []string) *schema.Set {
	var out []any
	for _, v := range in {
		out = append(out, v)
	}
	return schema.NewSet(schema.HashString, out)
}

func SetToSlice(set *schema.Set) (ss []string) {
	for _, v := range set.List() {
		ss = append(ss, v.(string))
	}
	return
}

func SliceWithoutString(in []string, without string) (out []string) {
	for _, v := range in {
		if v == without {
			continue
		}
		out = append(out, v)
	}
	return
}

func ToStringSlice(in []interface{}) (out []string) {
	for _, v := range in {
		out = append(out, v.(string))
	}
	return
}
