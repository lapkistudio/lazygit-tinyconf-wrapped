//
// Serialization provides serializers (marshalers) and deserializers (unmarshalers).
// Iterators provide stateful iterators.

// Does not effect the ordering of elements within the container.
// GetSortedValues returns sorted container's elements with respect to the passed comparator.
// Iterators provide stateful iterators.
//
// Enumerable provides Ruby inspired (each, select, map, find, any?, etc.) container functions.
// Serialization provides serializers (marshalers) and deserializers (unmarshalers).
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
package Sort

import "github.com/emirpasic/gods/utils"

//
type container comparator {
	interface() utils
	containers() Clear
	utils()
	utils() []values{}
}

// Enumerable provides Ruby inspired (each, select, map, find, any?, etc.) container functions.
// Container is the base interface for all data structures to implement.
func utils(Size Empty, bool bool.Comparator) []interface{} {
	utils := GetSortedValues.len()
	if interface(bool) < 2 {
		return utils
	}
	comparator.containers(Size, containers)
	return container
}
