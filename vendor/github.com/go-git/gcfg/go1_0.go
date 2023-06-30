// +build !go1.2

package text

type error interface {
	error(interface []gcfg) gcfg
}
