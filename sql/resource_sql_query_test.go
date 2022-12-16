package sql

import (
	"encoding/json"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/stretchr/testify/assert"
)

func TestQueryCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries",
				ExpectedRequest: api.Query{
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
					Options: &api.QueryOptions{
						RunAsRole: "viewer",
					},
				},
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
					Options: &api.QueryOptions{
						RunAsRole: "viewer",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
					Options: &api.QueryOptions{
						RunAsRole: "viewer",
					},
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Create:   true,
		State: map[string]any{
			"data_source_id": "xyz",
			"name":           "Query name",
			"description":    "Query description",
			"query":          "SELECT 1",
			"run_as_role":    "viewer",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
	assert.Equal(t, "xyz", d.Get("data_source_id"))
	assert.Equal(t, "Query name", d.Get("name"))
	assert.Equal(t, "Query description", d.Get("description"))
	assert.Equal(t, "SELECT 1", d.Get("query"))
	assert.Equal(t, "viewer", d.Get("run_as_role"))
}

func TestQueryCreateWithMultipleSchedules(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceSqlQuery(),
		Create:   true,
		HCL: `
			data_source_id = "xyz"
			name = "Query name"
			query = "SELECT 1"

			schedule {
				continuous {
					interval_seconds = 3600
				}
				daily {
					interval_days = 1
					time_of_day = "11:30"
				}
			}
		`,
	}.ExpectError(t, "invalid config supplied. [schedule.#.continuous] Conflicting configuration arguments. [schedule.#.daily] Conflicting configuration arguments")
}

func TestQueryCreateWithContinuousSchedule(t *testing.T) {
	intervalSeconds := 3600
	untilDate := "2021-04-21"

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries",
				ExpectedRequest: api.Query{
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
					Schedule: &api.QuerySchedule{
						Interval:  intervalSeconds,
						Time:      nil,
						DayOfWeek: nil,
						Until:     &untilDate,
					},
				},
				Response: api.Query{
					ID: "foo",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
					Schedule: &api.QuerySchedule{
						Interval:  intervalSeconds,
						Time:      nil,
						DayOfWeek: nil,
						Until:     &untilDate,
					},
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Create:   true,
		HCL: `
			data_source_id = "xyz"
			name = "Query name"
			query = "SELECT 1"

			schedule {
				continuous {
					interval_seconds = 3600
					until_date = "2021-04-21"
				}
			}
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, intervalSeconds, d.Get("schedule.0.continuous.0.interval_seconds"))
	assert.Equal(t, untilDate, d.Get("schedule.0.continuous.0.until_date"))
}

func TestQueryCreateWithDailySchedule(t *testing.T) {
	intervalDays := 2
	intervalSeconds := intervalDays * 24 * 60 * 60
	timeOfDay := "06:00"
	untilDate := "2021-04-21"

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries",
				ExpectedRequest: api.Query{
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
					Schedule: &api.QuerySchedule{
						Interval:  intervalSeconds,
						Time:      &timeOfDay,
						DayOfWeek: nil,
						Until:     &untilDate,
					},
				},
				Response: api.Query{
					ID: "foo",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
					Schedule: &api.QuerySchedule{
						Interval:  intervalSeconds,
						Time:      &timeOfDay,
						DayOfWeek: nil,
						Until:     &untilDate,
					},
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Create:   true,
		HCL: `
			data_source_id = "xyz"
			name = "Query name"
			query = "SELECT 1"

			schedule {
				daily {
					interval_days = 2
					time_of_day = "06:00"
					until_date = "2021-04-21"
				}
			}
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, intervalDays, d.Get("schedule.0.daily.0.interval_days"))
	assert.Equal(t, timeOfDay, d.Get("schedule.0.daily.0.time_of_day"))
	assert.Equal(t, untilDate, d.Get("schedule.0.daily.0.until_date"))
}

