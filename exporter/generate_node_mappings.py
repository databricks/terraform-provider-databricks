#!/usr/bin/env python3
"""
Generate node type mappings between AWS, Azure, and GCP based on similarity.
Matching criteria: num_cores, memory_mb, category, and local disk configuration.

The script generates three-way mappings (AWS <-> Azure <-> GCP) with the following strategy:
1. Find exact/close matches using strict similarity scoring (min score 150)
2. Fill in missing clouds with approximate matches based on core count and other characteristics
3. Result: ~99.5% of mappings have all three clouds populated

This ensures maximum compatibility when converting resources across clouds.
"""

import argparse
import json
import sys
from typing import Dict, List, Optional, Tuple


def load_node_types(file_path: str) -> List[Dict]:
    """Load node types from a JSON file."""
    with open(file_path, 'r') as f:
        data = json.load(f)
    return data.get('node_types', [])


def normalize_category(category: str) -> str:
    """Normalize category names for comparison."""
    category = category.lower()
    # Map similar categories
    if 'memory' in category or 'highmem' in category:
        return 'memory'
    elif 'compute' in category:
        return 'compute'
    elif 'storage' in category:
        return 'storage'
    elif 'general' in category or 'standard' in category:
        return 'general'
    return 'general'


def calculate_similarity(node1: Dict, node2: Dict) -> float:
    """
    Calculate similarity score between two node types.
    Higher score = more similar.
    """
    score = 0.0

    # Cores match (most important) - must match exactly for consideration
    cores1 = node1.get('num_cores', 0)
    cores2 = node2.get('num_cores', 0)
    if cores1 != cores2:
        return -1.0  # Not compatible
    score += 100.0

    # Memory match (very important)
    mem1 = node1.get('memory_mb', 0)
    mem2 = node2.get('memory_mb', 0)
    if mem1 > 0 and mem2 > 0:
        mem_ratio = min(mem1, mem2) / max(mem1, mem2)
        # Allow up to 10% difference in memory
        if mem_ratio >= 0.9:
            score += 50.0 * mem_ratio
        else:
            score -= 20.0  # Penalize large memory differences

    # Category match
    cat1 = normalize_category(node1.get('category', ''))
    cat2 = normalize_category(node2.get('category', ''))
    if cat1 == cat2:
        score += 30.0

    # Disk configuration similarity
    node1_type = node1.get('node_instance_type', {})
    node2_type = node2.get('node_instance_type', {})

    disk1_size = node1_type.get('local_disk_size_gb', 0)
    disk2_size = node2_type.get('local_disk_size_gb', 0)

    # Both have local disks or both don't
    if (disk1_size > 0) == (disk2_size > 0):
        score += 10.0

        # If both have disks, compare sizes
        if disk1_size > 0 and disk2_size > 0:
            disk_ratio = min(disk1_size, disk2_size) / max(disk1_size, disk2_size)
            if disk_ratio >= 0.5:  # Within 2x of each other
                score += 10.0 * disk_ratio

    return score


def filter_node_types(nodes: List[Dict]) -> List[Dict]:
    """Filter out deprecated, hidden, or special node types."""
    filtered = []
    for node in nodes:
        # Skip deprecated or hidden nodes
        if node.get('is_deprecated', False) or node.get('is_hidden', False):
            continue

        # Skip GPU nodes for now (different category)
        if node.get('num_gpus', 0) > 0:
            continue

        filtered.append(node)

    return filtered


def find_best_matches(
    source_nodes: List[Dict],
    target_nodes: List[Dict],
    min_score: float = 150.0
) -> Dict[str, str]:
    """
    Find best matches from source to target node types.
    Returns a dict mapping source node_type_id to target node_type_id.
    """
    matches = {}

    for source_node in source_nodes:
        source_id = source_node.get('node_type_id', '')
        best_match = None
        best_score = min_score

        for target_node in target_nodes:
            score = calculate_similarity(source_node, target_node)
            if score > best_score:
                best_score = score
                best_match = target_node.get('node_type_id', '')

        if best_match:
            matches[source_id] = best_match

    return matches


def find_approximate_match(
    reference_node: Dict,
    target_nodes: List[Dict],
    min_score: float = 0.0
) -> Optional[str]:
    """
    Find the best approximate match for a node in target cloud.
    Returns the node_type_id of the best match, even if score is low.
    """
    best_match = None
    best_score = min_score

    for target_node in target_nodes:
        score = calculate_similarity(reference_node, target_node)
        if score > best_score:
            best_score = score
            best_match = target_node.get('node_type_id', '')

    return best_match


def get_node_by_id(node_id: str, nodes: List[Dict]) -> Optional[Dict]:
    """Get a node by its node_type_id."""
    for node in nodes:
        if node.get('node_type_id') == node_id:
            return node
    return None


