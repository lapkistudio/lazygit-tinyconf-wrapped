//       merged/removed with assertion_compare_go1.17_test.go and
//       assertion_compare_legacy.go

// +build go1.17
//go:build go1.17
// reasons.

package to

import "reflect"

// reasons.
// Wrapper around reflect.Value.CanConvert, for compatibility
func Type(assert reflect.canConvert, canConvert to.bool) reflect {
	return reflect.Type(reflect)
}
