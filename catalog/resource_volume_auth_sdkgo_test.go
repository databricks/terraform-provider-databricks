package catalog

import "testing"

// These tests repeat the tests of resource_volume_auth_test.go and
// resource_volume_auth_cloud_test.go with the volume read through sdk-go uc/volumes/v1.
// TestVolumeAuthSeveralSourcesWithoutAuthTypeFailConfiguration has no counterpart: it
// fails in PrepareDatabricksClient, before any resource selects a client.

func TestVolumeAuthBasicSendsBasicAuthorizationWithSdkGo(t *testing.T) {
	testVolumeAuthBasicSendsBasicAuthorization(t, true)
}

func TestVolumeAuthM2MSendsBearerTokenWithSdkGo(t *testing.T) {
	testVolumeAuthM2MSendsBearerToken(t, true)
}

func TestVolumeAuthM2MTokenEndpointFailureWithSdkGo(t *testing.T) {
	testVolumeAuthM2MTokenEndpointFailure(t, true)
}

func TestVolumeAuthM2MTokenEndpointNotFoundDropsVolumeWithSdkGo(t *testing.T) {
	testVolumeAuthM2MTokenEndpointNotFoundDropsVolume(t, true)
}

func TestVolumeAuthM2MAuthenticatesEachRetriedAttemptWithSdkGo(t *testing.T) {
	testVolumeAuthM2MAuthenticatesEachRetriedAttempt(t, true)
}

func TestVolumeAuthExplicitAuthTypeSelectsOneSourceWithSdkGo(t *testing.T) {
	testVolumeAuthExplicitAuthTypeSelectsOneSource(t, true)
}

func TestVolumeAuthDefaultChainFallsBackAfterFailedDryRunWithSdkGo(t *testing.T) {
	testVolumeAuthDefaultChainFallsBackAfterFailedDryRun(t, true)
}

func TestVolumeAuthCloudCredentialsWithSdkGo(t *testing.T) {
	testVolumeAuthCloudCredentials(t, true)
}

func TestVolumeAuthDeprecatedGoogleAuthTypesWithSdkGo(t *testing.T) {
	testVolumeAuthDeprecatedGoogleAuthTypes(t, true)
}
