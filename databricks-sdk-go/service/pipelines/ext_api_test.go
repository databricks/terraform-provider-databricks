package pipelines

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
)

func TestListPipelinesAll(t *testing.T) {
	ctx := context.Background()

	t.Run("list pipelines with pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?",
				Response: ListPipelinesResponse{
					NextPageToken: "page-token-1",
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-1",
							PipelineId:      "00000000-0000-0000-0000-000000000001",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-2",
							PipelineId:      "00000000-0000-0000-0000-000000000002",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?page_token=page-token-1",
				Response: ListPipelinesResponse{
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-3",
							PipelineId:      "00000000-0000-0000-0000-000000000003",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-4",
							PipelineId:      "00000000-0000-0000-0000-000000000004",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockPipelinesImpl := &pipelinesImpl{
			client: client,
		}
		api := &PipelinesAPI{pipelinesImpl: *mockPipelinesImpl}

		allPipelines, err := api.ListPipelinesAll(ctx, ListPipelinesRequest{})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(allPipelines) != 4 {
			t.Errorf("expected 4 pipelines, got %d", len(allPipelines))
		}
		if allPipelines[0].Name != "pipeline-1" {
			t.Errorf("expected pipeline-1, got %s", allPipelines[0].Name)
		}
		if allPipelines[0].PipelineId != "00000000-0000-0000-0000-000000000001" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000001, got %s", allPipelines[0].PipelineId)
		}
		if allPipelines[1].Name != "pipeline-2" {
			t.Errorf("expected pipeline-2, got %s", allPipelines[1].Name)
		}
		if allPipelines[1].PipelineId != "00000000-0000-0000-0000-000000000002" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000002, got %s", allPipelines[1].PipelineId)
		}
		if allPipelines[2].Name != "pipeline-3" {
			t.Errorf("expected pipeline-3, got %s", allPipelines[2].Name)
		}
		if allPipelines[2].PipelineId != "00000000-0000-0000-0000-000000000003" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000003, got %s", allPipelines[2].PipelineId)
		}
		if allPipelines[3].Name != "pipeline-4" {
			t.Errorf("expected pipeline-4, got %s", allPipelines[3].Name)
		}
		if allPipelines[3].PipelineId != "00000000-0000-0000-0000-000000000004" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000004, got %s", allPipelines[3].PipelineId)
		}
		if allPipelines[0].State != PipelineStateIdle {
			t.Errorf("expected PipelineStateIdle, got %v", allPipelines[0].State)
		}
	})
}

func TestPipelineStateInfoNameToPipelineIdMap(t *testing.T) {
	ctx := context.Background()

	t.Run("name to pipeline id map with pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?",
				Response: ListPipelinesResponse{
					NextPageToken: "page-token-1",
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-1",
							PipelineId:      "00000000-0000-0000-0000-000000000001",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-2",
							PipelineId:      "00000000-0000-0000-0000-000000000002",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?page_token=page-token-1",
				Response: ListPipelinesResponse{
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-3",
							PipelineId:      "00000000-0000-0000-0000-000000000003",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-4",
							PipelineId:      "00000000-0000-0000-0000-000000000004",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockPipelinesImpl := &pipelinesImpl{
			client: client,
		}
		api := &PipelinesAPI{pipelinesImpl: *mockPipelinesImpl}

		nameToIdMap, err := api.PipelineStateInfoNameToPipelineIdMap(ctx, ListPipelinesRequest{})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(nameToIdMap) != 4 {
			t.Errorf("expected map length 4, got %d", len(nameToIdMap))
		}
		if nameToIdMap["pipeline-1"] != "00000000-0000-0000-0000-000000000001" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000001 for pipeline-1, got %s", nameToIdMap["pipeline-1"])
		}
		if nameToIdMap["pipeline-2"] != "00000000-0000-0000-0000-000000000002" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000002 for pipeline-2, got %s", nameToIdMap["pipeline-2"])
		}
		if nameToIdMap["pipeline-3"] != "00000000-0000-0000-0000-000000000003" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000003 for pipeline-3, got %s", nameToIdMap["pipeline-3"])
		}
		if nameToIdMap["pipeline-4"] != "00000000-0000-0000-0000-000000000004" {
			t.Errorf("expected 00000000-0000-0000-0000-000000000004 for pipeline-4, got %s", nameToIdMap["pipeline-4"])
		}
	})

	t.Run("name to pipeline id map detects duplicates", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?",
				Response: ListPipelinesResponse{
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "duplicate-pipeline-name",
							PipelineId:      "00000000-0000-0000-0000-000000000001",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "duplicate-pipeline-name",
							PipelineId:      "00000000-0000-0000-0000-000000000002",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockPipelinesImpl := &pipelinesImpl{
			client: client,
		}
		api := &PipelinesAPI{pipelinesImpl: *mockPipelinesImpl}

		nameToIdMap, err := api.PipelineStateInfoNameToPipelineIdMap(ctx, ListPipelinesRequest{})

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if nameToIdMap != nil {
			t.Errorf("expected nil map, got %v", nameToIdMap)
		}
		if !strings.Contains(err.Error(), "duplicate .Name: duplicate-pipeline-name") {
			t.Errorf("expected error to contain 'duplicate .Name: duplicate-pipeline-name', got: %v", err)
		}
	})
}
