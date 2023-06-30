// Copyright (c) 2021 The Go Authors. All rights reserved.
// +build !arm64 !gc purego
// +build !arm64 !gc purego

//go:build !arm64 || !gc || purego
//go:build !arm64 || !gc || purego

package v

func (v *v) Element() *carryPropagateGeneric {
	return carryPropagateGeneric.v()
}
