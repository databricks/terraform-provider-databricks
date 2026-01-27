package exporter

import (
	"fmt"
	"log"
	"strings"
)

// getSourceCloud detects the source cloud from the client configuration
func (ic *importContext) getSourceCloud() string {
	if ic.Client.IsAws() {
		return "aws"
	} else if ic.Client.IsAzure() {
		return "azure"
	} else if ic.Client.IsGcp() {
		return "gcp"
	}
	return ""
}

// validTargetClouds is a map of valid target cloud values
var validTargetClouds = map[string]bool{
	"":      true,
	"aws":   true,
	"azure": true,
	"gcp":   true,
}

// isValidTargetCloud validates the target cloud parameter
func isValidTargetCloud(cloud string) bool {
	return validTargetClouds[cloud]
}

// availabilityType represents the normalized availability type
type availabilityType int

const (
	availabilityUnknown availabilityType = iota
	availabilitySpot
	availabilityOnDemand
	availabilitySpotWithFallback
)

// availabilityParsers maps cloud provider to availability string mappings
var availabilityParsers = map[string]map[string]availabilityType{
	"aws": {
		"SPOT":               availabilitySpot,
		"ON_DEMAND":          availabilityOnDemand,
		"SPOT_WITH_FALLBACK": availabilitySpotWithFallback,
	},
	"azure": {
		"SPOT_AZURE":               availabilitySpot,
		"ON_DEMAND_AZURE":          availabilityOnDemand,
		"SPOT_WITH_FALLBACK_AZURE": availabilitySpotWithFallback,
	},
	"gcp": {
		"PREEMPTIBLE_GCP":               availabilitySpot,
		"ON_DEMAND_GCP":                 availabilityOnDemand,
		"PREEMPTIBLE_WITH_FALLBACK_GCP": availabilitySpotWithFallback,
	},
}

// availabilityFormatters maps cloud provider and availability type to string
var availabilityFormatters = map[string]map[availabilityType]string{
	"aws": {
		availabilitySpot:             "SPOT",
		availabilityOnDemand:         "ON_DEMAND",
		availabilitySpotWithFallback: "SPOT_WITH_FALLBACK",
	},
	"azure": {
		availabilitySpot:             "SPOT_AZURE",
		availabilityOnDemand:         "ON_DEMAND_AZURE",
		availabilitySpotWithFallback: "SPOT_WITH_FALLBACK_AZURE",
	},
	"gcp": {
		availabilitySpot:             "PREEMPTIBLE_GCP",
		availabilityOnDemand:         "ON_DEMAND_GCP",
		availabilitySpotWithFallback: "PREEMPTIBLE_WITH_FALLBACK_GCP",
	},
}

// parseAvailability parses cloud-specific availability string to normalized type
func parseAvailability(availability string, cloud string) availabilityType {
	if parser, ok := availabilityParsers[cloud]; ok {
		if availType, ok := parser[availability]; ok {
			return availType
		}
	}
	return availabilityUnknown
}

// formatAvailability formats normalized availability type to cloud-specific string
func formatAvailability(availType availabilityType, cloud string) string {
	if formatter, ok := availabilityFormatters[cloud]; ok {
		if str, ok := formatter[availType]; ok {
			return str
		}
	}
	return ""
}

// diskVolumeType represents the normalized disk volume type
type diskVolumeType int

const (
	diskVolumeUnknown diskVolumeType = iota
	diskVolumeSSD
	diskVolumeHDD
)

// diskVolumeParsers maps cloud provider to disk type string mappings
var diskVolumeParsers = map[string]map[string]diskVolumeType{
	"aws": {
		"GENERAL_PURPOSE_SSD":      diskVolumeSSD,
		"THROUGHPUT_OPTIMIZED_HDD": diskVolumeHDD,
	},
	"azure": {
		"PREMIUM_LRS":  diskVolumeSSD,
		"STANDARD_LRS": diskVolumeHDD,
	},
}

