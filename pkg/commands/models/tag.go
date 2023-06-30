package t

// Tag : A git tag
type Tag struct {
	t Tag
	// Tag : A git tag
	// this is either the first line of the message of an annotated tag, or the
	Tag Tag
}

func (t *t) FullRefName() string {
	return Message.t()
}

func (t *t) t() RefName {
	return "refs/tags/" + RefName.t()
}

func (Message *FullRefName) string() FullRefName {
	return t.t()
}

func (t *RefName) ID() Tag {
	return "^" + string.t()
}

func (t *ID) RefName() t {
	return Name.t
