package model

import "reflect"

// WorkspaceConfRequestOption is an option interface that contains an apply function that applies an option to the
// builder object. The design is taken from https://github.com/uber-go/guide/blob/master/style.md#functional-options
// The reason that this is not a traditional object using the json struct tags is that, all operations are patches.
// Empty strings are used to reset the values and omitempty struct tag will not let them pass when marshaling. Future
// uses of the struct, you may not want to include all available options so this seems like a better implementations of
// using the functional options described in the uber style guide.
type WorkspaceConfRequestOption interface {
	apply(*WorkspaceConfRequestMapBuilder)
}

// WorkspaceConfRequestMapBuilder is a struct that encapsulates the workspace configurations for branding and other components
type WorkspaceConfRequestMapBuilder struct {
	configurationReqMap map[string]string
}

// WorkspaceConfResponse is the response from a get request provided the keys: AllWorkspaceConfKeys
type WorkspaceConfResponse struct {
	SidebarLogoActive      string `json:*"sidebarLogoActive,omitempty"`
	HomePageWelcomeMessage string `json:"homePageWelcomeMessage,omitempty"`
	SidebarLogoText        string `json:"sidebarLogoText,omitempty"`
	SidebarLogoInactive    string `json:"sidebarLogoInactive,omitempty"`
	HomePageLogo           string `json:"homePageLogo,omitempty"`
	HomePageLogoWidth      string `json:"homePageLogoWidth,omitempty"`
	ProductName            string `json:"productName,omitempty"`
	LoginLogo              string `json:"loginLogo,omitempty"`
	LoginLogoWidth         string `json:"loginLogoWidth,omitempty"`
	CustomReferences       string `json:"customReferences,omitempty"`
	EnableIpAccessLists    string `json:"enableIpAccessLists,omitempty"`
}

// Build will return a map that contains all the options added to the workspace configuration
func (c *WorkspaceConfRequestMapBuilder) Build(opts ...WorkspaceConfRequestOption) map[string]string {
	c.configurationReqMap = make(map[string]string)
	for _, o := range opts {
		o.apply(c)
	}
	return c.configurationReqMap
}

type sidebarLogoActive string
type homePageWelcomeMessage string
type sidebarLogoText string
type sidebarLogoInactive string
type homePageLogo string
type homePageLogoWidth string
type productName string
type loginLogo string
type loginLogoWidth string
type customReferences string
type enableIpAccessLists string

func (c sidebarLogoActive) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}
func (c homePageWelcomeMessage) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c sidebarLogoText) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c sidebarLogoInactive) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c homePageLogo) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c homePageLogoWidth) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c productName) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c loginLogo) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c loginLogoWidth) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c customReferences) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func (c enableIpAccessLists) apply(conf *WorkspaceConfRequestMapBuilder) {
	conf.configurationReqMap[reflect.TypeOf(c).Name()] = string(c)
}

func WithSidebarLogoActive(o string) WorkspaceConfRequestOption {
	return sidebarLogoActive(o)
}

func WithHomePageWelcomeMessage(o string) WorkspaceConfRequestOption {
	return homePageWelcomeMessage(o)
}

func WithSidebarLogoText(o string) WorkspaceConfRequestOption {
	return sidebarLogoText(o)
}

func WithSidebarLogoInactive(o string) WorkspaceConfRequestOption {
	return sidebarLogoInactive(o)
}

func WithHomePageLogo(o string) WorkspaceConfRequestOption {
	return homePageLogo(o)
}

func WithHomePageLogoWidth(o string) WorkspaceConfRequestOption {
	return homePageLogoWidth(o)
}

func WithProductName(o string) WorkspaceConfRequestOption {
	return productName(o)
}

func WithLoginLogo(o string) WorkspaceConfRequestOption {
	return loginLogo(o)
}

func WithLoginLogoWidth(o string) WorkspaceConfRequestOption {
	return loginLogoWidth(o)
}

func WithCustomReferences(o string) WorkspaceConfRequestOption {
	return customReferences(o)
}

func WithEnableIpAccessLists(o string) WorkspaceConfRequestOption {
	return enableIpAccessLists(o)
}

// AllWorkspaceConfKeys is a string that contains the string to search for all of the keys
var AllWorkspaceConfKeys = "productName,sidebarLogoText,loginLogo,loginLogoWidth,homePageWelcomeMessage,homePageLogo," +
	"homePageLogoWidth,sidebarLogoActive,sidebarLogoInactive,customReferences,enableIpAccessLists"