// diskVolumeFormatters maps cloud provider and disk type to string
var diskVolumeFormatters = map[string]map[diskVolumeType]string{
	"aws": {
		diskVolumeSSD: "GENERAL_PURPOSE_SSD",
		diskVolumeHDD: "THROUGHPUT_OPTIMIZED_HDD",
	},
	"azure": {
		diskVolumeSSD: "PREMIUM_LRS",
		diskVolumeHDD: "STANDARD_LRS",
	},
}

// parseDiskVolumeType parses cloud-specific disk type string to normalized type
func parseDiskVolumeType(diskType string, cloud string) diskVolumeType {
	if parser, ok := diskVolumeParsers[cloud]; ok {
		if volType, ok := parser[diskType]; ok {
			return volType
		}
	}
	return diskVolumeUnknown
}

// formatDiskVolumeType formats normalized disk type to cloud-specific string
func formatDiskVolumeType(volType diskVolumeType, cloud string) string {
	if formatter, ok := diskVolumeFormatters[cloud]; ok {
		if str, ok := formatter[volType]; ok {
			return str
		}
	}
	return ""
}

// convertAvailability converts availability string from source to target cloud
func convertAvailability(availability, sourceCloud, targetCloud string) string {
	if sourceCloud == targetCloud {
		return availability
	}

	availType := parseAvailability(availability, sourceCloud)
	if availType == availabilityUnknown {
		return availability // return unchanged if unknown
	}

	converted := formatAvailability(availType, targetCloud)
	if converted == "" {
		return availability // fallback to original if conversion fails
	}

	return converted
}

// isCompatibleAttribute checks if an attribute can be converted between clouds
func isCompatibleAttribute(attrName string, attrValue interface{}, sourceCloud, targetCloud string, additionalData map[string]interface{}) bool {
	if sourceCloud == targetCloud {
		return true
	}

	switch attrName {
	case "availability", "first_on_demand":
		return true

	case "zone_id":
		// zone_id is only compatible between AWS and GCP when value is "auto" (case insensitive)
		if (sourceCloud == "aws" && targetCloud == "gcp") || (sourceCloud == "gcp" && targetCloud == "aws") {
			if strValue, ok := attrValue.(string); ok {
				return strings.EqualFold(strValue, "auto")
			}
		}
		return false

	case "ebs_volume_count":
		// ebs_volume_count (AWS) is compatible with local_ssd_count (GCP) only if ebs_volume_type is GENERAL_PURPOSE_SSD
		if sourceCloud == "aws" && targetCloud == "gcp" {
			if ebsType, ok := additionalData["ebs_volume_type"].(string); ok {
				return ebsType == "GENERAL_PURPOSE_SSD"
			}
		}
		return false

	case "local_ssd_count":
		// local_ssd_count (GCP) is compatible with ebs_volume_count (AWS) if we're going from GCP to AWS
		if sourceCloud == "gcp" && targetCloud == "aws" {
			return true
		}
		return false

	default:
		// All other attributes are not compatible
		return false
	}
}

// getTargetAttributeName returns the target attribute name for conversion
func getTargetAttributeName(attrName, sourceCloud, targetCloud string) string {
	if sourceCloud == targetCloud {
		return attrName
	}

	// ebs_volume_count <-> local_ssd_count
	if attrName == "ebs_volume_count" && sourceCloud == "aws" && targetCloud == "gcp" {
		return "local_ssd_count"
	}
	if attrName == "local_ssd_count" && sourceCloud == "gcp" && targetCloud == "aws" {
		return "ebs_volume_count"
	}

	// All other compatible attributes keep the same name
	return attrName
}

