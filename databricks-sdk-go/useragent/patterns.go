package useragent

import (
	"fmt"
	"regexp"
)

// Regular expression copied from https://semver.org/.
const (
	semVerCore          = `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`
	semVerPrerelease    = `(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?`
	semVerBuildmetadata = `(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`
)

var regexpSemVer = regexp.MustCompile(`^` + semVerCore + semVerPrerelease + semVerBuildmetadata + `$`)

// Sanitize replaces all non-alphanumeric characters with a hyphen. Use this to
// ensure that the user agent value is valid. This is useful when the value is not
// ensured to be valid at compile time.
//
// Example: You want to avoid having '/' and ' ' in the value because it will
// make downstream applications fail.
//
// Note: Semver strings are comprised of alphanumeric characters, hyphens, periods
// and plus signs. This function will not remove these characters.
// see:
// 1. https://semver.org/#spec-item-9
// 2. https://semver.org/#spec-item-10
var regexpAlphanum = regexp.MustCompile(`^[0-9A-Za-z_\.\+-]+$`)
var regexpAlphanumInverse = regexp.MustCompile(`[^0-9A-Za-z_\.\+-]`)

func Sanitize(s string) string {
	return regexpAlphanumInverse.ReplaceAllString(s, "-")
}

func matchSemVer(s string) error {
	if regexpSemVer.MatchString(s) {
		return nil
	}
	return fmt.Errorf("invalid semver string: %s", s)
}

func matchAlphanum(s string) error {
	if regexpAlphanum.MatchString(s) {
		return nil
	}
	return fmt.Errorf("invalid alphanumeric string: %s", s)
}

func matchAlphanumOrSemVer(s string) error {
	if regexpAlphanum.MatchString(s) || regexpSemVer.MatchString(s) {
		return nil
	}
	return fmt.Errorf("invalid alphanumeric or semver string: %s", s)
}
