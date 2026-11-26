// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apierr

import (
	"regexp"
)

var allOverrides = []errorOverride{
	{
		debugName:         "Clusters InvalidParameterValue=>ResourceDoesNotExist",
		pathRegex:         regexp.MustCompile(`^/api/2\.\d/clusters/get`),
		verb:              "GET",
		statusCodeMatcher: regexp.MustCompile(`^400$`),
		errorCodeMatcher:  regexp.MustCompile(`INVALID_PARAMETER_VALUE`),
		messageMatcher:    regexp.MustCompile(`Cluster .* does not exist`),
		customError:       ErrResourceDoesNotExist,
	}, {
		debugName:         "Jobs InvalidParameterValue=>ResourceDoesNotExist",
		pathRegex:         regexp.MustCompile(`^/api/2\.\d/jobs/get`),
		verb:              "GET",
		statusCodeMatcher: regexp.MustCompile(`^400$`),
		errorCodeMatcher:  regexp.MustCompile(`INVALID_PARAMETER_VALUE`),
		messageMatcher:    regexp.MustCompile(`Job .* does not exist`),
		customError:       ErrResourceDoesNotExist,
	}, {
		debugName:         "Job Runs InvalidParameterValue=>ResourceDoesNotExist",
		pathRegex:         regexp.MustCompile(`^/api/2\.\d/jobs/runs/get`),
		verb:              "GET",
		statusCodeMatcher: regexp.MustCompile(`^400$`),
		errorCodeMatcher:  regexp.MustCompile(`INVALID_PARAMETER_VALUE`),
		messageMatcher:    regexp.MustCompile(`(Run .* does not exist|Run: .* in job: .* doesn't exist)`),
		customError:       ErrResourceDoesNotExist,
	},
}

var allTransientErrors = []*regexp.Regexp{
	regexp.MustCompile(`com\.databricks\.backend\.manager\.util\.UnknownWorkerEnvironmentException`),
	regexp.MustCompile(`does not have any associated worker environments`),
	regexp.MustCompile(`There is no worker environment with id`),
	regexp.MustCompile(`Unknown worker environment`),
	regexp.MustCompile(`ClusterNotReadyException`),
	regexp.MustCompile(`worker env .* not found`),
	regexp.MustCompile(`Timed out after `),
	regexp.MustCompile(`deadline exceeded`),
}
