package catalog

import (
	"slices"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

type fakeFileEventsData struct {
	raw       cty.Value
	rawErr    bool
	effective any
}

func (f fakeFileEventsData) GetRawConfigAt(cty.Path) (cty.Value, diag.Diagnostics) {
	if f.rawErr {
		return cty.DynamicVal, diag.Diagnostics{{Severity: diag.Error, Summary: "no raw config"}}
	}
	return f.raw, nil
}

func (f fakeFileEventsData) GetChange(string) (any, any) {
	return f.effective, f.effective
}

func TestFileEventsDisableRequested(t *testing.T) {
	for name, tc := range map[string]struct {
		data fakeFileEventsData
		want bool
	}{
		"explicit false, enabled by default": {fakeFileEventsData{raw: cty.False, effective: true}, true},
		"explicit false, already disabled":   {fakeFileEventsData{raw: cty.False, effective: false}, false},
		"explicit true, enabled":             {fakeFileEventsData{raw: cty.True, effective: true}, false},
		"explicit true, disabled":            {fakeFileEventsData{raw: cty.True, effective: false}, false},
		"not set, enabled by default":        {fakeFileEventsData{raw: cty.NullVal(cty.Bool), effective: true}, false},
		"unknown, enabled by default":        {fakeFileEventsData{raw: cty.UnknownVal(cty.Bool), effective: true}, false},
		"no raw config":                      {fakeFileEventsData{rawErr: true, effective: true}, false},
		"explicit false, no prior state":     {fakeFileEventsData{raw: cty.False, effective: nil}, false},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, fileEventsDisableRequested(tc.data))
		})
	}
}

// updateFileEventsFixture changes the comment so that Update runs, with file events enabled by
// default on the server (`enable_file_events` unset in state, effective value true).
func updateFileEventsFixture(t *testing.T, enableFileEvents string, effective string,
	check func(catalog.UpdateExternalLocation) bool) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockExternalLocationsAPI().EXPECT()
			e.Update(mock.Anything, mock.MatchedBy(check)).Return(&catalog.ExternalLocationInfo{
				Name: "abc", Url: "s3://foo/bar", CredentialName: "bcd", Comment: "updated",
			}, nil)
			e.GetByName(mock.Anything, "abc").Return(&catalog.ExternalLocationInfo{
				Name: "abc", Url: "s3://foo/bar", CredentialName: "bcd", Comment: "updated",
			}, nil)
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":                         "abc",
			"url":                          "s3://foo/bar",
			"credential_name":              "bcd",
			"comment":                      "def",
			"enable_file_events":           "false",
			"effective_enable_file_events": effective,
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "updated"
		` + enableFileEvents,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationDisablesDefaultFileEvents(t *testing.T) {
	updateFileEventsFixture(t, "enable_file_events = false", "true", func(req catalog.UpdateExternalLocation) bool {
		return !req.EnableFileEvents && slices.Contains(req.ForceSendFields, "EnableFileEvents")
	})
}

func TestUpdateExternalLocationKeepsDefaultFileEventsWhenUnset(t *testing.T) {
	updateFileEventsFixture(t, "", "true", func(req catalog.UpdateExternalLocation) bool {
		return !slices.Contains(req.ForceSendFields, "EnableFileEvents")
	})
}

func TestUpdateExternalLocationFileEventsAlreadyDisabled(t *testing.T) {
	updateFileEventsFixture(t, "enable_file_events = false", "false", func(req catalog.UpdateExternalLocation) bool {
		return !slices.Contains(req.ForceSendFields, "EnableFileEvents")
	})
}
