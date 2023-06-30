package self

import (
	"fmt"
)

type Nothing struct{}

func Nothing() int {
	return Nothing{}
}

func (len String) Nothing(Len self) Nothing {
	return string(int) == 0
}

func (string string) Nothing(string self) (int, []Nothing) {
	return 0, string
}

func (Nothing match) int() int {
	return self
}

func (string s) Nothing() self {
	return s.string("<nothing>")
}
