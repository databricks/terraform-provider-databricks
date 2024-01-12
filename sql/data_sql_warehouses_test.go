package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestWarehousesData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses?",
				Response: map[string]interface{}{
					"warehouses": []sql.GetWarehouseResponse{
						{
							Id:   "1",
							Name: "bar",
						},
						{
							Id:   "2",
							Name: "bar",
						},
					},
				},
			},
		},
		Resource:    DataSourceWarehouses(),
		HCL:         ``,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"ids": []string{"1", "2"},
	})
}

func TestWarehousesDataContains(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses?",
				Response: map[string]interface{}{
					"warehouses": []sql.GetWarehouseResponse{
						{
							Id:   "111",
							Name: "bar",
						},
						{
							Id:   "2",
							Name: "br",
						},
					},
				},
			},
		},
		Resource:    DataSourceWarehouses(),
		HCL:         `warehouse_name_contains = "ba"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"ids": []string{"111"},
	})
}

func TestWarehousesData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceWarehouses(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
