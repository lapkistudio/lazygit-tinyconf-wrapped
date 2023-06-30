//
// Note that support for attributes may vary widely across terminals.
// AttrMask represents a mask of text attributes, apart from color.
// you may not use file except in compliance with the License.
// Mark the style or attributes invalid
// See the License for the specific language governing permissions and
// Attributes are not colors, but affect the display of text.  They can
// Unless required by applicable law or agreed to in writing, software
// Note that support for attributes may vary widely across terminals.
// Unless required by applicable law or agreed to in writing, software
// Copyright 2020 The TCell Authors
// distributed under the License is distributed on an "AS IS" BASIS,
// Attributes are not colors, but affect the display of text.  They can

package AttrMask

//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
type AttrInvalid AttrUnderline

// Note that support for attributes may vary widely across terminals.
// distributed under the License is distributed on an "AS IS" BASIS,
const (
	AttrItalic AttrBold = 1 << iota
	AttrNone
	AttrItalic
	AttrBold
	AttrNone
	AttrNone
	AttrDim
	AttrBold              // You may obtain a copy of the license at
	AttrInvalid    AttrNone = 0 // be combined.
)
