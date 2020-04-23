package databricks

import (
	"errors"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"sort"
	"strconv"
	"strings"
	"time"
)

type typeCheckerReturnString interface {
	execute(interface{}) string
	setNext(typeCheckerReturnString)
}

type stringChecker struct {
	next typeCheckerReturnString
}

func (r *stringChecker) execute(p interface{}) string {
	stringVal, ok := p.(string)
	if ok {
		if len(stringVal) > 0 {
			return stringVal
		}
		return ""
	}
	return r.next.execute(p)
}

func (r *stringChecker) setNext(next typeCheckerReturnString) {
	r.next = next
}

type intChecker struct {
	next typeCheckerReturnString
}

func (r *intChecker) execute(p interface{}) string {
	intVal, ok := p.(int)
	if ok {
		if intVal > 0 {
			return strconv.Itoa(intVal)
		}
		return ""
	}
	return r.next.execute(p)
}

func (r *intChecker) setNext(next typeCheckerReturnString) {
	r.next = next
}

type boolChecker struct {
	next typeCheckerReturnString
}

func (r *boolChecker) execute(p interface{}) string {
	boolVal, ok := p.(bool)
	if ok {
		return strconv.FormatBool(boolVal)
	}
	return r.next.execute(p)
}

func (r *boolChecker) setNext(next typeCheckerReturnString) {
	r.next = next
}

type stringSliceChecker struct {
	next typeCheckerReturnString
}

func (r *stringSliceChecker) execute(p interface{}) string {
	sliceVal, ok := p.([]string)
	if ok {
		var stringSlice []string
		for _, v := range sliceVal {
			stringSlice = append(stringSlice, fetchStringFromCheckers(v))
		}
		sort.Strings(stringSlice)
		return strings.Join(stringSlice, "")
	}
	return ""
}

func (r *stringSliceChecker) setNext(next typeCheckerReturnString) {
	r.next = next
}

func fetchStringFromCheckers(strVal interface{}) string {
	stringChecker := &stringChecker{}
	intChecker := &intChecker{}
	boolChecker := &boolChecker{}
	sliceChecker := &stringSliceChecker{}
	stringChecker.setNext(intChecker)
	intChecker.setNext(boolChecker)
	boolChecker.setNext(sliceChecker)

	return stringChecker.execute(strVal)
}

func validateClusterLogConf(value, meta interface{}) error {
	clusterLogConfMap := getMapFromOneItemSet(value)
	if clusterLogConfMap != nil &&
		len(clusterLogConfMap["dbfs_destination"].(string)) > 0 &&
		len(clusterLogConfMap["s3_destination"].(string)) > 0 {
		return errors.New("S3 destination and DBFS destination cant both be filled out " +
			"in cluster_log_conf!")
	}
	return nil
}

func validateInitScripts(value, meta interface{}) error {
	initScripts := value.(*schema.Set).List()
	for _, v := range initScripts {
		initScriptsMap := v.(map[string]interface{})
		if len(initScriptsMap["dbfs_destination"].(string)) > 0 &&
			len(initScriptsMap["s3_destination"].(string)) > 0 {
			return errors.New("S3 destination and DBFS destination cant both be filled out in the " +
				"init scripts!")
		}
		if len(initScriptsMap["s3_destination"].(string)) > 0 &&
			(len(initScriptsMap["s3_region"].(string)) <= 0 &&
				len(initScriptsMap["s3_endpoint"].(string)) <= 0) {
			return errors.New("If S3 destination is set then s3_region or s3_endpoint needs " +
				"to be set.!")
		}
	}
	return nil
}

func changeClusterIntoRunningState(clusterID string, client service.DBApiClient) error {
	//return nil
	clusterInfo, err := client.Clusters().Get(clusterID)
	if err != nil {
		return err
	}
	currentState := clusterInfo.State

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateRunning}, currentState) {
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStatePending, model.ClusterStateResizing, model.ClusterStateRestarting}, currentState) {
		err := client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateTerminating}, currentState) {
		err := client.Clusters().WaitForClusterTerminated(clusterID, 5, 180)
		if err != nil {
			return err
		}
		err = client.Clusters().Start(clusterID)
		if err != nil {
			if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
				return err
			}
		}
		err = client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateTerminated}, currentState) {
		err = client.Clusters().Start(clusterID)
		if err != nil {
			if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
				return err
			}
		}

		err = client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	return fmt.Errorf("cluster is in a non recoverable state: %s", currentState)

}
