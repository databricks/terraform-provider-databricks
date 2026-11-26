package useragent

import (
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
)

func TestLookupAgentProvider(t *testing.T) {
	tests := []struct {
		name   string
		envs   map[string]string
		expect string
	}{
		{
			name:   "no agent",
			envs:   nil,
			expect: "",
		},
		{
			name:   "claude code",
			envs:   map[string]string{"CLAUDECODE": "1"},
			expect: "claude-code",
		},
		{
			name:   "antigravity",
			envs:   map[string]string{"ANTIGRAVITY_AGENT": "1"},
			expect: "antigravity",
		},
		{
			name:   "cline",
			envs:   map[string]string{"CLINE_ACTIVE": "1"},
			expect: "cline",
		},
		{
			name:   "codex",
			envs:   map[string]string{"CODEX_CI": "1"},
			expect: "codex",
		},
		{
			name:   "copilot cli",
			envs:   map[string]string{"COPILOT_CLI": "1"},
			expect: "copilot-cli",
		},
		{
			name:   "cursor",
			envs:   map[string]string{"CURSOR_AGENT": "1"},
			expect: "cursor",
		},
		{
			name:   "gemini cli",
			envs:   map[string]string{"GEMINI_CLI": "1"},
			expect: "gemini-cli",
		},
		{
			name:   "opencode",
			envs:   map[string]string{"OPENCODE": "1"},
			expect: "opencode",
		},
		{
			name:   "openclaw",
			envs:   map[string]string{"OPENCLAW_SHELL": "exec"},
			expect: "openclaw",
		},
		{
			name:   "multiple agents stacked (e.g. Cursor CLI subagent invoked by Claude Code)",
			envs:   map[string]string{"CLAUDECODE": "1", "CURSOR_AGENT": "1"},
			expect: "multiple",
		},
		{
			name:   "three stacked agents also report multiple",
			envs:   map[string]string{"CLAUDECODE": "1", "CURSOR_AGENT": "1", "AUGMENT_AGENT": "1"},
			expect: "multiple",
		},
		{
			name:   "empty value still counts as set",
			envs:   map[string]string{"CLAUDECODE": ""},
			expect: "claude-code",
		},
		// New agent detections.
		{
			name:   "goose via GOOSE_TERMINAL",
			envs:   map[string]string{"GOOSE_TERMINAL": "1"},
			expect: "goose",
		},
		{
			name:   "goose via AGENT",
			envs:   map[string]string{"AGENT": "goose"},
			expect: "goose",
		},
		{
			name:   "goose via both GOOSE_TERMINAL and AGENT is not ambiguous",
			envs:   map[string]string{"GOOSE_TERMINAL": "1", "AGENT": "goose"},
			expect: "goose",
		},
		{
			name:   "amp via AMP_CURRENT_THREAD_ID",
			envs:   map[string]string{"AMP_CURRENT_THREAD_ID": "abc123"},
			expect: "amp",
		},
		{
			name:   "amp via AGENT",
			envs:   map[string]string{"AGENT": "amp"},
			expect: "amp",
		},
		{
			name:   "amp via both AMP_CURRENT_THREAD_ID and AGENT is not ambiguous",
			envs:   map[string]string{"AMP_CURRENT_THREAD_ID": "abc123", "AGENT": "amp"},
			expect: "amp",
		},
		{
			name:   "augment",
			envs:   map[string]string{"AUGMENT_AGENT": "1"},
			expect: "augment",
		},
		{
			name:   "vscode-agent",
			envs:   map[string]string{"VSCODE_AGENT": "1"},
			expect: "vscode-agent",
		},
		{
			name:   "kiro",
			envs:   map[string]string{"KIRO": "1"},
			expect: "kiro",
		},
		{
			name:   "windsurf",
			envs:   map[string]string{"WINDSURF_AGENT": "1"},
			expect: "windsurf",
		},
		{
			name:   "pi",
			envs:   map[string]string{"PI_CODING_AGENT": "true"},
			expect: "pi",
		},
		// AGENT fallback behavior.
		{
			name:   "AGENT=cursor falls back to cursor",
			envs:   map[string]string{"AGENT": "cursor"},
			expect: "cursor",
		},
		{
			name:   "AGENT with unrecognized value passes through (sanitized)",
			envs:   map[string]string{"AGENT": "someweirdthing"},
			expect: "someweirdthing",
		},
		{
			name:   "AGENT versioned variant passes through unchanged",
			envs:   map[string]string{"AGENT": "claude-code_2-1-141_agent"},
			expect: "claude-code_2-1-141_agent",
		},
		{
			name:   "AGENT with disallowed chars is sanitized to hyphens",
			envs:   map[string]string{"AGENT": "claude code/agent"},
			expect: "claude-code-agent",
		},
		{
			name:   "AGENT longer than the cap is truncated",
			envs:   map[string]string{"AGENT": strings.Repeat("a", 100)},
			expect: strings.Repeat("a", 64),
		},
		{
			name:   "AGENT empty string does not trigger fallback",
			envs:   map[string]string{"AGENT": ""},
			expect: "",
		},
		{
			name:   "known matcher wins over AGENT fallback",
			envs:   map[string]string{"AGENT": "somethingunknown", "CLAUDECODE": "1"},
			expect: "claude-code",
		},
		// Explicit env var always wins over the generic AGENT env var.
		{
			name:   "explicit CLAUDECODE wins over AGENT=goose",
			envs:   map[string]string{"AGENT": "goose", "CLAUDECODE": "1"},
			expect: "claude-code",
		},
		{
			name:   "explicit GOOSE_TERMINAL wins over AGENT=cursor",
			envs:   map[string]string{"GOOSE_TERMINAL": "1", "AGENT": "cursor"},
			expect: "goose",
		},
		// VSCODE_AGENT can legitimately stack with other agents (e.g. running
		// Copilot CLI from a VS Code agent terminal).
		{
			name:   "VSCODE_AGENT + COPILOT_CLI reports multiple",
			envs:   map[string]string{"VSCODE_AGENT": "1", "COPILOT_CLI": "1"},
			expect: "multiple",
		},
		// AI_AGENT fallback (Vercel @vercel/detect-agent convention).
		{
			name:   "AI_AGENT=cursor falls back to cursor",
			envs:   map[string]string{"AI_AGENT": "cursor"},
			expect: "cursor",
		},
		{
			name:   "AI_AGENT empty string does not trigger fallback",
			envs:   map[string]string{"AI_AGENT": ""},
			expect: "",
		},
		{
			name:   "known matcher wins over AI_AGENT fallback",
			envs:   map[string]string{"AI_AGENT": "somethingunknown", "CLAUDECODE": "1"},
			expect: "claude-code",
		},
		// AGENT vs AI_AGENT precedence: AGENT wins when both are non-empty.
		{
			name:   "AGENT wins over AI_AGENT when both are set to known products",
			envs:   map[string]string{"AGENT": "claude-code", "AI_AGENT": "cursor"},
			expect: "claude-code",
		},
		{
			name:   "AGENT set to unrecognized non-empty value still wins over AI_AGENT",
			envs:   map[string]string{"AGENT": "somethingunknown", "AI_AGENT": "cursor"},
			expect: "somethingunknown",
		},
		{
			name:   "AGENT set, AI_AGENT empty: AGENT value is used",
			envs:   map[string]string{"AGENT": "cursor", "AI_AGENT": ""},
			expect: "cursor",
		},
		{
			name:   "empty AGENT falls through to AI_AGENT",
			envs:   map[string]string{"AGENT": "", "AI_AGENT": "cursor"},
			expect: "cursor",
		},
		{
			name:   "both AGENT and AI_AGENT empty returns no agent",
			envs:   map[string]string{"AGENT": "", "AI_AGENT": ""},
			expect: "",
		},
		{
			name:   "explicit CLAUDECODE wins over AI_AGENT=cursor",
			envs:   map[string]string{"AI_AGENT": "cursor", "CLAUDECODE": "1"},
			expect: "claude-code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env.CleanupEnvironment(t)
			ClearCache()
			for k, v := range tt.envs {
				t.Setenv(k, v)
			}
			got := lookupAgentProvider()
			if got != tt.expect {
				t.Errorf("lookupAgentProvider() = %q, want %q", got, tt.expect)
			}
		})
	}
}

func TestAgentProviderCached(t *testing.T) {
	env.CleanupEnvironment(t)
	ClearCache()
	t.Setenv("CURSOR_AGENT", "1")
	got := AgentProvider()
	if got != "cursor" {
		t.Errorf("AgentProvider() = %q, want %q", got, "cursor")
	}
	// Change env after caching. Cached result should persist.
	t.Setenv("CURSOR_AGENT", "")
	t.Setenv("CLAUDECODE", "1")
	got = AgentProvider()
	if got != "cursor" {
		t.Errorf("AgentProvider() after cache = %q, want %q", got, "cursor")
	}
}
