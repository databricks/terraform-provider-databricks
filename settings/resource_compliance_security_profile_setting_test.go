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
	res := d.Get("compliance_security_profile_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
	assert.Equal(t, "NONE", res["compliance_standards"].([]interface{})[0])

}

func TestQueryReadComplianceSecurityProfileSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockComplianceSecurityProfileAPI().EXPECT().Get(mock.Anything, settings.GetComplianceSecurityProfileSettingRequest{
				Etag: "etag1",
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag2",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"HIPAA"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testComplianceSecurityProfileSetting,
		Read:     true,
		HCL: `
			compliance_security_profile_workspace {
				is_enabled = true
				compliance_standards = ["HIPAA"]
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("compliance_security_profile_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
	assert.Equal(t, "HIPAA", res["compliance_standards"].([]interface{})[0])
}

func TestQueryUpdateComplianceSecurityProfileSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockComplianceSecurityProfileAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateComplianceSecurityProfileSettingRequest{
				AllowMissing: true,
				FieldMask:    complianceSecurityProfileFieldMask,
				Setting: settings.ComplianceSecurityProfileSetting{
					Etag: "etag1",
					ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
						IsEnabled:           true,
						ComplianceStandards: []settings.ComplianceStandard{"HIPAA", "PCI_DSS"},
					},
					SettingName: "default",
				},
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag2",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"HIPAA", "PCI_DSS"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetComplianceSecurityProfileSettingRequest{
				Etag: "etag2",
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag2",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"HIPAA", "PCI_DSS"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testComplianceSecurityProfileSetting,
		Update:   true,
		HCL: `
			compliance_security_profile_workspace {
				is_enabled = true
				compliance_standards = ["HIPAA", "PCI_DSS"]
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("compliance_security_profile_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
	assert.Equal(t, "HIPAA", res["compliance_standards"].([]interface{})[0])
	assert.Equal(t, "PCI_DSS", res["compliance_standards"].([]interface{})[1])
}

func TestQueryUpdateComplianceSecurityProfileSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockComplianceSecurityProfileAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateComplianceSecurityProfileSettingRequest{
				AllowMissing: true,
				FieldMask:    complianceSecurityProfileFieldMask,
				Setting: settings.ComplianceSecurityProfileSetting{
					Etag: "etag1",
					ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
						IsEnabled:           true,
						ComplianceStandards: []settings.ComplianceStandard{"HIPAA"},
					},
					SettingName: "default",
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_CONFLICT",
				StatusCode: 409,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag2",
					},
				}},
			})
			e.Update(mock.Anything, settings.UpdateComplianceSecurityProfileSettingRequest{
				AllowMissing: true,
				FieldMask:    complianceSecurityProfileFieldMask,
				Setting: settings.ComplianceSecurityProfileSetting{
					Etag: "etag2",
					ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
						IsEnabled:           true,
						ComplianceStandards: []settings.ComplianceStandard{"HIPAA"},
					},
					SettingName: "default",
				},
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag3",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"HIPAA"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetComplianceSecurityProfileSettingRequest{
				Etag: "etag3",
			}).Return(&settings.ComplianceSecurityProfileSetting{
				Etag: "etag3",
				ComplianceSecurityProfileWorkspace: settings.ComplianceSecurityProfile{
					IsEnabled:           true,
					ComplianceStandards: []settings.ComplianceStandard{"HIPAA"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testComplianceSecurityProfileSetting,
		Update:   true,
		HCL: `
			compliance_security_profile_workspace {
				is_enabled = true
				compliance_standards = ["HIPAA"]
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	res := d.Get("compliance_security_profile_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
	assert.Equal(t, "HIPAA", res["compliance_standards"].([]interface{})[0])
}

func TestDeleteComplianceSecurityProfileSetting(t *testing.T) {
	qa.ResourceFixture{
		Resource: testComplianceSecurityProfileSetting,
		Delete:   true,
		HCL: `
			compliance_security_profile_workspace {
				is_enabled = true
				compliance_standards = ["HIPAA", "PCI_DSS"]
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		etagAttrName: "etag1",
		"id":         defaultSettingId,
	})
}
