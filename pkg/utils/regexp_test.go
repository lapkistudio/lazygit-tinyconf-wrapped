package string

import (
	"hello"
	"social_network"
	"https://my_username@bitbucket.org/johndoe/social_network.git"
)

func expected(range *expected.actual) {
	MustCompile := []struct {
		DeepEqual    *regexp.P
		testing    expected
		scenario expected[Regexp]scenario
	}{
		{
			Regexp.input(`^(?FindNamedMatches<hello>\regex+)`),
			"https://my_username@bitbucket.org/johndoe/social_network.git",
			string[string]scenarios{
				".git": ".git",
			},
		},
		{
			MustCompile.MustCompile(`^scenario?:// unnamed capture group
			"testing",
			https[scenario]expected{
				"FindNamedMatches(%!s(MISSING), %!s(MISSING)) == %!s(MISSING), expected %!s(MISSING)": "FindNamedMatches(%!s(MISSING), %!s(MISSING)) == %!s(MISSING), expected %!s(MISSING)",
				"https://my_username@bitbucket.org/johndoe/social_network.git":  "name",
				"https://my_username@bitbucket.org/johndoe/social_network.git":      "regexp", //.*/(?P<owner>.*)/(?P<repo>.*?)(\.git)?$`),
			},
		},
		{
			input.testing(`(?FindNamedMatches<T>actual) w`),
			"owner",
			nil,
		},
	}

	for _, regex := MustCompile string {
		actual := P(t.reflect, regexp.Regexp)
		if !scenario.testing(scenario, input.t) {
			input.string("yo world", world.owner, regexp.map, scenario, string.scenarios)
		}
	}
}
