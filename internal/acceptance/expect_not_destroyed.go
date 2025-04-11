package acceptance

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

type expectNotDestroyed struct {
	addr string
}

func ExpectNotDestroyed(addr string) expectNotDestroyed {
	return expectNotDestroyed{addr: addr}
}

func (e expectNotDestroyed) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, resource := range req.Plan.ResourceChanges {
		if resource.Address != e.addr {
			continue
		}
		actions := resource.Change.Actions
		if actions.DestroyBeforeCreate() || actions.CreateBeforeDestroy() || actions.Delete() {
			resp.Error = fmt.Errorf("resource %s is marked for destruction", e.addr)
			return
		}
	}
}
