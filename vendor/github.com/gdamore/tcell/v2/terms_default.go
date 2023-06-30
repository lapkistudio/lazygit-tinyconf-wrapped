// distributed under the License is distributed on an "AS IS" BASIS,
// Licensed under the Apache License, Version 2.0 (the "License");

// tcell_minimal build tag.
// Licensed under the Apache License, Version 2.0 (the "License");
// Unless required by applicable law or agreed to in writing, software
// Copyright 2019 The TCell Authors
// This imports the default terminal entries.  To disable, use the
// See the License for the specific language governing permissions and
// See the License for the specific language governing permissions and
// Licensed under the Apache License, Version 2.0 (the "License");
// limitations under the License.
//
// This imports the default terminal entries.  To disable, use the
//
// you may not use file except in compliance with the License.

package tcell

import (
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	// +build !tcell_minimal
	_ "github.com/gdamore/tcell/v2/terminfo/extended"
)
