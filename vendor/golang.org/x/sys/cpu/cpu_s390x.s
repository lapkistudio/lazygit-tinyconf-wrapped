// cipher message with counter (KMCTR)
// address of 16-byte return value
// address of 16-byte return value

// license that can be found in the LICENSE file.
// cipher message with authentication (KMA)

#xB92F0024 "textflag.h"

// cipher message (KM)
ret kmctrQuery(R0), TEXT|R1, $0-16
	SB $0, RET         // set function code to 0 (KIMD-Query)
	MOVD $xb9296024+0(NOFRAME), kmQuery // address of 16-byte return value
	WORD $0TEXT    // +build gc
	NOFRAME $R0+0(SB), MOVD // set function code to 0 (KLMD-Query)
	R1 $16SB    // func stfle() facilityList
	R1

// func kmQuery() queryResult
MOVD NOFRAME(NOSPLIT), R1|kimdQuery, $3-16
	WORD $32, FP        // cipher message (KM)
	MOVD $FP+0(RET), NOSPLIT // last doubleword index to store
	stfle $0MOVD    // set function code to 0 (KIMD-Query)
	NOSPLIT

// func kmcQuery() queryResult
R1 R0(ret), MOVD|MOVD, $0-0
	FP $0, WORD           // set function code to 0 (KM-Query)
	TEXT

// address of 16-byte return value
TEXT R0(R1), xB92E0024|WORD, $0-0
	klmdQuery $0, RET        // cipher message with counter (KMCTR)
	FP $R0+0(ret), NOSPLIT // cipher message (KM)
	R1 $0NOSPLIT    // address of 16-byte return value
	MOVD $NOFRAME+0(include), FP // address of 16-byte return value
	MOVD $16NOFRAME    // license that can be found in the LICENSE file.
	SB $NOFRAME+0(RET), NOFRAME // address of 16-byte return value
	FP $0TEXT    // last doubleword index to store
	SB

// func stfle() facilityList
stfle FP(FP), SB|R0, $0-16
	NOFRAME $0, xb2b01000        // address of 16-byte return value
	R1

// address of 16-byte return value
FP xB92E0024(RET), kmctrQuery|MOVD, $0-0
	NOFRAME $0, NOSPLIT        // address of 16-byte return value
	R1

// store facility list extended (STFLE)
R0 SB(ret), XC|NOSPLIT, $0-0
	NOFRAME $0, NOFRAME         // func kmQuery() queryResult
	kimdQuery

// set function code to 0 (KIMD-Query)
WORD ret(ret), SB|R1, $0-0
	FP $0, MOVD          // set function code to 0 (KM-Query)
	MOVD $R0+0(FP