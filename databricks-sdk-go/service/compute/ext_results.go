package compute

import (
	"errors"
	"html"
	"regexp"
	"strings"
)

var (
	// IPython's output prefixes
	outRE = regexp.MustCompile(`Out\[[\d\s]+\]:\s`)
	// HTML tags
	tagRE = regexp.MustCompile(`<[^>]*>`)
	// just exception content without exception name
	exceptionRE = regexp.MustCompile(`.*Exception:\s+(.*)`)
	// execution errors resulting from http errors are sometimes hidden in these keys
	executionErrorRE = regexp.MustCompile(`ExecutionError: ([\s\S]*)\n(StatusCode=[0-9]*)\n(StatusDescription=.*)\n`)
	// usual error message explanation is hidden in this key
	errorMessageRE = regexp.MustCompile(`ErrorMessage=(.+)\n`)
)

// Failed tells if command execution failed
func (r *Results) Failed() bool {
	return r.ResultType == "error"
}

// Text returns plain text results
func (r *Results) Text() string {
	if r.ResultType != "text" {
		return ""
	}
	return outRE.ReplaceAllLiteralString(r.Data.(string), "")
}

// Err returns error type
func (r *Results) Err() error {
	if !r.Failed() {
		return nil
	}
	return errors.New(r.Error())
}

// Error returns error in a bit more friendly way
func (r *Results) Error() string {
	if r.ResultType != "error" {
		return ""
	}
	summary := tagRE.ReplaceAllLiteralString(r.Summary, "")
	summary = html.UnescapeString(summary)

	exceptionMatches := exceptionRE.FindStringSubmatch(summary)
	if len(exceptionMatches) == 2 {
		summary = strings.ReplaceAll(exceptionMatches[1], "; nested exception is:", "")
		summary = strings.TrimRight(summary, " ")
		return summary
	}

	executionErrorMatches := executionErrorRE.FindStringSubmatch(r.Cause)
	if len(executionErrorMatches) == 4 {
		return strings.Join(executionErrorMatches[1:], "\n")
	}

	errorMessageMatches := errorMessageRE.FindStringSubmatch(r.Cause)
	if len(errorMessageMatches) == 2 {
		return errorMessageMatches[1]
	}

	return summary
}

// Scan scans for results
// TODO: change API, also in terraform (databricks_sql_permissions)
// for now we're adding `pos` field artificially. this must be removed
// before this repo is public.
func (r *Results) Scan(dest ...any) bool {
	if r.ResultType != ResultTypeTable {
		return false
	}
	if rows, ok := r.Data.([]any); ok {
		if r.Pos >= len(rows) {
			return false
		}
		if cols, ok := rows[r.Pos].([]any); ok {
			for i := range dest {
				switch d := dest[i].(type) {
				case *string:
					*d = cols[i].(string)
				case *int:
					*d = cols[i].(int)
				case *bool:
					*d = cols[i].(bool)
				}
			}
			r.Pos++
			return true
		}
	}
	return false
}
