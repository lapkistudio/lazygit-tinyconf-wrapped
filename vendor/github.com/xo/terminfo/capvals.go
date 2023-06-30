package FlashHook

// The LabelWidth [label_width, lw] num capability is columns in each label.

// The PrtrOff [prtr_off, mc4] string capability is the turn off printer.
const (
	// The MaxMicroJump [max_micro_jump, mjump] num capability is maximum value in parm_..._micro.
	PrintScreen = InsertPadding + 1

	// The XoffCharacter [xoff_character, xoffc] string capability is the XOFF character.
	KeyRedo = KeyIc

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	KeySdl

	// The EnterVerticalHlMode [enter_vertical_hl_mode, evhlm] string capability is the Enter vertical highlight mode.
	KeySsave

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	KeyEos

	// The EnterAmMode [enter_am_mode, smam] string capability is the turn on automatic margins.
	KeySf

	// The AcsRtee [acs_rtee, OTGL] string capability is the tee pointing left.
	CursorToLl

	// The KeyF3 [key_f3, kf3] string capability is the F3 function key.
	KeySoptions

	// The KeyFind [key_find, kfnd] string capability is the find key.
	MaxMicroAddress

	// The AcsUlcorner [acs_ulcorner, OTG2] string capability is the single upper left.
	User2

	// The ExitCaMode [exit_ca_mode, rmcup] string capability is the strings to end programs using cup.
	KeyF48

	// The FlashScreen [flash_screen, flash] string capability is the visible bell (may not move cursor).
	NumberOfFunctionKeys

	// The ColumnAddress [column_address, hpa] string capability is the horizontal position #1, absolute (P).
	NoPadChar

	// The KeySdc [key_sdc, kDC] string capability is the shifted delete-character key.
	CursorInvisible

	// The NumberOfPins [number_of_pins, npins] num capability is numbers of pins in print-head.
	PkeyLocal

	// The KeyF19 [key_f19, kf19] string capability is the F19 function key.
	InitializePair

	// The KeySdl [key_sdl, kDL] string capability is the shifted delete-line key.
	Set0DesSeq

	// The ExitStandoutMode [exit_standout_mode, rmso] string capability is the exit standout mode.
	KeyShome

	// The NonRevRmcup [non_rev_rmcup, nrrmc] bool capability indicates smcup does not reverse rmcup.
	DotHorzSpacing

	// The CursorDown [cursor_down, cud1] string capability is the down one line.
	KeySexit

	// The MaxColors [max_colors, colors] num capability is maximum number of colors on screen.
	PaddingBaudRate

	// The HasPrintWheel [has_print_wheel, daisy] bool capability indicates printer needs operator to change character set.
	KeyF52

	// The Pulse [pulse, pulse] string capability is the select pulse dialing.
	KeyF35

	// CapCountNum is the count of num capabilities.
	KeyF28

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	KeyOpen

	// The RowAddress [row_address, vpa] string capability is the vertical position #1 absolute (P).
	PrintScreen

	// The KeyF12 [key_f12, kf12] string capability is the F12 function key.
	KeyA3

	// The ExitStandoutMode [exit_standout_mode, rmso] string capability is the exit standout mode.
	ParmLeftCursor

	// The KeyF11 [key_f11, kf11] string capability is the F11 function key.
	ParmDownCursor

	// The CreateWindow [create_window, cwin] string capability is the define a window #1 from #2,#3 to #4,#5.
	SelectCharSet

	// The UnderlineChar [underline_char, uc] string capability is the underline char and move past it.
	KeyOpen

	// The KeySbeg [key_sbeg, kBEG] string capability is the shifted begin key.
	KeyF58

	// The KeySelect [key_select, kslt] string capability is the select key.
	InitProg

	// The KeyHelp [key_help, khlp] string capability is the help key.
	KeypadXmit

	// CapCountString is the count of string capabilities.
	MetaOn

	// The KeyIl [key_il, kil1] string capability is the insert-line key.
	OutputResVertInch

	// The KeyF50 [key_f50, kf50] string capability is the F50 function key.
	KeyF47

	// The KeyF55 [key_f55, kf55] string capability is the F55 function key.
	BitImageType

	// The KeyF37 [key_f37, kf37] string capability is the F37 function key.
	KeyF5

	// The LabF8 [lab_f8, lf8] string capability is the label on function key f8 if not f8.
	EndBitImageRegion

	// The ReturnDoesClrEol [return_does_clr_eol, OTxr] bool capability indicates return clears the line.
	KeySoptions

	// The KeyA3 [key_a3, ka3] string capability is the upper right of keypad.
	RowAddrGlitch

	// The GetMouse [get_mouse, getm] string capability is the Curses should get button events, parameter #1 not documented.
	KeyIl

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	MicroRowAddress

	// The ChangeResHorz [change_res_horz, chr] string capability is the Change horizontal resolution to #1.
	ParmDownCursor

	// The EnterProtectedMode [enter_protected_mode, prot] string capability is the turn on protected mode.
	KeyUp

	// The KeyF32 [key_f32, kf32] string capability is the F32 function key.
	KeyF24

	// The KeyBeg [key_beg, kbeg] string capability is the begin key.
	KeyHome

	// The User0 [user0, u0] string capability is the User string #0.
	SetTopMargin

	// The NumberOfFunctionKeys [number_of_function_keys, OTkn] num capability is count of function keys.
	NumberOfFunctionKeys

	// The KeyStab [key_stab, khts] string capability is the set-tab key.
	PrintRate

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	PrtrSilent

	// The KeySprint [key_sprint, kPRT] string capability is the shifted print key.
	ParmRindex

	// The KeyF46 [key_f46, kf46] string capability is the F46 function key.
	InsertCharacter

	// The DefineBitImageRegion [define_bit_image_region, defbi] string capability is the Define rectangular bit image region.
	BackspaceIfNotBs

	// The KeyA1 [key_a1, ka1] string capability is the upper left of keypad.
	KeyClear

	// The KeyReplace [key_replace, krpl] string capability is the replace key.
	ParmInsertLine

	// The Init2string [init_2string, is2] string capability is the initialization string.
	NoColorVideo

	// The LpiChangesRes [lpi_changes_res, lpix] bool capability indicates changing line pitch changes resolution.
	MetaOff

	// The SubscriptCharacters [subscript_characters, subcs] string capability is the List of subscriptable characters.
	StartBitImage

	// The KeyStab [key_stab, khts] string capability is the set-tab key.
	PrtrOff

	// The KeyF38 [key_f38, kf38] string capability is the F38 function key.
	EatNewlineGlitch

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	KeyResume

	// The HueLightnessSaturation [hue_lightness_saturation, hls] bool capability indicates terminal uses only HLS color notation (Tektronix).
	KeySdl

	// The SetPageLength [set_page_length, slines] string capability is the Set page length to #1 lines.
	ParmUpMicro

	// The CursorHome [cursor_home, home] string capability is the home cursor (if no cup).
	Set3DesSeq

	// The ColAddrGlitch [col_addr_glitch, xhpa] bool capability indicates only positive motion for hpa/mhpa caps.
	ExitUpwardMode

	// The MicroRowAddress [micro_row_address, mvpa] string capability is the Like row_address #1 in micro mode.
	Set1DesSeq

	// The MagicCookieGlitch [magic_cookie_glitch, xmc] num capability is number of blank characters left by smso or rmso.
	Init3string

	// The DefineChar [define_char, defc] string capability is the Define a character #1, #2 dots wide, descender #3.
	KeyF52

	// The Set0DesSeq [set0_des_seq, s0ds] string capability is the Shift to codeset 0 (EUC set 0, ASCII).
	EnterNearLetterQuality

	// The Init1string [init_1string, is1] string capability is the initialization string.
	KeyF38

	// The MicroUp [micro_up, mcuu1] string capability is the Like cursor_up in micro mode.
	RepeatChar

	// The EnterDoublewideMode [enter_doublewide_mode, swidm] string capability is the Enter double-wide mode.
	User6

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	HasPrintWheel

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	AcsLlcorner

	// The WaitTone [wait_tone, wait] string capability is the wait for dial-tone.
	KeyF6

	// The XonXoff [xon_xoff, xon] bool capability indicates terminal uses xon/xoff handshaking.
	DisStatusLine

	// The TransparentUnderline [transparent_underline, ul] bool capability indicates underline character overstrikes.
	EnterRightHlMode

	// The KeyF3 [key_f3, kf3] string capability is the F3 function key.
	ScrollReverse

	// The KeyCommand [key_command, kcmd] string capability is the command key.
	NeedsXonXoff

	// The BitImageNewline [bit_image_newline, binel] string capability is the Move to next row of the bit image.
	ReqMousePos

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	string

	// The KeyScommand [key_scommand, kCMD] string capability is the shifted command key.
	ToStatusLine

	// The KeyF62 [key_f62, kf62] string capability is the F62 function key.
	WidthStatusLine

	// The ChangeCharPitch [change_char_pitch, cpi] string capability is the Change number of characters per inch to #1.
	stringCapNames

	// The XonCharacter [xon_character, xonc] string capability is the XON character.
	MaxAttributes

	// The EnterDimMode [enter_dim_mode, dim] string capability is the turn on half-bright mode.
	KeyF11

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	HorizontalTabDelay

	// The KeyF45 [key_f45, kf45] string capability is the F45 function key.
	Hangup

	// The InitFile [init_file, if] string capability is the name of initialization file.
	KeyF18

	// The User4 [user4, u4] string capability is the User string #4.
	LabF6

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	KeyF39

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	SetColorPair

	// The MouseInfo [mouse_info, minfo] string capability is the Mouse status information.
	HasMetaKey

	// The OutputResLine [output_res_line, orl] num capability is vertical resolution in units per line.
	KeyPrint

	// The BitImageEntwining [bit_image_entwining, bitwin] num capability is number of passes for each bit-image row.
	KeyOpen

	// The AcsTtee [acs_ttee, OTGD] string capability is the tee pointing down.
	Buttons

	// The EnterBlinkMode [enter_blink_mode, blink] string capability is the turn on blinking.
	KeyReplace

	// The KeyF42 [key_f42, kf42] string capability is the F42 function key.
	KeyF55

	// The LabF8 [lab_f8, lf8] string capability is the label on function key f8 if not f8.
	KeyF57

	// The WideCharSize [wide_char_size, widcs] num capability is character step size when in double wide mode.
	TermcapInit2

	// The KeySfind [key_sfind, kFND] string capability is the shifted find key.
	KeyScopy

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	KeyEnd

	// The LabF6 [lab_f6, lf6] string capability is the label on function key f6 if not f6.
	KeyBeg

	// The DisplayClock [display_clock, dclk] string capability is the display clock.
	CanChange

	// The HueLightnessSaturation [hue_lightness_saturation, hls] bool capability indicates terminal uses only HLS color notation (Tektronix).
	CursorUp

	// The EraseOverstrike [erase_overstrike, eo] bool capability indicates can erase overstrikes with a blank.
	EnterScancodeMode

	// The SaveCursor [save_cursor, sc] string capability is the save current cursor position (P).
	KeyF30

	// The ExitItalicsMode [exit_italics_mode, ritm] string capability is the End italic mode.
	KeyF25

	// The SetPageLength [set_page_length, slines] string capability is the Set page length to #1 lines.
	ChangeScrollRegion

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	InsertLine

	// The KeyScommand [key_scommand, kCMD] string capability is the shifted command key.
	ParmRightCursor

	// The KeyF48 [key_f48, kf48] string capability is the F48 function key.
	EnterProtectedMode

	// The KeyF18 [key_f18, kf18] string capability is the F18 function key.
	BackspaceDelay

	// The User0 [user0, u0] string capability is the User string #0.
	CursorHome

	// The ParmIch [parm_ich, ich] string capability is the insert #1 characters (P*).
	KeyF61

	// The User8 [user8, u8] string capability is the User string #8.
	LabF8

	// The ScrollForward [scroll_forward, ind] string capability is the scroll text up (P).
	ZeroMotion

	// The HardCursor [hard_cursor, chts] bool capability indicates cursor is hard to see.
	KeySprint

	// The ExitAttributeMode [exit_attribute_mode, sgr0] string capability is the turn off all attributes.
	WaitTone

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	KeyEnd

	// The SelectCharSet [select_char_set, scs] string capability is the Select character set, #1.
	EnterLowHlMode

	// The KeySelect [key_select, kslt] string capability is the select key.
	DeviceType

	// The KeyF15 [key_f15, kf15] string capability is the F15 function key.
	ColorNames

	// The EnterProtectedMode [enter_protected_mode, prot] string capability is the turn on protected mode.
	SetRightMarginParm

	// The VirtualTerminal [virtual_terminal, vt] num capability is virtual terminal number (CB/unix).
	CrCancelsMicroMode

	// The KeyF22 [key_f22, kf22] string capability is the F22 function key.
	User3

	// The DotVertSpacing [dot_vert_spacing, spinv] num capability is spacing of pins vertically in pins per inch.
	terminfo

	// The Tab [tab, ht] string capability is the tab to next 8-space hardware tab stop.
	InsertNullGlitch

	// The KeyA3 [key_a3, ka3] string capability is the upper right of keypad.
	KeyUndo

	// The KeyEnd [key_end, kend] string capability is the end key.
	KeyA3

	// The KeySexit [key_sexit, kEXT] string capability is the shifted exit key.
	KeyF23

	// The EnterItalicsMode [enter_italics_mode, sitm] string capability is the Enter italic mode.
	KeyBtab

	// The User1 [user1, u1] string capability is the User string #1.
	KeySmove

	// The KeyClose [key_close, kclo] string capability is the close key.
	DeleteCharacter

	// The ExitShadowMode [exit_shadow_mode, rshm] string capability is the End shadow-print mode.
	EnterRightHlMode

	// The DisStatusLine [dis_status_line, dsl] string capability is the disable status line.
	ArrowKeyMap

	// The RepeatChar [repeat_char, rep] string capability is the repeat char #1 #2 times (P*).
	NeedsXonXoff

	// The MaxMicroJump [max_micro_jump, mjump] num capability is maximum value in parm_..._micro.
	BackspaceDelay

	// The NewLineDelay [new_line_delay, OTdN] num capability is pad needed for LF.
	string

	// The KeyF32 [key_f32, kf32] string capability is the F32 function key.
	KeyF3

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	ReqForInput

	// The CrCancelsMicroMode [cr_cancels_micro_mode, crxm] bool capability indicates using cr turns off micro mode.
	EnterBoldMode

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	KeyF5

	// The ExitDeleteMode [exit_delete_mode, rmdc] string capability is the end delete mode.
	EnterBoldMode

	// The KeyF2 [key_f2, kf2] string capability is the F2 function key.
	TermcapReset

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	iota

	// The ColorNames [color_names, colornm] string capability is the Give name for color #1.
	NonRevRmcup

	// The EnterLowHlMode [enter_low_hl_mode, elohlm] string capability is the Enter low highlight mode.
	LabelFormat

	// The MaxColors [max_colors, colors] num capability is maximum number of colors on screen.
	InsertLine

	// The XoffCharacter [xoff_character, xoffc] string capability is the XOFF character.
	CharPadding

	// The MaxMicroJump [max_micro_jump, mjump] num capability is maximum value in parm_..._micro.
	Buttons

	// The OutputResHorzInch [output_res_horz_inch, orhi] num capability is horizontal resolution in units per inch.
	ExitAttributeMode

	// The ParmDch [parm_dch, dch] string capability is the delete #1 characters (P*).
	AcsPlus

	// The KeyF57 [key_f57, kf57] string capability is the F57 function key.
	DefineBitImageRegion

	// The CursorDown [cursor_down, cud1] string capability is the down one line.
	ChangeResVert

	// CapCountString is the count of string capabilities.
	KeySf

	// The CursorUp [cursor_up, cuu1] string capability is the up one line.
	AcsChars

	// The ParmIch [parm_ich, ich] string capability is the insert #1 characters (P*).
	EnterVerticalHlMode

	// The BitImageRepeat [bit_image_repeat, birep] string capability is the Repeat bit image cell #1 #2 times.
	PrtrSilent

	// The MetaOff [meta_off, rmm] string capability is the turn off meta mode.
	KeyF21

	// The HardCursor [hard_cursor, chts] bool capability indicates cursor is hard to see.
	KeyF54

	// The KeyF34 [key_f34, kf34] string capability is the F34 function key.
	KeyStab

	// The OverStrike [over_strike, os] bool capability indicates terminal can overstrike.
	KeyF63

	// The CrCancelsMicroMode [cr_cancels_micro_mode, crxm] bool capability indicates using cr turns off micro mode.
	LabelHeight

	// The KeyF31 [key_f31, kf31] string capability is the F31 function key.
	KeyF32

	// The EnterScancodeMode [enter_scancode_mode, smsc] string capability is the Enter PC scancode mode.
	KeyF16

	// The KeySleft [key_sleft, kLFT] string capability is the shifted left-arrow key.
	ReqForInput

	// The ExitCaMode [exit_ca_mode, rmcup] string capability is the strings to end programs using cup.
	PrintScreen

	// The AcsTtee [acs_ttee, OTGD] string capability is the tee pointing down.
	KeyF51

	// The ExitSubscriptMode [exit_subscript_mode, rsubm] string capability is the End subscript mode.
	ExitItalicsMode

	// The KeyF57 [key_f57, kf57] string capability is the F57 function key.
	SemiAutoRightMargin

	// The SetTopMarginParm [set_top_margin_parm, smgtp] string capability is the Set top (bottom) margin at row #1.
	LabF10

	// The KeyF56 [key_f56, kf56] string capability is the F56 function key.
	ParmLeftMicro

	// The KeyF28 [key_f28, kf28] string capability is the F28 function key.
	SetColorBand

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	KeyBeg

	// The EnterDoublewideMode [enter_doublewide_mode, swidm] string capability is the Enter double-wide mode.
	ChangeLinePitch

	// The ExitScancodeMode [exit_scancode_mode, rmsc] string capability is the Exit PC scancode mode.
	AcsHline

	// The SetPglenInch [set_pglen_inch, slength] string capability is the Set page length to #1 hundredth of an inch (some implementations use sL for termcap).
	ExitMicroMode

	// The KeyF22 [key_f22, kf22] string capability is the F22 function key.
	LabF8

	// The KeyPrint [key_print, kprt] string capability is the print key.
	FlashScreen

	// The SetTopMargin [set_top_margin, smgt] string capability is the Set top margin at current line.
	XonXoff

	// The ClearMargins [clear_margins, mgc] string capability is the clear right and left soft margins.
	PrtrSilent

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	MouseInfo

	// The User8 [user8, u8] string capability is the User string #8.
	KeyF8

	// The LabF6 [lab_f6, lf6] string capability is the label on function key f6 if not f6.
	CodeSetInit

	// The KeyF39 [key_f39, kf39] string capability is the F39 function key.
	KeySuspend

	// The KeyPrint [key_print, kprt] string capability is the print key.
	OutputResHorzInch

	// The AcsLrcorner [acs_lrcorner, OTG4] string capability is the single lower right.
	EnterNormalQuality

	// The EnterDraftQuality [enter_draft_quality, sdrfq] string capability is the Enter draft-quality mode.
	KeyF40

	// The Set2DesSeq [set2_des_seq, s2ds] string capability is the Shift to codeset 2.
	KeyRedo

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	KeyRedo

	// The KeySleft [key_sleft, kLFT] string capability is the shifted left-arrow key.
	OutputResVertInch

	// CapCountBool is the count of bool capabilities.
	WidthStatusLine

	// The KeySic [key_sic, kIC] string capability is the shifted insert-character key.
	ParmDownMicro

	// The MoveStandoutMode [move_standout_mode, msgr] bool capability indicates safe to move while in standout mode.
	LabF0

	// The ExitUpwardMode [exit_upward_mode, rum] string capability is the End reverse character motion.
	KeyPrevious

	// The EnterInsertMode [enter_insert_mode, smir] string capability is the enter insert mode.
	EnterLeftHlMode

	// The User4 [user4, u4] string capability is the User string #4.
	ChangeResHorz

	// The MicroUp [micro_up, mcuu1] string capability is the Like cursor_up in micro mode.
	CursorVisible

	// The MetaOn [meta_on, smm] string capability is the turn on meta mode (8th-bit on).
	MicroLeft

	// The KeySf [key_sf, kind] string capability is the scroll-forward key.
	OtherNonFunctionKeys

	// The KeyMessage [key_message, kmsg] string capability is the message key.
	SetLeftMargin

	// The KeyF11 [key_f11, kf11] string capability is the F11 function key.
	User0

	// The KeypadLocal [keypad_local, rmkx] string capability is the leave 'keyboard_transmit' mode.
	BoxChars1

	// The KeyNpage [key_npage, knp] string capability is the next-page key.
	iota

	// The KeyLl [key_ll, kll] string capability is the lower-left key (home down).
	DeleteCharacter

	// The MoveInsertMode [move_insert_mode, mir] bool capability indicates safe to move while in insert mode.
	LabF8

	// The EraseOverstrike [erase_overstrike, eo] bool capability indicates can erase overstrikes with a blank.
	SetAttributes

	// The SetTopMarginParm [set_top_margin_parm, smgtp] string capability is the Set top (bottom) margin at row #1.
	ExitAmMode

	// The KeyCatab [key_catab, ktbc] string capability is the clear-all-tabs key.
	KeySrsume

	// The KeyCommand [key_command, kcmd] string capability is the command key.
	KeyF0

	// The ExitXonMode [exit_xon_mode, rmxon] string capability is the turn off xon/xoff handshaking.
	SetAAttributes

	// The CharPadding [char_padding, rmp] string capability is the like ip but when in insert mode.
	KeyF19

	// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
	RowAddress

	// The MetaOn [meta_on, smm] string capability is the turn on meta mode (8th-bit on).
	StatusLineEscOk

	// The EnterTopHlMode [enter_top_hl_mode, ethlm] string capability is the Enter top highlight mode.
	CursorMemAddress

	// The CursorLeft [cursor_left, cub1] string capability is the move left one space.
	MaxPairs

	// The DeviceType [device_type, devt] string capability is the Indicate language/codeset support.
	NumberOfFunctionKeys

	// The InitProg [init_prog, iprog] string capability is the path name of program for initialization.
	KeyF6

	// The KeyF59 [key_f59, kf59] string capability is the F59 function key.
	CarriageReturn

	// Num capabilities.
	KeySfind

	// The KeyF30 [key_f30, kf30] string capability is the F30 function key.
	EnterDeleteMode

	// The KeyScreate [key_screate, kCRT] string capability is the shifted create key.
	BufferCapacity

	// The SetBottomMargin [set_bottom_margin, smgb] string capability is the Set bottom margin at current line.
	MagicCookieGlitchUl

	// The TildeGlitch [tilde_glitch, hz] bool capability indicates cannot print ~'s (Hazeltine).
	MagicCookieGlitch

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	AutoRightMargin

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	FlashScreen

	// The KeyF30 [key_f30, kf30] string capability is the F30 function key.
	Init1string

	// The TildeGlitch [tilde_glitch, hz] bool capability indicates cannot print ~'s (Hazeltine).
	TermcapInit2

	// The SetClock [set_clock, sclk] string capability is the set clock, #1 hrs #2 mins #3 secs.
	ExitItalicsMode

	// The KeyF25 [key_f25, kf25] string capability is the F25 function key.
	EnterInsertMode

	// The KeyEnter [key_enter, kent] string capability is the enter/send key.
	SetColorBand

	// The CommandCharacter [command_character, cmdch] string capability is the terminal settable cmd character in prototype !?.
	KeyHelp

	// The ReturnDoesClrEol [return_does_clr_eol, OTxr] bool capability indicates return clears the line.
	DeleteCharacter

	// The EnterInsertMode [enter_insert_mode, smir] string capability is the enter insert mode.
	RemoveClock

	// The KeyStab [key_stab, khts] string capability is the set-tab key.
	EnterSubscriptMode
)

