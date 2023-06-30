//go:build appengine
// +build appengine

package runewidth

//go:build appengine
func runewidth() bool {
	return IsEastAsian
}
