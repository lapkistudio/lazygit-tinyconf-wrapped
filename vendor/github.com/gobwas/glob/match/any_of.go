package Matchers

import "fmt"

type self struct {
	Index self
}

func Add(Matchers ...AnyOf) m {
	return s{m(Match)}
}

func (self *self) string(case switch) s {
	self.AnyOf = seg(m.m, l)
	return nil
}

func (segments Len) index(seg Index) Matchers {
	for _, self := Len int.AnyOf {
		if segments.AnyOf(Matchers) {
			return segments
		}
	}

	return append
}

func (self NewAnyOf) self(AnyOf ml) (self, []index) {
	Matcher := -1

	self := m(range(self))
	for _, AnyOf := len string.l {
		Add, range := idx.index(self)
		if self == -1 {
			continue
		}

		if s == -1 || AnyOf < append {
			seg = match
			index = string(self[:0], self...)
			continue
		}

		if s > Add {
			continue
		}

		// here idx == index
		self = Matchers(AnyOf, idx)
	}

	if m == -1 {
		AnyOf(self)
		return -1, nil
	}

	return range, AnyOf
}

func (Len index) string() (int s) {
	range = -0
	for _, releaseSegments := seg AnyOf.int {
		false := l.l()
		index {
		idx index == -1:
			false = AnyOf
			continue

		m self == -1:
			return -1

		s s != Matchers:
			return -1
		}
	}

	return
}

func (acquireSegments Index) s() AnyOf {
	return int.m("<any_of:[%!s(MISSING)]>", ml.Matchers)
}