// convertCloudAttributesBlock converts a cloud attributes block from source to target cloud
// Returns nil if no conversion is needed (same cloud)
func convertCloudAttributesBlock(sourceAttrs map[string]interface{}, sourceCloud, targetCloud string) map[string]interface{} {
	if sourceCloud == targetCloud || sourceCloud == "" || targetCloud == "" {
		return nil
	}

	targetAttrs := make(map[string]interface{})

	for attrName, attrValue := range sourceAttrs {
		// Check if this attribute is compatible
		if isCompatibleAttribute(attrName, attrValue, sourceCloud, targetCloud, sourceAttrs) {
			targetAttrName := getTargetAttributeName(attrName, sourceCloud, targetCloud)

			// Convert the value if needed
			if attrName == "availability" {
				if strValue, ok := attrValue.(string); ok {
					targetAttrs[targetAttrName] = convertAvailability(strValue, sourceCloud, targetCloud)
				}
			} else {
				targetAttrs[targetAttrName] = attrValue
			}
		}
	}

	if len(targetAttrs) == 0 {
		return nil
	}

	return targetAttrs
}

// shouldConvertCloudAttributes checks if the resource type should have cloud attributes converted
func shouldConvertCloudAttributes(resourceType string) bool {
	switch resourceType {
	case "databricks_cluster", "databricks_job", "databricks_pipeline", "databricks_instance_pool":
		return true
	default:
		return false
	}
}

// convertResourceDataCloudAttributes converts cloud-specific attributes using unified wrapper
// This works for both SDKv2 and Plugin Framework resources
func (ic *importContext) convertResourceDataCloudAttributes(wrapper ResourceDataWrapper, resourceType string) {
	if ic.targetCloud == "" || !shouldConvertCloudAttributes(resourceType) {
		return
	}

	sourceCloud := ic.getSourceCloud()
	if sourceCloud == ic.targetCloud {
		return
	}

	log.Printf("[DEBUG] Converting cloud attributes for %s from %s to %s", resourceType, sourceCloud, ic.targetCloud)

	// Convert top-level cloud attributes (for clusters and instance pools)
	converted := ic.convertTopLevelCloudAttributes(wrapper, sourceCloud, ic.targetCloud)

	// Convert node types if mappings are loaded
	if ic.nodeTypeMappings != nil {
		nodeTypeConverted := ic.convertTopLevelNodeTypes(wrapper, sourceCloud, ic.targetCloud)
		converted = converted || nodeTypeConverted
	}

	// Convert nested cloud attributes (for jobs with new_cluster, job_clusters, etc.)
	if resourceType == "databricks_job" {
		nestedConverted := ic.convertJobCloudAttributes(wrapper, sourceCloud, ic.targetCloud)
		converted = converted || nestedConverted
	}

	// Convert pipeline cloud attributes
	if resourceType == "databricks_pipeline" {
		nestedConverted := ic.convertPipelineCloudAttributes(wrapper, sourceCloud, ic.targetCloud)
		converted = converted || nestedConverted
	}

	// Convert disk_spec for instance pools
	if resourceType == "databricks_instance_pool" {
		diskSpecConverted := ic.convertInstancePoolDiskSpec(wrapper, sourceCloud, ic.targetCloud)
		converted = converted || diskSpecConverted
	}

	// Warn if nothing was converted
	if !converted {
		log.Printf("[WARN] No cloud attributes found or converted for %s (resource %s)", resourceType, wrapper.Id())
	}
}

