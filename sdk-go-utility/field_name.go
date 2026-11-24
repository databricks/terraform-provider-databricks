package sdkgoutility

import (
	"reflect"
	"strings"
	"unicode"
)

// sdkModulePath is the import path prefix that all modular SDK packages share.
const sdkModulePath = "github.com/databricks/sdk-go/"

// FieldName returns the name that the Databricks API uses for field, a field
// of the modular SDK struct structType. It is the key of the field in the JSON
// that the MarshalJSON method of the struct writes.
//
// Modular SDK structs carry no json tag, so the name is the first of:
//
//  1. the entry for the field in fieldNameOverrides, which lists the fields
//     whose API name the next two steps get wrong;
//  2. the fieldmask tag of the field, which the SDK generator sets to the API
//     name on the messages that field masks address;
//  3. the snake_case form of the Go name, as snakeCase derives it.
//
// structType is the struct that declares field, or a pointer to it. It only
// selects the overrides: the same Go name can have different API names in
// different structs, and a reflect.StructField doesn't say which struct it
// comes from. With a nil structType, or one that is not a modular SDK struct,
// steps 2 and 3 still apply. FieldName returns the empty string for an
// unexported field, and it never panics.
//
// FieldName deliberately doesn't:
//
//   - read json or tf tags. Classic SDK structs and the provider's own structs
//     stay with the helpers in package common.
//   - tell whether the API has the field in its JSON at all. The interface
//     field of a oneof group, a request or response body stream and a field
//     that travels in an HTTP header get their snake_case name like any other
//     field, although MarshalJSON writes no such key. A helper that walks a
//     struct recognizes those shapes itself.
//   - look into the field's type. It names one field and never builds the path
//     of a nested one.
func FieldName(structType reflect.Type, field reflect.StructField) string {
	return fieldName(fieldNameOverrides, structType, field)
}

// fieldName is FieldName with the override table as a parameter, so that tests
// can pass overrides that fieldNameOverrides doesn't hold.
func fieldName(overrides map[string]map[string]string, structType reflect.Type, field reflect.StructField) string {
	if !field.IsExported() {
		return ""
	}
	if name, ok := overrides[overrideKey(structType)][field.Name]; ok {
		return name
	}
	if name := field.Tag.Get("fieldmask"); name != "" {
		return name
	}
	return snakeCase(field.Name)
}

// overrideKey returns the key of structType in an override table: the package
// path below sdkModulePath, a dot and the type name. It returns the empty
// string, which is never a key, for a type from outside the modular SDK.
func overrideKey(structType reflect.Type) string {
	for structType != nil && structType.Kind() == reflect.Pointer {
		structType = structType.Elem()
	}
	if structType == nil {
		return ""
	}
	pkg, ok := strings.CutPrefix(structType.PkgPath(), sdkModulePath)
	if !ok {
		return ""
	}
	return pkg + "." + structType.Name()
}

