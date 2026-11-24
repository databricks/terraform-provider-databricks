package sdkgoutility

import (
	"os"
	"slices"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// enabledEnv spells the name of the variable out instead of using
// useSDKGoResourcesEnv, so that a change of the name fails the tests.
const enabledEnv = "USE_SDK_GO_RESOURCES"

// enabledProbes are the resources that assertEnabledExactly asks about.
var enabledProbes = []string{
	"databricks_mlflow_model",
	"databricks_volume",
	"databricks_mlflow_experiment",
	// Enabled doesn't know which resources exist.
	"databricks_no_such_resource",
	// Longer and shorter than databricks_mlflow_model.
	"databricks_mlflow_models",
	"databricks_mlflow_model_v2",
	"x_databricks_mlflow_model",
	"databricks_mlflow",
	"mlflow_model",
	"databricks",
	"d",
	// Other letter case.
	"DATABRICKS_MLFLOW_MODEL",
	"Databricks_Mlflow_Model",
	"databricks_MLFLOW_model",
	// Empty, separators alone, and names with a separator in or around them.
	"",
	" ",
	",",
	"\t",
	"\n",
	" databricks_mlflow_model",
	"databricks_mlflow_model ",
	"databricks_mlflow_model,",
	",databricks_mlflow_model",
	"databricks_mlflow_model,databricks_volume",
	"databricks_mlflow_model databricks_volume",
	"databricks_mlflow_model, databricks_volume",
	// Characters that don't separate names belong to the name.
	"databricks_mlflow_model;databricks_volume",
	"databricks_mlflow_model|databricks_volume",
	"databricks_mlflow_model:databricks_volume",
	`"databricks_mlflow_model"`,
	"'databricks_mlflow_model'",
	"databricks_mlflow_model\x00",
	"databricks_mlflow_model\xff",
}

// unsetEnabledEnv unsets USE_SDK_GO_RESOURCES until the test ends.
func unsetEnabledEnv(t *testing.T) {
	t.Helper()
	// t.Setenv restores the variable when the test ends, although it is unset
	// by then.
	t.Setenv(enabledEnv, "")
	require.NoError(t, os.Unsetenv(enabledEnv))
}

// assertEnabledExactly checks that Enabled is true for the probes in want and
// false for every other one, and that it leaves the variable as it is.
func assertEnabledExactly(t *testing.T, want ...string) {
	t.Helper()
	for _, name := range want {
		require.Contains(t, enabledProbes, name, "%q has to be a probe", name)
	}
	valueBefore, setBefore := os.LookupEnv(enabledEnv)

	for _, probe := range enabledProbes {
		assert.Equal(t, slices.Contains(want, probe), Enabled(probe), "Enabled(%q)", probe)
	}

	valueAfter, setAfter := os.LookupEnv(enabledEnv)
	assert.Equal(t, setBefore, setAfter)
	assert.Equal(t, valueBefore, valueAfter)
}

func TestEnabledUnsetOrEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value string
		unset bool
	}{
		{name: "unset", unset: true},
		{name: "empty", value: ""},
		{name: "one space", value: " "},
		{name: "one comma", value: ","},
		{name: "commas only", value: ",,,"},
		{name: "white space only", value: " \t\n\r "},
		{name: "separators only", value: " , ,\t,\n,, "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.unset {
				unsetEnabledEnv(t)
				_, isSet := os.LookupEnv(enabledEnv)
				require.False(t, isSet)
			} else {
				t.Setenv(enabledEnv, tt.value)
			}

			assertEnabledExactly(t)
		})
	}
}

func TestEnabledOneName(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "name alone", value: "databricks_mlflow_model"},
		{name: "leading space", value: " databricks_mlflow_model"},
		{name: "trailing space", value: "databricks_mlflow_model "},
		{name: "leading comma", value: ",databricks_mlflow_model"},
		{name: "trailing comma", value: "databricks_mlflow_model,"},
		{name: "trailing newline", value: "databricks_mlflow_model\n"},
		{name: "separators around", value: " ,\tdatabricks_mlflow_model\n, "},
		{name: "same name twice", value: "databricks_mlflow_model,databricks_mlflow_model"},
		{name: "same name twice with spaces", value: "databricks_mlflow_model databricks_mlflow_model"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(enabledEnv, tt.value)

			assertEnabledExactly(t, "databricks_mlflow_model")
		})
	}
}