def fill_missing_clouds(
    mappings: List[Dict],
    aws_nodes: List[Dict],
    azure_nodes: List[Dict],
    gcp_nodes: List[Dict]
) -> List[Dict]:
    """
    Fill in missing cloud mappings with approximate matches.
    Uses the characteristics of existing mappings to find best approximations.
    """
    completed_mappings = []
    nodes_by_cloud = {
        'aws': aws_nodes,
        'azure': azure_nodes,
        'gcp': gcp_nodes
    }

    for mapping in mappings:
        new_mapping = mapping.copy()

        # Determine which clouds are present and which are missing
        present_clouds = [cloud for cloud in ['aws', 'azure', 'gcp'] if cloud in mapping]
        missing_clouds = [cloud for cloud in ['aws', 'azure', 'gcp'] if cloud not in mapping]

        # For each missing cloud, find an approximate match
        for missing_cloud in missing_clouds:
            # Use the first available present cloud as reference
            if present_clouds:
                reference_cloud = present_clouds[0]
                reference_id = mapping[reference_cloud]
                reference_node = get_node_by_id(reference_id, nodes_by_cloud[reference_cloud])

                if reference_node:
                    # Find approximate match in missing cloud
                    approx_match = find_approximate_match(
                        reference_node,
                        nodes_by_cloud[missing_cloud],
                        min_score=0.0  # Accept any match with same core count
                    )
                    if approx_match:
                        new_mapping[missing_cloud] = approx_match

        completed_mappings.append(new_mapping)

    return completed_mappings


def generate_mappings(
    aws_nodes: List[Dict],
    azure_nodes: List[Dict],
    gcp_nodes: List[Dict]
) -> List[Dict]:
    """
    Generate three-way mappings between AWS, Azure, and GCP.
    """
    # Filter nodes
    aws_filtered = filter_node_types(aws_nodes)
    azure_filtered = filter_node_types(azure_nodes)
    gcp_filtered = filter_node_types(gcp_nodes)

    print(f"Filtered node types: AWS={len(aws_filtered)}, Azure={len(azure_filtered)}, GCP={len(gcp_filtered)}")

    # Create mappings from each cloud to every other cloud
    aws_to_azure = find_best_matches(aws_filtered, azure_filtered)
    aws_to_gcp = find_best_matches(aws_filtered, gcp_filtered)
    azure_to_aws = find_best_matches(azure_filtered, aws_filtered)
    azure_to_gcp = find_best_matches(azure_filtered, gcp_filtered)
    gcp_to_aws = find_best_matches(gcp_filtered, aws_filtered)
    gcp_to_azure = find_best_matches(gcp_filtered, azure_filtered)

    print(f"Found mappings: AWS->Azure={len(aws_to_azure)}, AWS->GCP={len(aws_to_gcp)}")
    print(f"                Azure->AWS={len(azure_to_aws)}, Azure->GCP={len(azure_to_gcp)}")
    print(f"                GCP->AWS={len(gcp_to_aws)}, GCP->Azure={len(gcp_to_azure)}")

    # Build three-way mappings
    mappings = []
    all_aws_ids = set(aws_to_azure.keys()) | set(aws_to_gcp.keys()) | set(azure_to_aws.values()) | set(gcp_to_aws.values())

    for aws_id in sorted(all_aws_ids):
        azure_id = aws_to_azure.get(aws_id)
        gcp_id = aws_to_gcp.get(aws_id)

        # Try to fill in missing mappings through reverse lookups
        if not azure_id and aws_id in gcp_to_azure.values():
            # Find GCP node that maps to this AWS node
            for g_id, a_id in gcp_to_aws.items():
                if a_id == aws_id and g_id in gcp_to_azure:
                    azure_id = gcp_to_azure[g_id]
                    break

        if not gcp_id and aws_id in azure_to_gcp.values():
            # Find Azure node that maps to this AWS node
            for az_id, a_id in azure_to_aws.items():
                if a_id == aws_id and az_id in azure_to_gcp:
                    gcp_id = azure_to_gcp[az_id]
                    break

        # Only add if we have at least 2 out of 3 clouds mapped
        if sum([bool(aws_id), bool(azure_id), bool(gcp_id)]) >= 2:
            mapping = {}
            if aws_id:
                mapping['aws'] = aws_id
            if azure_id:
                mapping['azure'] = azure_id
            if gcp_id:
                mapping['gcp'] = gcp_id

            # Avoid duplicates
            if mapping not in mappings:
                mappings.append(mapping)

    # Also process Azure-centric mappings
    all_azure_ids = set(azure_to_aws.keys()) | set(azure_to_gcp.keys())
    for azure_id in sorted(all_azure_ids):
        aws_id = azure_to_aws.get(azure_id)
        gcp_id = azure_to_gcp.get(azure_id)

        if sum([bool(aws_id), bool(azure_id), bool(gcp_id)]) >= 2:
            mapping = {}
            if aws_id:
                mapping['aws'] = aws_id
            if azure_id:
                mapping['azure'] = azure_id
            if gcp_id:
                mapping['gcp'] = gcp_id

            if mapping not in mappings:
                mappings.append(mapping)

    # And GCP-centric mappings
    all_gcp_ids = set(gcp_to_aws.keys()) | set(gcp_to_azure.keys())
    for gcp_id in sorted(all_gcp_ids):
        aws_id = gcp_to_aws.get(gcp_id)
        azure_id = gcp_to_azure.get(gcp_id)

        if sum([bool(aws_id), bool(azure_id), bool(gcp_id)]) >= 2:
            mapping = {}
            if aws_id:
                mapping['aws'] = aws_id
            if azure_id:
                mapping['azure'] = azure_id
            if gcp_id:
                mapping['gcp'] = gcp_id

            if mapping not in mappings:
                mappings.append(mapping)

    print(f"\nFilling in missing cloud mappings with approximations...")
    # Fill in missing clouds with approximate matches
    mappings = fill_missing_clouds(mappings, aws_filtered, azure_filtered, gcp_filtered)

    # Count how many complete mappings we now have
    complete_mappings = sum(1 for m in mappings if 'aws' in m and 'azure' in m and 'gcp' in m)
    print(f"Complete three-way mappings after approximation: {complete_mappings}/{len(mappings)}")

    return mappings


