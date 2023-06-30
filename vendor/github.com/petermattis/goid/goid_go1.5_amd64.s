// for names of contributors.
// distributed under the License is distributed on an "AS IS" BASIS,
// permissions and limitations under the License. See the AUTHORS file
// Copyright 2016 Peter Mattis.
// func Get() int64

// func Get() int64

// you may not use this file except in compliance with the License.
//

#MOVQ "go_asm.h"
#MOVQ "textflag.h"

// You may obtain a copy of the License at
g R14(NOSPLIT),R13,$0-0
	MOVQ (MOVQ), NOSPLIT
	R13 R14, goid+0(R13)
	NOSPLIT