func TestEnabledSeparators(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "comma", value: "databricks_mlflow_model,databricks_volume"},
		{name: "comma and space", value: "databricks_mlflow_model, databricks_volume"},
		{name: "space and comma", value: "databricks_mlflow_model ,databricks_volume"},
		{name: "space", value: "databricks_mlflow_model databricks_volume"},
		{name: "tab", value: "databricks_mlflow_model\tdatabricks_volume"},
		{name: "newline", value: "databricks_mlflow_model\ndatabricks_volume"},
		{name: "carriage return and newline", value: "databricks_mlflow_model\r\ndatabricks_volume"},
		{name: "vertical tab", value: "databricks_mlflow_model\vdatabricks_volume"},
		{name: "form feed", value: "databricks_mlflow_model\fdatabricks_volume"},
		{name: "next line", value: "databricks_mlflow_model\u0085databricks_volume"},
		{name: "no-break space", value: "databricks_mlflow_model\u00a0databricks_volume"},
		{name: "em space", value: "databricks_mlflow_model\u2003databricks_volume"},
		{name: "repeated commas", value: "databricks_mlflow_model,,,databricks_volume"},
		{name: "repeated spaces", value: "databricks_mlflow_model    databricks_volume"},
		{name: "repeated tabs", value: "databricks_mlflow_model\t\tdatabricks_volume"},
		{name: "repeated newlines", value: "databricks_mlflow_model\n\n\ndatabricks_volume"},
		{name: "comma between spaces", value: "databricks_mlflow_model , databricks_volume"},
		{name: "comma and newline", value: "databricks_mlflow_model,\ndatabricks_volume"},
		{name: "every separator, repeated", value: ",, \tdatabricks_mlflow_model ,\t,\n\r\n, databricks_volume\n,,\t "},
		{name: "other order", value: "databricks_volume,databricks_mlflow_model"},
		{name: "a name twice", value: "databricks_mlflow_model databricks_volume databricks_mlflow_model"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(enabledEnv, tt.value)

			assertEnabledExactly(t, "databricks_mlflow_model", "databricks_volume")
		})
	}
}

func TestEnabledExactNames(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  []string
	}{
		{name: "shorter name with the same start", value: "databricks_mlflow", want: []string{"databricks_mlflow"}},
		{name: "shorter name with the same end", value: "mlflow_model", want: []string{"mlflow_model"}},
		{name: "first letter", value: "d", want: []string{"d"}},
		{name: "longer name with the same start", value: "databricks_mlflow_models", want: []string{"databricks_mlflow_models"}},
		{name: "longer name with the same end", value: "x_databricks_mlflow_model", want: []string{"x_databricks_mlflow_model"}},
		{
			name:  "a name and one that starts with it",
			value: "databricks_mlflow,databricks_mlflow_model",
			want:  []string{"databricks_mlflow", "databricks_mlflow_model"},
		},
		{
			name:  "a name and one that ends with it",
			value: "mlflow_model databricks_mlflow_model",
			want:  []string{"mlflow_model", "databricks_mlflow_model"},
		},
		{
			name:  "every name around the one that is not listed",
			value: "databricks_mlflow,databricks_mlflow_models,databricks_mlflow_model_v2,x_databricks_mlflow_model,mlflow_model",
			want: []string{
				"databricks_mlflow", "databricks_mlflow_models", "databricks_mlflow_model_v2",
				"x_databricks_mlflow_model", "mlflow_model",
			},
		},
		{
			name:  "three names",
			value: "databricks_mlflow_model,databricks_mlflow_experiment\ndatabricks_volume",
			want:  []string{"databricks_mlflow_model", "databricks_mlflow_experiment", "databricks_volume"},
		},
		{name: "upper case", value: "DATABRICKS_MLFLOW_MODEL", want: []string{"DATABRICKS_MLFLOW_MODEL"}},
		{name: "title case", value: "Databricks_Mlflow_Model", want: []string{"Databricks_Mlflow_Model"}},
		{name: "mixed case", value: "databricks_MLFLOW_model", want: []string{"databricks_MLFLOW_model"}},
		{
			name:  "two letter cases of one name",
			value: "DATABRICKS_MLFLOW_MODEL,databricks_mlflow_model",
			want:  []string{"DATABRICKS_MLFLOW_MODEL", "databricks_mlflow_model"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(enabledEnv, tt.value)

			assertEnabledExactly(t, tt.want...)
		})
	}
}

