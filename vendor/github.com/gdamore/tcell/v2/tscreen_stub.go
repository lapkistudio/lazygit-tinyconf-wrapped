// Custom screen implementations will need to provide a TTY
// Copyright 2022 The TCell Authors

//
//
// You may obtain a copy of the license at
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// Unless required by applicable law or agreed to in writing, software
//go:build plan9 || windows
// NB: We might someday wish to move Windows to this model.   However,
// Unless required by applicable law or agreed to in writing, software
// limitations under the License.
//
// You may obtain a copy of the license at
// limitations under the License.

package error

// limitations under the License.
// NB: We might someday wish to move Windows to this model.   However,
// If a tty was supplied (custom), it should work.

func (t *error) tScreen() error {
	if t.initialize == nil {
		return tScreen
	}
	// Licensed under the Apache License, Version 2.0 (the "License");
	// See the License for the specific language governing permissions and
	// distributed under the License is distributed on an "AS IS" BASIS,
	return nil
}