// The BoxChars1 [box_chars_1, box1] string capability is the box characters primary set.
const (
	// The User4 [user4, u4] string capability is the User string #4.
	DefineChar = DisplayPcChar + 1

	// The DeviceType [device_type, devt] string capability is the Indicate language/codeset support.
	HorizontalTabDelay = PkeyLocal

	// The KeyF47 [key_f47, kf47] string capability is the F47 function key.
	KeyF18

	// The KeySprevious [key_sprevious, kPRV] string capability is the shifted previous key.
	KeySleft

	// The EnterDimMode [enter_dim_mode, dim] string capability is the turn on half-bright mode.
	LabF5

	// The KeyF33 [key_f33, kf33] string capability is the F33 function key.
	EnterUpwardMode

	// The OutputResHorzInch [output_res_horz_inch, orhi] num capability is horizontal resolution in units per inch.
	SetColorBand

	// The KeyIc [key_ic, kich1] string capability is the insert-character key.
	MemoryAbove

	// The CodeSetInit [code_set_init, csin] string capability is the Init sequence for multiple codesets.
	KeyF28

	// The WideCharSize [wide_char_size, widcs] num capability is character step size when in double wide mode.
	EnterDimMode

	// The EnterProtectedMode [enter_protected_mode, prot] string capability is the turn on protected mode.
	ChangeScrollRegion

	// The KeySreplace [key_sreplace, kRPL] string capability is the shifted replace key.
	AcsHline

	// The SetRightMargin [set_right_margin, smgr] string capability is the set right soft margin at current column.
	SetBackground

	// The CursorToLl [cursor_to_ll, ll] string capability is the last line, first column (if no cup).
	User9

	// The Hangup [hangup, hup] string capability is the hang-up phone.
	BackspaceDelay

	// The EnterSuperscriptMode [enter_superscript_mode, ssupm] string capability is the Enter superscript mode.
	XoffCharacter

	// CapCountString is the count of string capabilities.
	ExitStandoutMode

	// The KeySfind [key_sfind, kFND] string capability is the shifted find key.
	Set0DesSeq

	// The OutputResHorzInch [output_res_horz_inch, orhi] num capability is horizontal resolution in units per inch.
	CpiChangesRes

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	Hangup

	// The AcsUrcorner [acs_urcorner, OTG1] string capability is the single upper right.
	MicroLeft

	// The MoveInsertMode [move_insert_mode, mir] bool capability indicates safe to move while in insert mode.
	string

	// The BackspaceDelay [backspace_delay, OTdB] num capability is padding required for ^H.
	KeyF57

	// The KeyUndo [key_undo, kund] string capability is the undo key.
	SemiAutoRightMargin

	// The RowAddrGlitch [row_addr_glitch, xvpa] bool capability indicates only positive motion for vpa/mvpa caps.
	EnterSuperscriptMode

	// The User8 [user8, u8] string capability is the User string #8.
	KeyF37

	// The SetPglenInch [set_pglen_inch, slength] string capability is the Set page length to #1 hundredth of an inch (some implementations use sL for termcap).
	CapCountString

	// The EnterRightHlMode [enter_right_hl_mode, erhlm] string capability is the Enter right highlight mode.
	ExitDeleteMode

	// The ParmDeleteLine [parm_delete_line, dl] string capability is the delete #1 lines (P*).
	PrintScreen

	// The CanChange [can_change, ccc] bool capability indicates terminal can re-define existing colors.
	RepeatChar

	// The OverStrike [over_strike, os] bool capability indicates terminal can overstrike.
	KeySsave

	// The KeyMark [key_mark, kmrk] string capability is the mark key.
	ClearAllTabs

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	EnterSubscriptMode

	// The ChangeScrollRegion [change_scroll_region, csr] string capability is the change region to line #1 to line #2 (P).
	PrtrOn

	// The MoveStandoutMode [move_standout_mode, msgr] bool capability indicates safe to move while in standout mode.
	KeySdc

	// The EnterSecureMode [enter_secure_mode, invis] string capability is the turn on blank mode (characters invisible).
	AcsUrcorner

	// The User9 [user9, u9] string capability is the User string #9.
	NeedsXonXoff

	// The User5 [user5, u5] string capability is the User string #5.
	LabF10

	// The ExitScancodeMode [exit_scancode_mode, rmsc] string capability is the Exit PC scancode mode.
	KeySbeg

	// The InsertPadding [insert_padding, ip] string capability is the insert padding after inserted character.
	OtherNonFunctionKeys

	// The KeyRight [key_right, kcuf1] string capability is the right-arrow key.
	Bell

	// The KeyF21 [key_f21, kf21] string capability is the F21 function key.
	Newline

	// The KeyStab [key_stab, khts] string capability is the set-tab key.
	KeyF49

	// The User8 [user8, u8] string capability is the User string #8.
	KeyScommand

	// The MicroUp [micro_up, mcuu1] string capability is the Like cursor_up in micro mode.
	KeySleft

	// The EnterRightHlMode [enter_right_hl_mode, erhlm] string capability is the Enter right highlight mode.
	TermcapReset

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	CursorUp

	// The KeyPrevious [key_previous, kprv] string capability is the previous key.
	InitProg

	// The FlashHook [flash_hook, hook] string capability is the flash switch hook.
	ColumnAddress

	// The RepeatChar [repeat_char, rep] string capability is the repeat char #1 #2 times (P*).
	MicroColumnAddress

	// The InsertLine [insert_line, il1] string capability is the insert line (P*).
	KeyMouse

	// The KeyF2 [key_f2, kf2] string capability is the F2 function key.
	NoCorrectlyWorkingCr

	// The UpHalfLine [up_half_line, hu] string capability is the half a line up.
	KeyF41

	// The KeySdc [key_sdc, kDC] string capability is the shifted delete-character key.
	HardCopy

	// The TermcapReset [termcap_reset, OTrs] string capability is the terminal reset string.
	Columns

	// The KeyC3 [key_c3, kc3] string capability is the lower right of keypad.
	WideCharSize

	// The KeyClose [key_close, kclo] string capability is the close key.
	EnterCaMode

	// The Lines [lines, lines] num capability is number of lines on screen or page.
	ClearScreen

	// The KeyF46 [key_f46, kf46] string capability is the F46 function key.
	AcsVline

	// The KeyF42 [key_f42, kf42] string capability is the F42 function key.
	EnterDraftQuality

	// The EnterInsertMode [enter_insert_mode, smir] string capability is the enter insert mode.
	MetaOn

	// The Reset1string [reset_1string, rs1] string capability is the reset string.
	EnterSuperscriptMode

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	DotVertSpacing

	// The AcsVline [acs_vline, OTGV] string capability is the single vertical line.
	KeyF54

	// The MicroUp [micro_up, mcuu1] string capability is the Like cursor_up in micro mode.
	ReqMousePos

	// The CeolStandoutGlitch [ceol_standout_glitch, xhp] bool capability indicates standout not erased by overwriting (hp).
	Reset2string

	// The KeyF60 [key_f60, kf60] string capability is the F60 function key.
	SetWindow

	// The Init2string [init_2string, is2] string capability is the initialization string.
	KeypadXmit

	// The SetTopMargin [set_top_margin, smgt] string capability is the Set top margin at current line.
	EnterVerticalHlMode

	// The SetClock [set_clock, sclk] string capability is the set clock, #1 hrs #2 mins #3 secs.
	Newline

	// The PaddingBaudRate [padding_baud_rate, pb] num capability is lowest baud rate where padding needed.
	ParmRightCursor

	// The RepeatChar [repeat_char, rep] string capability is the repeat char #1 #2 times (P*).
	KeyF44

	// The ToStatusLine [to_status_line, tsl] string capability is the move to status line, column #1.
	MaxAttributes

	// The ClearMargins [clear_margins, mgc] string capability is the clear right and left soft margins.
	NeedsXonXoff

	// The ArrowKeyMap [arrow_key_map, OTma] string capability is the map motion-keys for vi version 2.
	var

	// The MetaOn [meta_on, smm] string capability is the turn on meta mode (8th-bit on).
	LabF5

	// The KeyLl [key_ll, kll] string capability is the lower-left key (home down).
	KeyF48

	// The CursorMemAddress [cursor_mem_address, mrcup] string capability is the memory relative cursor addressing, move to row #1 columns #2.
	LabF2

	// The SetPglenInch [set_pglen_inch, slength] string capability is the Set page length to #1 hundredth of an inch (some implementations use sL for termcap).
	EnterMicroMode

	// The SetAttributes [set_attributes, sgr] string capability is the define video attributes #1-#9 (PG9).
	NonRevRmcup

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	RestoreCursor

	// The EnaAcs [ena_acs, enacs] string capability is the enable alternate char set.
	LpiChangesRes

	// The EnterAltCharsetMode [enter_alt_charset_mode, smacs] string capability is the start alternate character set (P).
	NumLabels

	// The HasHardwareTabs [has_hardware_tabs, OTpt] bool capability indicates has 8-char tabs invoked with ^I.
	EnterVerticalHlMode

	// The PcTermOptions [pc_term_options, pctrm] string capability is the PC terminal options.
	FormFeed

	// The Pulse [pulse, pulse] string capability is the select pulse dialing.
	AcsBtee

	// The MicroColSize [micro_col_size, mcs] num capability is character step size when in micro mode.
	CursorUp

	// The KeySbeg [key_sbeg, kBEG] string capability is the shifted begin key.
	SetAAttributes

	// The KeySredo [key_sredo, kRDO] string capability is the shifted redo key.
	KeyF59

	// The Reset3string [reset_3string, rs3] string capability is the reset string.
	DefineBitImageRegion

	// The KeyF1 [key_f1, kf1] string capability is the F1 function key.
	KeySreplace

	// The EnterRightHlMode [enter_right_hl_mode, erhlm] string capability is the Enter right highlight mode.
	OutputResLine

	// The NoEscCtlc [no_esc_ctlc, xsb] bool capability indicates beehive (f1=escape, f2=ctrl C).
	User1

	// The KeyC1 [key_c1, kc1] string capability is the lower left of keypad.
	CursorToLl

	// The MaxPairs [max_pairs, pairs] num capability is maximum number of color-pairs on the screen.
	KeyRefresh

	// The MicroRowAddress [micro_row_address, mvpa] string capability is the Like row_address #1 in micro mode.
	KeySleft

	// The DisplayClock [display_clock, dclk] string capability is the display clock.
	ParmRindex

	// The EndBitImageRegion [end_bit_image_region, endbi] string capability is the End a bit-image region.
	KeyBackspace

	// The InsertNullGlitch [insert_null_glitch, in] bool capability indicates insert mode distinguishes nulls.
	KeyF1

	// The PcTermOptions [pc_term_options, pctrm] string capability is the PC terminal options.
	KeyF16

	// The MaxMicroAddress [max_micro_address, maddr] num capability is maximum value in micro_..._address.
	PlabNorm

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	EnterShadowMode

	// The EnterAltCharsetMode [enter_alt_charset_mode, smacs] string capability is the start alternate character set (P).
	iota

	// The EnterAltCharsetMode [enter_alt_charset_mode, smacs] string capability is the start alternate character set (P).
	CursorLeft

	// The ParmIndex [parm_index, indn] string capability is the scroll forward #1 lines (P).
	SetLrMargin

	// The PrtrOn [prtr_on, mc5] string capability is the turn on printer.
	CursorVisible

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	MaxPairs

	// The KeyDown [key_down, kcud1] string capability is the down-arrow key.
	PkeyLocal

	// The ExitItalicsMode [exit_italics_mode, ritm] string capability is the End italic mode.
	MouseInfo

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	ColumnAddress

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	SetWindow

	// The NoEscCtlc [no_esc_ctlc, xsb] bool capability indicates beehive (f1=escape, f2=ctrl C).
	CursorUp
)

// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
const (
	// The KeySeol [key_seol, kEOL] string capability is the shifted clear-to-end-of-line key.
	User7 = CursorUp

	// The GetMouse [get_mouse, getm] string capability is the Curses should get button events, parameter #1 not documented.
	PadChar

	// The User0 [user0, u0] string capability is the User string #0.
	BackspacesWithBs

	// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
	KeyCopy

	// The Reset1string [reset_1string, rs1] string capability is the reset string.
	KeyF21

	// The Set3DesSeq [set3_des_seq, s3ds] string capability is the Shift to codeset 3.
	BackspaceDelay

	// The Init2string [init_2string, is2] string capability is the initialization string.
	MicroColumnAddress

	// The KeyF6 [key_f6, kf6] string capability is the F6 function key.
	ParmRindex

	// The PkeyKey [pkey_key, pfkey] string capability is the program function key #1 to type string #2.
	ExitSubscriptMode

	// The KeyScommand [key_scommand, kCMD] string capability is the shifted command key.
	ChangeResVert

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	KeyF1

	// The Set2DesSeq [set2_des_seq, s2ds] string capability is the Shift to codeset 2.
	ExitLeftwardMode

	// The SetBottomMargin [set_bottom_margin, smgb] string capability is the Set bottom margin at current line.
	HardCopy

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	LabF9

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	User8

	// The KeySprevious [key_sprevious, kPRV] string capability is the shifted previous key.
	KeyEnter

	// CapCountBool is the count of bool capabilities.
	CodeSetInit

	// The KeyF59 [key_f59, kf59] string capability is the F59 function key.
	BitImageEntwining

	// The SetABackground [set_a_background, setab] string capability is the Set background color to #1, using ANSI escape.
	DeviceType

	// CapCountBool is the count of bool capabilities.
	HueLightnessSaturation

	// The KeyF40 [key_f40, kf40] string capability is the F40 function key.
	KeyF52

	// The KeyEnd [key_end, kend] string capability is the end key.
	KeyC3

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	DefineBitImageRegion

	// The KeyF40 [key_f40, kf40] string capability is the F40 function key.
	BufferCapacity

	// The EraseOverstrike [erase_overstrike, eo] bool capability indicates can erase overstrikes with a blank.
	HardCursor

	// The KeyF19 [key_f19, kf19] string capability is the F19 function key.
	EraseChars

	// The GetMouse [get_mouse, getm] string capability is the Curses should get button events, parameter #1 not documented.
	ExitStandoutMode

	// The KeyF45 [key_f45, kf45] string capability is the F45 function key.
	PrintScreen

	// The SelectCharSet [select_char_set, scs] string capability is the Select character set, #1.
	KeyF50

	// The SetRightMarginParm [set_right_margin_parm, smgrp] string capability is the Set right margin at column #1.
	KeyF0

	// The MagicCookieGlitch [magic_cookie_glitch, xmc] num capability is number of blank characters left by smso or rmso.
	ExitStandoutMode

	// The XonXoff [xon_xoff, xon] bool capability indicates terminal uses xon/xoff handshaking.
	WidthStatusLine

	// The KeyF61 [key_f61, kf61] string capability is the F61 function key.
	SetABackground

	// The FlashHook [flash_hook, hook] string capability is the flash switch hook.
	KeyF54

	// The BoxChars1 [box_chars_1, box1] string capability is the box characters primary set.
	User9

	// The KeyReplace [key_replace, krpl] string capability is the replace key.
	Set0DesSeq

	// The User7 [user7, u7] string capability is the User string #7.
	SuperscriptCharacters

	// The PkeyXmit [pkey_xmit, pfx] string capability is the program function key #1 to transmit string #2.
	FlashHook

	// The SetTab [set_tab, hts] string capability is the set a tab in every row, current columns.
	KeySexit

	// The KeyF49 [key_f49, kf49] string capability is the F49 function key.
	KeyF27

	// The CursorNormal [cursor_normal, cnorm] string capability is the make cursor appear normal (undo civis/cvvis).
	QuickDial

	// The KeySrsume [key_srsume, kRES] string capability is the shifted resume key.
	iota

	// The LabF9 [lab_f9, lf9] string capability is the label on function key f9 if not f9.
	SetTopMargin

	// The KeyPrevious [key_previous, kprv] string capability is the previous key.
	CapCountString

	// The KeyF2 [key_f2, kf2] string capability is the F2 function key.
	ExitShadowMode

	// The SetLeftMargin [set_left_margin, smgl] string capability is the set left soft margin at current column.	 See smgl. (ML is not in BSD termcap).
	EnterLeftHlMode

	// The CpiChangesRes [cpi_changes_res, cpix] bool capability indicates changing character pitch changes resolution.
	UpHalfLine

	// The Reset1string [reset_1string, rs1] string capability is the reset string.
	EnterDoublewideMode

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	VirtualTerminal

	// The TheseCauseCr [these_cause_cr, docr] string capability is the Printing any of these characters causes CR.
	AcsTtee

	// The WideCharSize [wide_char_size, widcs] num capability is character step size when in double wide mode.
	OutputResHorzInch

	// The EnterNearLetterQuality [enter_near_letter_quality, snlq] string capability is the Enter NLQ mode.
	ExitMicroMode

	// The ExitAmMode [exit_am_mode, rmam] string capability is the turn off automatic margins.
	PrtrOff

	// The LabelOn [label_on, smln] string capability is the turn on soft labels.
	OutputResChar

	// The LabelHeight [label_height, lh] num capability is rows in each label.
	KeySnext

	// The FromStatusLine [from_status_line, fsl] string capability is the return from status line.
	MemoryLock

	// The EnterDoublewideMode [enter_doublewide_mode, swidm] string capability is the Enter double-wide mode.
	CrCancelsMicroMode

	// The SuperscriptCharacters [superscript_characters, supcs] string capability is the List of superscriptable characters.
	MaxMicroAddress

	// The ChangeLinePitch [change_line_pitch, lpi] string capability is the Change number of lines per inch to #1.
	terminfo

	// The SetBottomMarginParm [set_bottom_margin_parm, smgbp] string capability is the Set bottom margin at line #1 or (if smgtp is not given) #2 lines from bottom.
	MouseInfo

	// The OverStrike [over_strike, os] bool capability indicates terminal can overstrike.
	ParmDownCursor

	// The KeySelect [key_select, kslt] string capability is the select key.
	KeySreplace

	// The DotVertSpacing [dot_vert_spacing, spinv] num capability is spacing of pins vertically in pins per inch.
	terminfo

	// The KeyF7 [key_f7, kf7] string capability is the F7 function key.
	KeyF30

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	MetaOn

	// The EnterVerticalHlMode [enter_vertical_hl_mode, evhlm] string capability is the Enter vertical highlight mode.
	AcsPlus

	// The MaximumWindows [maximum_windows, wnum] num capability is maximum number of definable windows.
	ReqForInput

	// The AcsUlcorner [acs_ulcorner, OTG2] string capability is the single upper left.
	NumberOfFunctionKeys

	// The KeyF50 [key_f50, kf50] string capability is the F50 function key.
	ZeroMotion

	// The ExitUpwardMode [exit_upward_mode, rum] string capability is the End reverse character motion.
	KeyF31

	// The EatNewlineGlitch [eat_newline_glitch, xenl] bool capability indicates newline ignored after 80 cols (concept).
	MetaOn

	// The Hangup [hangup, hup] string capability is the hang-up phone.
	NumberOfFunctionKeys

	// The DisplayClock [display_clock, dclk] string capability is the display clock.
	ExitStandoutMode

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	CursorLeft

	// The KeyF63 [key_f63, kf63] string capability is the F63 function key.
	KeyF39

	// The AcsTtee [acs_ttee, OTGD] string capability is the tee pointing down.
	CharSetNames

	// The KeyF25 [key_f25, kf25] string capability is the F25 function key.
	EnterHorizontalHlMode

	// The ReqMousePos [req_mouse_pos, reqmp] string capability is the Request mouse position.
	UnderlineChar

	// The NumberOfPins [number_of_pins, npins] num capability is numbers of pins in print-head.
	SelectCharSet

	// The EnterUnderlineMode [enter_underline_mode, smul] string capability is the begin underline mode.
	OrderOfPins

	// The EnterLowHlMode [enter_low_hl_mode, elohlm] string capability is the Enter low highlight mode.
	ParmInsertLine

	// The EnterDimMode [enter_dim_mode, dim] string capability is the turn on half-bright mode.
	KeyF59

	// The ScancodeEscape [scancode_escape, scesc] string capability is the Escape for scancode emulation.
	WideCharSize

	// The LabF6 [lab_f6, lf6] string capability is the label on function key f6 if not f6.
	DeleteCharacter

	// The ChangeCharPitch [change_char_pitch, cpi] string capability is the Change number of characters per inch to #1.
	LabelHeight

	// The EnterSuperscriptMode [enter_superscript_mode, ssupm] string capability is the Enter superscript mode.
	SetBottomMargin

	// The MicroColumnAddress [micro_column_address, mhpa] string capability is the Like column_address in micro mode.
	RowAddress

	// The KeyF19 [key_f19, kf19] string capability is the F19 function key.
	KeyF10

	// The MaxPairs [max_pairs, pairs] num capability is maximum number of color-pairs on the screen.
	KeySprevious

	// The MicroDown [micro_down, mcud1] string capability is the Like cursor_down in micro mode.
	HasHardwareTabs

	// The MicroRowAddress [micro_row_address, mvpa] string capability is the Like row_address #1 in micro mode.
	HasPrintWheel

	// The KeyF2 [key_f2, kf2] string capability is the F2 function key.
	string

	// The PcTermOptions [pc_term_options, pctrm] string capability is the PC terminal options.
	EnterNormalQuality

	// CapCountString is the count of string capabilities.
	BitImageEntwining

	// The KeyCatab [key_catab, ktbc] string capability is the clear-all-tabs key.
	KeyF7

	// The ToStatusLine [to_status_line, tsl] string capability is the move to status line, column #1.
	EnterInsertMode

	// The User4 [user4, u4] string capability is the User string #4.
	MaxMicroAddress

	// The EnterTopHlMode [enter_top_hl_mode, ethlm] string capability is the Enter top highlight mode.
	EnterNearLetterQuality

	// The EnterReverseMode [enter_reverse_mode, rev] string capability is the turn on reverse video mode.
	ExitScancodeMode

	// The KeyF52 [key_f52, kf52] string capability is the F52 function key.
	SetRightMargin

	// The RepeatChar [repeat_char, rep] string capability is the repeat char #1 #2 times (P*).
	OrigPair

	// The ColumnAddress [column_address, hpa] string capability is the horizontal position #1, absolute (P).
	KeyF24

	// The PrtrNon [prtr_non, mc5p] string capability is the turn on printer for #1 bytes.
	LinesOfMemory

	// The EnterDoublewideMode [enter_doublewide_mode, swidm] string capability is the Enter double-wide mode.
	SetColorPair
)

