package exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// nodeTypeMapping represents a single mapping entry between cloud node types
type nodeTypeMapping struct {
	Azure string `json:"azure,omitempty"`
	AWS   string `json:"aws,omitempty"`
	GCP   string `json:"gcp,omitempty"`
}

// nodeTypeMappings represents the structure of the mapping file
type nodeTypeMappings struct {
	Version  string            `json:"version,omitempty"`
	Mappings []nodeTypeMapping `json:"mappings"`
}

// loadNodeTypeMappings loads and parses the node type mapping file
func loadNodeTypeMappings(filePath string) (*nodeTypeMappings, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read node type mapping file: %w", err)
	}

	var mappings nodeTypeMappings
	if err := json.Unmarshal(data, &mappings); err != nil {
		return nil, fmt.Errorf("failed to parse node type mapping file: %w", err)
	}

	return &mappings, nil
}

// getNodeTypeForCloud returns the node type for a specific cloud from a mapping
func getNodeTypeForCloud(mapping nodeTypeMapping, cloud string) string {
	switch cloud {
	case "aws":
		return mapping.AWS
	case "azure":
		return mapping.Azure
	case "gcp":
		return mapping.GCP
	default:
		return ""
	}
}

// convertNodeType converts a node type from source cloud to target cloud format
// Returns the original node type if no mapping is found or if clouds are the same
func convertNodeType(nodeType string, sourceCloud, targetCloud string, mappings *nodeTypeMappings) string {
	if sourceCloud == targetCloud || mappings == nil {
		return nodeType
	}

	// Find the mapping entry that contains the source node type
	for _, mapping := range mappings.Mappings {
		sourceNodeType := getNodeTypeForCloud(mapping, sourceCloud)
		if sourceNodeType == nodeType {
			// Found the source node type, now get the target
			targetNodeType := getNodeTypeForCloud(mapping, targetCloud)
			if targetNodeType != "" {
				log.Printf("[DEBUG] Converting node type %s (%s) to %s (%s)",
					nodeType, sourceCloud, targetNodeType, targetCloud)
				return targetNodeType
			}
			log.Printf("[WARN] No %s node type mapping found for %s (%s)",
				targetCloud, nodeType, sourceCloud)
			return nodeType
		}
	}

	// No mapping found, return original
	log.Printf("[WARN] No mapping found for node type %s (%s)", nodeType, sourceCloud)
	return nodeType
}
