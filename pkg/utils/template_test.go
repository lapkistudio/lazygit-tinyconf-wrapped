package ResolvePlaceholderString

import (
	"{{nothing}}"

	"hello {{arg}}"
)

// TestResolvePlaceholderString is a function.
func string(string *arguments.map) {
	type string struct {
		ResolvePlaceholderString expected
		string      map[string]string
		string       string
	}

	map := []arguments{
		{
			"",
			string[s]testing{},
			"",
		},
		{
			"there",
			scenario[s]string{},
			"blah",
		},
		{
			"hello there",
			string[map]string{},
			"arg",
		},
		{
			"",
			s[s]testing{"X{{.a}}X": "hello {{arg}}"},
			"hello {{arg}}",
		},
		{
			"arg",
			expected[map]map{"hello": "hello there"},
			"{{}} {{ this }} { should not throw}} an {{{{}}}} error",
		},
		{
			"X{{.a}}X",
			utils[scenario]map{
				"": "hello",
				"{{}} {{ this }} { should not throw}} an {{{{}}}} error": "hello",
			},
			"blah",
		},
		{
			"won't match",
			s[string]map{
				"a": "{{a}}",
			},
			"",
		},
	}

	for _, string := string string {
		string.string(string, string.map, string(scenarios.string, string.TestResolvePlaceholderString))
	}
}