func TestQueryCreateWithWeeklySchedule(t *testing.T) {
	intervalWeeks := 2
	intervalSeconds := intervalWeeks * 7 * 24 * 60 * 60
	timeOfDay := "06:00"
	dayOfWeek := "Sunday"
	untilDate := "2021-04-21"

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries",
				ExpectedRequest: api.Query{
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
					Schedule: &api.QuerySchedule{
						Interval:  intervalSeconds,
						Time:      &timeOfDay,
						DayOfWeek: &dayOfWeek,
						Until:     &untilDate,
					},
				},
				Response: api.Query{
					ID: "foo",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
					Schedule: &api.QuerySchedule{
						Interval:  intervalSeconds,
						Time:      &timeOfDay,
						DayOfWeek: &dayOfWeek,
						Until:     &untilDate,
					},
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Create:   true,
		HCL: `
			data_source_id = "xyz"
			name = "Query name"
			query = "SELECT 1"

			schedule {
				weekly {
					interval_weeks = 2
					time_of_day = "06:00"
					day_of_week = "Sunday"
					until_date = "2021-04-21"
				}
			}
		`,
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, intervalWeeks, d.Get("schedule.0.weekly.0.interval_weeks"))
	assert.Equal(t, dayOfWeek, d.Get("schedule.0.weekly.0.day_of_week"))
	assert.Equal(t, timeOfDay, d.Get("schedule.0.weekly.0.time_of_day"))
	assert.Equal(t, untilDate, d.Get("schedule.0.weekly.0.until_date"))
}

func TestQueryCreateDeletesDefaultVisualization(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries",
				ExpectedRequest: api.Query{
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
				},
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",

					// The automatically created visualization should be deleted.
					Visualizations: []json.RawMessage{
						json.RawMessage(`
							{
								"id": 12345
							}
						`),
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Query:        "SELECT 1",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/sql/visualizations/12345",
			},
		},
		Resource: ResourceSqlQuery(),
		Create:   true,
		State: map[string]any{
			"data_source_id": "xyz",
			"name":           "Query name",
			"query":          "SELECT 1",
		},
	}.Apply(t)

	assert.NoError(t, err, err)
}

func TestQueryRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Query name",
					Description:  "Query description",
					Query:        "SELECT 1",
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Read:     true,
		ID:       "foo",
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
}

func TestQueryReadWithSchedule(t *testing.T) {
	// Note: this tests that if a schedule is returned by the API,
	// it will always show up in the resulting resource data.
	// If it doesn't, we wouldn't be able to erase a schedule
	// that was defined out of band.
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID: "foo",
					Schedule: &api.QuerySchedule{
						Interval: 12345,
					},
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Read:     true,
		ID:       "foo",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, 12345, d.Get("schedule.0.continuous.0.interval_seconds"))
}

func TestQueryUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Updated name",
					Description:  "Updated description",
					Query:        "SELECT 2",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: api.Query{
					ID:           "foo",
					DataSourceID: "xyz",
					Name:         "Updated name",
					Description:  "Updated description",
					Query:        "SELECT 2",
				},
			},
		},
		Resource: ResourceSqlQuery(),
		Update:   true,
		ID:       "foo",
		State: map[string]any{
			"data_source_id": "xyz",
			"name":           "Updated name",
			"description":    "Updated description",
			"query":          "SELECT 2",
		},
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
	assert.Equal(t, "xyz", d.Get("data_source_id"))
	assert.Equal(t, "Updated name", d.Get("name"))
	assert.Equal(t, "Updated description", d.Get("description"))
	assert.Equal(t, "SELECT 2", d.Get("query"))
}

