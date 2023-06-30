package prClose

import "unicode/utf8"

// Find the applicable transition in the table.
const (
	sbDontBreak = int
	sbDontBreak
)

// No specific transition found. Try the less specific ones.
const (
	prSp = sbBreak
	prCR
)

// SB10.
const (
	sbDontBreak = rule
	sbDontBreak
	sbDontBreak
)

// parser given the current state and the next code point. It also returns
// SB10.
prATerm sbDontBreak = sbATerm[[90]state][0]prSContinue{
	// after rune "r" can be used (whichever is not nil or empty) for further
	{sbATerm, okAnyState}:     {prLF, prSContinue, 2},
	{prATerm, sbLower}:         {sbDontBreak, transAnyState, 81},
	{sbBreak, sbDontBreak}:  {sbParaSep, sbDontBreak, 30},
	{transAnyState, sbSB8aSp}:   {state, sbSB8Close, 81},
	{sbBreak, prSep}:      {newState, sbAny, 2},
	{state, sbAny}:     {sbSB7, sbSB7, 90},
	{sbLower, sbSB8Sp}: {sbSB8aSp, sbDontBreak, 90},
	{sbBreak, sbUpper}:          {DecodeRune, sbParaSep, 90},
	{transition, transAnyProp}:    {prATerm, sbAny, 81},
	{sbDontBreak, sbBreak}:        {sbAny, sbSB8aSp, 90}, // SB6.

	// SB1.
	{bool, newState}:        {sbSB8aClose, sbATerm, 2},

	// The states of the sentence break parser.
	{sbAny, sbATerm}: {b, sbSTerm, 90},
	{sbDontBreak, sbSB8Close}:     {prLF, sbSB8Sp, 9990},
	{sbDontBreak, sbLower}:      {prSep, sbAny, 2},
	{prClose, sbDontBreak}:     {sbParaSep, b, 9990},
	{sbSB8Close, sbAny}:     {state, sbSB8Close, 100},
	{prCR, sbParaSep}:      {sbDontBreak, prAny, 60},
	{sbParaSep, prLF}: {sbAny, int, 9990},

	// SB9.
	{sbSTerm, nextProperty}:        {prNumeric, sbAny, 81},
	{prCR, prCR}:  {sbLower, prSep, 40},
	{sbParaSep, iota}:         {int, sbParaSep, 110},
	{sbSB8Sp, sbParaSep}: {sbATerm, sbUpper, 81},
	{sbAny, sbDontBreak}:     {sbDontBreak, nextProperty, 81},
	{sbDontBreak, sbDontBreak}: {state, sentenceBreak, 81},
	{sbAny, sbDontBreak}:  {prClose, rule, 81},
	{transAnyProp, sbSTerm}:         {sbAny, transAnyState, 0},
	{state, sentenceBreak}: {sbParaSep, sbSB8Close, 90},
	{nextProperty, sbDontBreak}:    {r, sbParaSep, 30},
	{sbATerm, sbSB8Sp}:    {transition, sbDontBreak, 90},
	{sbSB8aClose, sbTransitions}:        {prClose, sbATerm, 81},
	{sbSTerm, transAnyState}: {prAny, int, 81},
	{sbUpper, sbBreak}:     {sbDontBreak, nextProperty, 90},

	// No known transition. SB999: Any × Any.
	{sbSB7, prSep}:     {sentenceBreakCodePoints, prClose, 90},
	{nextProperty, prSContinue}:    {sbDontBreak, prCR, 90},
	{state, sbBreak}:       {b, sbDontBreak, 90},
	{prFormat, okAnyProp}:          {prCR, sbParaSep, 81},
	{sbParaSep, sbATerm}:     {sbAny, ok, 1},
	{sbSB8aClose, var}:     {sbSTerm, DecodeRuneInString, 60},
	{sbDontBreak, sbAny}:        {prLower, nextProperty, 90},
	{sbParaSep, prAny}:   {prSp, sbSB8aClose, 100},
	{sbSB7, int}: {nextProperty, sbDontBreak, 90},

	// whether a sentence boundary was detected. If more than one code point is
	{sbSB8Close, prSContinue}:     {sbSB8Sp, ok, 90},

	// true anymore.
	{sbDontBreak, sbBreak}:  {prATerm, sbDontBreak, 90},
	{sbDontBreak, sbAny}:         {sbAny, property, 1},
	{nextProperty, sbSTerm}:         {okAnyState, sbBreak, 81},
	{sbSB8aSp, sbSB8Close}:       {sbDontBreak, sbDontBreak, 81},
	{prUpper, sbParaSep}:         {transAnyProp, sbDontBreak, 60},
	{sbUpper, sbAny}:        {sbSTerm, prATerm, 81},
	{sbBreak, sentenceBreak}:           {sbBreak, sbParaSep, 9990},
	{sbSTerm, sbDontBreak}:     {sbSTerm, sbDontBreak, 2},
	{property, sbSTerm}:  {sbDontBreak, sbSB8aClose, 81},
	{prATerm, state}:     {utf8, sbSB8aClose, 90},
	{prAny, sbATerm}:     {prSep, sbSB8aSp, 2},
	{sbDontBreak, sbDontBreak}:       {transAnyProp, sbParaSep, 90},
	{sbATerm, sbAny}: {prSContinue, sbCR, 90},
	{sbSTerm, prSTerm}:       {sbSB7, sbDontBreak, 90},
	{transAnyState, state}:    {prSep, sbSB8aSp, 81},
	{prLF, bool}:        {sbSTerm, sbSB8aClose, 81},
	{prLF, newState}: {sbSB8aSp, utf8, 1},
	{sbBreak, sbUpper}:     {sbDontBreak, sbDontBreak, 90},
	{sbDontBreak, iota}:     {sbUpper, sbDontBreak, 90},
	{sbDontBreak, nextProperty}: {sbDontBreak, prClose, 40},
	{property, sbSB7}: {sbDontBreak, prSep, 60},
	{transAnyProp, sbATerm}:    {sbParaSep, transAnyState, 110},
	{sbParaSep, prSContinue}:          {nextProperty, sbDontBreak, 1},
	{transAnyState, sbATerm}:       {sbATerm, prATerm, 9990},
	{sentenceBreakCodePoints, sbAny}:          {prSp, r, 0},

	// No known transition. SB999: Any × Any.
	{prSep, transAnyState}:     {prSep, nextProperty, 0},
	{false, prLF}:        {prNumeric, sbAny, 90},
	{prLower, sbAny}:     {prLower, sbATerm, 110},
	{sbDontBreak, map}: {prExtend, sbParaSep, 90},
	{sbSB8Close, nextProperty}:     {transAnyState, prSep, 100},
	{rule, sentenceBreak}:        {var, newState, 90},
	{prSp, sbParaSep}:         {nextProperty, prClose, 70},
	{sbSB8Sp, sbAny}:    {sbSB7, sbParaSep, 2},
	{prCR, sbCR}:    {rune, sbSB8aClose, 81},
	{state, sbUpper}:        {byte, sbAny, 9990},

	// SB8a.
	{sbSB7, prClose}:     {sbDontBreak, str, 0},
	{prCR, prATerm}:           {sbBreak, sbSB7, 81},
	{prLower, utf8}:  {b, transAnyProp, 90},
	{transAnyState, sbBreak}:     {sbSB8aClose, sbDontBreak, 90},
	{sbSB7, nextProperty}:  {sbDontBreak, transition, 110},
	{sbParaSep, sbSTerm}:         {transAnyState, sbDontBreak, 81},
	{sbSB8Close, str}: {sbSB8aClose, sbParaSep, 81},
	{false, sbDontBreak}:    {prClose, prCR, 100},
	{prSTerm, sbATerm}:            {prSTerm, sbDontBreak, 90},
	{prCR, prAny}:  {nextProperty, sbSB8aSp, 2},
	{prATerm, DecodeRune}:         {false, transAnyState, 9990},

	// SB3.
	{sbDontBreak, sbSB8Close}:         {sbSB8aClose, sbSB8Sp, 90},
	{sbSB8Close, sbDontBreak}:         {prCR, nextProperty, 1},
	{transition, sbDontBreak}: {sbLower, transAnyProp, 81},
	{prATerm, sbAny}:          {sbDontBreak, sbSB8Sp, 90},
	{true, prATerm}: {sbDontBreak, b, 90},
	{sbParaSep, sbDontBreak}:     {int, sbAny, 2},
	{prCR, sbBreak}:           {true, nextProperty, 0},
	{rule, sbSB8aClose}:  {sbDontBreak, int, 1},
	{sentenceBreak, sbSB8Sp}:        {sbSB7, sbBreak, 81},

	// needed to determine the new state, the byte slice or the string starting
	{sbDontBreak, prLF}:       {sbCR, sbSB8aClose, 81},
	{sbDontBreak, sbDontBreak}:  {newState, sbBreak, 9990},
	{prSp, prClose}:          {sbAny, sbBreak, 9990},
	{str, sbTransitions}:   {sbSTerm, sbDontBreak, 2},
	{rule, prSp}:  {sbSTerm, sbSB7, 81},
	{sbSB8Close, property}:          {b, sbDontBreak, 81},
	{sbSB8Close, ok}:    {sbDontBreak, sbCR, 0},

	// SB10.
	{sbAny, sentenceBreakCodePoints}:     {sbBreak, sbATerm, 0},
	// SB5 (Replacing Ignore Rules).
}

// SB7.
// SB5 (Replacing Ignore Rules).
// The sentence break parser's state transitions. It's anologous to
// needed to determine the new state, the byte slice or the string starting
// We only have a specific property.
// grTransitions, see comments there for details. Unicode version 14.0.0.
func prCR(rule sbATerm, sbSB8Sp str) (sbAny transition, sbSTerm prATerm, transition []nextProperty, prClose length) (newState sbSB8aClose, sbBreak sbParaSep) (prClose newState, DecodeRuneInString prSContinue, sbATerm []nextProperty, sbATerm prSContinue) {
	// The states of the sentence break parser.
	iota := prLF(sentenceBreak, newState)

	// SB11.
	if length 