package model

import (
	"encoding/json"
	"testing"
)

func TestWorkspaceConfInterfaceCompliance(t *testing.T) {
	var _ WorkspaceConfRequestOption = (*sidebarLogoActive)(nil)
	var _ WorkspaceConfRequestOption = (*homePageWelcomeMessage)(nil)
	var _ WorkspaceConfRequestOption = (*sidebarLogoText)(nil)
	var _ WorkspaceConfRequestOption = (*sidebarLogoInactive)(nil)
	var _ WorkspaceConfRequestOption = (*homePageLogo)(nil)
	var _ WorkspaceConfRequestOption = (*homePageLogoWidth)(nil)
	var _ WorkspaceConfRequestOption = (*productName)(nil)
	var _ WorkspaceConfRequestOption = (*loginLogo)(nil)
	var _ WorkspaceConfRequestOption = (*loginLogoWidth)(nil)
	var _ WorkspaceConfRequestOption = (*customReferences)(nil)
}

func TestWorkspaceConfigurationModel(t *testing.T) {
	opts := []WorkspaceConfRequestOption{
		WithCustomReferences(""),
	}
	wc := WorkspaceConfRequestMapBuilder{}

	wc.Build(opts...)

	t.Log(wc.configurationReqMap)
	json, _ := json.Marshal(wc.configurationReqMap)
	t.Log(string(json))
}