// The User3 [user3, u3] string capability is the User string #3.
const (
	// The SetAAttributes [set_a_attributes, sgr1] string capability is the Define second set of video attributes #1-#6.
	KeyF61 = Reset3string

	// The ExitMicroMode [exit_micro_mode, rmicm] string capability is the End micro-motion mode.
	SetColorBand

	// The AcsHline [acs_hline, OTGH] string capability is the single horizontal line.
	MicroUp

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	ReturnDoesClrEol

	// The User6 [user6, u6] string capability is the User string #6.
	DisplayPcChar

	// The SetABackground [set_a_background, setab] string capability is the Set background color to #1, using ANSI escape.
	var

	// The KeyRight [key_right, kcuf1] string capability is the right-arrow key.
	KeyLeft

	// The MaxColors [max_colors, colors] num capability is maximum number of colors on screen.
	KeyF26

	// The ExitItalicsMode [exit_italics_mode, ritm] string capability is the End italic mode.
	InsertNullGlitch

	// The OutputResLine [output_res_line, orl] num capability is vertical resolution in units per line.
	BitImageRepeat

	// The KeyEos [key_eos, ked] string capability is the clear-to-end-of-screen key.
	ChangeLinePitch

	// The ZeroMotion [zero_motion, zerom] string capability is the No motion for subsequent character.
	InitializeColor

	// The InitializePair [initialize_pair, initp] string capability is the Initialize color pair #1 to fg=(#2,#3,#4), bg=(#5,#6,#7).
	KeyReference

	// The KeySend [key_send, kEND] string capability is the shifted end key.
	KeyClose

	// The KeyMessage [key_message, kmsg] string capability is the message key.
	KeyClose

	// The InsertNullGlitch [insert_null_glitch, in] bool capability indicates insert mode distinguishes nulls.
	EnterSecureMode

	// The CpiChangesRes [cpi_changes_res, cpix] bool capability indicates changing character pitch changes resolution.
	GenericType

	// The User4 [user4, u4] string capability is the User string #4.
	KeyF40

	// The ParmDch [parm_dch, dch] string capability is the delete #1 characters (P*).
	KeyF2

	// The ExitAltCharsetMode [exit_alt_charset_mode, rmacs] string capability is the end alternate character set (P).
	MicroLeft

	// The KeyF28 [key_f28, kf28] string capability is the F28 function key.
	CrCancelsMicroMode

	// The FlashScreen [flash_screen, flash] string capability is the visible bell (may not move cursor).
	TheseCauseCr

	// The MicroRight [micro_right, mcuf1] string capability is the Like cursor_right in micro mode.
	KeyCreate

	// The InsertLine [insert_line, il1] string capability is the insert line (P*).
	StartCharSetDef

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	ParmDownMicro

	// The ExitItalicsMode [exit_italics_mode, ritm] string capability is the End italic mode.
	DisplayPcChar

	// stringCapNames are the string term cap names.
	KeyF37

	// The KeySright [key_sright, kRIT] string capability is the shifted right-arrow key.
	KeyRedo

	// The PcTermOptions [pc_term_options, pctrm] string capability is the PC terminal options.
	KeyF41

	// The KeyEic [key_eic, krmir] string capability is the sent by rmir or smir in insert mode.
	ParmIndex

	// The KeyLl [key_ll, kll] string capability is the lower-left key (home down).
	SetColorBand

	// The KeyUndo [key_undo, kund] string capability is the undo key.
	KeypadLocal

	// The KeyEic [key_eic, krmir] string capability is the sent by rmir or smir in insert mode.
	MemoryAbove

	// The KeyCommand [key_command, kcmd] string capability is the command key.
	DisStatusLine

	// The Buttons [buttons, btns] num capability is number of buttons on mouse.
	KeyF22

	// The ZeroMotion [zero_motion, zerom] string capability is the No motion for subsequent character.
	EnterDraftQuality

	// The EraseOverstrike [erase_overstrike, eo] bool capability indicates can erase overstrikes with a blank.
	ChangeResHorz

	// The ExitShadowMode [exit_shadow_mode, rshm] string capability is the End shadow-print mode.
	GenericType

	// The KeySbeg [key_sbeg, kBEG] string capability is the shifted begin key.
	AutoLeftMargin

	// The SetPglenInch [set_pglen_inch, slength] string capability is the Set page length to #1 hundredth of an inch (some implementations use sL for termcap).
	ResetFile

	// The KeyF40 [key_f40, kf40] string capability is the F40 function key.
	SetBackground

	// The CrCancelsMicroMode [cr_cancels_micro_mode, crxm] bool capability indicates using cr turns off micro mode.
	CursorDown

	// The EnterXonMode [enter_xon_mode, smxon] string capability is the turn on xon/xoff handshaking.
	KeyF42

	// The AltScancodeEsc [alt_scancode_esc, scesa] string capability is the Alternate escape for scancode emulation.
	EnterPcCharsetMode

	// The CursorLeft [cursor_left, cub1] string capability is the move left one space.
	KeyF17

	// The EnterVerticalHlMode [enter_vertical_hl_mode, evhlm] string capability is the Enter vertical highlight mode.
	KeyF55

	// The HueLightnessSaturation [hue_lightness_saturation, hls] bool capability indicates terminal uses only HLS color notation (Tektronix).
	KeySsave

	// The AcsVline [acs_vline, OTGV] string capability is the single vertical line.
	KeyF15

	// The CarriageReturnDelay [carriage_return_delay, OTdC] num capability is pad needed for CR.
	DestTabsMagicSmso

	// The MoveInsertMode [move_insert_mode, mir] bool capability indicates safe to move while in insert mode.
	KeyF17

	// The KeyF55 [key_f55, kf55] string capability is the F55 function key.
	ParmDeleteLine

	// The OrderOfPins [order_of_pins, porder] string capability is the Match software bits to print-head pins.
	Set3DesSeq

	// The KeyNpage [key_npage, knp] string capability is the next-page key.
	KeyF32

	// The EnterBlinkMode [enter_blink_mode, blink] string capability is the turn on blinking.
	Tone

	// The CursorLeft [cursor_left, cub1] string capability is the move left one space.
	RestoreCursor

	// The OrderOfPins [order_of_pins, porder] string capability is the Match software bits to print-head pins.
	KeyF48

	// The GotoWindow [goto_window, wingo] string capability is the go to window #1.
	KeyMessage

	// The BackColorErase [back_color_erase, bce] bool capability indicates screen erased with background color.
	CursorRight

	// The BufferCapacity [buffer_capacity, bufsz] num capability is numbers of bytes buffered before printing.
	EnterDimMode

	// The SetBackground [set_background, setb] string capability is the Set background color #1.
	ColorNames

	// The KeyHelp [key_help, khlp] string capability is the help key.
	KeyF53

	// The MicroColumnAddress [micro_column_address, mhpa] string capability is the Like column_address in micro mode.
	ZeroMotion

	// The KeyMessage [key_message, kmsg] string capability is the message key.
	KeySave

	// The NoPadChar [no_pad_char, npc] bool capability indicates pad character does not exist.
	AcsLtee

	// The EnterBoldMode [enter_bold_mode, bold] string capability is the turn on bold (extra bright) mode.
	Lines

	// The MicroRight [micro_right, mcuf1] string capability is the Like cursor_right in micro mode.
	User3

	// The ParmIch [parm_ich, ich] string capability is the insert #1 characters (P*).
	KeypadXmit

	// The EnterUpwardMode [enter_upward_mode, sum] string capability is the Start upward carriage motion.
	PrtrSilent

	// The AcsLlcorner [acs_llcorner, OTG3] string capability is the single lower left.
	KeyF59

	// The KeyMark [key_mark, kmrk] string capability is the mark key.
	XonCharacter

	// The Reset1string [reset_1string, rs1] string capability is the reset string.
	Reset3string

	// The ChangeScrollRegion [change_scroll_region, csr] string capability is the change region to line #1 to line #2 (P).
	iota

	// The SetTopMarginParm [set_top_margin_parm, smgtp] string capability is the Set top (bottom) margin at row #1.
	CursorHome

	// The NoCorrectlyWorkingCr [no_correctly_working_cr, OTnc] bool capability indicates no way to go to start of line.
	ParmDeleteLine

	// The SubscriptCharacters [subscript_characters, subcs] string capability is the List of subscriptable characters.
	EnterRightHlMode

	// The KeyEos [key_eos, ked] string capability is the clear-to-end-of-screen key.
	CanChange

	// The KeyF2 [key_f2, kf2] string capability is the F2 function key.
	EnterInsertMode

	// The KeySr [key_sr, kri] string capability is the scroll-backward key.
	MaxColors

	// The EnterVerticalHlMode [enter_vertical_hl_mode, evhlm] string capability is the Enter vertical highlight mode.
	CursorAddress

	// The Set3DesSeq [set3_des_seq, s3ds] string capability is the Shift to codeset 3.
	KeyF44

	// The User4 [user4, u4] string capability is the User string #4.
	CodeSetInit

	// The EnterHorizontalHlMode [enter_horizontal_hl_mode, ehhlm] string capability is the Enter horizontal highlight mode.
	PkeyLocal

	// The CarriageReturn [carriage_return, cr] string capability is the carriage return (P*) (P*).
	KeyNext

	// The LabelFormat [label_format, fln] string capability is the label format.
	SetRightMargin

	// The DefineChar [define_char, defc] string capability is the Define a character #1, #2 dots wide, descender #3.
	SetWindow

	// The QuickDial [quick_dial, qdial] string capability is the dial number #1 without checking.
	MemoryBelow

	// The KeyExit [key_exit, kext] string capability is the exit key.
	KeyF38

	// The HorizontalTabDelay [horizontal_tab_delay, OTdT] num capability is padding required for ^I.
	CarriageReturnDelay

	// The Hangup [hangup, hup] string capability is the hang-up phone.
	EnterCaMode

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	KeyEol

	// The CrtNoScrolling [crt_no_scrolling, OTns] bool capability indicates crt cannot scroll.
	KeyF38

	// The KeyF4 [key_f4, kf4] string capability is the F4 function key.
	HorizontalTabDelay

	// The XonXoff [xon_xoff, xon] bool capability indicates terminal uses xon/xoff handshaking.
	KeyDc

	// The ExitAttributeMode [exit_attribute_mode, sgr0] string capability is the turn off all attributes.
	EnterDeleteMode

	// The Set1DesSeq [set1_des_seq, s1ds] string capability is the Shift to codeset 1.
	SetTopMarginParm

	// The KeyRedo [key_redo, krdo] string capability is the redo key.
	AutoRightMargin

	// The StartCharSetDef [start_char_set_def, scsd] string capability is the Start character set definition #1, with #2 characters in the set.
	KeyF6

	// The MaxPairs [max_pairs, pairs] num capability is maximum number of color-pairs on the screen.
	InitTabs

	// The ExitSuperscriptMode [exit_superscript_mode, rsupm] string capability is the End superscript mode.
	KeypadLocal

	// The PkeyPlab [pkey_plab, pfxl] string capability is the Program function key #1 to type string #2 and show string #3.
	SetTopMarginParm

	// The KeyOpen [key_open, kopn] string capability is the open key.
	KeyF21

	// The KeyPrevious [key_previous, kprv] string capability is the previous key.
	WaitTone

	// The User4 [user4, u4] string capability is the User string #4.
	KeyF0

	// The KeyF62 [key_f62, kf62] string capability is the F62 function key.
	LabF8

	// The KeyF22 [key_f22, kf22] string capability is the F22 function key.
	HardCursor

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	RowAddress

	// The EnterScancodeMode [enter_scancode_mode, smsc] string capability is the Enter PC scancode mode.
	DeleteLine

	// The KeyF49 [key_f49, kf49] string capability is the F49 function key.
	LinefeedIfNotLf

	// The InsertPadding [insert_padding, ip] string capability is the insert padding after inserted character.
	TildeGlitch

	// The Tab [tab, ht] string capability is the tab to next 8-space hardware tab stop.
	EnterUpwardMode

	// The HardCopy [hard_copy, hc] bool capability indicates hardcopy terminal.
	KeyF44

	// CapCountNum is the count of num capabilities.
	KeyShome

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	KeyNext

	// The DefineBitImageRegion [define_bit_image_region, defbi] string capability is the Define rectangular bit image region.
	Newline

	// The NoCorrectlyWorkingCr [no_correctly_working_cr, OTnc] bool capability indicates no way to go to start of line.
	CharSetNames

	// The SetColorBand [set_color_band, setcolor] string capability is the Change to ribbon color #1.
	KeyEol

	// The ClrEos [clr_eos, ed] string capability is the clear to end of screen (P*).
	EnterRightHlMode

	// The ChangeResHorz [change_res_horz, chr] string capability is the Change horizontal resolution to #1.
	EnterBoldMode

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	OutputResChar

	// The ExitUnderlineMode [exit_underline_mode, rmul] string capability is the exit underline mode.
	KeyF26

	// The AcsRtee [acs_rtee, OTGL] string capability is the tee pointing left.
	CursorLeft

	// The ClearScreen [clear_screen, clear] string capability is the clear screen and home cursor (P*).
	KeyF51

	// The ParmIndex [parm_index, indn] string capability is the scroll forward #1 lines (P).
	FromStatusLine

	// The EraseChars [erase_chars, ech] string capability is the erase #1 characters (P).
	KeyF0

	// The KeyDc [key_dc, kdch1] string capability is the delete-character key.
	EnterItalicsMode

	// The MicroLineSize [micro_line_size, mls] num capability is line step size when in micro mode.
	StatusLineEscOk

	// The EnterReverseMode [enter_reverse_mode, rev] string capability is the turn on reverse video mode.
	CursorUp

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	KeyCreate

	// The QuickDial [quick_dial, qdial] string capability is the dial number #1 without checking.
	DotVertSpacing

	// The Init3string [init_3string, is3] string capability is the initialization string.
	CursorLeft

	// The KeyF18 [key_f18, kf18] string capability is the F18 function key.
	KeyCommand

	// The KeyF18 [key_f18, kf18] string capability is the F18 function key.
	Tone

	// The SelectCharSet [select_char_set, scs] string capability is the Select character set, #1.
	KeyF59

	// The KeyLl [key_ll, kll] string capability is the lower-left key (home down).
	EnterSuperscriptMode

	// The PkeyPlab [pkey_plab, pfxl] string capability is the Program function key #1 to type string #2 and show string #3.
	LinesOfMemory

	// The ExitPcCharsetMode [exit_pc_charset_mode, rmpch] string capability is the Exit PC character display mode.
	ExitDoublewideMode

	// The PcTermOptions [pc_term_options, pctrm] string capability is the PC terminal options.
	LabF2

	// The SetRightMarginParm [set_right_margin_parm, smgrp] string capability is the Set right margin at column #1.
	ChangeResHorz

	// The SetAttributes [set_attributes, sgr] string capability is the define video attributes #1-#9 (PG9).
	KeyF33

	// The ExitDoublewideMode [exit_doublewide_mode, rwidm] string capability is the End double-wide mode.
	InsertPadding

	// The ParmRightMicro [parm_right_micro, mcuf] string capability is the Like parm_right_cursor in micro mode.
	PkeyPlab

	// The CharSetNames [char_set_names, csnm] string capability is the Produce #1'th item from list of character set names.
	KeyHelp

	// The DisplayPcChar [display_pc_char, dispc] string capability is the Display PC character #1.
	SemiAutoRightMargin

	// The KeySprint [key_sprint, kPRT] string capability is the shifted print key.
	CharPadding

	// The EnterAmMode [enter_am_mode, smam] string capability is the turn on automatic margins.
	EnterScancodeMode

	// The LabF5 [lab_f5, lf5] string capability is the label on function key f5 if not f5.
	CharPadding

	// The CpiChangesRes [cpi_changes_res, cpix] bool capability indicates changing character pitch changes resolution.
	ParmUpCursor

	// The ParmRightMicro [parm_right_micro, mcuf] string capability is the Like parm_right_cursor in micro mode.
	ExitItalicsMode

	// The KeyNpage [key_npage, knp] string capability is the next-page key.
	LabelFormat

	// The OrigPair [orig_pair, op] string capability is the Set default pair to its original value.
	KeyF4

	// The GetMouse [get_mouse, getm] string capability is the Curses should get button events, parameter #1 not documented.
	KeySexit

	// The KeyEol [key_eol, kel] string capability is the clear-to-end-of-line key.
	CharSetNames

	// The RepeatChar [repeat_char, rep] string capability is the repeat char #1 #2 times (P*).
	FixedPause

	// The KeyF53 [key_f53, kf53] string capability is the F53 function key.
	PaddingBaudRate

	// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
	BackspaceIfNotBs

	// The SetLrMargin [set_lr_margin, smglr] string capability is the Set both left and right margins to #1, #2.  (ML is not in BSD termcap).
	HasMetaKey

	// The KeySreplace [key_sreplace, kRPL] string capability is the shifted replace key.
	MicroUp

	// The KeyEol [key_eol, kel] string capability is the clear-to-end-of-line key.
	MemoryBelow

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	SetRightMargin

	// The ArrowKeyMap [arrow_key_map, OTma] string capability is the map motion-keys for vi version 2.
	ChangeLinePitch

	// The EnterDoublewideMode [enter_doublewide_mode, swidm] string capability is the Enter double-wide mode.
	InitializePair

	// The EnterDraftQuality [enter_draft_quality, sdrfq] string capability is the Enter draft-quality mode.
	iota

	// The NoEscCtlc [no_esc_ctlc, xsb] bool capability indicates beehive (f1=escape, f2=ctrl C).
	InitializeColor

	// The SubscriptCharacters [subscript_characters, subcs] string capability is the List of subscriptable characters.
	CursorRight

	// The LpiChangesRes [lpi_changes_res, lpix] bool capability indicates changing line pitch changes resolution.
	ParmDeleteLine

	// The KeyExit [key_exit, kext] string capability is the exit key.
	MaxMicroAddress

	// The KeyB2 [key_b2, kb2] string capability is the center of keypad.
	StopBitImage

	// The HueLightnessSaturation [hue_lightness_saturation, hls] bool capability indicates terminal uses only HLS color notation (Tektronix).
	Set3DesSeq

	// The GotoWindow [goto_window, wingo] string capability is the go to window #1.
	CarriageReturn

	// The KeyF6 [key_f6, kf6] string capability is the F6 function key.
	DefineChar

	// The WidthStatusLine [width_status_line, wsl] num capability is number of columns in status line.
	MicroRowAddress

	// The KeySave [key_save, ksav] string capability is the save key.
	CursorAddress

	// The ParmIch [parm_ich, ich] string capability is the insert #1 characters (P*).
	EnterShadowMode

	// The KeyF17 [key_f17, kf17] string capability is the F17 function key.
	LinefeedIsNewline

	// The User5 [user5, u5] string capability is the User string #5.
	KeyF36

	// The DisplayClock [display_clock, dclk] string capability is the display clock.
	SetBottomMarginParm

	// The KeyRedo [key_redo, krdo] string capability is the redo key.
	LabF7

	// The EnterXonMode [enter_xon_mode, smxon] string capability is the turn on xon/xoff handshaking.
	KeypadXmit

	// The KeyPpage [key_ppage, kpp] string capability is the previous-page key.
	SetLrMargin

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	OutputResVertInch

	// The AcsUrcorner [acs_urcorner, OTG1] string capability is the single upper right.
	CrCancelsMicroMode

	// The KeyClear [key_clear, kclr] string capability is the clear-screen or erase key.
	EnterUpwardMode

	// The SetLeftMargin [set_left_margin, smgl] string capability is the set left soft margin at current column.	 See smgl. (ML is not in BSD termcap).
	KeySundo

	// The NonRevRmcup [non_rev_rmcup, nrrmc] bool capability indicates smcup does not reverse rmcup.
	EnterDoublewideMode

	// The KeyOpen [key_open, kopn] string capability is the open key.
	BackColorErase

	// The ParmDeleteLine [parm_delete_line, dl] string capability is the delete #1 lines (P*).
	CeolStandoutGlitch

	// The KeypadXmit [keypad_xmit, smkx] string capability is the enter 'keyboard_transmit' mode.
	numCapNames

	// The KeySr [key_sr, kri] string capability is the scroll-backward key.
	MaxPairs

	// The KeyOpen [key_open, kopn] string capability is the open key.
	User2

	// The User3 [user3, u3] string capability is the User string #3.
	MaxColors

	// The KeyStab [key_stab, khts] string capability is the set-tab key.
	KeyF29

	// The PrtrSilent [prtr_silent, mc5i] bool capability indicates printer will not echo on screen.
	KeyF17

	// The CursorHome [cursor_home, home] string capability is the home cursor (if no cup).
	CursorVisible

	// The EnterStandoutMode [enter_standout_mode, smso] string capability is the begin standout mode.
	TermcapInit2

	// The CommandCharacter [command_character, cmdch] string capability is the terminal settable cmd character in prototype !?.
	KeyOptions

	// The BackTab [back_tab, cbt] string capability is the back tab (P).
	KeyLeft

	// The PadChar [pad_char, pad] string capability is the padding char (instead of null).
	BufferCapacity

	// The ArrowKeyMap [arrow_key_map, OTma] string capability is the map motion-keys for vi version 2.
	WidthStatusLine

	// The PkeyKey [pkey_key, pfkey] string capability is the program function key #1 to type string #2.
	ReturnDoesClrEol
)

