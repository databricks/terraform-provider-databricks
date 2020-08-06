package model

import (
	"fmt"
	"testing"
)

func TestClusterState_CanReach(t *testing.T) {
	tests := []struct {
		from ClusterState
		to   ClusterState
		want bool
	}{
		{ClusterStatePending, ClusterStatePending, true},
		{ClusterStatePending, ClusterStateRunning, true},
		{ClusterStatePending, ClusterStateRestarting, true},
		{ClusterStatePending, ClusterStateResizing, true},
		{ClusterStatePending, ClusterStateTerminating, true},
		{ClusterStatePending, ClusterStateTerminated, true},
		{ClusterStatePending, ClusterStateError, false},
		{ClusterStatePending, ClusterStateUnknown, false},

		{ClusterStateRunning, ClusterStatePending, false},
		{ClusterStateRunning, ClusterStateRunning, true},
		{ClusterStateRunning, ClusterStateRestarting, true},
		{ClusterStateRunning, ClusterStateResizing, true},
		{ClusterStateRunning, ClusterStateTerminating, true},
		{ClusterStateRunning, ClusterStateTerminated, true},
		{ClusterStateRunning, ClusterStateError, false},
		{ClusterStateRunning, ClusterStateUnknown, false},

		{ClusterStateRestarting, ClusterStatePending, false},
		{ClusterStateRestarting, ClusterStateRunning, true},
		{ClusterStateRestarting, ClusterStateRestarting, true},
		{ClusterStateRestarting, ClusterStateResizing, true},
		{ClusterStateRestarting, ClusterStateTerminating, true},
		{ClusterStateRestarting, ClusterStateTerminated, true},
		{ClusterStateRestarting, ClusterStateError, false},
		{ClusterStateRestarting, ClusterStateUnknown, false},

		{ClusterStateResizing, ClusterStatePending, false},
		{ClusterStateResizing, ClusterStateRunning, true},
		{ClusterStateResizing, ClusterStateRestarting, true},
		{ClusterStateResizing, ClusterStateResizing, true},
		{ClusterStateResizing, ClusterStateTerminating, true},
		{ClusterStateResizing, ClusterStateTerminated, true},
		{ClusterStateResizing, ClusterStateError, false},
		{ClusterStateResizing, ClusterStateUnknown, false},

		{ClusterStateTerminating, ClusterStatePending, false},
		{ClusterStateTerminating, ClusterStateRunning, false},
		{ClusterStateTerminating, ClusterStateRestarting, false},
		{ClusterStateTerminating, ClusterStateResizing, false},
		{ClusterStateTerminating, ClusterStateTerminating, true},
		{ClusterStateTerminating, ClusterStateTerminated, true},
		{ClusterStateTerminating, ClusterStateError, false},
		{ClusterStateTerminating, ClusterStateUnknown, false},

		{ClusterStateTerminated, ClusterStatePending, false},
		{ClusterStateTerminated, ClusterStateRunning, false},
		{ClusterStateTerminated, ClusterStateRestarting, false},
		{ClusterStateTerminated, ClusterStateResizing, false},
		{ClusterStateTerminated, ClusterStateTerminating, false},
		{ClusterStateTerminated, ClusterStateTerminated, true},
		{ClusterStateTerminated, ClusterStateError, false},
		{ClusterStateTerminated, ClusterStateUnknown, false},

		{ClusterStateError, ClusterStatePending, false},
		{ClusterStateError, ClusterStateRunning, false},
		{ClusterStateError, ClusterStateRestarting, false},
		{ClusterStateError, ClusterStateResizing, false},
		{ClusterStateError, ClusterStateTerminating, false},
		{ClusterStateError, ClusterStateTerminated, false},
		{ClusterStateError, ClusterStateError, true},
		{ClusterStateError, ClusterStateUnknown, false},

		{ClusterStateUnknown, ClusterStatePending, false},
		{ClusterStateUnknown, ClusterStateRunning, false},
		{ClusterStateUnknown, ClusterStateRestarting, false},
		{ClusterStateUnknown, ClusterStateResizing, false},
		{ClusterStateUnknown, ClusterStateTerminating, false},
		{ClusterStateUnknown, ClusterStateTerminated, false},
		{ClusterStateUnknown, ClusterStateError, false},
		{ClusterStateUnknown, ClusterStateUnknown, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s to %s", tt.from, tt.to), func(t *testing.T) {
			if got := tt.from.CanReach(tt.to); got != tt.want {
				t.Errorf("ClusterState.CanReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
