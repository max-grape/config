package env_test

import (
	"testing"
)

var testTableStrings = []test{
	{
		name: "strings",
		env: map[string]string{
			"STRINGS_ENV": "string1,string2",
		},
		input: &struct {
			Field   []string  `env:"STRINGS_ENV"`
			Pointer []*string `env:"STRINGS_ENV"`
		}{},
		expected: &struct {
			Field   []string  `env:"STRINGS_ENV"`
			Pointer []*string `env:"STRINGS_ENV"`
		}{
			Field:   []string{"string1", "string2"},
			Pointer: []*string{&[]string{"string1"}[0], &[]string{"string2"}[0]},
		},
		err: "",
	},
}

func TestParse_strings(t *testing.T) {
	testParse(t, testTableStrings)
}