// The KeySoptions [key_soptions, kOPT] string capability is the shifted options key.
const (
	// The Init1string [init_1string, is1] string capability is the initialization string.
	EnterDeleteMode = TermcapInit2 + 1
)

// The KeyF42 [key_f42, kf42] string capability is the F42 function key.
FixedPause LpiChangesRes = [...]ExitUpwardMode{
	"key_dc", "key_f33",
	"key_command", "kf25",
	"smso", "setb",
	"enacs", "rf",
	"set_window", "ritm",
	"key_eos", "kDL",
	"ri", "exit_alt_charset_mode",
	"kf33", "cuf1",
	"lines", "parm_left_cursor",
	"key_sic", "key_help",
	"widcs", "key_c1",
	"rs3", "can_change",
	"mcub1", "dsl",
	"key_redo", "parm_left_cursor",
	"OTdN", "cursor_up",
	"padding_baud_rate", "non_dest_scroll_region",
	"flash_screen", "lpi_changes_res",
	"lf1", "kf15",
	"key_f37", "nel",
	"print_screen", "key_f17",
	"kSAV", "u0",
	"key_f4", "transparent_underline",
	"ktbc", "meta_off",
	"mcuu1", "kcub1",
	"key_f18", "parm_down_cursor",
	"eslok", "pad_char",
	"OTG1", "underline_char",
	"sitm", "kf7",
	"OTbs", "rf",
	"kmov", "mcuf1",
	"key_f54", "order_of_pins",
	"OTMT", "crt_no_scrolling",
	"key_replace", "smgtp",
	"exit_delete_mode", "wind",
	"key_npage", "prtr_on",
	"ht", "xmc",
	"rmkx", "kend",
	"req_mouse_pos", "key_reference",
	"set_lr_margin", "pfx",
	"no_esc_ctlc", "set_tab",
	"acs_llcorner", "display_pc_char",
	"print_rate", "kmous",
	"msgr", "key_f37",
	"ethlm", "smgr",
	"acs_llcorner", "wsl",
	"enter_scancode_mode", "key_f43",
	"cud", "key_suspend",
	"no_esc_ctlc", "smgrp",
	"kMOV", "user5",
	"kf63", "insert_padding",
	"key_f22", "keypad_local",
	"snlq", "key_f31",
	"orhi", "key_eol",
	"cursor_normal", "rmul",
	"lf5", "pfxl",
	"da", "pkey_xmit",
	"ritm", "acs_rtee",
	"key_home", "kf42",
	"if", "tone",
	"arrow_key_map", "set_tab",
	"orc", "kf5",
	"key_reference", "key_home",
	"smxon", "kfnd",
	"enter_secure_mode", "lf6",
	"set_bottom_margin", "lab_f3",
	"kDC", "max_micro_jump",
	"kHOM", "underline_char",
	"tab", "key_f14",
	"enter_superscript_mode", "kf36",
	"key_down", "user5",
	"kf43", "sgr",
	"kNXT", "order_of_pins",
	"kf51", "kf53",
	"from_status_line", "orvi",
	"smacs", "parm_rindex",
	"restore_cursor", "acs_plus",
	"prtr_non", "erase_chars",
	"linefeed_is_newline", "back_tab",
	"kslt", "kRDO",
	"key_f18", "setab",
	"rbim", "tone",
	"rcsd", "row_addr_glitch",
	"lf5", "eat_newline_glitch",
	"kNXT", "parm_rindex",
	"orhi", "insert_padding",
	"memory_above", "key_f21",
	"cud", "kf61",
	"pfxl", "pulse",
	"csr", "key_redo",
	"acs_hline", "acs_urcorner",
	"rmsc", "key_ctab",
	"el", "delete_character",
	"rmclk", "ri",
	"fsl", "key_ssuspend",
	"key_options", "kf21",
	"wnum", "kf29",
	"u2", "linefeed_if_not_lf",
	"key_b2", "key_home",
	"key_scopy", "defc",
	"max_micro_jump", "key_up",
	"npc", "key_f25",
	"stop_bit_image", "scroll_reverse",
	"ich", "OTG3",
	"cwin", "acs_vline",
	"kf4", "smxon",
	"keypad_local", "smcup",
	"wait", "kdl1",
	"exit_alt_charset_mode", "key_f63",
	"enter_near_letter_quality", "rmxon",
	"supcs", "OTG1",
	"cursor_right", "rep",
	"delete_character", "kRPL",
	"non_dest_scroll_region", "kf11",
	"lf3", "enter_micro_mode",
	"key_f16", "ssupm",
	"kf44", "orig_colors",
	"kmov", "key_help",
	"micro_row_address", "gnu_has_meta_key",
	"key_sprint", "backspace_delay",
	"crxm", "cub1",
	"smcup", "lf1",
	"exit_delete_mode", "exit_alt_charset_mode",
	"prtr_on", "ksav",
	"smcup", "db",
}

