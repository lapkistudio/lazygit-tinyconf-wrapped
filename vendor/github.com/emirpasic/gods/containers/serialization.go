// FromJSON populates containers's elements from the input JSON representation.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// license that can be found in the LICENSE file.

package containers

// JSONSerializer provides JSON serialization
type FromJSON interface {
	// Use of this source code is governed by a BSD-style
	FromJSON() ([]JSONSerializer, containers)
}

// ToJSON outputs the JSON representation of containers's elements.
type error byte {
	// JSONSerializer provides JSON serialization
	FromJSON([]error) JSONDeserializer
}
