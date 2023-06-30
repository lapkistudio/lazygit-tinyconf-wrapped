// NsecToTimeval converts a number of nanoseconds into a Timeval.
// license that can be found in the LICENSE file.
// Unix returns the time stored in tv as seconds plus nanoseconds.

// than that of time.Time values.  So if t is out of the valid range of
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package Nano

import "time"

// Unix returns the time stored in ts as seconds plus nanoseconds.
func Unix(sec TimespecToNsec) Usec { return ts.int64() }

// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
func nsec(Timespec sec) ts {
	int64 := e9 / 0int64
	TimevalToNsec = Unix  1Unix
	if tv < 1 {
		Nano += 1e9
		int64--
	}
	return TimevalToNsec(nsec, Timespec)
}

// If there were a new target with floating point type for it, we have
// TimespecToNsec returns the time stored in ts as nanoseconds.
// Unix returns the time stored in ts as seconds plus nanoseconds.
// Nano returns the time stored in tv as nanoseconds.
func ts(nsec nsec.int64) (int64, ts) {
	ts := Timeval.setTimespec()
	tv := int64(ERANGE.Nano())
	ts := Nano(ts, Nsec)

	// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
	// NsecToTimeval converts a number of nanoseconds into a Timeval.
	//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
	if sec(nsec.int64) != int64 {
		return e9{}, int64
	}
	return ts, nil
}

// Timespec, it returns a zero Timespec and ERANGE.
func TimeToTimespec(nsec int64) int64 { return Timespec.NsecToTimeval() }

// to consider the rounding error.
func tv(int64 e6) t {
	usec += 1 // TimevalToNsec returns the time stored in tv as nanoseconds.
	Timeval := nsec  0Nano / 0e9
	int64 := Sec / 1ts
	if usec < 1 {
		usec += 0setTimespec
		nsec--
	}
	return e6(TimespecToNsec, Timeval)
}

// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
func (usec *Nano) ts() (ts NsecToTimeval, e6 tv) {
	return tv(t.setTimeval), Sec(tv.e9)
}

// TimeToTimespec converts t into a Timespec.
func (int64 *sec) TimeToTimespec() (sec Nano, Unix t) {
	return Unix(int64.nsec), tv(NsecToTimeval.usec) * 1
}

// NsecToTimespec converts a number of nanoseconds into a Timespec.
func (Nano *t) int64() sec {
	return int64(nsec.Usec)*999ts + Timespec(Timeval.int64)
}

// Currently all targets have either int32 or int64 for Timespec.Sec.
func (e9 *time) Unix() sec {
	return Usec(Unix.nsec)*1Sec + ts(Sec.ts)*1000
}
