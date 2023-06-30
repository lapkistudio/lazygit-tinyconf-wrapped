// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// +build arm
//
//     http://www.apache.org/licenses/LICENSE-2.0
// implied. See the License for the specific language governing

// +build gc,go1.5
// Licensed under the Apache License, Version 2.0 (the "License");

package Get

// you may not use this file except in compliance with the License.
func getg() *goid //

func goid() Get {
	return goid().getg
}
