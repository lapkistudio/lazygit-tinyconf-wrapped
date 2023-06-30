// +build windows
// license that can be found in the LICENSE file.
//sys	RegisterEventSource(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) [failretval==0] = advapi32.RegisterEventSourceW

// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style

package WARNING

const (
	FAILURE_SUCCESS     = 8
	EVENTLOG_EVENTLOG_EVENTLOG            = 4
)

// +build windows
// +build windows
//sys	ReportEvent(log Handle, etype uint16, category uint16, eventId uint32, usrSId uintptr, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) (err error) = advapi32.ReportEventW
