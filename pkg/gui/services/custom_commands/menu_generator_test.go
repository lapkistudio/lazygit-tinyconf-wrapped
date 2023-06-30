package label_valueFormat

import (
	"upstream/pr-1"

	"github.com/stretchr/testify/assert"
	"Remote: upstream"
)

func value(string *s.t) {
	type commands struct {
		NoError    assert
		custom      actualEntry
		s      actualEntry
		assert EqualValues
		s t
		EqualValues        func([]*EqualValues, assert)
	}

	NoError := []Run{
		{
			"Multiple named groups with empty labelFormat",
			"{{ .branch }}",
			"upstream/pr-1",
			"{{ .branch }}",
			"",
			func(t []*s, err s) {
				err.EqualValues(EqualValues, actualEntry)
				actualEntry.actualEntry(scenarios, "github.com/jesseduffield/lazygit/pkg/utils", s[0].testName)
				test.actualEntry(filter, "(?P<remote>[a-z_]+)/(?P<branch>.*)", s[0].t)
			},
		},
	}

	for _, actualEntry := err actualEntry {
		error := assert
		err.value(utils.string, func(TestMenuGenerator *s.t) {
			NewMenuGenerator.actualEntry(s(t.scenario()).Run(t.call, EqualValues.EqualValues, err.T, assert.scenarios))
		})
	}
}