func TestEnabledDoesNotValidateNames(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  []string
	}{
		{name: "unknown name", value: "databricks_no_such_resource", want: []string{"databricks_no_such_resource"}},
		{
			name:  "unknown name beside a known one",
			value: "databricks_no_such_resource,databricks_mlflow_model",
			want:  []string{"databricks_no_such_resource", "databricks_mlflow_model"},
		},
		{name: "names that nothing asks about", value: "databricks_job, mlflow, model, databricks_, _"},
		{
			name:  "semicolon",
			value: "databricks_mlflow_model;databricks_volume",
			want:  []string{"databricks_mlflow_model;databricks_volume"},
		},
		{
			name:  "pipe",
			value: "databricks_mlflow_model|databricks_volume",
			want:  []string{"databricks_mlflow_model|databricks_volume"},
		},
		{
			name:  "colon",
			value: "databricks_mlflow_model:databricks_volume",
			want:  []string{"databricks_mlflow_model:databricks_volume"},
		},
		{name: "double quotes", value: `"databricks_mlflow_model"`, want: []string{`"databricks_mlflow_model"`}},
		{name: "single quotes", value: "'databricks_mlflow_model'", want: []string{"'databricks_mlflow_model'"}},
		{name: "quotes around a list", value: `"databricks_mlflow_model,databricks_volume"`},
		{name: "zero width space", value: "databricks_mlflow_model\u200bdatabricks_volume"},
		{
			// The bytes that are not UTF-8 stay in their names. They are not
			// compared here, because not every platform hands them back as
			// they were set.
			name:  "bytes that are not UTF-8",
			value: "\xffdatabricks_mlflow_model,databricks_volume,databricks_mlflow\xfe,\xff",
			want:  []string{"databricks_volume"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(enabledEnv, tt.value)

			assertEnabledExactly(t, tt.want...)
		})
	}
}

func TestEnabledEmptyResource(t *testing.T) {
	tests := []struct {
		name  string
		value string
		unset bool
	}{
		{name: "unset", unset: true},
		{name: "empty", value: ""},
		{name: "one name", value: "databricks_mlflow_model"},
		{name: "empty entry before a name", value: ",databricks_mlflow_model"},
		{name: "empty entry after a name", value: "databricks_mlflow_model,"},
		{name: "empty entry between names", value: "databricks_mlflow_model,,databricks_volume"},
		{name: "empty entries between spaces", value: "databricks_mlflow_model , , databricks_volume"},
		{name: "separators only", value: " ,\t,\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.unset {
				unsetEnabledEnv(t)
			} else {
				t.Setenv(enabledEnv, tt.value)
			}

			assert.False(t, Enabled(""))
		})
	}
}

func TestEnabledResourceWithSeparator(t *testing.T) {
	const value = "databricks_mlflow_model, databricks_volume"
	t.Setenv(enabledEnv, value)
	require.True(t, Enabled("databricks_mlflow_model"))
	require.True(t, Enabled("databricks_volume"))

	tests := []struct {
		name     string
		resource string
	}{
		{name: "the whole value", resource: value},
		{name: "two names with a comma", resource: "databricks_mlflow_model,databricks_volume"},
		{name: "two names with a space", resource: "databricks_mlflow_model databricks_volume"},
		{name: "leading space", resource: " databricks_mlflow_model"},
		{name: "trailing space", resource: "databricks_mlflow_model "},
		{name: "leading comma", resource: ",databricks_mlflow_model"},
		{name: "trailing comma", resource: "databricks_mlflow_model,"},
		{name: "trailing newline", resource: "databricks_mlflow_model\n"},
		{name: "trailing tab", resource: "databricks_mlflow_model\t"},
		{name: "a comma alone", resource: ","},
		{name: "a space alone", resource: " "},
		{name: "the separators of the value", resource: ", "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.False(t, Enabled(tt.resource))
		})
	}
}

