package settings

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testComplianceSecurityProfileSetting = AllSettingsResources()["compliance_security_profile_workspace"]

func TestQueryCreateComplianceSecurityProfileSettingWithNoneStandard(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockComplianceSecurityProfileAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateComplianceSecurityProfileSettingRequest{
				AllowMissing: true,
				FieldMask:    complianceSecurityProfileFieldMask,
				Setting: settings.ComplianceSecurityProfileSetting{
					Etag: "",
					ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
						IsEnabled:           true,
						ComplianceStandards: []settings.ComplianceStandard{"NONE"},
					},
					SettingName: "default",
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag1",
					},
				}},
			})
			e.Update(mock.Anything, settings.UpdateComplianceSecurityProfileSettingRequest{
				AllowMissing: true,
				FieldMask:    complianceSecurityProfileFieldMask,
				Setting: settings.ComplianceSecurityProfileSetting{
					Etag: "etag1",
					ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
						IsEnabled:           true,
						ComplianceStandards: []settings.ComplianceStandard{"NONE"},
					},
					SettingName: "default",
				},
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag2",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"NONE"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetComplianceSecurityProfileSettingRequest{
				Etag: "etag2",
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag2",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"NONE"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testComplianceSecurityProfileSetting,
		Create:   true,
		HCL: `
			compliance_security_profile_workspace {
				is_enabled = true
				compliance_standards = ["NONE"]
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("compliance_security_profile_workspace.0.is_enabled"))
}
