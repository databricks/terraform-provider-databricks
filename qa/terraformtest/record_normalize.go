package terraformtest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"gopkg.in/yaml.v3"
)

type recordingID struct{ replay, live string }

const recordingReplayHost = "example.databricks.com"

type recordingNormalizer struct {
	ids      []recordingID
	liveHost string
	secrets  []string
}

func newRecordingNormalizer(d *testDefinition, cfg *config.Config) (*recordingNormalizer, error) {
	host, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("cannot normalize recorder profile host")
	}
	n := &recordingNormalizer{liveHost: host.Host}
	for _, field := range []struct{ name, live, fallback string }{
		{"workspace_id", cfg.WorkspaceID, "12345"}, {"account_id", cfg.AccountID, "replay-account"},
	} {
		stable, err := replayRoutingID(d, field.name, field.fallback)
		if err != nil {
			return nil, err
		}
		if field.live == "" {
			if _, configured := d.HostMetadata[field.name]; configured || stable != field.fallback {
				return nil, fmt.Errorf("recorder profile must provide %s for this test", field.name)
			}
			continue
		}
		if stable == field.live {
			return nil, fmt.Errorf("%s in YAML must use a replay ID, different from the recorder profile", field.name)
		}
		n.ids = append(n.ids, recordingID{stable, field.live})
	}
	for _, attr := range config.ConfigAttributes {
		if attr.Sensitive && !attr.IsZero(cfg) {
			n.secrets = append(n.secrets, attr.GetString(cfg))
		}
	}
	return n, nil
}

// Prefer the fixture's explicit routing IDs, including literal provider fields.
// Distinct routing IDs need multiple profiles/hosts and cannot be refreshed by
// this recorder. Reject ambiguity before any API request.
func replayRoutingID(d *testDefinition, name, fallback string) (string, error) {
	values := map[string]bool{}
	if value, ok := d.HostMetadata[name]; ok {
		values[fmt.Sprint(value)] = true
	}
	for _, step := range d.Steps {
		file, diags := hclsyntax.ParseConfig([]byte(step.Config), "test.tf", hcl.InitialPos)
		if diags.HasErrors() {
			continue // Invalid HCL can itself be the expected diagnostic.
		}
		for _, block := range file.Body.(*hclsyntax.Body).Blocks {
			if block.Type != "provider" || len(block.Labels) != 1 || block.Labels[0] != "databricks" {
				continue
			}
			if attr := block.Body.Attributes[name]; attr != nil {
				value, diags := attr.Expr.Value(nil)
				if !diags.HasErrors() && value.IsKnown() && !value.IsNull() && value.Type().FriendlyName() == "string" {
					values[value.AsString()] = true
				}
			}
		}
	}
	delete(values, "")
	if len(values) > 1 {
		return "", fmt.Errorf("recording requires one replay %s; metadata and literal provider IDs must agree", name)
	}
	for value := range values {
		return value, nil
	}
	return fallback, nil
}

func recordingDefinition(path string, original []byte, n *recordingNormalizer) (*testDefinition, error) {
	live, err := n.transform(original, true, false)
	if err != nil {
		return nil, err
	}
	d, err := decodeDefinition(bytes.NewReader(live))
	if err == nil {
		err = d.loadFiles(filepath.Dir(path))
	}
	return d, err
}

// The selected profile host and routing IDs round-trip between live and replay.
// Boundaries keep unrelated hosts and IDs intact; URL paths and schemes remain.
func (n *recordingNormalizer) text(value string, toLive bool) string {
	if n.liveHost != "" {
		from, to := n.liveHost, recordingReplayHost
		if toLive {
			from, to = to, from
		}
		value = replaceRecordingValue(value, from, to, hostCharacter)
	}
	for _, id := range n.ids {
		from, to := id.live, id.replay
		if toLive {
			from, to = to, from
		}
		value = replaceRecordingValue(value, from, to, idCharacter)
	}
	return value
}

func replaceRecordingValue(value, from, to string, part func(byte) bool) string {
	var out strings.Builder
	cursor := 0
	for _, index := range allIndices(value, from) {
		end := index + len(from)
		if (index != 0 && part(value[index-1])) || (end != len(value) && part(value[end])) {
			continue
		}
		out.WriteString(value[cursor:index])
		out.WriteString(to)
		cursor = end
	}
	out.WriteString(value[cursor:])
	return out.String()
}

func allIndices(value, part string) []int {
	var result []int
	for offset := 0; part != "" && offset < len(value); {
		index := strings.Index(value[offset:], part)
		if index < 0 {
			break
		}
		result = append(result, offset+index)
		offset += index + len(part)
	}
	return result
}

func idCharacter(b byte) bool {
	return b >= '0' && b <= '9' || b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b == '_' || b == '-'
}

func hostCharacter(b byte) bool {
	return idCharacter(b) || b == '.' || b == ':' || b == '[' || b == ']'
}

func (n *recordingNormalizer) transform(data []byte, toLive, redact bool) ([]byte, error) {
	var document yaml.Node
	if err := yaml.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	if err := n.visit(&document, "", toLive, redact, false); err != nil {
		return nil, err
	}
	return yaml.Marshal(&document)
}

func (n *recordingNormalizer) visit(node *yaml.Node, key string, toLive, redact, inJSON bool) error {
	if node.Kind == yaml.ScalarNode {
		// A JSON property named base64 is ordinary data, not a body selector.
		return n.scalar(node, key, toLive, redact, key == "base64" && !inJSON)
	}
	for i, child := range node.Content {
		childKey := key
		if node.Kind == yaml.MappingNode {
			if i%2 == 0 {
				continue
			}
			childKey = node.Content[i-1].Value
		}
		if err := n.visit(child, childKey, toLive, redact, inJSON || key == "json"); err != nil {
			return err
		}
	}
	return nil
}

func (n *recordingNormalizer) scalar(node *yaml.Node, key string, toLive, redact, binary bool) error {
	value := node.Value
	if binary {
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return err
		}
		value = string(decoded)
	}
	value = n.text(value, toLive)
	if redact {
		for _, secret := range n.secrets {
			value = strings.ReplaceAll(value, secret, "[REDACTED]")
		}
		if sensitiveRecordingKey(key) && node.Tag != "!!null" {
			value = "[REDACTED]"
			node.Tag = "!!str"
		}
	}
	if binary {
		value = base64.StdEncoding.EncodeToString([]byte(value))
	}
	if node.Tag == "!!int" && node.Value != value {
		if _, err := strconv.ParseUint(value, 10, 64); err != nil {
			if _, signedErr := strconv.ParseInt(value, 10, 64); signedErr != nil {
				return fmt.Errorf("recorded numeric ID cannot map to a nonnumeric replay ID")
			}
		}
	}
	node.Value = value
	return nil
}

func sensitiveRecordingKey(key string) bool {
	key = strings.ToLower(strings.NewReplacer("-", "", "_", "").Replace(key))
	switch key {
	case "authorization", "cookie", "setcookie", "token", "tokenvalue", "accesstoken", "refreshtoken", "idtoken",
		"password", "clientsecret", "secret", "secretvalue", "stringvalue", "bytesvalue", "privatekey",
		"secretaccesskey", "awssecretkey", "awssecretaccesskey", "awssessiontoken", "azureclientsecret", "gcpserviceaccountkey":
		return true
	}
	return false
}