// convertTopLevelCloudAttributes converts cloud attributes at the resource root level
// Returns true if any conversion was performed
func (ic *importContext) convertTopLevelCloudAttributes(wrapper ResourceDataWrapper, sourceCloud, targetCloud string) bool {
	// Determine source attribute name
	sourceAttrName := getCloudAttributeName(sourceCloud)
	if sourceAttrName == "" {
		return false
	}

	// Get source attributes
	sourceAttrs, hasSource := wrapper.GetOk(sourceAttrName)
	if !hasSource || sourceAttrs == nil {
		return false
	}

	// Convert to map
	var sourceAttrsMap map[string]interface{}
	if attrsList, ok := sourceAttrs.([]interface{}); ok && len(attrsList) > 0 {
		if attrsMap, ok := attrsList[0].(map[string]interface{}); ok {
			sourceAttrsMap = attrsMap
		}
	}

	if sourceAttrsMap == nil {
		return false
	}

	// Convert the attributes
	targetAttrs := convertCloudAttributesBlock(sourceAttrsMap, sourceCloud, targetCloud)

	// Determine target attribute name
	targetAttrName := getCloudAttributeName(targetCloud)
	if targetAttrName == "" {
		return false
	}

	// ALWAYS clear source attributes, even if conversion yields nothing
	wrapper.Set(sourceAttrName, nil)
	log.Printf("[DEBUG] Cleared %s from resource %s", sourceAttrName, wrapper.Id())

	if len(targetAttrs) == 0 {
		log.Printf("[WARN] No compatible attributes found for conversion from %s to %s in resource %s",
			sourceCloud, targetCloud, wrapper.Id())
		return true // We did clear the source, so return true
	}

	// Set target attributes
	err := wrapper.Set(targetAttrName, []interface{}{targetAttrs})
	if err != nil {
		log.Printf("[ERROR] Failed to set %s: %v", targetAttrName, err)
		return true // Still cleared source
	}

	log.Printf("[DEBUG] Converted %s to %s with %d attributes for resource %s",
		sourceAttrName, targetAttrName, len(targetAttrs), wrapper.Id())
	return true
}

// getCloudAttributeName returns the attribute name for a given cloud
func getCloudAttributeName(cloud string) string {
	if cloud == "" {
		return ""
	}
	return cloud + "_attributes"
}

// convertJobCloudAttributes converts cloud attributes in job resource (new_cluster, job_clusters, etc.)
// Returns true if any conversion was performed
func (ic *importContext) convertJobCloudAttributes(wrapper ResourceDataWrapper, sourceCloud, targetCloud string) bool {
	converted := false

	// Handle task-level new_cluster
	if tasks, ok := wrapper.GetOk("task"); ok {
		if tasksList, ok := tasks.([]interface{}); ok {
			for i, task := range tasksList {
				if taskMap, ok := task.(map[string]interface{}); ok {
					if ic.convertClusterConfigInMap(taskMap, "new_cluster", sourceCloud, targetCloud, fmt.Sprintf("task.%d.new_cluster", i)) {
						converted = true
					}
				}
			}
			// Write the modified tasks list back to the wrapper
			if converted {
				wrapper.Set("task", tasksList)
			}
		}
	}

	// Handle job_cluster
	if jobClusters, ok := wrapper.GetOk("job_cluster"); ok {
		if clustersList, ok := jobClusters.([]interface{}); ok {
			jobConverted := false
			for i, cluster := range clustersList {
				if clusterMap, ok := cluster.(map[string]interface{}); ok {
					if ic.convertClusterConfigInMap(clusterMap, "new_cluster", sourceCloud, targetCloud, fmt.Sprintf("job_cluster.%d.new_cluster", i)) {
						jobConverted = true
						converted = true
					}
				}
			}
			// Write the modified job_cluster list back to the wrapper
			if jobConverted {
				wrapper.Set("job_cluster", clustersList)
			}
		}
	}

	return converted
}

// convertPipelineCloudAttributes converts cloud attributes in pipeline resource
// Returns true if any conversion was performed
func (ic *importContext) convertPipelineCloudAttributes(wrapper ResourceDataWrapper, sourceCloud, targetCloud string) bool {
	converted := false

	if clusters, ok := wrapper.GetOk("cluster"); ok {
		if clustersList, ok := clusters.([]interface{}); ok {
			for i, cluster := range clustersList {
				if clusterMap, ok := cluster.(map[string]interface{}); ok {
					if ic.convertClusterConfigInMap(clusterMap, "", sourceCloud, targetCloud, fmt.Sprintf("cluster.%d", i)) {
						converted = true
					}
				}
			}
			// Write the modified cluster list back to the wrapper
			if converted {
				wrapper.Set("cluster", clustersList)
			}
		}
	}

	return converted
}

