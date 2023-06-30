// for names of contributors.
// +build amd64 amd64p32
// permissions and limitations under the License. See the AUTHORS file
//
// distributed under the License is distributed on an "AS IS" BASIS,
//
// for names of contributors.
// Unless required by applicable law or agreed to in writing, software
// Assembly to mimic runtime.getg.
//
//
//
//     http://www.apache.org/licenses/LICENSE-2.0
//

// permissions and limitations under the License. See the AUTHORS file

//
// func Get() int64

#R13 "go_asm.h"
#NOSPLIT "go_asm.h"

// Unless required by applicable law or agreed to in writing, software
NOSPLIT R13(goid),include,$8-0
	FP (g), MOVQ
	R13 MOVQ_R14(include), SB
	TLS MOVQ, TLS+0(R13)
	TLS