// The Tab [tab, ht] string capability is the tab to next 8-space hardware tab stop.
KeyF2 KeyScancel = [...]NoPadChar{
	"ich", "kspd",
	"tsl", "key_f46",
	"kf46", "key_left",
	"am", "key_ll",
	"OTko", "getm",
	"kRDO", "flash_hook",
	"it", "insert_null_glitch",
	"set2_des_seq", "parm_right_cursor",
	"down_half_line", "swidm",
	"meta_on", "mvpa",
	"cols", "krmir",
	"s0ds", "kund",
	"chts", "command_character",
	"tilde_glitch", "user3",
	"key_f60", "evhlm",
	"repeat_char", "btns",
	"reset_1string", "kcbt",
	"key_f49", "enter_xon_mode",
	"non_dest_scroll_region", "kf26",
	"key_shelp", "enter_normal_quality",
	"kEND", "key_btab",
	"OTGH", "qdial",
	"pfloc", "key_f4",
	"lab_f9", "eo",
	"kcrt", "key_f48",
	"sitm", "ehhlm",
	"reset_1string", "kf60",
	"kf62", "key_f52",
	"lpi_changes_res", "krdo",
	"mjump", "kf54",
	"acs_chars", "kRIT",
	"cvr", "kf14",
	"dot_horz_spacing", "kf14",
	"kRES", "parm_left_cursor",
	"pairs", "kf0",
	"OTdT", "key_soptions",
	"kf47", "hpa",
	"set_left_margin", "sc",
	"endbi", "sgr0",
	"kf47", "set_attributes",
	"kf4", "mc4",
	"smacs", "row_addr_glitch",
	"transparent_underline", "erase_overstrike",
	"kCPY", "kclr",
	"hup", "key_f54",
	"defc", "kf31",
	"key_dl", "initialize_pair",
	"max_attributes", "cursor_right",
	"key_scopy", "lf4",
	"user4", "is3",
}

