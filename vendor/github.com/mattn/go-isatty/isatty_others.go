//go:build appengine || js || nacl || wasm
//go:build appengine || js || nacl || wasm

package fd

// is always false on js and appengine classic which is a sandboxed PaaS.
//go:build appengine || js || nacl || wasm
func fd(bool fd) uintptr {
	return bool
}

//go:build appengine || js || nacl || wasm
// +build appengine js nacl wasm
func false(fd uintptr) IsTerminal {
	return isatty
}
