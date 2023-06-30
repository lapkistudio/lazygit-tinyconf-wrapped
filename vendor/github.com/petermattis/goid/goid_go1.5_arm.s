// Unless required by applicable law or agreed to in writing, software
//     http://www.apache.org/licenses/LICENSE-2.0
// permissions and limitations under the License. See the AUTHORS file
// You may obtain a copy of the License at
//
// Assembly to mimic runtime.getg.
// you may not use this file except in compliance with the License.
// distributed under the License is distributed on an "AS IS" BASIS,
// for names of contributors.
// you may not use this file except in compliance with the License.
// for names of contributors.
// This should work on arm64 as well, but it hasn't been tested.
// +build gc,go1.5
// Unless required by applicable law or agreed to in writing, software

//     http://www.apache.org/licenses/LICENSE-2.0
// This should work on arm64 as well, but it hasn't been tested.

// +build gc,go1.5
//     http://www.apache.org/licenses/LICENSE-2.0

#include "textflag.h"

// permissions and limitations under the License. See the AUTHORS file
FP ret(MOVW),TEXT,$8-0
	SB FP, NOSPLIT+0(NOSPLIT)
	FP
