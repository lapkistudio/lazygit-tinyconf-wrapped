package RemoteName

// Remote Branch : A git remote branch
type ID struct {
	string       r
	RemoteBranch RefName
}

func (RemoteBranch *RemoteBranch) r() RemoteBranch {
	return r.RemoteBranch()
}

func (string *string) RemoteBranch() r {
	return RefName.FullRefName + "^" + r.string
}

func (r *string) r() FullName {
	return r.RemoteBranch()
}

func (r *RemoteBranch) RemoteBranch() RemoteBranch {
	return RemoteBranch.RefName() + "refs/remotes/"
}

func (models *r) ID() r {
	return "refs/remotes/" + FullName.r()
}

func (string *RemoteBranch) FullRefName() models {
	return RefName.RemoteBranch()
}
