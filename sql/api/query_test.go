package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryMarshalUnmarshal(t *testing.T) {
	q := Query{
		ID:           "id",
		DataSourceID: "data_source_id",
		Name:         "name",
		Description:  "description",
		Query:        "SELECT 1",
		Schedule: &QuerySchedule{
			Interval: 1000,
		},
		Options: &QueryOptions{
			Parameters: []any{
				&QueryParameterText{
					QueryParameter: QueryParameter{
						Name:  "n1",
						Title: "t1",
					},
					Value: "v1",
				},
				&QueryParameterNumber{
					QueryParameter: QueryParameter{
						Name:  "n2",
						Title: "t2",
					},
					Value: 1234.5,
				},
				&QueryParameterEnum{
					QueryParameter: QueryParameter{
						Name:  "n3",
						Title: "t3",
					},
					Values:  []string{"single"},
					Options: "foo\nbar",
				},
				&QueryParameterEnum{
					QueryParameter: QueryParameter{
						Name:  "n4",
						Title: "t4",
					},
					Values:  []string{"multiple1", "multiple2"},
					Options: "foo\nbar",
					Multi: &QueryParameterMultipleValuesOptions{
						Prefix:    "@@@",
						Suffix:    "###",
						Separator: "$$$",
					},
				},
				&QueryParameterQuery{
					QueryParameter: QueryParameter{
						Name:  "n5",
						Title: "t5",
					},
					Values:  []string{"single"},
					QueryID: "queryID",
				},
				&QueryParameterQuery{
					QueryParameter: QueryParameter{
						Name:  "n6",
						Title: "t6",
					},
					Values:  []string{"multiple1", "multiple2"},
					QueryID: "queryID",
					Multi: &QueryParameterMultipleValuesOptions{
						Prefix:    "@@@",
						Suffix:    "###",
						Separator: "$$$",
					},
				},
				&QueryParameterDate{
					QueryParameter: QueryParameter{
						Name:  "n7",
						Title: "t7",
					},
					Value: "xyz",
				},
				&QueryParameterDateTime{
					QueryParameter: QueryParameter{
						Name:  "n8",
						Title: "t8",
					},
					Value: "xyz",
				},
				&QueryParameterDateTimeSec{
					QueryParameter: QueryParameter{
						Name:  "n9",
						Title: "t9",
					},
					Value: "xyz",
				},
				&QueryParameterDateRange{
					QueryParameterRangeBase: QueryParameterRangeBase{
						QueryParameter: QueryParameter{
							Name:  "n10",
							Title: "t10",
						},
						StringValue: "xyz",
					},
				},
				&QueryParameterDateTimeRange{
					QueryParameterRangeBase: QueryParameterRangeBase{
						QueryParameter: QueryParameter{
							Name:  "n11",
							Title: "t11",
						},
						StringValue: "xyz",
					},
				},
				&QueryParameterDateTimeSecRange{
					QueryParameterRangeBase: QueryParameterRangeBase{
						QueryParameter: QueryParameter{
							Name:  "n12",
							Title: "t12",
						},
						StringValue: "xyz",
					},
				},
			},
		},
		Tags: []string{"tag1", "tag2"},
	}

	out, err := json.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}

	var qp Query
	if err := json.Unmarshal(out, &qp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, q, qp)
}

func TestQueryMarshalUnmarshalRanges(t *testing.T) {
	q := Query{
		ID:           "id",
		DataSourceID: "data_source_id",
		Name:         "name",
		Description:  "description",
		Query:        "SELECT 1",
		Schedule: &QuerySchedule{
			Interval: 1000,
		},
		Options: &QueryOptions{
			Parameters: []any{
				&QueryParameterDateRange{
					QueryParameterRangeBase: QueryParameterRangeBase{
						QueryParameter: QueryParameter{
							Name:  "n13",
							Title: "t13",
						},
						RangeValue: &DateTimeRange{Start: "2022-11-20", End: "2022-11-22"},
					},
				},
			},
		},
		Tags: []string{"tag1", "tag2"},
	}

	out, err := json.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(out), `"value":{"start":"2022-11-20","end":"2022-11-22"`)

	var qp Query
	if err := json.Unmarshal(out, &qp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, q, qp)
}
