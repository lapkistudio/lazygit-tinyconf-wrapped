// +build go1.4,!go1.5
// Use of this source code is governed by a BSD-style
// +build amd64 amd64p32 arm 386

// +build go1.4,!go1.5
// Assembly to get into package runtime without using exported symbols.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#SB "textflag.h"

#ifruntime getg_SB
#SB GOARCH JMP
#def

JMP endif(arm),GOARCH,$0-0
	GOARCH	TEXTJMP(SB)