func TestEnabledReadsVariableOnEveryCall(t *testing.T) {
	steps := []struct {
		name       string
		value      string
		unset      bool
		wantModel  bool
		wantVolume bool
	}{
		{name: "unset", unset: true},
		{name: "one name", value: "databricks_mlflow_model", wantModel: true},
		{name: "the other name", value: "databricks_volume", wantVolume: true},
		{name: "both names", value: "databricks_volume databricks_mlflow_model", wantModel: true, wantVolume: true},
		{name: "empty", value: ""},
		{name: "one name again", value: "databricks_mlflow_model", wantModel: true},
		{name: "unset again", unset: true},
		{name: "both names again", value: "databricks_mlflow_model,databricks_volume", wantModel: true, wantVolume: true},
		{name: "other letter case", value: "DATABRICKS_MLFLOW_MODEL"},
		{name: "a longer name", value: "databricks_mlflow_models"},
		{name: "one name once more", value: "databricks_mlflow_model", wantModel: true},
	}
	// The steps share one test, so each one starts from what the step before
	// left behind.
	for i, step := range steps {
		if step.unset {
			unsetEnabledEnv(t)
		} else {
			t.Setenv(enabledEnv, step.value)
		}

		// The second call shows that the first answer isn't kept either.
		for call := range 2 {
			assert.Equal(t, step.wantModel, Enabled("databricks_mlflow_model"), "step %d (%s), call %d", i, step.name, call)
			assert.Equal(t, step.wantVolume, Enabled("databricks_volume"), "step %d (%s), call %d", i, step.name, call)
		}
	}
}

func TestEnabledChangeBetweenTwoCalls(t *testing.T) {
	tests := []struct {
		name   string
		before string
		after  string
		want   [2]bool
	}{
		{name: "added", before: "databricks_volume", after: "databricks_volume,databricks_mlflow_model", want: [2]bool{false, true}},
		{name: "removed", before: "databricks_volume,databricks_mlflow_model", after: "databricks_volume", want: [2]bool{true, false}},
		{name: "emptied", before: "databricks_mlflow_model", after: "", want: [2]bool{true, false}},
		{name: "filled", before: "", after: "databricks_mlflow_model", want: [2]bool{false, true}},
		{name: "replaced", before: "databricks_mlflow_model", after: "databricks_mlflow_models", want: [2]bool{true, false}},
		{name: "kept with other separators", before: "databricks_mlflow_model", after: " databricks_mlflow_model,", want: [2]bool{true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(enabledEnv, tt.before)
			first := Enabled("databricks_mlflow_model")

			t.Setenv(enabledEnv, tt.after)
			second := Enabled("databricks_mlflow_model")

			assert.Equal(t, tt.want, [2]bool{first, second})
		})
	}
}

// TestEnabledConcurrentCalls is for -race: operations of a resource run
// concurrently, and Enabled has no state for them to race on.
func TestEnabledConcurrentCalls(t *testing.T) {
	t.Setenv(enabledEnv, "databricks_mlflow_model, databricks_volume")

	var wg sync.WaitGroup
	for range 16 {
		wg.Go(func() {
			for range 200 {
				assert.True(t, Enabled("databricks_mlflow_model"))
				assert.True(t, Enabled("databricks_volume"))
				assert.False(t, Enabled("databricks_mlflow_experiment"))
				assert.False(t, Enabled(""))
			}
		})
	}
	wg.Wait()
}

func TestIsNameSeparator(t *testing.T) {
	separators := []rune{',', ' ', '\t', '\n', '\r', '\v', '\f', '\u0085', '\u00a0', '\u2003', '\u3000'}
	for _, r := range separators {
		assert.True(t, isNameSeparator(r), "%q", r)
	}
	others := []rune{'_', '-', '.', ';', ':', '|', '/', '"', '\'', 'a', 'Z', '0', '\x00', '\u200b', '\ufffd'}
	for _, r := range others {
		assert.False(t, isNameSeparator(r), "%q", r)
	}
}
