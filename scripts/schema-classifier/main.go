// schema-classifier walks two `terraform providers schema -json` dumps and
// reports which deltas are breaking. Exits 1 on any breaking change unless
// --allow-breaking is set. See scripts/schema-classifier/README.md for the
// rule taxonomy.
//
// Known limitation: ForceNew is NOT emitted by `terraform providers schema
// -json`. Flipping an attribute to ForceNew silently destroys/recreates
// resources and is INVISIBLE to this classifier. Reviewers must catch those
// changes manually.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	switch err := run(os.Args[1:], os.Stdout, os.Stderr); err.(type) {
	case nil:
		return
	case errBreaking:
		os.Exit(1)
	default:
		fmt.Fprintln(os.Stderr, "schema-classifier:", err)
		os.Exit(2)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("schema-classifier", flag.ContinueOnError)
	fs.SetOutput(stderr)
	var (
		basePath      = fs.String("base", "", "path to base schema JSON")
		headPath      = fs.String("head", "", "path to head schema JSON")
		format        = fs.String("format", "text", "output format: text or markdown")
		allowBreaking = fs.Bool("allow-breaking", false, "do not exit non-zero on breaking changes (for local dev)")
	)
	fs.Usage = func() {
		fmt.Fprintln(stderr, "Usage: schema-classifier --base BASE.json --head HEAD.json [--format text|markdown] [--allow-breaking]")
		fmt.Fprintln(stderr, "")
		fmt.Fprintln(stderr, "Classifies Terraform provider schema deltas as breaking or non-breaking and")
		fmt.Fprintln(stderr, "exits non-zero when any breaking change is present.")
		fmt.Fprintln(stderr, "")
		fmt.Fprintln(stderr, "Known limitation: ForceNew flips are NOT detected (the Terraform schema")
		fmt.Fprintln(stderr, "dump does not expose ForceNew). Reviewers must catch those manually.")
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *basePath == "" || *headPath == "" {
		fs.Usage()
		return fmt.Errorf("--base and --head are required")
	}

	base, err := loadSchema(*basePath)
	if err != nil {
		return fmt.Errorf("load base: %w", err)
	}
	head, err := loadSchema(*headPath)
	if err != nil {
		return fmt.Errorf("load head: %w", err)
	}

	changes := Classify(base, head)

	switch *format {
	case "text":
		writeText(stdout, changes)
	case "markdown":
		writeMarkdown(stdout, changes)
	default:
		return fmt.Errorf("unknown --format %q (expected text or markdown)", *format)
	}

	if HasBreaking(changes) && !*allowBreaking {
		return errBreaking{}
	}
	return nil
}

type errBreaking struct{}

func (errBreaking) Error() string { return "breaking schema changes detected" }

func loadSchema(path string) (*ProviderSchemas, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s ProviderSchemas
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return &s, nil
}

func writeText(w io.Writer, changes []Change) {
	if len(changes) == 0 {
		fmt.Fprintln(w, "No schema changes detected.")
		return
	}
	var breaking, nonBreaking []Change
	for _, c := range changes {
		if c.Severity == Breaking {
			breaking = append(breaking, c)
		} else {
			nonBreaking = append(nonBreaking, c)
		}
	}
	if len(breaking) > 0 {
		fmt.Fprintf(w, "BREAKING CHANGES (%d):\n", len(breaking))
		for _, c := range breaking {
			fmt.Fprintf(w, "  - [%s] %s\n      %s\n", c.Kind, c.Path, c.Message)
		}
	}
	if len(nonBreaking) > 0 {
		if len(breaking) > 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "Non-breaking changes (%d):\n", len(nonBreaking))
		for _, c := range nonBreaking {
			fmt.Fprintf(w, "  - [%s] %s\n      %s\n", c.Kind, c.Path, c.Message)
		}
	}
}

func writeMarkdown(w io.Writer, changes []Change) {
	if len(changes) == 0 {
		fmt.Fprintln(w, "_No schema changes detected._")
		return
	}
	var breaking, nonBreaking []Change
	for _, c := range changes {
		if c.Severity == Breaking {
			breaking = append(breaking, c)
		} else {
			nonBreaking = append(nonBreaking, c)
		}
	}
	if len(breaking) > 0 {
		fmt.Fprintf(w, "### Breaking changes (%d)\n\n", len(breaking))
		writeMarkdownTable(w, breaking)
		fmt.Fprintln(w)
	}
	if len(nonBreaking) > 0 {
		fmt.Fprintf(w, "<details><summary>Non-breaking changes (%d)</summary>\n\n", len(nonBreaking))
		writeMarkdownTable(w, nonBreaking)
		fmt.Fprintln(w, "</details>")
	}
}

func writeMarkdownTable(w io.Writer, changes []Change) {
	fmt.Fprintln(w, "| Kind | Path | Description |")
	fmt.Fprintln(w, "|---|---|---|")
	for _, c := range changes {
		fmt.Fprintf(w, "| `%s` | `%s` | %s |\n",
			c.Kind, c.Path, strings.ReplaceAll(c.Message, "|", "\\|"))
	}
}
