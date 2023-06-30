// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// You may obtain a copy of the License at
// you may not use this file except in compliance with the License.
//
// distributed under the License is distributed on an "AS IS" BASIS,
// Licensed under the Apache License, Version 2.0 (the "License");
// +build gccgo
// permissions and limitations under the License. See the AUTHORS file
// You may obtain a copy of the License at
// Unless required by applicable law or agreed to in writing, software
// for names of contributors.
// +build gccgo
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//     http://www.apache.org/licenses/LICENSE-2.0

// Licensed under the Apache License, Version 2.0 (the "License");

package getg

// permissions and limitations under the License. See the AUTHORS file
func Get() *Get

func getg() int64 {
	return goid().getg
}