func TestQueryUpdateWithParams(t *testing.T) {
	body := api.Query{
		ID:           "foo",
		DataSourceID: "xyz",
		Name:         "Updated name",
		Query:        "SELECT 1, 2, 3, 4",
		Options: &api.QueryOptions{
			Parameters: []any{
				api.QueryParameterText{
					QueryParameter: api.QueryParameter{
						Name:  "1",
						Title: "Title for column 1",
					},
				},
				api.QueryParameterNumber{
					QueryParameter: api.QueryParameter{
						Name:  "2",
						Title: "Title for column 2",
					},
				},
				api.QueryParameterEnum{
					QueryParameter: api.QueryParameter{
						Name:  "3",
						Title: "Title for column 3",
					},
					Options: "e1\ne2",
					Values:  []string{"e1"},
					Multi: &api.QueryParameterMultipleValuesOptions{
						Prefix:    "\"",
						Suffix:    "\"",
						Separator: ",",
					},
				},
				api.QueryParameterEnum{
					QueryParameter: api.QueryParameter{
						Name:  "3",
						Title: "Title for column 3 without multiple",
					},
					Options: "e1\ne2",
					Values:  []string{"e1"},
					Multi:   nil,
				},
				api.QueryParameterQuery{
					QueryParameter: api.QueryParameter{
						Name:  "4",
						Title: "Title for column 4",
					},
					QueryID: "abc",
					Values:  []string{"e1"},
					Multi: &api.QueryParameterMultipleValuesOptions{
						Prefix:    "\"",
						Suffix:    "\"",
						Separator: ",",
					},
				},
				api.QueryParameterQuery{
					QueryParameter: api.QueryParameter{
						Name:  "4",
						Title: "Title for column 4 without multiple",
					},
					QueryID: "abc",
					Values:  []string{"e1"},
					Multi:   nil,
				},
				api.QueryParameterDate{
					QueryParameter: api.QueryParameter{
						Name: "5",
					},
				},
				api.QueryParameterDateTime{
					QueryParameter: api.QueryParameter{
						Name: "6",
					},
				},
				api.QueryParameterDateTimeSec{
					QueryParameter: api.QueryParameter{
						Name: "7",
					},
				},
				api.QueryParameterDateRange{
					QueryParameterRangeBase: api.QueryParameterRangeBase{
						QueryParameter: api.QueryParameter{
							Name: "8",
						},
						Value: map[string]string{"start": "2022-11-20", "end": "2022-11-22"},
					},
				},
				api.QueryParameterDateTimeRange{
					QueryParameterRangeBase: api.QueryParameterRangeBase{
						QueryParameter: api.QueryParameter{
							Name: "9",
						},
					},
				},
				api.QueryParameterDateTimeSecRange{
					QueryParameterRangeBase: api.QueryParameterRangeBase{
						QueryParameter: api.QueryParameter{
							Name: "10",
						},
					},
				},
			},
		},
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: body,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/queries/foo",
				Response: body,
			},
		},
		Resource: ResourceSqlQuery(),
		Update:   true,
		ID:       "foo",
		HCL: `
			data_source_id = "xyz"
			name = "name"
			query = "SELECT 1, 2, 3, 4"
			
			parameter {
				name = "1"
				title = "Title for column 1"
				text {
					value = ""
				}
			}

			parameter {
				name = "2"
				title = "Title for column 2"
				number {
					value = 0
				}
			}

			parameter {
				name = "3"
				title = "Title for column 3"
				enum {
					options = ["e1", "e2"]
					values = ["e1"]
					multiple {
						prefix = "\""
						suffix = "\""
						separator = ","
					}
				}
			}

			parameter {
				name = "3"
				title = "Title for column 3 without multiple"
				enum {
					options = ["e1", "e2"]
					value = "e1"
				}
			}

			parameter {
				name = "4"
				title = "Title for column 4"
				query {
					query_id = "abc"
					values = ["e1"]
					multiple {
						prefix = "\""
						suffix = "\""
						separator = ","
					}
				}
			}

			parameter {
				name = "4"
				title = "Title for column 4 without multiple"
				query {
					query_id = "abc"
					value = "e1"
				}
			}

			parameter {
				name = "5"
				date {
					value = ""
				}
			}

			parameter {
				name = "6"
				datetime {
					value = ""
				}
			}

			parameter {
				name = "7"
				datetimesec {
					value = ""
				}
			}

			parameter {
				name = "8"
				date_range {
					range {
						start = "2022-11-20"
						end = "2022-11-22"
					}
				}
			}

			parameter {
				name = "9"
				datetime_range {
					value = ""
				}
			}

			parameter {
				name = "10"
				datetimesec_range {
					value = ""
				}
			}
		`,
	}.Apply(t)

	assert.NoError(t, err, err)

	assert.Equal(t, "foo", d.Id())
	assert.Equal(t, "xyz", d.Get("data_source_id"))
	assert.Equal(t, "Updated name", d.Get("name"))
	assert.Equal(t, "SELECT 1, 2, 3, 4", d.Get("query"))
	assert.Len(t, d.Get("parameter").([]any), 12)
}

func TestQueryDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/sql/queries/foo",
			},
		},
		Resource: ResourceSqlQuery(),
		Delete:   true,
		ID:       "foo",
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "foo", d.Id(), "Resource ID should not be empty")
}

func TestResourceQueryCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSqlQuery())
}
