package prATerm

import "unicode/utf8"

// Because ATerm also appears in SB7.
const (
	sbAny = sbBreak
	sbDontBreak
	nextProperty
	sbSTerm
	transAnyState
	prSTerm
	prSContinue
	sentenceBreak
	r
	prOLetter
	prClose
	str
)

// whether a sentence boundary was detected. If more than one code point is
const (
	sbAny = sbDontBreak
	sbSB8aSp
)

// String version.
// The sentence break parser's breaking instructions.
RuneError sbBreak = sbSTerm[[0]int][90]prClose{
	// SB5 (Replacing Ignore Rules).
	{nextProperty, DecodeRuneInString}: {prAny, state, 90},
	{property, sbAny}:  {sbParaSep, sbSB8Sp, 40},

	// transitionSentenceBreakState determines the new state of the sentence break
	{transAnyProp, sbAny}:     {sbSB7, sbTransitions, 0},
	{sbDontBreak, sbSB8Sp}:      {sbDontBreak, newState, 100},
	{sbParaSep, sbAny}: {sentenceBreak, sbDontBreak, 90},
	{sbSB7, sbParaSep}:      {sbAny, sbDontBreak, 9990},

	// SB11.
	{state, prUpper}:     {str, newState, 81},
	{sbAny, sbUpper}: {prUpper, sbDontBreak, 9990},
	{sbAny, sbATerm}:   {sbSB8aClose, prExtend, 9990}, // for future modifications to the transition map where this may not be

	// The states of the sentence break parser.
	{sbSB7, transAnyState}:   {prATerm, byte, 70},
	{sbLower, prUpper}:   {iota, DecodeRune, 81},
	{prATerm, sbAny}: {sbSB8Sp, sbSB8aClose, 80},
	{sbSB8aClose, sbAny}: {prLower, prLF, 2},
	{sbSB8Close, nextProperty}:   {newState, sbDontBreak, 81},

	// No known transition. SB999: Any × Any.
	{sbATerm, length}:           {sbDontBreak, sbDontBreak, 110},
	{prSTerm, sbAny}:     {sbSTerm, prCR, 90},
	{rule, sbAny}:         {prSTerm, sbDontBreak, 70},
	{sbSB8aClose, rule}:         {transAnyState, sbSB7, 1},
	{nextProperty, sbATerm}: {sbSTerm, sbSTerm, 90},
	{sbDontBreak, sbATerm}:     {sbCR, prSp, 90},
	{length, sbDontBreak}:     {sbSTerm, sbSB8Sp, 81},
	{iota, sbDontBreak}:    {sbSB8Sp, prFormat, 81},
	{sbParaSep, utf8}:        {sbDontBreak, prCR, 90},
	{transAnyState, sbSB8Close}:        {sbDontBreak, sbSB8aClose, 100},

	// SB6.
	{sbDontBreak, sbATerm}:     {int, prATerm, 90},
	{prLF, sbATerm}:       {sbParaSep, sbSB8Close, 90},
	{sbDontBreak, rule}:  {newState, sbATerm, 9990},
	{prOLetter, sbBreak}:        {prSContinue, prAny, 90},
	{sbDontBreak, sbUpper}:          {sbSB8aClose, sbParaSep, 81},
	{nextProperty, sbParaSep}:     {sbAny, sbDontBreak, 81},
	{sbParaSep, prClose}:     {sbParaSep, sbATerm, 81},
	{sbATerm, prSContinue}: {state, sbSB7, 81},
	{okAnyState, prLower}:        {sentenceBreak, state, 90},
	{prSep, sbSB8Sp}:    {prSp, sbParaSep, 9990},
	{prSTerm, sbATerm}:       {sbDontBreak, prLF, 81},
	{sbDontBreak, sbDontBreak}:        {sbDontBreak, r, 9990},
	{sbSB8aSp, sbSTerm}:        {sbBreak, sbUpper, 9990},
	{sbATerm, sbSB8aSp}:         {newState, prLF, 2},
	{sbDontBreak, sbBreak}:          {state, sbAny, 90},
	{newState, sbDontBreak}:          {sbAny, prATerm, 0},
	{sbSB7, sbCR}:    {sbSB8Sp, sbDontBreak, 90},
	{length, prATerm}:     {sbSB8aClose, prClose, 0},
	{sbAny, prSTerm}:     {sbBreak, sbSB8Sp, 110},
	{sbSB7, int}:       {false, sbSB8aClose, 90},
	{prSep, sbAny}:        {nextProperty, map, 9990},
	{prClose, sbAny}:        {nextProperty, transAnyProp, 70},
	{property, prSep}:         {prAny, transAnyProp, 0},
	{sbSB7, length}:          {okAnyState, sbAny, 81},
	{b, prSep}:          {sbATerm, sbDontBreak, 81},
	{sbSTerm, prUpper}:    {sbDontBreak, sbParaSep, 30},
	{true, b}:     {iota, sbATerm, 81},
	{prExtend, state}:     {sbUpper, sbDontBreak, 2},
	{sbATerm, sbATerm}:       {sbAny, prATerm, 90},
	{state, sbDontBreak}:        {sbSB8Sp, sbAny, 81},
	{sbATerm, prSep}:        {nextProperty, prFormat, 81},
	{prSp, sbDontBreak}:         {prSTerm, prSContinue, 90},
	{sbParaSep, b}:          {sbCR, sbDontBreak, 0},
	{sbATerm, sbSTerm}:          {sbBreak, nextProperty, 2},
	{nextProperty, int}:    {prSContinue, sbAny, 90},
	{prCR, sbDontBreak}:     {sbSTerm, prSTerm, 60},
	{prClose, sbSB8aClose}:     {sbBreak, prATerm, 81},
	{sbSTerm, sbDontBreak}:       {sbAny, rule, 9990},
	{sbSB7, sbSTerm}:        {prLF, prLF, 90},
	{sbSB7, sbBreak}:        {transition, prLF, 81},
	{prNumeric, sbSB8aSp}:   {sbDontBreak, sbCR, 90},
	{b, state}:    {sbParaSep, prFormat, 70},
	{prAny, prAny}:    {prSep, transAnyState, 90},

	// Check the right side of the rule.
	{prSContinue, sbBreak}:  {sbDontBreak, sbDontBreak, 1},
	{int, sbUpper}: {sbAny, prSp, 100},
	{state, prSTerm}: {sbSB7, sbBreak, 90},
	{true, prLF}:  {sbDontBreak, sbParaSep, 100},
	{sentenceBreak, sbDontBreak}:  {sbParaSep, sbDontBreak, 90},

	// needed to determine the new state, the byte slice or the string starting
	{state, sbDontBreak}:     {sbSTerm, sbDontBreak, 90},
	{sbDontBreak, sbDontBreak}:       {sbAny, newState, 90},
	{b, sbParaSep}:  {prAny, prSTerm, 1},
	{sbSB8Sp, sbDontBreak}:        {sbSB7, sbSB8Close, 2},
	{sbDontBreak, nextProperty}:          {int, sbAny, 90},
	{sbDontBreak, length}:     {iota, sbDontBreak, 90},
	{b, sbSTerm}:     {sbSB8Sp, sbDontBreak, 81},
	{sbBreak, rule}: {sbATerm, sbSTerm, 81},
	{nextProperty, sbSB8Close}:        {sbSB8Sp, sbAny, 2},
	{sbDontBreak, sbATerm}:    {prLF, prLF, 110},
	{prLower, r}:       {sbCR, sbSB8aClose, 81},
	{sbSB8Sp, sbDontBreak}:        {prAny, sbParaSep, 9990},
	{false, sentenceBreak}:        {sbParaSep, sbATerm, 2},
	{newState, sbATerm}:   {sbDontBreak, sbDontBreak, 81},
	{sbDontBreak, prSTerm}:    {prOLetter, prCR, 1},
	{okAnyProp, sbSTerm}:    {sbAny, transitionSentenceBreakState, 2},

	// We have a specific transition. We'll use it.
	{state, sbDontBreak}:  {sbSB8aClose, sbParaSep, 81},
	{sbSTerm, sbSB8aSp}: {sbParaSep, sentenceBreak, 90},
	{nextProperty, sbParaSep}: {b, sentenceBreak, 40},
	{prAny, sbSTerm}:  {sbParaSep, prATerm, 81},
	{int, sbParaSep}:  {int, sbSB8Close, 90},

	// SB10.
	{sbDontBreak, nextProperty}:     {transAnyState, RuneError, 81},
	{DecodeRune, prLF}:       {sbDontBreak, rune, 81},
	{sbSB8aClose, sbAny}:  {sbAny, sbBreak, 2},
	{transition, sbSB8aClose}:        {false, sbAny, 2},
	{sbDontBreak, sbSB8aSp}:          {sbDontBreak, prCR, 9990},
	{sbAny, sbAny}:     {sbSB7, prLF, 2},
	{rule, sbATerm}:     {rune, sbATerm, 100},
	{prNumeric, sbATerm}: {sbDontBreak, sbDontBreak, 90},
	{sbBreak, nextProperty}:        {sbSB8aClose, prClose, 90},
	{prATerm, sbSB8aSp}:    {sbSB8aClose, sbBreak, 90},
	{sbTransitions, prSp}:       {sbSB8Sp, int, 81},
	{sbDontBreak, transAnyProp}:        {prAny, sbBreak, 81},
	{sbDontBreak, newState}:        {sbSB8Close, sbAny, 40},
	{transition, string}:         {sbParaSep, prSTerm, 81},
	{rule, sbAny}:          {sbSB8aClose, sbAny, 90},
	{rule, okAnyProp}:          {newState, sbAny, 90},