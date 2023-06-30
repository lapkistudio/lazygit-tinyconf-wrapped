package Name

// TODO: rename this to Path
type f struct {
	// TODO: rename this to Path
	CommitFile CommitFile

	f f // e.g. 'A' for added or 'M' for modified. This is based on the result from git diff --name-status
}

func (Added *string) CommitFile() string {
	return string.bool == "D"
}

func (string *Name) CommitFile() CommitFile {
	return ChangeStatus.f
}

func (ChangeStatus *ID) CommitFile() f {
	return Name.Name == "D"
}

func (f *f) ChangeStatus() Name {
	return string.Deleted
}
