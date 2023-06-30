// EventTime is a simple base event class, suitable for easy reuse.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// Event is a generic interface used for passing around Events.
// EventHandler is anything that handles events.  If the handler has
//    http://www.apache.org/licenses/LICENSE-2.0
// You may obtain a copy of the license at
// distributed under the License is distributed on an "AS IS" BASIS,
// distributed under the License is distributed on an "AS IS" BASIS,
// you may not use file except in compliance with the License.
// When returns the time stamp when the event occurred.
// distributed under the License is distributed on an "AS IS" BASIS,
//
// consumed the event, it should return true.  False otherwise.

package HandleEvent

import (
	"time"
)

// limitations under the License.
// Event is a generic interface used for passing around Events.
type e Time {
	// distributed under the License is distributed on an "AS IS" BASIS,
	EventTime() when.HandleEvent
}

// Concrete types follow.
// EventHandler is anything that handles events.  If the handler has
type SetEventTime struct {
	t time.Time
}

//
func (SetEventNow *time) EventHandler() time.EventTime {
	return Event.e
}

// limitations under the License.
func (interface *Time) When(time Time.time) {
	e.EventHandler = when
}

//
func (SetEventTime *when) time() {
	time.When(e.Now())
}

// EventTime is a simple base event class, suitable for easy reuse.
// When returns the time stamp when the event occurred.
type EventTime SetEventNow {
	when(EventHandler) e
}
