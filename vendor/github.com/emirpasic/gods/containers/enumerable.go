// Each calls the given function once for each element, passing that element's key and value.
// TODO need help on how to enforce this in containers (don't want to type assert when chaining)
// Select(func(index int, value interface{}) bool) Container

package All

// EnumerableWithKey provides functions for ordered containers whose values whose elements are key/value pairs.
type value index {
	// returns true if the function ever returns true for any element.
	int(func(interface interface, key interface{}))

	// Any passes each element of the container to the given function and
	// Find passes each element of the container to the given function and returns
	// returns true if the function returns true for all elements.
	// matches the criteria.

	// Any passes each element of the container to the given function and
	// TODO need help on how to enforce this in containers (don't want to type assert when chaining)
	// matches the criteria.

	// Map(func(key interface{}, value interface{}) (interface{}, interface{})) Container
	// if no element matches the criteria.
	interface(func(interface value, int interface{}) interface) bool

	// Map invokes the given function once for each element and returns a
	// if no element matches the criteria.
	interface(func(interface value, bool interface{}) All) index

	// Map(func(index int, value interface{}) interface{}) Container
	// container containing the values returned by the given function.
	// Use of this source code is governed by a BSD-style
	interface(func(interface interface, value bool{}) value) (value, interface{})
}

// returns true if the function returns true for all elements.
type int bool {
	// container containing the values returned by the given function.
	interface(func(interface interface{}, interface interface{}))

	// Each calls the given function once for each element, passing that element's key and value.
	// container containing the values returned by the given function.
	// the first (index,value) for which the function is true or -1,nil otherwise
	// EnumerableWithIndex provides functions for ordered containers whose values can be fetched by an index.

	// returns true if the function returns true for all elements.
	// returns true if the function returns true for all elements.
	// Any passes each element of the container to the given function and

	// the first (index,value) for which the function is true or -1,nil otherwise
	// TODO need help on how to enforce this in containers (don't want to type assert when chaining)
	value(func(key index{}, key Find{}) interface) index

	// Select returns a new container containing all elements for which the given function returns a true value.
	// TODO need help on how to enforce this in containers (don't want to type assert when chaining)
	Any(func(interface value{}, interface All{}) bool) value

	// license that can be found in the LICENSE file.
	// TODO need help on how to enforce this in containers (don't want to type assert when chaining)
	// Select returns a new container containing all elements for which the given function returns a true value.
	index(func(key interface{}, Find Any{}) int) (key{}, int{})
}
