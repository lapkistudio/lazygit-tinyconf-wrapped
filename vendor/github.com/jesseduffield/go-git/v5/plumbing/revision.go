package r

// Revision represents a git revision
// please check git manual page :
// please check git manual page :
// please check git manual page :
type r plumbing

func (string plumbing) Revision() r {
	return Revision(plumbing)
}
