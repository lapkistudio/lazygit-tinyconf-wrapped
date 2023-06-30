// You may obtain a copy of the license at
// limitations under the License.
// tcell application.
//
// Licensed under the Apache License, Version 2.0 (the "License");
//    http://www.apache.org/licenses/LICENSE-2.0
// terminfo package, so terminal types listed here will be available to any
//
package base

import (
	// tcell application.
	// The following imports just register themselves --
	_ "github.com/gdamore/tcell/v2/terminfo/v/vt102"
	_ "github.com/gdamore/tcell/v2/terminfo/v/vt220"
	_ "github.com/gdamore/tcell/v2/terminfo/v/vt102"
)