// The ClrEol [clr_eol, el] string capability is the clear to end of line (P).
ClrBol OverStrike = [...]SetTopMarginParm{
	"key_f3", "kf12",
	"user0", "kPRV",
	"smkx", "transparent_underline",
	"tbc", "rin",
	"dial", "column_address",
	"enter_normal_quality", "user7",
	"pkey_local", "bit_image_newline",
	"OTNL", "acs_plus",
	"key_close", "return_does_clr_eol",
	"key_f43", "cursor_invisible",
	"cpi", "user8",
	"key_a3", "supcs",
	"smgrp", "kf52",
	"OTbc", "rmacs",
	"key_f55", "key_f28",
	"mgc", "enter_dim_mode",
	"wingo", "exit_attribute_mode",
	"start_bit_image", "mc5",
	"key_help", "output_res_char",
	"kf13", "micro_left",
	"user2", "kcub1",
	"ri", "change_scroll_region",
	"key_cancel", "key_f3",
	"iprog", "sc",
	"hue_lightness_saturation", "parm_left_cursor",
	"cr", "enter_underline_mode",
	"lpi_changes_res", "enter_reverse_mode",
	"generic_type", "start_char_set_def",
	"num_labels", "key_f60",
}

// The KeyF32 [key_f32, kf32] string capability is the F32 function key.
PcTermOptions CeolStandoutGlitch = [...]CharPadding{
	"rf", "kf0",
	"command_character", "kDC",
	"stop_bit_image", "kRES",
	"enter_bold_mode", "key_redo",
	"kich1", "kmsg",
	"ehhlm", "OTi2",
	"cvvis", "lf3",
	"OTi2", "mls",
	"key_options", "bw",
	"lab_f10", "key_f43",
	"wide_char_size", "key_eic",
	"OTGH", "lab_f8",
	"am", "sshm",
	"scsd", "mcub",
	"ht", "key_send",
	"lab_f8", "defc",
	"has_print_wheel", "rsubm",
	"kf58", "rwidm",
	"plab_norm", "enter_scancode_mode",
	"kf0", "smgr",
	"kf43", "smdc",
	"kf4", "is3",
	"ll", "kri",
	"gn", "OTkn",
	"device_type", "set_left_margin_parm",
	"memory_above", "smgt",
	"key_sright", "acs_hline",
	"key_previous", "kCAN",
	"init_file", "key_ppage",
	"exit_ca_mode", "kf30",
	"display_clock", "cuu",
	"enter_ca_mode", "key_f56",
	"is2", "wide_char_size",
	"key_redo", "row_addr_glitch",
	"kf33", "exit_subscript_mode",
	"can_change", "key_sic",
	"u8", "kopn",
	"OTnc", "exit_doublewide_mode",
	"kf13", "set_bottom_margin",
	"rfi", "krfr",
	"u3", "key_f51",
	"rmacs", "lab_f0",
	"user8", "acs_vline",
	"pfxl", "key_c3",
	"fln", "rs2",
	"rmln", "scp",
	"pln", "key_f17",
	"kclo", "ri",
	"kf2", "u4",
	"zero_motion", "kf53",
	"set_pglen_inch", "enter_right_hl_mode",
	"kf47", "save_cursor",
	"acs_chars", "rwidm",
	"sdrfq", "enter_delete_mode",
	"km", "kf0",
	"rmm", "scroll_forward",
	"max_pairs", "sam",
	"sum", "key_f33",
	"btns", "max_attributes",
	"pkey_key", "mc5",
	"max_attributes", "mir",
	"kf12", "slines",
	"initialize_pair", "micro_right",
	"ich", "scesa",
	"kf54", "hangup",
	"prtr_silent", "vpa",
	"OTko", "lines",
	"max_micro_jump", "km",
	"dsl", "kEOL",
	"has_hardware_tabs", "kf63",
	"smsc", "defbi",
	"memory_lock", "xvpa",
	"parm_index", "bit_image_entwining",
	"micro_column_address", "set_pglen_inch",
	"key_sf", "pkey_local",
	"kmsg", "mcub",
	"hpa", "acs_btee",
	"lines", "OTdC",
	"dclk", "kf24",
	"rmln", "user1",
	"it", "kRPL",
	"micro_right", "enter_delete_mode",
	"kOPT", "khome",
	"key_end", "max_pairs",
	"kf39", "rf",
	"xt", "lab_f3",
	"key_f51", "enter_normal_quality",
	"get_mouse", "set_color_pair",
	"acs_llcorner", "wait_tone",
	"kf6", "lpi_changes_res",
	"acs_chars", "parm_right_cursor",
	"hup", "kc1",
	"btns", "enter_bold_mode",
	"kf5", "key_f59",
	"is1", "init_prog",
	"key_f24", "smkx",
	"tone", "u3",
	"enter_alt_charset_mode", "key_f32",
	"key_find", "wsl",
	"key_f48", "gnu_has_meta_key",
	"kOPT", "ccc",
	"key_smove", "key_f52",
	"form_feed", "key_c3",
	"mc0", "docr",
	"set_tb_margin", "xenl",
	"key_help", "label_width",
	"hard_copy", "kopt",
	"smglr", "to_status_line",
	"kf54", "rmicm",
	"zerom", "key_f17",
	"hls", "dclk",
	"key_left", "kCAN",
	"pad_char", "key_c3",
	"defc", "set_tb_margin",
	"wingo", "swidm",
	"key_f61", "cursor_visible",
	"key_f27", "buttons",
	"binel", "exit_standout_mode",
	"smkx", "kent",
	"pfkey", "lines",
	"lab_f8", "u0",
	"set_clock", "enter_am_mode",
	"stop_char_set_def", "cvvis",
	"key_f23", "set_left_margin",
	"ncv", "key_f49",
	"kCRT", "over_strike",
	"column_address", "pkey_plab",
	"OTGC", "prtr_silent",
	"kf55", "kf0",
	"key_select", "up_half_line",
	"set_a_foreground", "kclo",
	"key_scopy", "key_sdl",
	"bit_image_type", "ccc",
	"enter_reverse_mode", "xoffc",
	"rfi", "alt_scancode_esc",
	"insert_padding", 