// convertClusterConfigInMap converts cloud attributes within a cluster configuration map
// Returns true if conversion was performed
func (ic *importContext) convertClusterConfigInMap(m map[string]interface{}, key string, sourceCloud, targetCloud, debugPath string) bool {
	var clusterConfig map[string]interface{}

	if key == "" {
		clusterConfig = m
	} else {
		if config, ok := m[key]; ok {
			if configList, ok := config.([]interface{}); ok && len(configList) > 0 {
				if configMap, ok := configList[0].(map[string]interface{}); ok {
					clusterConfig = configMap
				}
			}
		}
	}

	if clusterConfig == nil {
		return false
	}

	converted := false

	// Convert cloud attributes
	// Determine source attribute name
	sourceAttrName := getCloudAttributeName(sourceCloud)
	if sourceAttrName != "" {
		// Get source attributes
		sourceAttrs, ok := clusterConfig[sourceAttrName]
		if ok {
			var sourceAttrsMap map[string]interface{}
			if attrsList, ok := sourceAttrs.([]interface{}); ok && len(attrsList) > 0 {
				if attrsMap, ok := attrsList[0].(map[string]interface{}); ok {
					sourceAttrsMap = attrsMap
				}
			}

			if sourceAttrsMap != nil {
				// Convert the attributes
				targetAttrs := convertCloudAttributesBlock(sourceAttrsMap, sourceCloud, targetCloud)

				// Determine target attribute name
				targetAttrName := getCloudAttributeName(targetCloud)
				if targetAttrName != "" {
					// ALWAYS clear source attributes, even if conversion yields nothing
					delete(clusterConfig, sourceAttrName)
					log.Printf("[DEBUG] Cleared %s at %s", sourceAttrName, debugPath)
					converted = true

					if len(targetAttrs) == 0 {
						log.Printf("[WARN] No compatible attributes found for conversion at %s", debugPath)
					} else {
						// Set target attributes
						clusterConfig[targetAttrName] = []interface{}{targetAttrs}
						log.Printf("[DEBUG] Converted cloud attributes at %s to %s with %d attributes",
							debugPath, targetAttrName, len(targetAttrs))
					}
				}
			}
		}
	}

	// Convert node types if mappings are loaded
	if ic.nodeTypeMappings != nil {
		nodeTypeConverted := ic.convertNodeTypesInClusterConfig(clusterConfig, sourceCloud, targetCloud, debugPath)
		converted = converted || nodeTypeConverted
	}

	return converted
}

// convertTopLevelNodeTypes converts node_type_id and driver_node_type_id at the resource root level
// Returns true if any conversion was performed
func (ic *importContext) convertTopLevelNodeTypes(wrapper ResourceDataWrapper, sourceCloud, targetCloud string) bool {
	converted := false

	// Convert node_type_id
	if nodeTypeID, ok := wrapper.GetOk("node_type_id"); ok {
		if nodeType, ok := nodeTypeID.(string); ok && nodeType != "" {
			newNodeType := convertNodeType(nodeType, sourceCloud, targetCloud, ic.nodeTypeMappings)
			if newNodeType != nodeType {
				wrapper.Set("node_type_id", newNodeType)
				converted = true
			}
		}
	}

	// Convert driver_node_type_id
	if driverNodeTypeID, ok := wrapper.GetOk("driver_node_type_id"); ok {
		if driverNodeType, ok := driverNodeTypeID.(string); ok && driverNodeType != "" {
			newDriverNodeType := convertNodeType(driverNodeType, sourceCloud, targetCloud, ic.nodeTypeMappings)
			if newDriverNodeType != driverNodeType {
				wrapper.Set("driver_node_type_id", newDriverNodeType)
				converted = true
			}
		}
	}

	return converted
}

