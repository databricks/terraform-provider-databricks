package exporter

import (
	"regexp"
	"strings"
)

const (
	sqlQuotedIdentifier = "`((?:[^`]|``)+)`"
	sqlPlainIdentifier  = `([A-Za-z_][A-Za-z0-9_]*)`
	sqlIdentifierPart   = `(?:` + sqlQuotedIdentifier + `|` + sqlPlainIdentifier + `)`
)

var (
	// sqlTableReferenceRegex matches three-part names (`catalog.schema.object`) that follow FROM or
	// JOIN keywords (optionally with STREAM) in SQL text. Each part could be either plain or quoted
	// with backticks. Groups: 1 - quoted catalog, 2 - plain catalog, 3 - quoted schema, 4 - plain
	// schema, 5/6 - object.
	sqlTableReferenceRegex = regexp.MustCompile(`(?i:\b(?:FROM|JOIN)\s+(?:STREAM\s*\(\s*|STREAM\s+)?)` +
		sqlIdentifierPart + `\.` + sqlIdentifierPart + `\.` + sqlIdentifierPart)
	sqlPlainIdentifierRegex = regexp.MustCompile(`^` + sqlPlainIdentifier + `$`)
)

// sqlSchemaName describes the `catalog.schema` part of a three-part name found in SQL text.
type sqlSchemaName struct {
	// Start and End of the `catalog.schema` part (including backticks) in the SQL text
	Start, End int
	// Catalog and Schema names in lower case, without backticks
	Catalog, Schema string
}

// FullName returns the full name of the schema in the `catalog.schema` format
func (s sqlSchemaName) FullName() string {
	return s.Catalog + "." + s.Schema
}

// sqlNonCodeRanges returns the ranges of comments, string literals, and content of quoted
// identifiers in SQL text, so we don't match names inside them.
func sqlNonCodeRanges(sql string) [][2]int {
	var ranges [][2]int
	for i := 0; i < len(sql); i++ {
		start := i
		switch {
		case strings.HasPrefix(sql[i:], "--"):
			if end := strings.IndexByte(sql[i:], '\n'); end >= 0 {
				i += end
			} else {
				i = len(sql)
			}
			ranges = append(ranges, [2]int{start, i})
		case strings.HasPrefix(sql[i:], "/*"):
			if end := strings.Index(sql[i+2:], "*/"); end >= 0 {
				i += end + 3
			} else {
				i = len(sql)
			}
			ranges = append(ranges, [2]int{start, i})
		case sql[i] == '\'' || sql[i] == '"':
			quote := sql[i]
			for i++; i < len(sql) && sql[i] != quote; i++ {
				if sql[i] == '\\' {
					i++
				}
			}
			ranges = append(ranges, [2]int{start, i})
		case sql[i] == '`':
			if end := strings.IndexByte(sql[i+1:], '`'); end >= 0 {
				i += end + 1
			} else {
				i = len(sql)
			}
			// only the content of quoted identifier, so names starting with backticks are matched
			ranges = append(ranges, [2]int{start + 1, i - 1})
		}
	}
	return ranges
}

func isInRanges(pos int, ranges [][2]int) bool {
	for _, r := range ranges {
		if pos >= r[0] && pos <= r[1] {
			return true
		}
	}
	return false
}

// sqlIdentifierAt returns an unquoted identifier from one of the two groups of the match. It
// returns false if the name can't be used without backticks.
func sqlIdentifierAt(sql string, match []int, quotedGroup, plainGroup int) (string, bool) {
	if match[2*plainGroup] >= 0 {
		return strings.ToLower(sql[match[2*plainGroup]:match[2*plainGroup+1]]), true
	}
	if match[2*quotedGroup] >= 0 {
		name := sql[match[2*quotedGroup]:match[2*quotedGroup+1]]
		if sqlPlainIdentifierRegex.MatchString(name) {
			return strings.ToLower(name), true
		}
	}
	return "", false
}

// parseSqlSchemaName converts a match of sqlTableReferenceRegex into a schema name. It returns
// false if the match is inside a comment or string literal, or if the names require backticks
// (so we can't replace them with a reference that is rendered without backticks).
func parseSqlSchemaName(sql string, match []int, nonCode [][2]int) (sqlSchemaName, bool) {
	if len(match) < 14 {
		return sqlSchemaName{}, false
	}
	// the span includes backticks of the quoted parts
	start := match[4]
	if match[2] >= 0 {
		start = match[2] - 1
	}
	end := match[9]
	if match[7] >= 0 {
		end = match[7] + 1
	}
	if isInRanges(start, nonCode) {
		return sqlSchemaName{}, false
	}
	catalog, ok := sqlIdentifierAt(sql, match, 1, 2)
	if !ok {
		return sqlSchemaName{}, false
	}
	schema, ok := sqlIdentifierAt(sql, match, 3, 4)
	if !ok {
		return sqlSchemaName{}, false
	}
	return sqlSchemaName{Start: start, End: end, Catalog: catalog, Schema: schema}, true
}

// findSqlSchemaNames returns schemas of three-part names used in FROM and JOIN clauses of SQL
// text, ignoring comments and string literals.
func findSqlSchemaNames(sql string) []sqlSchemaName {
	nonCode := sqlNonCodeRanges(sql)
	var names []sqlSchemaName
	for _, match := range sqlTableReferenceRegex.FindAllStringSubmatchIndex(sql, -1) {
		if name, ok := parseSqlSchemaName(sql, match, nonCode); ok {
			names = append(names, name)
		}
	}
	return names
}

// sqlSchemaNameCapture is a MultiMatchCaptureFunc that substitutes the `catalog.schema` part
// of three-part names in SQL text with a reference to the schema's ID.
func sqlSchemaNameCapture(sql string, match []int) (int, int, string, bool) {
	name, ok := parseSqlSchemaName(sql, match, sqlNonCodeRanges(sql))
	if !ok {
		return 0, 0, "", false
	}
	return name.Start, name.End, name.FullName(), true
}
