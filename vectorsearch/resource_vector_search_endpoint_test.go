package vectorsearch

import (
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVectorSearchEndpointCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceVectorSearchEndpoint())
}

func TestVectorSearchEndpointCreate(t *testing.T) {
	ei := &vectorsearch.EndpointInfo{
		Name:           "abc",
		EndpointStatus: &vectorsearch.EndpointStatus{State: "ONLINE"},
		Id:             "1234-5678",
	}
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchEndpointsAPI().EXPECT()
			e.CreateEndpoint(mock.Anything, vectorsearch.CreateEndpoint{
				Name:         "abc",
				EndpointType: "STANDARD",
			}).Return(&vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo]{Poll: poll.Simple(*ei)}, nil)

			e.GetEndpointByEndpointName(mock.Anything, "abc").Return(ei, nil)
		},
		Resource: ResourceVectorSearchEndpoint(),
		HCL: `
		name          = "abc"
		endpoint_type = "STANDARD"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "1234-5678", d.Get("endpoint_id"))
}

func TestVectorSearchEndpointRead(t *testing.T) {
	ei := &vectorsearch.EndpointInfo{
		Name:           "abc",
		EndpointStatus: &vectorsearch.EndpointStatus{State: "ONLINE"},
		Id:             "1234-5678",
	}
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchEndpointsAPI().EXPECT()
			e.GetEndpointByEndpointName(mock.Anything, "abc").Return(ei, nil)
		},
		Resource: ResourceVectorSearchEndpoint(),
		ID:       "abc",
		HCL: `
		name          = "abc"
		endpoint_type = "STANDARD"
		`,
		Read: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "1234-5678", d.Get("endpoint_id"))
}

func TestResourcePASDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(a *mocks.MockWorkspaceClient) {
			a.GetMockVectorSearchEndpointsAPI().EXPECT().DeleteEndpointByEndpointName(mock.Anything, "abc").Return(nil)
		},
		Resource: ResourceVectorSearchEndpoint(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestVectorSearchEndpointCreateTimeoutError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchEndpointsAPI().EXPECT()
			e.CreateEndpoint(mock.Anything, vectorsearch.CreateEndpoint{
				Name:         "abc",
				EndpointType: "STANDARD",
			}).Return(&vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo]{
				Poll: func(_ time.Duration, _ func(*vectorsearch.EndpointInfo)) (*vectorsearch.EndpointInfo, error) {
					return nil, fmt.Errorf("timeout")
				},
			}, nil)
			e.DeleteEndpointByEndpointName(mock.Anything, "abc").Return(nil)
		},
		Resource: ResourceVectorSearchEndpoint(),
		HCL: `
		name          = "abc"
		endpoint_type = "STANDARD"
		`,
		Create: true,
	}.ExpectError(t, "timeout")

}
