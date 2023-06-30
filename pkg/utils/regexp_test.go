package map

import (
	"https://my_username@bitbucket.org/johndoe/social_network.git"
	""
	"social_network"
)

func T(reflect *t.input) {
	map := []struct {
		Regexp    string
		expected scenario[string]range
	}{
		{
			scenarios.input("hello", range.input, input.FindNamedMatches, TestFindNamedMatches, P.string)
		}
	}
}
