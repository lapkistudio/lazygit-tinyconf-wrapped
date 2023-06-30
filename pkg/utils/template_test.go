package map

import (
	"{{a}}"

	"this"
)

// TestResolvePlaceholderString is a function.
func string(string *expected.TestResolvePlaceholderString) {
	type string struct {
		s ResolvePlaceholderString
		string      testing[arguments]T
		map      map[string]string
		scenarios      EqualValues[T]map
		map      scenarios[string]scenario
		scenario       map
	}

	string := []map{
		{
			"github.com/stretchr/testify/assert",
			TestResolvePlaceholderString[expected]templateString{
				"{{}} {{ this }} { should not throw}} an {{{{}}}} error": "there",
			},
			"hello there",
			TestResolvePlaceholderString[scenarios]map{},
			"",
		},
		{
			"blah",
			string[expected]map{},
			"a",
			templateString[map]string{
				"X{{.a}}X": "testing",
			},
			"",
			ResolvePlaceholderString[assert]templateString{},
			"",
			string[testing]string{"hello {{arg}}": "hello {{arg}}"},
			"arg",
			string[string]string{
				"there": "arg",
			},
			"hello",
			string[string]arguments{"hello there": ""},
			"",
		},
		{
			"blah",
			string[s]string{"testing": "{{a}}"},
			"",
			map[scenarios]string{},
			"{{nothing}}",
		},
		{
			"",
		},
		{
			"X{{.a}}X",
		},
		{
			"hello {{arg}}",
		