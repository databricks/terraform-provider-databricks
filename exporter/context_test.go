package exporter

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestNoServicesSkipsRun(t *testing.T) {
	assert.EqualError(t, (&importContext{}).Run(), "no services to import")
}

func TestMatchesName(t *testing.T) {
	assert.False(t, (&importContext{match: "x"}).MatchesName("y"))
}

func TestImportContextFindSkips(t *testing.T) {
	_, traversal := (&importContext{
		State: stateApproximation{
			resources: []resourceApproximation{
				{
					Type: "a",
					Instances: []instanceApproximation{
						{
							Attributes: map[string]any{
								"b": nil,
							},
						},
					},
				},
			},
		},
	}).Find(&resource{
		Resource:  "a",
		Attribute: "b",
		Name:      "c",
	}, "x", reference{})
	assert.Nil(t, traversal)
}

func TestImportContextHas(t *testing.T) {
	assert.True(t, (&importContext{
		State: stateApproximation{
			resources: []resourceApproximation{
				{
					Type: "a",
					Instances: []instanceApproximation{
						{
							Attributes: map[string]any{
								"b": "d",
							},
						},
					},
				},
			},
		},
	}).Has(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	}))
}

func TestEmitNaResource(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoImportable(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoSearchAvail(t *testing.T) {
	ch := make(resourceChannel)
	ic := &importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
			},
		},
		services:  "e",
		waitGroup: &sync.WaitGroup{},
		channels: map[string]resourceChannel{
			"a": ch,
		},
	}
	go func() {
		for r := range ch {
			r.ImportResource(ic)
		}
	}()
	ic.Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
	ic.waitGroup.Wait()
	close(ch)
}

func TestEmitNoSearchFails(t *testing.T) {
	ch := make(resourceChannel, 10)
	ic := &importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
				Search: func(ic *importContext, r *resource) error {
					return fmt.Errorf("just fails")
				},
			},
		},
		services:  "e",
		waitGroup: &sync.WaitGroup{},
		channels: map[string]resourceChannel{
			"a": ch,
		},
	}
	go func() {
		for r := range ch {
			r.ImportResource(ic)
		}
	}()
	ic.Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
	ic.waitGroup.Wait()
	close(ch)
}

func TestEmitNoSearchNoId(t *testing.T) {
	ch := make(resourceChannel, 10)
	ic := &importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
				Search: func(ic *importContext, r *resource) error {
					return nil
				},
			},
		},
		services:  "e",
		waitGroup: &sync.WaitGroup{},
		channels: map[string]resourceChannel{
			"a": ch,
		},
	}
	go func() {
		for r := range ch {
			r.ImportResource(ic)
		}
	}()
	ic.Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
	ic.waitGroup.Wait()
	close(ch)
}

func TestEmitNoSearchSucceedsImportFails(t *testing.T) {
	ch := make(resourceChannel, 10)
	ic := &importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
				Search: func(ic *importContext, r *resource) error {
					r.ID = "some"
					return nil
				},
				Import: func(ic *importContext, r *resource) error {
					return fmt.Errorf("fails")
				},
			},
		},
		services:  "e",
		waitGroup: &sync.WaitGroup{},
		channels: map[string]resourceChannel{
			"a": ch,
		},
	}
	go func() {
		for r := range ch {
			r.ImportResource(ic)
		}
	}()
	ic.Emit(&resource{
		Data:      &schema.ResourceData{},
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
	ic.waitGroup.Wait()
	close(ch)
}

func TestLoadingLastRun(t *testing.T) {
	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	defer os.RemoveAll(tmpDir)

	fname := tmpDir + "1.json"
	// no file yet
	s := getLastRunString(fname)
	assert.Equal(t, "", s)

	_ = os.WriteFile(fname, []byte("{"), 0755)
	s = getLastRunString(fname)
	assert.Equal(t, "", s)

	// no required field
	_ = os.WriteFile(fname, []byte("{}"), 0755)
	s = getLastRunString(fname)
	assert.Equal(t, "", s)

	_ = os.WriteFile(fname, []byte(`{"startTime": "2023-07-24T00:00:00Z"}`), 0755)
	s = getLastRunString(fname)
	assert.Equal(t, "2023-07-24T00:00:00Z", s)
}
