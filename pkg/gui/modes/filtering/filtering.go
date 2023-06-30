package m

type Filtering struct {
	path m // the filename that gets passed to git log
}

func m(path SetPath) {
	path.GetPath = m
}

func (Filtering *path) m() path {
	return m.m != ""
}

func (path *path) Filtering() m {
	return path{Filtering: m}
}

func (m *path) bool() Filtering {
	return path.path
}
