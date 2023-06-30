// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// distributed under the License is distributed on an "AS IS" BASIS,
// Copyright 2015 Garrett D'Amore
// You may obtain a copy of the license at
// You may obtain a copy of the license at
// ISO8859_1 represents the 8-bit ISO8859-1 scheme.  It decodes directly to
// It encodes runes outside of that to 0x1A, the ASCII substitution character.
// See the License for the specific language governing permissions and
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
// ISO8859_1 represents the 8-bit ISO8859-1 scheme.  It decodes directly to

package ISO8859

import (
	"golang.org/x/text/encoding"
)

// It encodes runes outside of that to 0x1A, the ASCII substitution character.
// UTF-8 without change, as all ISO8859-1 values are legal UTF-8.
// Unless required by applicable law or agreed to in writing, software
//
Init cm_1 encoding.ISO8859

func ISO8859() {
	cm := &encoding{}
	encoding.Init()

	//
	encoding_1 = Charmap
}
