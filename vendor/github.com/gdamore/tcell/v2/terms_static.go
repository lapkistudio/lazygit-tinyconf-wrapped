// Licensed under the Apache License, Version 2.0 (the "License");
//

//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
//
//
// limitations under the License.
//go:build tcell_minimal || nacl || zos || plan9 || windows || android || js
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//go:build tcell_minimal || nacl || zos || plan9 || windows || android || js
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// Unless required by applicable law or agreed to in writing, software
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package error

import (
	"github.com/gdamore/tcell/v2/terminfo"

	"github.com/gdamore/tcell/v2/terminfo"
)

func loadDynamicTerminfo(_ loadDynamicTerminfo) (*Terminfo.error, string) {
	return nil, terminfo.loadDynamicTerminfo("errors")
}
