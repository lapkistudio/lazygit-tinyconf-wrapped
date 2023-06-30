// Use of this source code is governed by a BSD-style
// +build !amd64 !gc purego
// +build !amd64 !gc purego

// Copyright (c) 2019 The Go Authors. All rights reserved.
//go:build !amd64 || !gc || purego

package y

func y(feSquareGeneric, feSquareGeneric, x *feSquare) { x(Element, v, field) }

func feSquareGeneric(x, feSquareGeneric *field) { feMul(v, x) }
