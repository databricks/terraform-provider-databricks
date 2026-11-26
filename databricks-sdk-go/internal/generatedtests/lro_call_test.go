// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package generated_tests

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/experimental/api"
	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/lrotesting"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestLRO_CreateTestResource_Wait(t *testing.T) {
	tests := []struct {
		name       string
		fixtures   qa.HTTPFixtures
		wantResult *lrotesting.TestResource
		wantErr    bool
	}{

		{
			name: "Success",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.0/lro-testing/operations/operations/test-resource-create-12345?",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 75\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.0/lro-testing/operations/operations/test-resource-create-12345?",
				Response: lrotesting.Operation{
					Done:     true,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 100\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
					Response: json.RawMessage("{\n\t\t\t\t\t\"id\":          \"test-resource-123\",\n\t\t\t\t\t\"name\":        \"test-resource\"\n\t\t\t\t}"),
				},
			}},
			wantResult: &lrotesting.TestResource{
				Id:   "test-resource-123",
				Name: "test-resource",
			},
			wantErr: false,
		},

		{
			name: "Error",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.0/lro-testing/operations/operations/test-resource-create-12345?",
				Response: lrotesting.Operation{
					Done: true,
					Error: &lrotesting.DatabricksServiceExceptionWithDetailsProto{
						ErrorCode: lrotesting.ErrorCodeInternalError,
						Message:   "Test error message",
					},
					Name: "operations/test-resource-create-12345",
				},
			}},
			wantResult: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fixtures.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
				service := lrotesting.NewLroTesting(client)
				lroOp, err := service.CreateTestResource(ctx, lrotesting.CreateTestResourceRequest{
					Resource: lrotesting.TestResource{},
				})
				if err != nil {
					t.Fatalf("CreateTestResource failed: %v", err)
				}
				result, err := lroOp.Wait(ctx, api.WithTimeout(1*time.Minute))
				if diff := cmp.Diff(tt.wantResult, result, cmpopts.IgnoreFields(lrotesting.TestResource{}, "ForceSendFields")); diff != "" {
					t.Errorf("result mismatch (-expected +actual):\n%s", diff)
				}
				if tt.wantErr && err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !tt.wantErr && err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
			})
		})
	}
}

func TestLRO_CancelTestResource_Cancel(t *testing.T) {
	tests := []struct {
		name     string
		fixtures qa.HTTPFixtures
		wantErr  bool
	}{

		{
			name: "Success",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}, {
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/operations/operations/test-resource-create-12345/cancel",
				Response: lrotesting.Operation{
					Done: true,
					Name: "operations/test-resource-create-12345",
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fixtures.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
				service := lrotesting.NewLroTesting(client)
				lroOp, err := service.CreateTestResource(ctx, lrotesting.CreateTestResourceRequest{
					Resource: lrotesting.TestResource{},
				})
				if err != nil {
					t.Fatalf("CreateTestResource failed: %v", err)
				}
				err = lroOp.Cancel(ctx)
				if tt.wantErr && err == nil {
					t.Fatal("Cancel should have failed")
				}
				if !tt.wantErr && err != nil {
					t.Fatalf("Cancel failed: %v", err)
				}
			})
		})
	}
}
func TestLRO_CreateTestResource_Name(t *testing.T) {
	tests := []struct {
		name     string
		fixtures qa.HTTPFixtures
		wantName string
	}{

		{
			name: "Success",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}},
			wantName: "operations/test-resource-create-12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fixtures.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
				service := lrotesting.NewLroTesting(client)
				lroOp, err := service.CreateTestResource(ctx, lrotesting.CreateTestResourceRequest{
					Resource: lrotesting.TestResource{},
				})
				if err != nil {
					t.Fatalf("CreateTestResource failed: %v", err)
				}
				name := lroOp.Name()
				if name != tt.wantName {
					t.Errorf("name mismatch: expected %q, got %q", tt.wantName, name)
				}
			})
		})
	}
}
func TestLRO_CreateTestResource_Metadata(t *testing.T) {
	tests := []struct {
		name         string
		fixtures     qa.HTTPFixtures
		wantMetadata *lrotesting.TestResourceOperationMetadata
		wantErr      bool
	}{

		{
			name: "Success",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}},
			wantMetadata: &lrotesting.TestResourceOperationMetadata{
				ProgressPercent: 5,
				ResourceId:      "test-resource-123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fixtures.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
				service := lrotesting.NewLroTesting(client)
				lroOp, err := service.CreateTestResource(ctx, lrotesting.CreateTestResourceRequest{
					Resource: lrotesting.TestResource{},
				})
				if err != nil {
					t.Fatalf("CreateTestResource failed: %v", err)
				}
				metadata, err := lroOp.Metadata()
				if tt.wantErr && err == nil {
					t.Fatal("Metadata should have failed")
				}
				if !tt.wantErr && err != nil {
					t.Fatalf("Metadata failed: %v", err)
				}
				if diff := cmp.Diff(tt.wantMetadata, metadata, cmpopts.IgnoreFields(lrotesting.TestResourceOperationMetadata{}, "ForceSendFields")); diff != "" {
					t.Errorf("metadata mismatch (-expected +actual):\n%s", diff)
				}
			})
		})
	}
}
func TestLRO_CreateTestResource_Done(t *testing.T) {
	tests := []struct {
		name     string
		fixtures qa.HTTPFixtures
		wantDone bool
		wantErr  bool
	}{

		{
			name: "True",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.0/lro-testing/operations/operations/test-resource-create-12345?",
				Response: lrotesting.Operation{
					Done:     true,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 100\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
					Response: json.RawMessage("{\n\t\t\t\t\t\"id\":          \"test-resource-123\",\n\t\t\t\t\t\"name\":        \"test-resource\"\n\t\t\t\t}"),
				},
			}},
			wantDone: true,
			wantErr:  false,
		},

		{
			name: "False",
			fixtures: qa.HTTPFixtures{{
				Method:   "POST",
				Resource: "/api/2.0/lro-testing/resources",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 5\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.0/lro-testing/operations/operations/test-resource-create-12345?",
				Response: lrotesting.Operation{
					Done:     false,
					Metadata: json.RawMessage("{\n\t\t\t\t\t\"resource_id\":      \"test-resource-123\",\n\t\t\t\t\t\"progress_percent\": 75\n\t\t\t\t}"),
					Name:     "operations/test-resource-create-12345",
				},
			}},
			wantDone: false,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fixtures.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
				service := lrotesting.NewLroTesting(client)
				lroOp, err := service.CreateTestResource(ctx, lrotesting.CreateTestResourceRequest{
					Resource: lrotesting.TestResource{},
				})
				if err != nil {
					t.Fatalf("CreateTestResource failed: %v", err)
				}
				done, err := lroOp.Done()
				if tt.wantErr && err == nil {
					t.Fatal("Done should have failed")
				}
				if !tt.wantErr && err != nil {
					t.Fatalf("Done failed: %v", err)
				}
				if done != tt.wantDone {
					t.Errorf("done mismatch: expected %v, got %v", tt.wantDone, done)
				}
			})
		})
	}
}