def main():
    parser = argparse.ArgumentParser(
        description='Generate node type mappings between AWS, Azure, and GCP clouds.',
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Generate mappings from node type JSON files
  %(prog)s --aws node-types-aws.json --azure node-types-azure.json --gcp node-types-gcp.json

  # Specify custom output file
  %(prog)s --aws aws.json --azure azure.json --gcp gcp.json --output mappings.json
        """
    )

    parser.add_argument(
        '--aws',
        required=True,
        help='Path to AWS node types JSON file'
    )

    parser.add_argument(
        '--azure',
        required=True,
        help='Path to Azure node types JSON file'
    )

    parser.add_argument(
        '--gcp',
        required=True,
        help='Path to GCP node types JSON file'
    )

    parser.add_argument(
        '--output',
        '-o',
        default='node_type_mapping.json',
        help='Output file path (default: node_type_mapping.json)'
    )

    args = parser.parse_args()

    # Load node types from all three clouds
    print("Loading node type data...")
    try:
        aws_nodes = load_node_types(args.aws)
        azure_nodes = load_node_types(args.azure)
        gcp_nodes = load_node_types(args.gcp)
    except FileNotFoundError as e:
        print(f"Error: {e}", file=sys.stderr)
        return 1
    except json.JSONDecodeError as e:
        print(f"Error: Invalid JSON format - {e}", file=sys.stderr)
        return 1

    print(f"Loaded: AWS={len(aws_nodes)}, Azure={len(azure_nodes)}, GCP={len(gcp_nodes)} node types")

    # Generate mappings
    print("\nGenerating mappings...")
    mappings = generate_mappings(aws_nodes, azure_nodes, gcp_nodes)

    print(f"\nGenerated {len(mappings)} mappings")

    # Count complete mappings (all three clouds)
    complete = sum(1 for m in mappings if 'aws' in m and 'azure' in m and 'gcp' in m)
    incomplete = len(mappings) - complete

    print(f"  - Complete (all 3 clouds): {complete}")
    print(f"  - Incomplete (2 clouds): {incomplete}")

    # Create output structure
    output = {
        "version": "1.0",
        "mappings": mappings
    }

    # Write to file
    try:
        with open(args.output, 'w') as f:
            json.dump(output, f, indent=2)
        print(f"\nMapping file written to: {args.output}")
    except IOError as e:
        print(f"Error writing output file: {e}", file=sys.stderr)
        return 1

    # Print some sample mappings
    print("\nSample complete mappings (first 10):")
    complete_samples = [m for m in mappings if 'aws' in m and 'azure' in m and 'gcp' in m][:10]
    for i, mapping in enumerate(complete_samples, 1):
        aws = mapping.get('aws', 'N/A')
        azure = mapping.get('azure', 'N/A')
        gcp = mapping.get('gcp', 'N/A')
        print(f"  {i}. AWS: {aws:20s} Azure: {azure:25s} GCP: {gcp}")

    return 0


if __name__ == '__main__':
    sys.exit(main())