// snakeCase inverts the naming of the SDK generator, which builds the Go name
// of a field by capitalizing the first letter of every underscore-separated
// word of its API name: registered_model_id becomes RegisteredModelId. So a
// capital letter starts a new word, and nothing else does:
//
//   - Every capital in a run of capitals is a word of its own. The generator
//     capitalizes only the first letter of a word (http_status is HttpStatus
//     and t_shirt_size is TShirtSize), so snakeCase doesn't guess acronyms:
//     HTTPStatus gives h_t_t_p_status, not http_status.
//   - A digit never starts a word and stays with the letters before it:
//     Ipv6Cidr gives ipv6_cidr, S3Uri gives s3_uri and Oauth2Token gives
//     oauth2_token. An API word that starts with a digit can't be told apart
//     from that (visualization_color_01 is VisualizationColor01 in Go), so
//     such fields are in fieldNameOverrides.
//   - An underscore that the Go name already has is kept, and the capital
//     after it adds no second one.
func snakeCase(goName string) string {
	var b strings.Builder
	var previous rune
	for i, r := range goName {
		if unicode.IsUpper(r) && i > 0 && previous != '_' {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
		previous = r
	}
	return b.String()
}

// fieldNameOverrides holds the API names that the fieldmask tag and snakeCase
// get wrong, by overrideKey of the struct and then by Go field name. A
// comparison of every public field of every message with the json name of its
// wire struct field, over all packages of github.com/databricks/sdk-go at
// commit 7a95446, found these three kinds:
//
//   - APIs whose JSON names aren't snake_case, such as SCIM (displayName,
//     Resources, $ref) and command execution (clusterId);
//   - API words that start with a digit (visualization_color_01);
//   - fields that the generator renamed (PipelineTaskParameters is parameters).
//
// The table is not checked against the SDK in this repository, because the
// provider doesn't depend on these packages yet. When a resource starts to use
// a package, test its messages against MarshalJSON, as the test for
// modelregistry does, and add an entry for each field that the test reports.
var fieldNameOverrides = map[string]map[string]string{
	"commandexecution/v2.CancelCommandRequest":          {"ClusterId": "clusterId", "CommandId": "commandId", "ContextId": "contextId"},
	"commandexecution/v2.CreateContextRequest":          {"ClusterId": "clusterId"},
	"commandexecution/v2.DestroyContextRequest":         {"ClusterId": "clusterId", "ContextId": "contextId"},
	"commandexecution/v2.ExecuteCommandRequest":         {"ClusterId": "clusterId", "ContextId": "contextId"},
	"commandexecution/v2.GetCommandStatusRequest":       {"ClusterId": "clusterId", "CommandId": "commandId", "ContextId": "contextId"},
	"commandexecution/v2.GetContextStatusRequest":       {"ClusterId": "clusterId", "ContextId": "contextId"},
	"commandexecution/v2.Results":                       {"FileName": "fileName", "FileNames": "fileNames", "IsJsonSchema": "isJsonSchema", "ResultType": "resultType"},
	"jobs/v2.PipelineTask":                              {"PipelineTaskParameters": "parameters"},
	"jobs/v2.ResolvedValues_PipelineTaskResolvedValues": {"PipelineTaskParameters": "parameters"},
	"marketplaces/v1.ListingSummary":                    {"ListingType": "listingType"},
	"marketplaces/v1.TokenDetail":                       {"BearerToken": "bearerToken", "ExpirationTime": "expirationTime", "ShareCredentialsVersion": "shareCredentialsVersion"},
	"modelservingquery/v1.V1ResponseChoiceElement":      {"FinishReason": "finishReason"},
	"oauth/v1.CredentialPartitionId":                    {"WorkspaceId": "workspaceId"},
	"oauth/v1.ExchangeToken":                            {"CredentialEolTime": "credentialEolTime", "OwnerId": "ownerId", "TokenType": "tokenType"},
	"oauth/v1.ExchangeTokenExchangeRequest":             {"PartitionId": "partitionId", "TokenType": "tokenType"},
	"sandbox/v1.DownloadFileRequest":                    {"Range": "byte_range"},
	"scim/v1.AccountComplexValue":                       {"Ref": "$ref"},
	"scim/v1.AccountGroup":                              {"DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.AccountName":                               {"FamilyName": "familyName", "GivenName": "givenName"},
	"scim/v1.AccountResourceMeta":                       {"ResourceType": "resourceType"},
	"scim/v1.AccountServicePrincipal":                   {"ApplicationId": "applicationId", "DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.AccountUser":                               {"DisplayName": "displayName", "ExternalId": "externalId", "UserName": "userName"},
	"scim/v1.ComplexValue":                              {"Ref": "$ref"},
	"scim/v1.CreateAccountGroupRequest":                 {"DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.CreateAccountServicePrincipalRequest":      {"ApplicationId": "applicationId", "DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.CreateAccountUserRequest":                  {"DisplayName": "displayName", "ExternalId": "externalId", "UserName": "userName"},
	"scim/v1.CreateGroupRequest":                        {"DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.CreateServicePrincipalRequest":             {"ApplicationId": "applicationId", "DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.CreateUserRequest":                         {"DisplayName": "displayName", "ExternalId": "externalId", "UserName": "userName"},
	"scim/v1.GetAccountUserRequest":                     {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.GetUserRequest":                            {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.Group":                                     {"DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.ListAccountGroupsRequest":                  {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.ListAccountGroupsResponse":                 {"ItemsPerPage": "itemsPerPage", "Resources": "Resources", "StartIndex": "startIndex", "TotalResults": "totalResults"},
	"scim/v1.ListAccountServicePrincipalsRequest":       {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.ListAccountServicePrincipalsResponse":      {"ItemsPerPage": "itemsPerPage", "Resources": "Resources", "StartIndex": "startIndex", "TotalResults": "totalResults"},
	"scim/v1.ListAccountUsersRequest":                   {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.ListAccountUsersResponse":                  {"ItemsPerPage": "itemsPerPage", "Resources": "Resources", "StartIndex": "startIndex", "TotalResults": "totalResults"},
	"scim/v1.ListGroupsRequest":                         {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.ListGroupsResponse":                        {"ItemsPerPage": "itemsPerPage", "Resources": "Resources", "StartIndex": "startIndex", "TotalResults": "totalResults"},
	"scim/v1.ListServicePrincipalResponse":              {"ItemsPerPage": "itemsPerPage", "Resources": "Resources", "StartIndex": "startIndex", "TotalResults": "totalResults"},
	"scim/v1.ListServicePrincipalsRequest":              {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.ListUsersRequest":                          {"ExcludedAttributes": "excludedAttributes", "SortBy": "sortBy", "SortOrder": "sortOrder", "StartIndex": "startIndex"},
	"scim/v1.ListUsersResponse":                         {"ItemsPerPage": "itemsPerPage", "Resources": "Resources", "StartIndex": "startIndex", "TotalResults": "totalResults"},
	"scim/v1.MeRequest":                                 {"ExcludedAttributes": "excludedAttributes"},
	"scim/v1.Name":                                      {"FamilyName": "familyName", "GivenName": "givenName"},
	"scim/v1.PatchAccountGroupRequest":                  {"Operations": "Operations"},
	"scim/v1.PatchAccountServicePrincipalRequest":       {"Operations": "Operations"},
	"scim/v1.PatchAccountUserRequest":                   {"Operations": "Operations"},
	"scim/v1.PatchGroupRequest":                         {"Operations": "Operations"},
	"scim/v1.PatchServicePrincipalRequest":              {"Operations": "Operations"},
	"scim/v1.PatchUserRequest":                          {"Operations": "Operations"},
	"scim/v1.ResourceMeta":                              {"ResourceType": "resourceType"},
	"scim/v1.ServicePrincipal":                          {"ApplicationId": "applicationId", "DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.UpdateAccountGroupRequest":                 {"DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.UpdateAccountServicePrincipalRequest":      {"ApplicationId": "applicationId", "DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.UpdateAccountUserRequest":                  {"DisplayName": "displayName", "ExternalId": "externalId", "UserName": "userName"},
	"scim/v1.UpdateGroupRequest":                        {"DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.UpdateServicePrincipalRequest":             {"ApplicationId": "applicationId", "DisplayName": "displayName", "ExternalId": "externalId"},
	"scim/v1.UpdateUserRequest":                         {"DisplayName": "displayName", "ExternalId": "externalId", "UserName": "userName"},
	"scim/v1.User":                                      {"DisplayName": "displayName", "ExternalId": "externalId", "UserName": "userName"},
	"settings/v2.DashboardTheme": {
		"VisualizationColor01": "visualization_color_01", "VisualizationColor02": "visualization_color_02", "VisualizationColor03": "visualization_color_03", "VisualizationColor04": "visualization_color_04",
		"VisualizationColor05": "visualization_color_05", "VisualizationColor06": "visualization_color_06", "VisualizationColor07": "visualization_color_07", "VisualizationColor08": "visualization_color_08",
		"VisualizationColor09": "visualization_color_09", "VisualizationColor10": "visualization_color_10", "VisualizationColor11": "visualization_color_11", "VisualizationColor12": "visualization_color_12",
		"VisualizationColor13": "visualization_color_13", "VisualizationColor14": "visualization_color_14", "VisualizationColor15": "visualization_color_15", "VisualizationColor16": "visualization_color_16",
		"VisualizationColor17": "visualization_color_17", "VisualizationColor18": "visualization_color_18", "VisualizationColor19": "visualization_color_19", "VisualizationColor20": "visualization_color_20",
	},
	"sharing/v1.RetrieveTokenResponse":                    {"BearerToken": "bearerToken", "ExpirationTime": "expirationTime", "ShareCredentialsVersion": "shareCredentialsVersion"},
	"uc/credentials/v1.ValidateCredentialResponse":        {"IsDir": "isDir"},
	"uc/credentials/v1.ValidateStorageCredentialResponse": {"IsDir": "isDir"},
}
