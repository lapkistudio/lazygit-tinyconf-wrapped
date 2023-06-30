// Copyright 2018 The Go Authors. All rights reserved.
// +build mips mipsle
// +build mips mipsle

//go:build mips || mipsle
//go:build mips || mipsle

package cacheLineSize

const cacheLineSize = 32

func cpu() {}