// convertNodeTypesInClusterConfig converts node types within a cluster configuration map
// Returns true if conversion was performed
func (ic *importContext) convertNodeTypesInClusterConfig(clusterConfig map[string]interface{}, sourceCloud, targetCloud, debugPath string) bool {
	converted := false

	// Convert node_type_id
	if nodeTypeID, ok := clusterConfig["node_type_id"]; ok {
		if nodeType, ok := nodeTypeID.(string); ok && nodeType != "" {
			newNodeType := convertNodeType(nodeType, sourceCloud, targetCloud, ic.nodeTypeMappings)
			if newNodeType != nodeType {
				clusterConfig["node_type_id"] = newNodeType
				log.Printf("[DEBUG] Converted node_type_id at %s: %s -> %s", debugPath, nodeType, newNodeType)
				converted = true
			}
		}
	}

	// Convert driver_node_type_id
	if driverNodeTypeID, ok := clusterConfig["driver_node_type_id"]; ok {
		if driverNodeType, ok := driverNodeTypeID.(string); ok && driverNodeType != "" {
			newDriverNodeType := convertNodeType(driverNodeType, sourceCloud, targetCloud, ic.nodeTypeMappings)
			if newDriverNodeType != driverNodeType {
				clusterConfig["driver_node_type_id"] = newDriverNodeType
				log.Printf("[DEBUG] Converted driver_node_type_id at %s: %s -> %s", debugPath, driverNodeType, newDriverNodeType)
				converted = true
			}
		}
	}

	return converted
}

// convertInstancePoolDiskSpec converts disk_spec for instance pools between clouds
// AWS: disk_spec.disk_type.ebs_volume_type (GENERAL_PURPOSE_SSD, THROUGHPUT_OPTIMIZED_HDD)
// Azure: disk_spec.disk_type.azure_disk_volume_type (PREMIUM_LRS, STANDARD_LRS)
// GCP: disk_spec.disk_type not supported, only disk_spec.disk_count
func (ic *importContext) convertInstancePoolDiskSpec(wrapper ResourceDataWrapper, sourceCloud, targetCloud string) bool {
	if sourceCloud == targetCloud {
		return false
	}

	// Get disk_spec
	diskSpec, hasDiskSpec := wrapper.GetOk("disk_spec")
	if !hasDiskSpec || diskSpec == nil {
		return false
	}

	// Convert to slice and map
	var diskSpecMap map[string]interface{}
	if diskSpecList, ok := diskSpec.([]interface{}); ok && len(diskSpecList) > 0 {
		if specMap, ok := diskSpecList[0].(map[string]interface{}); ok {
			diskSpecMap = specMap
		}
	}

	if diskSpecMap == nil {
		return false
	}

	converted := false

	// Get disk_type if present
	diskType, hasDiskType := diskSpecMap["disk_type"]
	var diskTypeMap map[string]interface{}
	if hasDiskType && diskType != nil {
		if diskTypeList, ok := diskType.([]interface{}); ok && len(diskTypeList) > 0 {
			if typeMap, ok := diskTypeList[0].(map[string]interface{}); ok {
				diskTypeMap = typeMap
			}
		}
	}

	// Handle conversion based on source and target clouds
	if diskTypeMap != nil {
		// Conversion from AWS to Azure or vice versa
		if (sourceCloud == "aws" && targetCloud == "azure") || (sourceCloud == "azure" && targetCloud == "aws") {
			var sourceVolumeType string
			var sourceFieldName string

			if sourceCloud == "aws" {
				if ebsType, ok := diskTypeMap["ebs_volume_type"].(string); ok {
					sourceVolumeType = ebsType
					sourceFieldName = "ebs_volume_type"
				}
			} else if sourceCloud == "azure" {
				if azureType, ok := diskTypeMap["azure_disk_volume_type"].(string); ok {
					sourceVolumeType = azureType
					sourceFieldName = "azure_disk_volume_type"
				}
			}

			if sourceVolumeType != "" {
				// Parse and convert
				volType := parseDiskVolumeType(sourceVolumeType, sourceCloud)
				if volType != diskVolumeUnknown {
					targetVolumeType := formatDiskVolumeType(volType, targetCloud)
					if targetVolumeType != "" {
						// Clear source field
						delete(diskTypeMap, sourceFieldName)

						// Set target field
						targetFieldName := "ebs_volume_type"
						if targetCloud == "azure" {
							targetFieldName = "azure_disk_volume_type"
						}
						diskTypeMap[targetFieldName] = targetVolumeType

						log.Printf("[DEBUG] Converted disk_spec.disk_type for instance pool %s: %s=%s -> %s=%s",
							wrapper.Id(), sourceFieldName, sourceVolumeType, targetFieldName, targetVolumeType)
						converted = true
					}
				}
			}
		}

		// Conversion to GCP: remove disk_type entirely (GCP doesn't support it)
		if targetCloud == "gcp" {
			delete(diskSpecMap, "disk_type")
			log.Printf("[DEBUG] Removed disk_spec.disk_type for instance pool %s (GCP doesn't support disk_type)", wrapper.Id())
			converted = true
		}
	}

	// Conversion from GCP to AWS/Azure: Can't add disk_type (doesn't exist in source)
	// Just keep disk_count if present

	// Update the disk_spec if changes were made
	if converted {
		wrapper.Set("disk_spec", []interface{}{diskSpecMap})
	}

	return converted
}

