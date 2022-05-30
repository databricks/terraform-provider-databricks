package acceptance

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/commands"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

// For writing a unit test to intercept the errors (t.Fatalf literally ends the test in failure)
func environmentTemplate(t *testing.T, template string, otherVars ...map[string]string) (string, error) {
	vars := map[string]string{
		"RANDOM": qa.RandomName("t"),
	}
	if len(otherVars) > 1 {
		return "", errors.New("cannot have more than one custom variable map")
	}
	if len(otherVars) == 1 {
		for k, v := range otherVars[0] {
			vars[k] = v
		}
	}
	// pullAll otherVars
	missing := 0
	var varType, varName, value string
	r := regexp.MustCompile(`{(env|var).([^{}]*)}`)
	for _, variableMatch := range r.FindAllStringSubmatch(template, -1) {
		value = ""
		varType = variableMatch[1]
		varName = variableMatch[2]
		switch varType {
		case "env":
			value = os.Getenv(varName)
		case "var":
			value = vars[varName]
		}
		if value == "" {
			t.Logf("Missing %s %s variable.", varType, varName)
			missing++
			continue
		}
		template = strings.ReplaceAll(template, `{`+varType+`.`+varName+`}`, value)
	}
	if missing > 0 {
		return "", fmt.Errorf("please set %d variables and restart", missing)
	}
	return commands.TrimLeadingWhitespace(template), nil
}

// EnvironmentTemplate asserts existence and fills in {env.VAR} & {var.RANDOM} placeholders in template
func EnvironmentTemplate(t *testing.T, template string, otherVars ...map[string]string) string {
	resp, err := environmentTemplate(t, template, otherVars...)
	if err != nil {
		t.Skipf(err.Error())
	}
	return resp
}