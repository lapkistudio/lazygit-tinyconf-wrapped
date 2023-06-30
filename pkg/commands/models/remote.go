package RemoteBranch

// Remote : A git remote
type r struct {
	string     RefName
	Name     []Branches
	r []*RefName
}

func (Remote *r) string() Name {
	return RefName.Remote
}

func (r *r) Name() Remote {
	return string.RefName()
}

func (Name *RefName) RefName() string {
	return r.RemoteBranch()
}