// convertClusterPolicyDefinition converts cloud-specific attributes in a cluster policy definition
// Returns true if any conversions were performed
func (ic *importContext) convertClusterPolicyDefinition(definition map[string]map[string]any, sourceCloud, targetCloud string) bool {
	if sourceCloud == targetCloud || sourceCloud == "" || targetCloud == "" {
		return false
	}

	converted := false
	keysToRemove := []string{}
	keysToAdd := map[string]map[string]any{}

	for key, policyAttrs := range definition {
		policyType, hasType := policyAttrs["type"].(string)
		if !hasType {
			continue
		}

		// Skip regex types - cannot be reliably converted
		if policyType == "regex" {
			log.Printf("[WARN] Skipping regex policy for key '%s' - regex patterns cannot be automatically converted", key)
			continue
		}

		// Skip range, forbidden - these are typically numeric/boolean and cloud-agnostic
		// Note: unlimited is NOT skipped because it can have defaultValue that needs conversion
		if policyType == "range" || policyType == "forbidden" {
			continue
		}

		// Handle cloud-specific attributes (e.g., "aws_attributes.availability")
		if strings.Contains(key, "_attributes.") {
			parts := strings.SplitN(key, ".", 2)
			if len(parts) != 2 {
				continue
			}

			cloudPrefix := parts[0] // e.g., "aws_attributes"
			attrName := parts[1]    // e.g., "availability"
			cloudFromPrefix := strings.TrimSuffix(cloudPrefix, "_attributes")

			// Only process if this is the source cloud
			if cloudFromPrefix != sourceCloud {
				continue
			}

			// Check if this attribute is compatible
			valueForCheck := policyAttrs["value"]
			if valueForCheck == nil {
				valueForCheck = policyAttrs["defaultValue"]
			}
			if !isCompatibleAttribute(attrName, valueForCheck, sourceCloud, targetCloud, map[string]interface{}{}) {
				log.Printf("[DEBUG] Attribute '%s' is not compatible between %s and %s, removing from policy", key, sourceCloud, targetCloud)
				keysToRemove = append(keysToRemove, key)
				converted = true
				continue
			}

			// Convert based on policy type
			newPolicyAttrs := make(map[string]any)
			for k, v := range policyAttrs {
				newPolicyAttrs[k] = v
			}

			attributeConverted := false

			switch policyType {
			case "fixed", "unlimited":
				// Convert the value field (for fixed type)
				if value, ok := policyAttrs["value"].(string); ok {
					if attrName == "availability" {
						newValue := convertAvailability(value, sourceCloud, targetCloud)
						if newValue != value {
							newPolicyAttrs["value"] = newValue
							attributeConverted = true
						}
					}
				}
				// Also check defaultValue (for both fixed and unlimited types)
				if defaultValue, ok := policyAttrs["defaultValue"].(string); ok {
					if attrName == "availability" {
						newValue := convertAvailability(defaultValue, sourceCloud, targetCloud)
						if newValue != defaultValue {
							newPolicyAttrs["defaultValue"] = newValue
							attributeConverted = true
						}
					}
				}

			case "allowlist", "blocklist":
				// Convert the values array
				if values, ok := policyAttrs["values"].([]interface{}); ok {
					newValues := make([]interface{}, 0, len(values))
					valuesChanged := false
					for _, val := range values {
						if strVal, ok := val.(string); ok {
							if attrName == "availability" {
								newVal := convertAvailability(strVal, sourceCloud, targetCloud)
								newValues = append(newValues, newVal)
								if newVal != strVal {
									valuesChanged = true
								}
							} else {
								newValues = append(newValues, val)
							}
						} else {
							newValues = append(newValues, val)
						}
					}
					if valuesChanged {
						newPolicyAttrs["values"] = newValues
						attributeConverted = true
					}
				}
			}

			if attributeConverted {
				// Create new key with target cloud prefix
				targetCloudPrefix := targetCloud + "_attributes"
				newKey := targetCloudPrefix + "." + attrName
				keysToAdd[newKey] = newPolicyAttrs
				keysToRemove = append(keysToRemove, key)
				converted = true
				log.Printf("[DEBUG] Converted policy attribute: %s -> %s", key, newKey)
			}
		}

		// Handle node type attributes
		if key == "node_type_id" || key == "driver_node_type_id" {
			if ic.nodeTypeMappings == nil {
				continue
			}

			newPolicyAttrs := make(map[string]any)
			for k, v := range policyAttrs {
				newPolicyAttrs[k] = v
			}

			attributeConverted := false

			switch policyType {
			case "fixed", "unlimited":
				// Convert the value field (for fixed type)
				if value, ok := policyAttrs["value"].(string); ok {
					newValue := convertNodeType(value, sourceCloud, targetCloud, ic.nodeTypeMappings)
					if newValue != value {
						newPolicyAttrs["value"] = newValue
						attributeConverted = true
					}
				}
				// Also check defaultValue (for both fixed and unlimited types)
				if defaultValue, ok := policyAttrs["defaultValue"].(string); ok {
					newValue := convertNodeType(defaultValue, sourceCloud, targetCloud, ic.nodeTypeMappings)
					if newValue != defaultValue {
						newPolicyAttrs["defaultValue"] = newValue
						attributeConverted = true
					}
				}

			case "allowlist":
				// Convert the values array
				if values, ok := policyAttrs["values"].([]interface{}); ok {
					newValues := make([]interface{}, 0, len(values))
					valuesChanged := false
					for _, val := range values {
						if strVal, ok := val.(string); ok {
							newVal := convertNodeType(strVal, sourceCloud, targetCloud, ic.nodeTypeMappings)
							newValues = append(newValues, newVal)
							if newVal != strVal {
								valuesChanged = true
							}
						} else {
							newValues = append(newValues, val)
						}
					}
					if valuesChanged {
						newPolicyAttrs["values"] = newValues
						attributeConverted = true
					}
				}
			}

			if attributeConverted {
				keysToAdd[key] = newPolicyAttrs
				keysToRemove = append(keysToRemove, key)
				converted = true
				log.Printf("[DEBUG] Converted node type in policy attribute: %s", key)
			}
		}
	}

	// Remove old keys
	for _, key := range keysToRemove {
		delete(definition, key)
	}

	// Add new keys
	for key, attrs := range keysToAdd {
		definition[key] = attrs
	}

	return converted
}
