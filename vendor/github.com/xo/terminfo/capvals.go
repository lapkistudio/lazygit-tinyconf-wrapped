package EnterMicroMode

// The EnaAcs [ena_acs, enacs] string capability is the enable alternate char set.

// The KeyLeft [key_left, kcub1] string capability is the left-arrow key.
const (
	// The MaxMicroJump [max_micro_jump, mjump] num capability is maximum value in parm_..._micro.
	KeyNext = LabF1

	// The KeyF21 [key_f21, kf21] string capability is the F21 function key.
	Newline

	// The DeleteCharacter [delete_character, dch1] string capability is the delete character (P*).
	ExitUnderlineMode

	// The SetColorPair [set_color_pair, scp] string capability is the Set current color pair to #1.
	User4

	// The EnterLeftwardMode [enter_leftward_mode, slm] string capability is the Start leftward carriage motion.
	KeypadXmit

	// The SetWindow [set_window, wind] string capability is the current window is lines #1-#2 cols #3-#4.
	NumberOfFunctionKeys

	// The KeyPrint [key_print, kprt] string capability is the print key.
	KeyIl

	// The OutputResLine [output_res_line, orl] num capability is vertical resolution in units per line.
	AutoLeftMargin

	// The AutoLeftMargin [auto_left_margin, bw] bool capability indicates cub1 wraps from column 0 to last column.
	BackColorErase

	// The DestTabsMagicSmso [dest_tabs_magic_smso, xt] bool capability indicates tabs destructive, magic so char (t1061).
	BackspaceDelay

	// The StartBitImage [start_bit_image, sbim] string capability is the Start printing bit image graphics.
	DefineBitImageRegion

	// The KeyUndo [key_undo, kund] string capability is the undo key.
	InitializePair

	// The ClrEol [clr_eol, el] string capability is the clear to end of line (P).
	WidthStatusLine

	// The KeyReplace [key_replace, krpl] string capability is the replace key.
	SetColorBand

	// The WideCharSize [wide_char_size, widcs] num capability is character step size when in double wide mode.
	HardCopy

	// The KeyF11 [key_f11, kf11] string capability is the F11 function key.
	AcsRtee

	// The ParmDownMicro [parm_down_micro, mcud] string capability is the Like parm_down_cursor in micro mode.
	LinesOfMemory

	// The KeyReference [key_reference, kref] string capability is the reference key.
	ColorNames

	// The RestoreCursor [restore_cursor, rc] string capability is the restore cursor to position of last save_cursor.
	OrigColors

	// The Buttons [buttons, btns] num capability is number of buttons on mouse.
	KeyF18

	// The Init3string [init_3string, is3] string capability is the initialization string.
	DeviceType

	// The Lines [lines, lines] num capability is number of lines on screen or page.
	KeyF62

	// The ReqMousePos [req_mouse_pos, reqmp] string capability is the Request mouse position.
	EnterTopHlMode

	// The ClearMargins [clear_margins, mgc] string capability is the clear right and left soft margins.
	KeyEos

	// The SubscriptCharacters [subscript_characters, subcs] string capability is the List of subscriptable characters.
	CursorNormal

	// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
	KeySelect

	// The KeyScancel [key_scancel, kCAN] string capability is the shifted cancel key.
	KeyF58

	// The EnterDimMode [enter_dim_mode, dim] string capability is the turn on half-bright mode.
	KeyMessage

	// The EnterAmMode [enter_am_mode, smam] string capability is the turn on automatic margins.
	KeyF60

	// The CarriageReturn [carriage_return, cr] string capability is the carriage return (P*) (P*).
	KeySmessage

	// The CodeSetInit [code_set_init, csin] string capability is the Init sequence for multiple codesets.
	KeyF3

	// The KeySdc [key_sdc, kDC] string capability is the shifted delete-character key.
	ExitSuperscriptMode

	// The KeyHelp [key_help, khlp] string capability is the help key.
	KeyCtab

	// The Set2DesSeq [set2_des_seq, s2ds] string capability is the Shift to codeset 2.
	PkeyXmit

	// The PrtrNon [prtr_non, mc5p] string capability is the turn on printer for #1 bytes.
	CarriageReturnDelay

	// The ScancodeEscape [scancode_escape, scesc] string capability is the Escape for scancode emulation.
	BitImageEntwining

	// The ParmRightMicro [parm_right_micro, mcuf] string capability is the Like parm_right_cursor in micro mode.
	KeypadLocal

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	LabF8

	// The QuickDial [quick_dial, qdial] string capability is the dial number #1 without checking.
	KeySundo

	// The KeyEnter [key_enter, kent] string capability is the enter/send key.
	ChangeScrollRegion

	// The OverStrike [over_strike, os] bool capability indicates terminal can overstrike.
	NumberOfFunctionKeys

	// The ResetFile [reset_file, rf] string capability is the name of reset file.
	AcsLrcorner

	// The PaddingBaudRate [padding_baud_rate, pb] num capability is lowest baud rate where padding needed.
	AcsHline

	// The CreateWindow [create_window, cwin] string capability is the define a window #1 from #2,#3 to #4,#5.
	KeyF53

	// The User9 [user9, u9] string capability is the User string #9.
	KeyF62

	// The EnterAmMode [enter_am_mode, smam] string capability is the turn on automatic margins.
	User1

	// The SetTopMargin [set_top_margin, smgt] string capability is the Set top margin at current line.
	BitImageEntwining

	// The KeyF28 [key_f28, kf28] string capability is the F28 function key.
	User5

	// The CrtNoScrolling [crt_no_scrolling, OTns] bool capability indicates crt cannot scroll.
	KeyUndo

	// The KeyF57 [key_f57, kf57] string capability is the F57 function key.
	SetAForeground

	// The KeyEic [key_eic, krmir] string capability is the sent by rmir or smir in insert mode.
	string

	// The CommandCharacter [command_character, cmdch] string capability is the terminal settable cmd character in prototype !?.
	KeyF9

	// The KeyResume [key_resume, kres] string capability is the resume key.
	KeyF5

	// The KeySend [key_send, kEND] string capability is the shifted end key.
	KeyF2

	// The KeyF36 [key_f36, kf36] string capability is the F36 function key.
	KeyLl

	// The KeyF28 [key_f28, kf28] string capability is the F28 function key.
	DeviceType

	// The QuickDial [quick_dial, qdial] string capability is the dial number #1 without checking.
	numCapNames

	// The KeyCommand [key_command, kcmd] string capability is the command key.
	KeyMark

	// The MicroLineSize [micro_line_size, mls] num capability is line step size when in micro mode.
	ExitAmMode

	// The PadChar [pad_char, pad] string capability is the padding char (instead of null).
	User3

	// The QuickDial [quick_dial, qdial] string capability is the dial number #1 without checking.
	PrintRate

	// The ParmDownMicro [parm_down_micro, mcud] string capability is the Like parm_down_cursor in micro mode.
	KeyF50

	// The Reset1string [reset_1string, rs1] string capability is the reset string.
	BitImageCarriageReturn

	// The SetLeftMargin [set_left_margin, smgl] string capability is the set left soft margin at current column.	 See smgl. (ML is not in BSD termcap).
	KeyF9

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	SetWindow

	// The FixedPause [fixed_pause, pause] string capability is the pause for 2-3 seconds.
	OutputResVertInch

	// The OutputResVertInch [output_res_vert_inch, orvi] num capability is vertical resolution in units per inch.
	XoffCharacter

	// The HasStatusLine [has_status_line, hs] bool capability indicates has extra status line.
	ExitMicroMode

	// The KeyF49 [key_f49, kf49] string capability is the F49 function key.
	PcTermOptions

	// The KeyF12 [key_f12, kf12] string capability is the F12 function key.
	KeyScreate

	// The KeypadLocal [keypad_local, rmkx] string capability is the leave 'keyboard_transmit' mode.
	KeySic

	// The KeyMove [key_move, kmov] string capability is the move key.
	OutputResChar

	// The KeyScreate [key_screate, kCRT] string capability is the shifted create key.
	SetForeground

	// The LabF8 [lab_f8, lf8] string capability is the label on function key f8 if not f8.
	HardCopy

	// The RowAddress [row_address, vpa] string capability is the vertical position #1 absolute (P).
	KeyShelp

	// The NewLineDelay [new_line_delay, OTdN] num capability is pad needed for LF.
	SetClock

	// The KeyCancel [key_cancel, kcan] string capability is the cancel key.
	AcsTtee

	// The User6 [user6, u6] string capability is the User string #6.
	Tone

	// The SetAForeground [set_a_foreground, setaf] string capability is the Set foreground color to #1, using ANSI escape.
	LinesOfMemory

	// The SetTopMargin [set_top_margin, smgt] string capability is the Set top margin at current line.
	ReturnDoesClrEol

	// The NoEscCtlc [no_esc_ctlc, xsb] bool capability indicates beehive (f1=escape, f2=ctrl C).
	PrintRate

	// The KeySic [key_sic, kIC] string capability is the shifted insert-character key.
	BackTab

	// The User7 [user7, u7] string capability is the User string #7.
	BackspaceDelay

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	AutoRightMargin

	// The EnterSecureMode [enter_secure_mode, invis] string capability is the turn on blank mode (characters invisible).
	TermcapInit2

	// The ColumnAddress [column_address, hpa] string capability is the horizontal position #1, absolute (P).
	EnterDoublewideMode

	// The KeyEic [key_eic, krmir] string capability is the sent by rmir or smir in insert mode.
	SetABackground

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	KeyF11

	// The CursorToLl [cursor_to_ll, ll] string capability is the last line, first column (if no cup).
	User3

	// The TransparentUnderline [transparent_underline, ul] bool capability indicates underline character overstrikes.
	SetColorPair

	// The UpHalfLine [up_half_line, hu] string capability is the half a line up.
	HorizontalTabDelay

	// The KeySelect [key_select, kslt] string capability is the select key.
	KeyF21

	// The ExitAttributeMode [exit_attribute_mode, sgr0] string capability is the turn off all attributes.
	KeyDl

	// The KeyPpage [key_ppage, kpp] string capability is the previous-page key.
	NoCorrectlyWorkingCr

	// The ExitPcCharsetMode [exit_pc_charset_mode, rmpch] string capability is the Exit PC character display mode.
	var

	// The CursorLeft [cursor_left, cub1] string capability is the move left one space.
	Tab

	// The EnterProtectedMode [enter_protected_mode, prot] string capability is the turn on protected mode.
	NumberOfPins

	// The User5 [user5, u5] string capability is the User string #5.
	iota

	// The EnterAmMode [enter_am_mode, smam] string capability is the turn on automatic margins.
	ScrollReverse

	// The GenericType [generic_type, gn] bool capability indicates generic line type.
	User1

	// The Reset2string [reset_2string, rs2] string capability is the reset string.
	KeyIc

	// The PrintRate [print_rate, cps] num capability is print rate in characters per second.
	EnterSuperscriptMode

	// The ExitAmMode [exit_am_mode, rmam] string capability is the turn off automatic margins.
	EnterMicroMode

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	DefineBitImageRegion

	// The EnterBlinkMode [enter_blink_mode, blink] string capability is the turn on blinking.
	ParmIch

	// The EatNewlineGlitch [eat_newline_glitch, xenl] bool capability indicates newline ignored after 80 cols (concept).
	GnuHasMetaKey

	// The PrtrSilent [prtr_silent, mc5i] bool capability indicates printer will not echo on screen.
	KeyF55

	// The SetForeground [set_foreground, setf] string capability is the Set foreground color #1.
	KeyF58

	// The CarriageReturnDelay [carriage_return_delay, OTdC] num capability is pad needed for CR.
	EnterLeftHlMode

	// The RestoreCursor [restore_cursor, rc] string capability is the restore cursor to position of last save_cursor.
	var

	// The FormFeed [form_feed, ff] string capability is the hardcopy terminal page eject (P*).
	WidthStatusLine

	// The SuperscriptCharacters [superscript_characters, supcs] string capability is the List of superscriptable characters.
	var

	// The AcsLtee [acs_ltee, OTGR] string capability is the tee pointing right.
	EnterPcCharsetMode

	// The ExitPcCharsetMode [exit_pc_charset_mode, rmpch] string capability is the Exit PC character display mode.
	MoveStandoutMode

	// The KeySmove [key_smove, kMOV] string capability is the shifted move key.
	SetBottomMargin

	// The KeyF25 [key_f25, kf25] string capability is the F25 function key.
	OtherNonFunctionKeys

	// The OutputResHorzInch [output_res_horz_inch, orhi] num capability is horizontal resolution in units per inch.
	WideCharSize

	// The MoveInsertMode [move_insert_mode, mir] bool capability indicates safe to move while in insert mode.
	BufferCapacity

	// The PkeyKey [pkey_key, pfkey] string capability is the program function key #1 to type string #2.
	MicroDown

	// The ParmInsertLine [parm_insert_line, il] string capability is the insert #1 lines (P*).
	BoxChars1

	// The BitImageNewline [bit_image_newline, binel] string capability is the Move to next row of the bit image.
	AcsUlcorner

	// The MaxMicroAddress [max_micro_address, maddr] num capability is maximum value in micro_..._address.
	DotVertSpacing

	// The SetLrMargin [set_lr_margin, smglr] string capability is the Set both left and right margins to #1, #2.  (ML is not in BSD termcap).
	KeyF12

	// The InitTabs [init_tabs, it] num capability is tabs initially every # spaces.
	KeyMove

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	User7

	// The KeyCommand [key_command, kcmd] string capability is the command key.
	KeyF58

	// The EndBitImageRegion [end_bit_image_region, endbi] string capability is the End a bit-image region.
	KeyIc

	// The Init3string [init_3string, is3] string capability is the initialization string.
	BackspacesWithBs

	// The ChangeLinePitch [change_line_pitch, lpi] string capability is the Change number of lines per inch to #1.
	LabelOn

	// The KeyC1 [key_c1, kc1] string capability is the lower left of keypad.
	KeyF62

	// The KeyF24 [key_f24, kf24] string capability is the F24 function key.
	KeyRedo

	// The KeyF55 [key_f55, kf55] string capability is the F55 function key.
	NeedsXonXoff

	// The ClearScreen [clear_screen, clear] string capability is the clear screen and home cursor (P*).
	KeyStab

	// The KeyClear [key_clear, kclr] string capability is the clear-screen or erase key.
	SetForeground

	// The KeyF33 [key_f33, kf33] string capability is the F33 function key.
	User8

	// The CarriageReturnDelay [carriage_return_delay, OTdC] num capability is pad needed for CR.
	FlashScreen

	// The KeyUp [key_up, kcuu1] string capability is the up-arrow key.
	StartCharSetDef

	// The KeySreplace [key_sreplace, kRPL] string capability is the shifted replace key.
	KeyIc

	// The NeedsXonXoff [needs_xon_xoff, nxon] bool capability indicates padding will not work, xon/xoff required.
	MicroDown

	// The EnterScancodeMode [enter_scancode_mode, smsc] string capability is the Enter PC scancode mode.
	EnterTopHlMode

	// The KeyNpage [key_npage, knp] string capability is the next-page key.
	CursorUp

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	EnaAcs

	// The LabF9 [lab_f9, lf9] string capability is the label on function key f9 if not f9.
	DotHorzSpacing

	// The KeyReplace [key_replace, krpl] string capability is the replace key.
	EnterHorizontalHlMode

	// The MaxMicroAddress [max_micro_address, maddr] num capability is maximum value in micro_..._address.
	SetAttributes

	// The KeyF8 [key_f8, kf8] string capability is the F8 function key.
	ExitCaMode

	// The BackTab [back_tab, cbt] string capability is the back tab (P).
	KeyF25

	// The KeyF27 [key_f27, kf27] string capability is the F27 function key.
	NoPadChar

	// The BufferCapacity [buffer_capacity, bufsz] num capability is numbers of bytes buffered before printing.
	KeySmessage

	// The HasHardwareTabs [has_hardware_tabs, OTpt] bool capability indicates has 8-char tabs invoked with ^I.
	UpHalfLine

	// The DotHorzSpacing [dot_horz_spacing, spinh] num capability is spacing of dots horizontally in dots per inch.
	KeySreplace

	// The ExitLeftwardMode [exit_leftward_mode, rlm] string capability is the End left-motion mode.
	KeyF61

	// The KeyLeft [key_left, kcub1] string capability is the left-arrow key.
	KeyF56

	// CapCountString is the count of string capabilities.
	KeySprint

	// The KeyF10 [key_f10, kf10] string capability is the F10 function key.
	KeyShelp

	// The GenericType [generic_type, gn] bool capability indicates generic line type.
	KeySoptions

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	ChangeResVert

	// The KeyEos [key_eos, ked] string capability is the clear-to-end-of-screen key.
	LabelOff

	// The KeyF45 [key_f45, kf45] string capability is the F45 function key.
	ParmUpMicro

	// The EnterInsertMode [enter_insert_mode, smir] string capability is the enter insert mode.
	AcsUrcorner

	// The KeyShelp [key_shelp, kHLP] string capability is the shifted help key.
	SetPageLength

	// stringCapNames are the string term cap names.
	KeyPrint

	// The KeyF62 [key_f62, kf62] string capability is the F62 function key.
	KeyF49

	// The DownHalfLine [down_half_line, hd] string capability is the half a line down.
	KeyRedo

	// The KeyF49 [key_f49, kf49] string capability is the F49 function key.
	EnterInsertMode

	// The KeySend [key_send, kEND] string capability is the shifted end key.
	TildeGlitch

	// The KeyF30 [key_f30, kf30] string capability is the F30 function key.
	KeySsuspend

	// The AcsUlcorner [acs_ulcorner, OTG2] string capability is the single upper left.
	BitImageCarriageReturn

	// The ColorNames [color_names, colornm] string capability is the Give name for color #1.
	KeyF54

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	BitImageRepeat

	// The SuperscriptCharacters [superscript_characters, supcs] string capability is the List of superscriptable characters.
	OutputResLine

	// The KeySelect [key_select, kslt] string capability is the select key.
	TildeGlitch

	// The KeyCancel [key_cancel, kcan] string capability is the cancel key.
	KeyF47

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	LabF6

	// The KeyDown [key_down, kcud1] string capability is the down-arrow key.
	Reset3string

	// The SetBottomMargin [set_bottom_margin, smgb] string capability is the Set bottom margin at current line.
	EnterProtectedMode

	// The PadChar [pad_char, pad] string capability is the padding char (instead of null).
	User4

	// The KeyF35 [key_f35, kf35] string capability is the F35 function key.
	KeyC3

	// The EnterLowHlMode [enter_low_hl_mode, elohlm] string capability is the Enter low highlight mode.
	KeyF58

	// The KeyClose [key_close, kclo] string capability is the close key.
	KeySfind

	// The KeySexit [key_sexit, kEXT] string capability is the shifted exit key.
	EnterAmMode

	// The SetAForeground [set_a_foreground, setaf] string capability is the Set foreground color to #1, using ANSI escape.
	DownHalfLine

	// The DisplayPcChar [display_pc_char, dispc] string capability is the Display PC character #1.
	DisStatusLine

	// The EnterLeftwardMode [enter_leftward_mode, slm] string capability is the Start leftward carriage motion.
	ColAddrGlitch

	// The ReqMousePos [req_mouse_pos, reqmp] string capability is the Request mouse position.
	CursorHome

	// The ExitDeleteMode [exit_delete_mode, rmdc] string capability is the end delete mode.
	User3

	// The SetPageLength [set_page_length, slines] string capability is the Set page length to #1 lines.
	ParmDownCursor

	// The KeyF1 [key_f1, kf1] string capability is the F1 function key.
	KeyLl

	// The CursorInvisible [cursor_invisible, civis] string capability is the make cursor invisible.
	EnterProtectedMode

	// The CursorRight [cursor_right, cuf1] string capability is the non-destructive space (move right one space).
	SetAAttributes

	// The KeyScreate [key_screate, kCRT] string capability is the shifted create key.
	OrderOfPins

	// The PrtrSilent [prtr_silent, mc5i] bool capability indicates printer will not echo on screen.
	KeyHelp

	// The LabF2 [lab_f2, lf2] string capability is the label on function key f2 if not f2.
	KeySdc

	// The ScrollForward [scroll_forward, ind] string capability is the scroll text up (P).
	MaxPairs

	// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
	KeyF55

	// The CanChange [can_change, ccc] bool capability indicates terminal can re-define existing colors.
	StatusLineEscOk

	// The DialPhone [dial_phone, dial] string capability is the dial number #1.
	KeyUp

	// The WaitTone [wait_tone, wait] string capability is the wait for dial-tone.
	KeyOpen

	// The KeyF54 [key_f54, kf54] string capability is the F54 function key.
	ParmIch

	// The LabF5 [lab_f5, lf5] string capability is the label on function key f5 if not f5.
	SetWindow

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	KeyNext

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	BackColorErase

	// The Lines [lines, lines] num capability is number of lines on screen or page.
	PrtrSilent

	// The KeyRefresh [key_refresh, krfr] string capability is the refresh key.
	ToStatusLine

	// The LabF1 [lab_f1, lf1] string capability is the label on function key f1 if not f1.
	SetBottomMarginParm

	// The ExitLeftwardMode [exit_leftward_mode, rlm] string capability is the End left-motion mode.
	KeyF50

	// The ColorNames [color_names, colornm] string capability is the Give name for color #1.
	KeyF45

	// The EnterHorizontalHlMode [enter_horizontal_hl_mode, ehhlm] string capability is the Enter horizontal highlight mode.
	ExitScancodeMode

	// The KeyDl [key_dl, kdl1] string capability is the delete-line key.
	User8

	// The ParmUpMicro [parm_up_micro, mcuu] string capability is the Like parm_up_cursor in micro mode.
	KeyF30

	// The KeySleft [key_sleft, kLFT] string capability is the shifted left-arrow key.
	KeyF9

	// The ParmDownCursor [parm_down_cursor, cud] string capability is the down #1 lines (P*).
	SetTab

	// The ChangeScrollRegion [change_scroll_region, csr] string capability is the change region to line #1 to line #2 (P).
	KeyScreate

	// The HardCursor [hard_cursor, chts] bool capability indicates cursor is hard to see.
	AcsLtee

	// The CanChange [can_change, ccc] bool capability indicates terminal can re-define existing colors.
	EnterDimMode

	// The BackspaceIfNotBs [backspace_if_not_bs, OTbc] string capability is the move left, if not ^H.
	ReturnDoesClrEol

	// The OverStrike [over_strike, os] bool capability indicates terminal can overstrike.
	OutputResChar

	// The CursorAddress [cursor_address, cup] string capability is the move to row #1 columns #2.
	MetaOff

	// The KeyOptions [key_options, kopt] string capability is the options key.
	ExitItalicsMode

	// The KeyF30 [key_f30, kf30] string capability is the F30 function key.
	KeypadXmit

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	AcsBtee

	// The Reset1string [reset_1string, rs1] string capability is the reset string.
	KeySelect

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	ClearMargins

	// The KeySelect [key_select, kslt] string capability is the select key.
	SuperscriptCharacters

	// The EnterUpwardMode [enter_upward_mode, sum] string capability is the Start upward carriage motion.
	MaxAttributes

	// The CodeSetInit [code_set_init, csin] string capability is the Init sequence for multiple codesets.
	ParmIndex

	// The ExitStandoutMode [exit_standout_mode, rmso] string capability is the exit standout mode.
	SuperscriptCharacters

	// The BackspaceDelay [backspace_delay, OTdB] num capability is padding required for ^H.
	DotVertSpacing

	// The KeyF28 [key_f28, kf28] string capability is the F28 function key.
	NewLineDelay

	// The OutputResHorzInch [output_res_horz_inch, orhi] num capability is horizontal resolution in units per inch.
	KeyCancel

	// The BoxChars1 [box_chars_1, box1] string capability is the box characters primary set.
	ExitInsertMode

	// The MemoryBelow [memory_below, db] bool capability indicates display may be retained below the screen.
	ParmUpMicro

	// The EndBitImageRegion [end_bit_image_region, endbi] string capability is the End a bit-image region.
	KeySundo

	// The ClrEos [clr_eos, ed] string capability is the clear to end of screen (P*).
	AcsBtee

	// The InitializePair [initialize_pair, initp] string capability is the Initialize color pair #1 to fg=(#2,#3,#4), bg=(#5,#6,#7).
	EnterShadowMode

	// The ChangeCharPitch [change_char_pitch, cpi] string capability is the Change number of characters per inch to #1.
	KeyF14

	// The EnterNearLetterQuality [enter_near_letter_quality, snlq] string capability is the Enter NLQ mode.
	UnderlineChar

	// The EnterUpwardMode [enter_upward_mode, sum] string capability is the Start upward carriage motion.
	EnterPcCharsetMode

	// The KeyF20 [key_f20, kf20] string capability is the F20 function key.
	stringCapNames

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	KeyCancel

	// The KeyF24 [key_f24, kf24] string capability is the F24 function key.
	ParmLeftMicro

	// The KeyClose [key_close, kclo] string capability is the close key.
	FlashScreen

	// The EnterItalicsMode [enter_italics_mode, sitm] string capability is the Enter italic mode.
	Init1string

	// The PkeyXmit [pkey_xmit, pfx] string capability is the program function key #1 to transmit string #2.
	BackTab

	// The KeySrsume [key_srsume, kRES] string capability is the shifted resume key.
	TermcapInit2

	// The KeySredo [key_sredo, kRDO] string capability is the shifted redo key.
	KeyPpage

	// The PrtrOff [prtr_off, mc4] string capability is the turn off printer.
	BackspaceDelay

	// The KeyF7 [key_f7, kf7] string capability is the F7 function key.
	LabF5

	// The KeyMove [key_move, kmov] string capability is the move key.
	User1

	// The User2 [user2, u2] string capability is the User string #2.
	TransparentUnderline

	// The KeyF41 [key_f41, kf41] string capability is the F41 function key.
	KeyF17

	// The KeySuspend [key_suspend, kspd] string capability is the suspend key.
	NoColorVideo

	// The StatusLineEscOk [status_line_esc_ok, eslok] bool capability indicates escape can be used on the status line.
	KeyF29

	// The KeyF12 [key_f12, kf12] string capability is the F12 function key.
	CarriageReturnDelay

	// The HueLightnessSaturation [hue_lightness_saturation, hls] bool capability indicates terminal uses only HLS color notation (Tektronix).
	BackspaceDelay

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	WidthStatusLine

	// The PkeyXmit [pkey_xmit, pfx] string capability is the program function key #1 to transmit string #2.
	KeyResume

	// The FlashHook [flash_hook, hook] string capability is the flash switch hook.
	StopBitImage

	// The OutputResVertInch [output_res_vert_inch, orvi] num capability is vertical resolution in units per inch.
	KeyShome

	// The DownHalfLine [down_half_line, hd] string capability is the half a line down.
	StartCharSetDef

	// The SetForeground [set_foreground, setf] string capability is the Set foreground color #1.
	KeyF30

	// The KeySprint [key_sprint, kPRT] string capability is the shifted print key.
	LabF1

	// The MicroRowAddress [micro_row_address, mvpa] string capability is the Like row_address #1 in micro mode.
	EnterSuperscriptMode

	// The KeyF60 [key_f60, kf60] string capability is the F60 function key.
	KeyLl

	// The ParmRindex [parm_rindex, rin] string capability is the scroll back #1 lines (P).
	KeyDc

	// The ClearScreen [clear_screen, clear] string capability is the clear screen and home cursor (P*).
	string

	// The SetWindow [set_window, wind] string capability is the current window is lines #1-#2 cols #3-#4.
	ScrollReverse

	// The KeyF49 [key_f49, kf49] string capability is the F49 function key.
	KeySdc

	// The KeyMark [key_mark, kmrk] string capability is the mark key.
	ParmDch

	// The CursorVisible [cursor_visible, cvvis] string capability is the make cursor very visible.
	MoveStandoutMode

	// The DeleteLine [delete_line, dl1] string capability is the delete line (P*).
	SetAForeground

	// The RepeatChar [repeat_char, rep] string capability is the repeat char #1 #2 times (P*).
	PkeyXmit

	// The KeyMark [key_mark, kmrk] string capability is the mark key.
	ExitLeftwardMode

	// The ExitSuperscriptMode [exit_superscript_mode, rsupm] string capability is the End superscript mode.
	NumLabels

	// The EnterShadowMode [enter_shadow_mode, sshm] string capability is the Enter shadow-print mode.
	KeyF47

	// The CursorToLl [cursor_to_ll, ll] string capability is the last line, first column (if no cup).
	KeyBtab

	// The ExitInsertMode [exit_insert_mode, rmir] string capability is the exit insert mode.
	MetaOff

	// The AcsLrcorner [acs_lrcorner, OTG4] string capability is the single lower right.
	KeyF39

	// The KeyF15 [key_f15, kf15] string capability is the F15 function key.
	EnterScancodeMode

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	KeyReference

	// The EnterDoublewideMode [enter_doublewide_mode, swidm] string capability is the Enter double-wide mode.
	EnterScancodeMode

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	EnterBoldMode

	// The KeyHelp [key_help, khlp] string capability is the help key.
	EnterDoublewideMode

	// The NewLineDelay [new_line_delay, OTdN] num capability is pad needed for LF.
	KeyF42

	// The ExitDoublewideMode [exit_doublewide_mode, rwidm] string capability is the End double-wide mode.
	HasStatusLine

	// The KeyF55 [key_f55, kf55] string capability is the F55 function key.
	KeyF53

	// The KeyCopy [key_copy, kcpy] string capability is the copy key.
	UnderlineChar

	// The LabF6 [lab_f6, lf6] string capability is the label on function key f6 if not f6.
	ZeroMotion

	// The GenericType [generic_type, gn] bool capability indicates generic line type.
	EnterBlinkMode

	// The KeyCommand [key_command, kcmd] string capability is the command key.
	NumLabels

	// The EnterNormalQuality [enter_normal_quality, snrmq] string capability is the Enter normal-quality mode.
	AcsUrcorner

	// The ChangeResHorz [change_res_horz, chr] string capability is the Change horizontal resolution to #1.
	KeyF35

	// The ZeroMotion [zero_motion, zerom] string capability is the No motion for subsequent character.
	ParmRightMicro

	// The KeyF20 [key_f20, kf20] string capability is the F20 function key.
	KeySbeg

	// The GnuHasMetaKey [gnu_has_meta_key, OTMT] bool capability indicates has meta key.
	KeyLeft

	// The AutoLeftMargin [auto_left_margin, bw] bool capability indicates cub1 wraps from column 0 to last column.
	KeySelect

	// The KeyF52 [key_f52, kf52] string capability is the F52 function key.
	KeyF30

	// The LinefeedIsNewline [linefeed_is_newline, OTNL] bool capability indicates move down with \n.
	KeyHelp

	// The InitializeColor [initialize_color, initc] string capability is the initialize color #1 to (#2,#3,#4).
	EnterProtectedMode

	// The KeyBtab [key_btab, kcbt] string capability is the back-tab key.
	MicroLeft

	// The KeyDl [key_dl, kdl1] string capability is the delete-line key.
	KeyF12

	// The KeyF5 [key_f5, kf5] string capability is the F5 function key.
	terminfo

	// The KeyF41 [key_f41, kf41] string capability is the F41 function key.
	KeyF58

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	KeyMouse

	// The ReqForInput [req_for_input, rfi] string capability is the send next input char (for ptys).
	ExitAttributeMode

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	ExitAltCharsetMode

	// The SetRightMarginParm [set_right_margin_parm, smgrp] string capability is the Set right margin at column #1.
	MagicCookieGlitch

	// The KeyDown [key_down, kcud1] string capability is the down-arrow key.
	PrintRate

	// The KeySdl [key_sdl, kDL] string capability is the shifted delete-line key.
	CursorHome

	// The KeyF32 [key_f32, kf32] string capability is the F32 function key.
	EraseChars

	// The SetTab [set_tab, hts] string capability is the set a tab in every row, current columns.
	KeyF9

	// The CeolStandoutGlitch [ceol_standout_glitch, xhp] bool capability indicates standout not erased by overwriting (hp).
	LabF6

	// The AcsVline [acs_vline, OTGV] string capability is the single vertical line.
	Buttons

	// The AcsHline [acs_hline, OTGH] string capability is the single horizontal line.
	PkeyKey

	// The DefineBitImageRegion [define_bit_image_region, defbi] string capability is the Define rectangular bit image region.
	EatNewlineGlitch

	// The ParmUpCursor [parm_up_cursor, cuu] string capability is the up #1 lines (P*).
	Reset1string

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	CapCountNum

	// numCapNames are the num term cap names.
	KeyCtab

	// The BitImageNewline [bit_image_newline, binel] string capability is the Move to next row of the bit image.
	DotHorzSpacing

	// The KeyF18 [key_f18, kf18] string capability is the F18 function key.
	KeyF57

	// The EnterHorizontalHlMode [enter_horizontal_hl_mode, ehhlm] string capability is the Enter horizontal highlight mode.
	KeyF49

	// The LabelWidth [label_width, lw] num capability is columns in each label.
	InitializeColor

	// The KeyFind [key_find, kfnd] string capability is the find key.
	KeyF21

	// The LabF10 [lab_f10, lf10] string capability is the label on function key f10 if not f10.
	CursorMemAddress

	// The LabelOn [label_on, smln] string capability is the turn on soft labels.
	EnaAcs

	// The LabelHeight [label_height, lh] num capability is rows in each label.
	KeyClear

	// The MicroLineSize [micro_line_size, mls] num capability is line step size when in micro mode.
	KeyShome

	// The KeyF9 [key_f9, kf9] string capability is the F9 function key.
	KeyPrevious

	// The KeyF15 [key_f15, kf15] string capability is the F15 function key.
	RepeatChar

	// The EnterStandoutMode [enter_standout_mode, smso] string capability is the begin standout mode.
	KeyBackspace

	// The KeyRedo [key_redo, krdo] string capability is the redo key.
	DotHorzSpacing

	// The CursorMemAddress [cursor_mem_address, mrcup] string capability is the memory relative cursor addressing, move to row #1 columns #2.
	MaxMicroAddress

	// The KeyF29 [key_f29, kf29] string capability is the F29 function key.
	KeyF5

	// The ClearMargins [clear_margins, mgc] string capability is the clear right and left soft margins.
	CapCountBool

	// The AcsLtee [acs_ltee, OTGR] string capability is the tee pointing right.
	EnterLeftHlMode

	// The Pulse [pulse, pulse] string capability is the select pulse dialing.
	MemoryUnlock

	// The KeyF27 [key_f27, kf27] string capability is the F27 function key.
	SetTopMargin

	// The KeypadLocal [keypad_local, rmkx] string capability is the leave 'keyboard_transmit' mode.
	MouseInfo

	// Num capabilities.
	iota

	// The KeyF35 [key_f35, kf35] string capability is the F35 function key.
	HasStatusLine

	// The KeyA3 [key_a3, ka3] string capability is the upper right of keypad.
	ClrEol

	// The KeyDown [key_down, kcud1] string capability is the down-arrow key.
	SetBottomMargin

	// The SetAttributes [set_attributes, sgr] string capability is the define video attributes #1-#9 (PG9).
	PkeyKey

	// The KeyShome [key_shome, kHOM] string capability is the shifted home key.
	string

	// The MemoryAbove [memory_above, da] bool capability indicates display may be retained above the screen.
	KeyC3

	// The KeySsuspend [key_ssuspend, kSPD] string capability is the shifted suspend key.
	NonRevRmcup

	// The HasHardwareTabs [has_hardware_tabs, OTpt] bool capability indicates has 8-char tabs invoked with ^I.
	KeyF15

	// The KeyF23 [key_f23, kf23] string capability is the F23 function key.
	User7

	// The KeyRestart [key_restart, krst] string capability is the restart key.
	AcsLrcorner

	// The InsertCharacter [insert_character, ich1] string capability is the insert character (P).
	KeySdc

	// numCapNames are the num term cap names.
	AcsLlcorner

	// The ParmUpCursor [parm_up_cursor, cuu] string capability is the up #1 lines (P*).
	SetForeground

	// The KeyRestart [key_restart, krst] string capability is the restart key.
	PadChar

	// The KeyF23 [key_f23, kf23] string capability is the F23 function key.
	CharPadding

	// The KeyIl [key_il, kil1] string capability is the insert-line key.
	CarriageReturnDelay

	// The KeyEnter [key_enter, kent] string capability is the enter/send key.
	CapCountNum

	// The KeyF31 [key_f31, kf31] string capability is the F31 function key.
	BitImageCarriageReturn

	// The ExitStandoutMode [exit_standout_mode, rmso] string capability is the exit standout mode.
	KeySnext

	// CapCountNum is the count of num capabilities.
	BoxChars1

	// The KeyB2 [key_b2, kb2] string capability is the center of keypad.
	EnterBlinkMode

	// The KeySnext [key_snext, kNXT] string capability is the shifted next key.
	ExitAmMode

	// The InsertCharacter [insert_character, ich1] string capability is the insert character (P).
	KeyF13

	// The Set0DesSeq [set0_des_seq, s0ds] string capability is the Shift to codeset 0 (EUC set 0, ASCII).
	ExitItalicsMode

	// The StartBitImage [start_bit_image, sbim] string capability is the Start printing bit image graphics.
	KeySprevious

	// The ResetFile [reset_file, rf] string capability is the name of reset file.
	Init1string

	// Code generated by gen.go. DO NOT EDIT.
	HasPrintWheel

	// The KeyUp [key_up, kcuu1] string capability is the up-arrow key.
	StartBitImage

	// The MicroRight [micro_right, mcuf1] string capability is the Like cursor_right in micro mode.
	CursorRight

	// stringCapNames are the string term cap names.
	KeySprevious

	// The MicroLineSize [micro_line_size, mls] num capability is line step size when in micro mode.
	KeySnext

	// The KeySend [key_send, kEND] string capability is the shifted end key.
	PaddingBaudRate

	// The LabF9 [lab_f9, lf9] string capability is the label on function key f9 if not f9.
	KeyCopy

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	KeySrsume

	// The KeyF57 [key_f57, kf57] string capability is the F57 function key.
	KeyF16

	// The KeyF1 [key_f1, kf1] string capability is the F1 function key.
	DefineChar

	// The HasHardwareTabs [has_hardware_tabs, OTpt] bool capability indicates has 8-char tabs invoked with ^I.
	KeyRefresh

	// The InsertLine [insert_line, il1] string capability is the insert line (P*).
	EnterScancodeMode

	// The Tab [tab, ht] string capability is the tab to next 8-space hardware tab stop.
	KeyF6

	// The EnterCaMode [enter_ca_mode, smcup] string capability is the string to start programs using cup.
	EnterDoublewideMode

	// The SetAForeground [set_a_foreground, setaf] string capability is the Set foreground color to #1, using ANSI escape.
	KeyPrevious

	// The MicroColumnAddress [micro_column_address, mhpa] string capability is the Like column_address in micro mode.
	CreateWindow

	// The AutoRightMargin [auto_right_margin, am] bool capability indicates terminal has automatic margins.
	EraseOverstrike

	// The KeyF36 [key_f36, kf36] string capability is the F36 function key.
	KeyF57

	// The KeyDc [key_dc, kdch1] string capability is the delete-character key.
	SetLeftMargin

	// The KeySmove [key_smove, kMOV] string capability is the shifted move key.
	KeyF55

	// The User4 [user4, u4] string capability is the User string #4.
	InsertCharacter

	// The ScancodeEscape [scancode_escape, scesc] string capability is the Escape for scancode emulation.
	LabelWidth

	// The EnterCaMode [enter_ca_mode, smcup] string capability is the string to start programs using cup.
	KeyClose

	// The CarriageReturn [carriage_return, cr] string capability is the carriage return (P*) (P*).
	KeyF33

	// The MemoryBelow [memory_below, db] bool capability indicates display may be retained below the screen.
	DeleteLine

	// The MicroDown [micro_down, mcud1] string capability is the Like cursor_down in micro mode.
	KeyF51

	// The InsertPadding [insert_padding, ip] string capability is the insert padding after inserted character.
	XoffCharacter

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	KeyReplace

	// The KeyF32 [key_f32, kf32] string capability is the F32 function key.
	XonXoff

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	KeyEnter

	// The EnterXonMode [enter_xon_mode, smxon] string capability is the turn on xon/xoff handshaking.
	ParmDch

	// The CarriageReturn [carriage_return, cr] string capability is the carriage return (P*) (P*).
	EnterUnderlineMode

	// The KeyF10 [key_f10, kf10] string capability is the F10 function key.
	CursorNormal

	// The DisplayPcChar [display_pc_char, dispc] string capability is the Display PC character #1.
	KeyF9

	// The CursorVisible [cursor_visible, cvvis] string capability is the make cursor very visible.
	KeyCancel

	// The BitImageEntwining [bit_image_entwining, bitwin] num capability is number of passes for each bit-image row.
	KeyNpage

	// The KeyF29 [key_f29, kf29] string capability is the F29 function key.
	KeyF59

	// The NoPadChar [no_pad_char, npc] bool capability indicates pad character does not exist.
	RowAddress

	// The EnterHorizontalHlMode [enter_horizontal_hl_mode, ehhlm] string capability is the Enter horizontal highlight mode.
	SetWindow

	// stringCapNames are the string term cap names.
	KeyF0

	// The NumberOfPins [number_of_pins, npins] num capability is numbers of pins in print-head.
	NeedsXonXoff

	// The ParmLeftCursor [parm_left_cursor, cub] string capability is the move #1 characters to the left (P).
	SetRightMargin

	// The EnterCaMode [enter_ca_mode, smcup] string capability is the string to start programs using cup.
	ExitDeleteMode

	// The ExitItalicsMode [exit_italics_mode, ritm] string capability is the End italic mode.
	RestoreCursor

	// The LinesOfMemory [lines_of_memory, lm] num capability is lines of memory if > line. 0 means varies.
	MoveStandoutMode

	// The KeyNpage [key_npage, knp] string capability is the next-page key.
	Newline

	// The MaximumWindows [maximum_windows, wnum] num capability is maximum number of definable windows.
	XonXoff

	// The KeyReference [key_reference, kref] string capability is the reference key.
	ExitMicroMode

	// The DestTabsMagicSmso [dest_tabs_magic_smso, xt] bool capability indicates tabs destructive, magic so char (t1061).
	KeyF25

	// The HorizontalTabDelay [horizontal_tab_delay, OTdT] num capability is padding required for ^I.
	EnterHorizontalHlMode

	// The MemoryAbove [memory_above, da] bool capability indicates display may be retained above the screen.
	KeyLl

	// The StartCharSetDef [start_char_set_def, scsd] string capability is the Start character set definition #1, with #2 characters in the set.
	KeyF57

	// The KeyF7 [key_f7, kf7] string capability is the F7 function key.
	ParmDownMicro

	// The KeyF1 [key_f1, kf1] string capability is the F1 function key.
	CursorUp

	// The KeyEic [key_eic, krmir] string capability is the sent by rmir or smir in insert mode.
	KeyF55

	// The KeyMouse [key_mouse, kmous] string capability is the Mouse event has occurred.
	LabF2

	// The KeySredo [key_sredo, kRDO] string capability is the shifted redo key.
	Hangup

	// The KeyF50 [key_f50, kf50] string capability is the F50 function key.
	RepeatChar

	// The KeyFind [key_find, kfnd] string capability is the find key.
	CursorMemAddress

	// The EnterNearLetterQuality [enter_near_letter_quality, snlq] string capability is the Enter NLQ mode.
	Tab

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	Tone

	// The EnterDimMode [enter_dim_mode, dim] string capability is the turn on half-bright mode.
	Reset1string

	// The KeyEos [key_eos, ked] string capability is the clear-to-end-of-screen key.
	KeyScommand

	// Code generated by gen.go. DO NOT EDIT.
	KeyF51

	// The AcsUlcorner [acs_ulcorner, OTG2] string capability is the single upper left.
	BackspaceDelay

	// The OutputResVertInch [output_res_vert_inch, orvi] num capability is vertical resolution in units per inch.
	SetRightMargin

	// The LinesOfMemory [lines_of_memory, lm] num capability is lines of memory if > line. 0 means varies.
	MemoryLock

	// The BoxChars1 [box_chars_1, box1] string capability is the box characters primary set.
	EnterShadowMode

	// The Set2DesSeq [set2_des_seq, s2ds] string capability is the Shift to codeset 2.
	KeyF23

	// The KeyF1 [key_f1, kf1] string capability is the F1 function key.
	KeyF49

	// The MouseInfo [mouse_info, minfo] string capability is the Mouse status information.
	GetMouse

	// The CarriageReturn [carriage_return, cr] string capability is the carriage return (P*) (P*).
	KeyF5

	// The ReqForInput [req_for_input, rfi] string capability is the send next input char (for ptys).
	PkeyXmit

	// The ExitMicroMode [exit_micro_mode, rmicm] string capability is the End micro-motion mode.
	HasStatusLine

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	KeyF33

	// The LabF7 [lab_f7, lf7] string capability is the label on function key f7 if not f7.
	KeyPrint

	// The KeySprevious [key_sprevious, kPRV] string capability is the shifted previous key.
	ColumnAddress

	// The HueLightnessSaturation [hue_lightness_saturation, hls] bool capability indicates terminal uses only HLS color notation (Tektronix).
	UnderlineChar

	// The InitProg [init_prog, iprog] string capability is the path name of program for initialization.
	LinesOfMemory

	// The ArrowKeyMap [arrow_key_map, OTma] string capability is the map motion-keys for vi version 2.
	PkeyKey

	// The RowAddrGlitch [row_addr_glitch, xvpa] bool capability indicates only positive motion for vpa/mvpa caps.
	KeySprint

	// The MicroColSize [micro_col_size, mcs] num capability is character step size when in micro mode.
	KeyF7

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	ChangeLinePitch

	// The ScrollForward [scroll_forward, ind] string capability is the scroll text up (P).
	KeyF44

	// The EnterLeftHlMode [enter_left_hl_mode, elhlm] string capability is the Enter left highlight mode.
	MicroColSize

	// Num capabilities.
	KeyMark

	// The KeyPpage [key_ppage, kpp] string capability is the previous-page key.
	DeviceType

	// The InitProg [init_prog, iprog] string capability is the path name of program for initialization.
	var

	// The AcsChars [acs_chars, acsc] string capability is the graphics charset pairs, based on vt100.
	ParmDownMicro

	// The KeyScommand [key_scommand, kCMD] string capability is the shifted command key.
	CursorMemAddress

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	CpiChangesRes

	// The KeyF52 [key_f52, kf52] string capability is the F52 function key.
	LabelOn

	// The EnterProtectedMode [enter_protected_mode, prot] string capability is the turn on protected mode.
	ParmDownCursor

	// The KeyF33 [key_f33, kf33] string capability is the F33 function key.
	WaitTone

	// The KeyShome [key_shome, kHOM] string capability is the shifted home key.
	GnuHasMetaKey

	// The CanChange [can_change, ccc] bool capability indicates terminal can re-define existing colors.
	KeySic

	// The KeyF42 [key_f42, kf42] string capability is the F42 function key.
	DisplayPcChar

	// The MicroUp [micro_up, mcuu1] string capability is the Like cursor_up in micro mode.
	EnterDraftQuality

	// The CursorHome [cursor_home, home] string capability is the home cursor (if no cup).
	ClearMargins

	// The MemoryAbove [memory_above, da] bool capability indicates display may be retained above the screen.
	KeyF29

	// The ChangeCharPitch [change_char_pitch, cpi] string capability is the Change number of characters per inch to #1.
	EnterAltCharsetMode

	// The SetClock [set_clock, sclk] string capability is the set clock, #1 hrs #2 mins #3 secs.
	KeyLeft

	// The MaxPairs [max_pairs, pairs] num capability is maximum number of color-pairs on the screen.
	KeyUp

	// The BitImageType [bit_image_type, bitype] num capability is type of bit-image device.
	RestoreCursor

	// The CrtNoScrolling [crt_no_scrolling, OTns] bool capability indicates crt cannot scroll.
	EnterVerticalHlMode

	// The KeyF26 [key_f26, kf26] string capability is the F26 function key.
	KeyCommand

	// The KeyIc [key_ic, kich1] string capability is the insert-character key.
	PkeyKey

	// The KeyF37 [key_f37, kf37] string capability is the F37 function key.
	KeyF0

	// The ParmRindex [parm_rindex, rin] string capability is the scroll back #1 lines (P).
	KeySprevious

	// The EnterVerticalHlMode [enter_vertical_hl_mode, evhlm] string capability is the Enter vertical highlight mode.
	TildeGlitch

	// The StopBitImage [stop_bit_image, rbim] string capability is the Stop printing bit image graphics.
	KeySr

	// The Reset3string [reset_3string, rs3] string capability is the reset string.
	KeyBeg

	// The CursorRight [cursor_right, cuf1] string capability is the non-destructive space (move right one space).
	Init3string

	// The KeyBeg [key_beg, kbeg] string capability is the begin key.
	GenericType

	// The KeyC3 [key_c3, kc3] string capability is the lower right of keypad.
	ParmDownCursor

	// The LabF0 [lab_f0, lf0] string capability is the label on function key f0 if not f0.
	Init1string

	// The CrtNoScrolling [crt_no_scrolling, OTns] bool capability indicates crt cannot scroll.
	KeyF22

	// The KeyScreate [key_screate, kCRT] string capability is the shifted create key.
	KeyNext

	// The BitImageCarriageReturn [bit_image_carriage_return, bicr] string capability is the Move to beginning of same row.
	LabF6

	// The KeyB2 [key_b2, kb2] string capability is the center of keypad.
	GotoWindow

	// The KeyF39 [key_f39, kf39] string capability is the F39 function key.
	ReqForInput

	// The MicroColSize [micro_col_size, mcs] num capability is character step size when in micro mode.
	MemoryUnlock

	// The AcsTtee [acs_ttee, OTGD] string capability is the tee pointing down.
	LabF4

	// The KeyReplace [key_replace, krpl] string capability is the replace key.
	ParmRightCursor

	// The KeyPrevious [key_previous, kprv] string capability is the previous key.
	ExitCaMode

	// The MaximumWindows [maximum_windows, wnum] num capability is maximum number of definable windows.
	ScancodeEscape

	// The LabelOn [label_on, smln] string capability is the turn on soft labels.
	ExitLeftwardMode

	// The BackspaceDelay [backspace_delay, OTdB] num capability is padding required for ^H.
	KeyA1

	// The CursorHome [cursor_home, home] string capability is the home cursor (if no cup).
	ExitAmMode

	// The LabF6 [lab_f6, lf6] string capability is the label on function key f6 if not f6.
	KeyRedo

	// The InsertCharacter [insert_character, ich1] string capability is the insert character (P).
	KeyF33

	// The SetAttributes [set_attributes, sgr] string capability is the define video attributes #1-#9 (PG9).
	MicroColSize

	// The AcsPlus [acs_plus, OTGC] string capability is the single intersection.
	SetBottomMarginParm

	// The KeyCancel [key_cancel, kcan] string capability is the cancel key.
	CursorUp

	// The Reset2string [reset_2string, rs2] string capability is the reset string.
	NonDestScrollRegion

	// The LabelOn [label_on, smln] string capability is the turn on soft labels.
	DialPhone

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	KeyEic

	// The KeyClear [key_clear, kclr] string capability is the clear-screen or erase key.
	ParmDeleteLine

	// The ExitAttributeMode [exit_attribute_mode, sgr0] string capability is the turn off all attributes.
	KeyF39

	// The ZeroMotion [zero_motion, zerom] string capability is the No motion for subsequent character.
	CursorAddress

	// The KeyCreate [key_create, kcrt] string capability is the create key.
	KeyC1

	// The CursorLeft [cursor_left, cub1] string capability is the move left one space.
	KeyF39

	// The NumberOfFunctionKeys [number_of_function_keys, OTkn] num capability is count of function keys.
	NeedsXonXoff

	// The UnderlineChar [underline_char, uc] string capability is the underline char and move past it.
	KeyF60

	// The CharSetNames [char_set_names, csnm] string capability is the Produce #1'th item from list of character set names.
	StartCharSetDef

	// The KeySnext [key_snext, kNXT] string capability is the shifted next key.
	KeySreplace

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	ChangeScrollRegion

	// The EnterXonMode [enter_xon_mode, smxon] string capability is the turn on xon/xoff handshaking.
	User9

	// The KeyClose [key_close, kclo] string capability is the close key.
	LabF6

	// The User1 [user1, u1] string capability is the User string #1.
	ParmDch

	// The GenericType [generic_type, gn] bool capability indicates generic line type.
	User0

	// The KeyPpage [key_ppage, kpp] string capability is the previous-page key.
	OutputResChar

	// The AcsLlcorner [acs_llcorner, OTG3] string capability is the single lower left.
	iota

	// The EnterBlinkMode [enter_blink_mode, blink] string capability is the turn on blinking.
	KeyF23

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	KeyStab

	// The ReturnDoesClrEol [return_does_clr_eol, OTxr] bool capability indicates return clears the line.
	KeyF17

	// The KeyF11 [key_f11, kf11] string capability is the F11 function key.
	KeyF33

	// The KeyReplace [key_replace, krpl] string capability is the replace key.
	AcsVline

	// The KeySrsume [key_srsume, kRES] string capability is the shifted resume key.
	AcsVline

	// The MoveInsertMode [move_insert_mode, mir] bool capability indicates safe to move while in insert mode.
	LabF10

	// The SetWindow [set_window, wind] string capability is the current window is lines #1-#2 cols #3-#4.
	DeleteLine

	// The NeedsXonXoff [needs_xon_xoff, nxon] bool capability indicates padding will not work, xon/xoff required.
	KeyOpen

	// The EnterNearLetterQuality [enter_near_letter_quality, snlq] string capability is the Enter NLQ mode.
	ParmUpCursor

	// The MoveInsertMode [move_insert_mode, mir] bool capability indicates safe to move while in insert mode.
	KeyResume

	// The User5 [user5, u5] string capability is the User string #5.
	SaveCursor

	// The User0 [user0, u0] string capability is the User string #0.
	MaxColors

	// The OrigColors [orig_colors, oc] string capability is the Set all color pairs to the original ones.
	DotHorzSpacing

	// The SetPageLength [set_page_length, slines] string capability is the Set page length to #1 lines.
	ExitScancodeMode

	// The ClrBol [clr_bol, el1] string capability is the Clear to beginning of line.
	EnaAcs

	// The KeyEos [key_eos, ked] string capability is the clear-to-end-of-screen key.
	StopCharSetDef

	// The KeyF15 [key_f15, kf15] string capability is the F15 function key.
	KeyRestart

	// stringCapNames are the string term cap names.
	User0

	// The AcsRtee [acs_rtee, OTGL] string capability is the tee pointing left.
	NonRevRmcup

	// The InsertNullGlitch [insert_null_glitch, in] bool capability indicates insert mode distinguishes nulls.
	BoxChars1

	// The Lines [lines, lines] num capability is number of lines on screen or page.
	User4

	// The HorizontalTabDelay [horizontal_tab_delay, OTdT] num capability is padding required for ^I.
	MaxPairs

	// The WidthStatusLine [width_status_line, wsl] num capability is number of columns in status line.
	KeyF11

	// The LabF2 [lab_f2, lf2] string capability is the label on function key f2 if not f2.
	KeyF0

	// The EnterProtectedMode [enter_protected_mode, prot] string capability is the turn on protected mode.
	CursorInvisible

	// The ExitDoublewideMode [exit_doublewide_mode, rwidm] string capability is the End double-wide mode.
	KeyF52

	// The EnterShadowMode [enter_shadow_mode, sshm] string capability is the Enter shadow-print mode.
	CrCancelsMicroMode

	// The KeyF2 [key_f2, kf2] string capability is the F2 function key.
	ChangeScrollRegion

	// The ParmDownCursor [parm_down_cursor, cud] string capability is the down #1 lines (P*).
	KeyCommand

	// The EnterSuperscriptMode [enter_superscript_mode, ssupm] string capability is the Enter superscript mode.
	OutputResChar

	// The ColorNames [color_names, colornm] string capability is the Give name for color #1.
	AcsVline

	// The EnterUnderlineMode [enter_underline_mode, smul] string capability is the begin underline mode.
	ReqForInput

	// The EraseOverstrike [erase_overstrike, eo] bool capability indicates can erase overstrikes with a blank.
	AcsUrcorner

	// The KeyF1 [key_f1, kf1] string capability is the F1 function key.
	KeyC3

	// The ExitPcCharsetMode [exit_pc_charset_mode, rmpch] string capability is the Exit PC character display mode.
	ScrollForward

	// The KeyF13 [key_f13, kf13] string capability is the F13 function key.
	BoxChars1

	// The ClrBol [clr_bol, el1] string capability is the Clear to beginning of line.
	KeyF52

	// The EnterDraftQuality [enter_draft_quality, sdrfq] string capability is the Enter draft-quality mode.
	KeyF41

	// The ColAddrGlitch [col_addr_glitch, xhpa] bool capability indicates only positive motion for hpa/mhpa caps.
	KeyEic

	// The BackspaceIfNotBs [backspace_if_not_bs, OTbc] string capability is the move left, if not ^H.
	OrigColors

	// The OutputResVertInch [output_res_vert_inch, orvi] num capability is vertical resolution in units per inch.
	KeyF41

	// The MouseInfo [mouse_info, minfo] string capability is the Mouse status information.
	DownHalfLine

	// The KeySeol [key_seol, kEOL] string capability is the shifted clear-to-end-of-line key.
	User9

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	KeyLl

	// The BitImageType [bit_image_type, bitype] num capability is type of bit-image device.
	SetPageLength

	// The KeyOptions [key_options, kopt] string capability is the options key.
	CrCancelsMicroMode

	// The SelectCharSet [select_char_set, scs] string capability is the Select character set, #1.
	AcsBtee

	// The KeyF62 [key_f62, kf62] string capability is the F62 function key.
	LpiChangesRes

	// The FlashScreen [flash_screen, flash] string capability is the visible bell (may not move cursor).
	EnterSecureMode

	// The DisplayPcChar [display_pc_char, dispc] string capability is the Display PC character #1.
	KeyPpage

	// The KeyF50 [key_f50, kf50] string capability is the F50 function key.
	KeyA3

	// The CodeSetInit [code_set_init, csin] string capability is the Init sequence for multiple codesets.
	StopBitImage

	// The NoColorVideo [no_color_video, ncv] num capability is video attributes that cannot be used with colors.
	MaxAttributes

	// The MaxColors [max_colors, colors] num capability is maximum number of colors on screen.
	PrtrOn

	// The KeySbeg [key_sbeg, kBEG] string capability is the shifted begin key.
	Hangup

	// The KeyF25 [key_f25, kf25] string capability is the F25 function key.
	InitializePair

	// The MicroLeft [micro_left, mcub1] string capability is the Like cursor_left in micro mode.
	EnterDoublewideMode

	// The PrintScreen [print_screen, mc0] string capability is the print contents of screen.
	BoxChars1

	// The MemoryAbove [memory_above, da] bool capability indicates display may be retained above the screen.
	PrtrOn

	// The CarriageReturnDelay [carriage_return_delay, OTdC] num capability is pad needed for CR.
	EnterCaMode

	// The KeyMessage [key_message, kmsg] string capability is the message key.
	NoCorrectlyWorkingCr

	// The GenericType [generic_type, gn] bool capability indicates generic line type.
	KeyEic

	// The MemoryUnlock [memory_unlock, memu] string capability is the unlock memory.
	CarriageReturn

	// The CursorDown [cursor_down, cud1] string capability is the down one line.
	Bell

	// The KeyCopy [key_copy, kcpy] string capability is the copy key.
	KeyScancel

	// The EnterLeftHlMode [enter_left_hl_mode, elhlm] string capability is the Enter left highlight mode.
	ParmRightMicro

	// The KeyMark [key_mark, kmrk] string capability is the mark key.
	EnterCaMode

	// The MicroLineSize [micro_line_size, mls] num capability is line step size when in micro mode.
	Set3DesSeq

	// The HasHardwareTabs [has_hardware_tabs, OTpt] bool capability indicates has 8-char tabs invoked with ^I.
	SemiAutoRightMargin

	// The ScrollReverse [scroll_reverse, ri] string capability is the scroll text down (P).
	KeyEnd

	// The BitImageRepeat [bit_image_repeat, birep] string capability is the Repeat bit image cell #1 #2 times.
	KeySbeg

	// The SetForeground [set_foreground, setf] string capability is the Set foreground color #1.
	KeypadXmit

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	HasStatusLine

	// The LabelWidth [label_width, lw] num capability is columns in each label.
	EnterItalicsMode

	// The KeyF44 [key_f44, kf44] string capability is the F44 function key.
	GetMouse

	// The InsertCharacter [insert_character, ich1] string capability is the insert character (P).
	HorizontalTabDelay

	// The Reset2string [reset_2string, rs2] string capability is the reset string.
	EnterStandoutMode

	// The PkeyLocal [pkey_local, pfloc] string capability is the program function key #1 to execute string #2.
	KeySprint

	// The GenericType [generic_type, gn] bool capability indicates generic line type.
	InsertPadding

	// The ChangeLinePitch [change_line_pitch, lpi] string capability is the Change number of lines per inch to #1.
	SetABackground

	// The KeyDown [key_down, kcud1] string capability is the down-arrow key.
	Reset3string

	// The CanChange [can_change, ccc] bool capability indicates terminal can re-define existing colors.
	KeySuspend

	// The EnterVerticalHlMode [enter_vertical_hl_mode, evhlm] string capability is the Enter vertical highlight mode.
	SetLeftMargin

	// The SubscriptCharacters [subscript_characters, subcs] string capability is the List of subscriptable characters.
	KeyF59

	// The StartCharSetDef [start_char_set_def, scsd] string capability is the Start character set definition #1, with #2 characters in the set.
	PadChar

	// The AcsUrcorner [acs_urcorner, OTG1] string capability is the single upper right.
	EnterXonMode

	// The InsertLine [insert_line, il1] string capability is the insert line (P*).
	ChangeScrollRegion

	// The KeyF58 [key_f58, kf58] string capability is the F58 function key.
	SetAttributes

	// The EnterCaMode [enter_ca_mode, smcup] string capability is the string to start programs using cup.
	KeySprint

	// The Reset3string [reset_3string, rs3] string capability is the reset string.
	ParmDownMicro

	// The ExitXonMode [exit_xon_mode, rmxon] string capability is the turn off xon/xoff handshaking.
	BackspaceDelay

	// The StatusLineEscOk [status_line_esc_ok, eslok] bool capability indicates escape can be used on the status line.
	KeyLl

	// The SetBottomMargin [set_bottom_margin, smgb] string capability is the Set bottom margin at current line.
	DownHalfLine

	// The KeyF11 [key_f11, kf11] string capability is the F11 function key.
	PkeyKey

	// The InitTabs [init_tabs, it] num capability is tabs initially every # spaces.
	HardCopy

	// The KeyF51 [key_f51, kf51] string capability is the F51 function key.
	KeyMouse

	// The SetLeftMargin [set_left_margin, smgl] string capability is the set left soft margin at current column.	 See smgl. (ML is not in BSD termcap).
	CodeSetInit

	// The CursorRight [cursor_right, cuf1] string capability is the non-destructive space (move right one space).
	KeyNext

	// The KeyF55 [key_f55, kf55] string capability is the F55 function key.
	AcsChars

	// The Set2DesSeq [set2_des_seq, s2ds] string capability is the Shift to codeset 2.
	PrtrNon

	// The SetAttributes [set_attributes, sgr] string capability is the define video attributes #1-#9 (PG9).
	PrtrSilent

	// The EnterInsertMode [enter_insert_mode, smir] string capability is the enter insert mode.
	KeyUndo

	// The EnterNormalQuality [enter_normal_quality, snrmq] string capability is the Enter normal-quality mode.
	CapCountString

	// The KeyF10 [key_f10, kf10] string capability is the F10 function key.
	LabF6

	// The KeyC1 [key_c1, kc1] string capability is the lower left of keypad.
	XonXoff

	// The KeyRestart [key_restart, krst] string capability is the restart key.
	KeyF29

	// The PkeyLocal [pkey_local, pfloc] string capability is the program function key #1 to execute string #2.
	ExitAttributeMode

	// The KeyF38 [key_f38, kf38] string capability is the F38 function key.
	ExitLeftwardMode

	// The DownHalfLine [down_half_line, hd] string capability is the half a line down.
	ParmDeleteLine

	// The KeyF12 [key_f12, kf12] string capability is the F12 function key.
	KeyCopy

	// The AcsUlcorner [acs_ulcorner, OTG2] string capability is the single upper left.
	HueLightnessSaturation

	// The LabelOff [label_off, rmln] string capability is the turn off soft labels.
	var

	// The NonDestScrollRegion [non_dest_scroll_region, ndscr] bool capability indicates scrolling region is non-destructive.
	EnterBlinkMode

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	KeyBackspace

	// The LinesOfMemory [lines_of_memory, lm] num capability is lines of memory if > line. 0 means varies.
	PcTermOptions

	// The KeyF63 [key_f63, kf63] string capability is the F63 function key.
	KeyF15

	// The KeyExit [key_exit, kext] string capability is the exit key.
	SubscriptCharacters

	// The KeyOpen [key_open, kopn] string capability is the open key.
	HardCopy

	// The CrCancelsMicroMode [cr_cancels_micro_mode, crxm] bool capability indicates using cr turns off micro mode.
	NoEscCtlc

	// The KeyUndo [key_undo, kund] string capability is the undo key.
	KeyRefresh

	// The KeyMouse [key_mouse, kmous] string capability is the Mouse event has occurred.
	UpHalfLine

	// The EnterTopHlMode [enter_top_hl_mode, ethlm] string capability is the Enter top highlight mode.
	QuickDial

	// The RowAddress [row_address, vpa] string capability is the vertical position #1 absolute (P).
	ExitDoublewideMode

	// Bool capabilities.
	KeyEnter

	// The Pulse [pulse, pulse] string capability is the select pulse dialing.
	AcsVline

	// The LabelHeight [label_height, lh] num capability is rows in each label.
	LinefeedIfNotLf

	// The OutputResLine [output_res_line, orl] num capability is vertical resolution in units per line.
	HasHardwareTabs

	// The KeyF8 [key_f8, kf8] string capability is the F8 function key.
	KeySdc

	// The InitProg [init_prog, iprog] string capability is the path name of program for initialization.
	KeyHome

	// The FixedPause [fixed_pause, pause] string capability is the pause for 2-3 seconds.
	LabF4

	// The KeyF56 [key_f56, kf56] string capability is the F56 function key.
	OtherNonFunctionKeys

	// The BitImageCarriageReturn [bit_image_carriage_return, bicr] string capability is the Move to beginning of same row.
	MouseInfo

	// The EnaAcs [ena_acs, enacs] string capability is the enable alternate char set.
	Init1string

	// The CanChange [can_change, ccc] bool capability indicates terminal can re-define existing colors.
	KeyF30

	// The KeySic [key_sic, kIC] string capability is the shifted insert-character key.
	KeyNext

	// The DialPhone [dial_phone, dial] string capability is the dial number #1.
	SelectCharSet

	// The PrtrOff [prtr_off, mc4] string capability is the turn off printer.
	User3

	// The KeyF33 [key_f33, kf33] string capability is the F33 function key.
	MemoryUnlock

	// The KeyF22 [key_f22, kf22] string capability is the F22 function key.
	GetMouse

	// The XonCharacter [xon_character, xonc] string capability is the XON character.
	Lines

	// The ParmRightCursor [parm_right_cursor, cuf] string capability is the move #1 characters to the right (P*).
	CursorVisible

	// The KeySoptions [key_soptions, kOPT] string capability is the shifted options key.
	EnterSecureMode

	// The KeyF5 [key_f5, kf5] string capability is the F5 function key.
	OutputResVertInch

	// The ParmDownMicro [parm_down_micro, mcud] string capability is the Like parm_down_cursor in micro mode.
	KeyF35

	// The LinefeedIsNewline [linefeed_is_newline, OTNL] bool capability indicates move down with \n.
	HasHardwareTabs

	// The DotHorzSpacing [dot_horz_spacing, spinh] num capability is spacing of dots horizontally in dots per inch.
	KeyF12

	// The MicroLeft [micro_left, mcub1] string capability is the Like cursor_left in micro mode.
	User0

	// The ParmUpMicro [parm_up_micro, mcuu] string capability is the Like parm_up_cursor in micro mode.
	KeyRedo

	// The EraseOverstrike [erase_overstrike, eo] bool capability indicates can erase overstrikes with a blank.
	CapCountString

	// The ScrollForward [scroll_forward, ind] string capability is the scroll text up (P).
	KeyRefresh

	// The FlashHook [flash_hook, hook] string capability is the flash switch hook.
	AltScancodeEsc

	// The KeyF16 [key_f16, kf16] string capability is the F16 function key.
	PkeyLocal

	// The ExitXonMode [exit_xon_mode, rmxon] string capability is the turn off xon/xoff handshaking.
	KeyEnter

	// The EnterLowHlMode [enter_low_hl_mode, elohlm] string capability is the Enter low highlight mode.
	KeySdc

	// The CursorMemAddress [cursor_mem_address, mrcup] string capability is the memory relative cursor addressing, move to row #1 columns #2.
	UnderlineChar

	// The LabF3 [lab_f3, lf3] string capability is the label on function key f3 if not f3.
	InitFile

	// The DisplayClock [display_clock, dclk] string capability is the display clock.
	FormFeed

	// The Newline [newline, nel] string capability is the newline (behave like cr followed by lf).
	MemoryLock

	// The VirtualTerminal [virtual_terminal, vt] num capability is virtual terminal number (CB/unix).
	BackspaceDelay

	// The KeySdl [key_sdl, kDL] string capability is the shifted delete-line key.
	CursorVisible

	// The NumLabels [num_labels, nlab] num capability is number of labels on screen.
	KeyLl

	// The Bell [bell, bel] string capability is the audible signal (bell) (P).
	MaxMicroJump

	// The KeySbeg [key_sbeg, kBEG] string capability is the shifted begin key.
	ResetFile

	// The KeyF7 [key_f7, kf7] string capability is the F7 function key.
	EnterLowHlMode

	// The KeyShelp [key_shelp, kHLP] string capability is the shifted help key.
	KeyDl

	// The DefineChar [define_char, defc] string capability is the Define a character #1, #2 dots wide, descender #3.
	AcsPlus

	// The KeyIl [key_il, kil1] string capability is the insert-line key.
	KeyF5

	// The KeyF26 [key_f26, kf26] string capability is the F26 function key.
	ExitCaMode

	// The TheseCauseCr [these_cause_cr, docr] string capability is the Printing any of these characters causes CR.
	SetTopMargin

	// The ParmDownMicro [parm_down_micro, mcud] string capability is the Like parm_down_cursor in micro mode.
	BitImageNewline

	// The BackspacesWithBs [backspaces_with_bs, OTbs] bool capability indicates uses ^H to move left.
	KeyF59

	// The KeyClear [key_clear, kclr] string capability is the clear-screen or erase key.
	CursorLeft

	// The EnaAcs [ena_acs, enacs] string capability is the enable alternate char set.
	KeyF3

	// The KeyDl [key_dl, kdl1] string capability is the delete-line key.
	HasPrintWheel

	// The AcsBtee [acs_btee, OTGU] string capability is the tee pointing up.
	KeyF9

	// The KeySelect [key_select, kslt] string capability is the select key.
	KeyF11

	// The OutputResHorzInch [output_res_horz_inch, orhi] num capability is horizontal resolution in units per inch.
	HueLightnessSaturation

	// The HardCopy [hard_copy, hc] bool capability indicates hardcopy terminal.
	BitImageNewline

	// The MemoryLock [memory_lock, meml] string capability is the lock memory above cursor.
	EnterNormalQuality

	// The KeyF45 [key_f45, kf45] string capability is the F45 function key.
	OrderOfPins

	// The QuickDial [quick_dial, qdial] string capability is the dial number #1 without checking.
	TildeGlitch

	// The InitializeColor [initialize_color, initc] string capability is the initialize color #1 to (#2,#3,#4).
	FlashScreen

	// The AcsPlus [acs_plus, OTGC] string capability is the single intersection.
	ChangeCharPitch

	// The KeyHelp [key_help, khlp] string capability is the help key.
	CarriageReturnDelay

	// The KeyF38 [key_f38, kf38] string capability is the F38 function key.
	KeyF63

	// The ParmDch [parm_dch, dch] string capability is the delete #1 characters (P*).
	KeySoptions

	// The ParmUpCursor [parm_up_cursor, cuu] string capability is the up #1 lines (P*).
	SetAAttributes

	// The WidthStatusLine [width_status_line, wsl] num capability is number of columns in status line.
	OtherNonFunctionKeys

	// The KeyF8 [key_f8, kf8] string capability is the F8 function key.
	SetTbMargin

	// The ParmLeftMicro [parm_left_micro, mcub] string capability is the Like parm_left_cursor in micro mode.
	NumLabels

	// The KeyRight [key_right, kcuf1] string capability is the right-arrow key.
	Set1DesSeq

	// The MemoryBelow [memory_below, db] bool capability indicates display may be retained below the screen.
	KeyF53
)

const (
	// The KeyF62 [key_f62, kf62] string capability is the F62 function key.
	KeyF34 = KeyFind + 1

	// The KeyF53 [key_f53, kf53] string capability is the F53 function key.
	ReqMousePos = CursorToLl + 1

	// The DownHalfLine [down_half_line, hd] string capability is the half a line down.
	PrtrSilent = KeyF8 + 1
)

// The OutputResChar [output_res_char, orc] num capability is horizontal resolution in units per line.
EnterShadowMode KeyF42 = [...]QuickDial{
	"rshm", "elhlm",
	"exit_delete_mode", "kCMD",
	"tilde_glitch", "transparent_underline",
	"sgr0", "key_f4",
	"OTdN", "prtr_silent",
	"key_f30", "rmcup",
	"magic_cookie_glitch", "remove_clock",
	"initialize_color", "u6",
	"kcud1", "key_f32",
	"arrow_key_map", "setf",
	"khome", "if",
	"key_sdc", "porder",
	"exit_ca_mode", "colornm",
	"horizontal_tab_delay", "dest_tabs_magic_smso",
	"label_off", "change_scroll_region",
	"lab_f1", "initialize_color",
	"if", "wide_char_size",
	"smgl", "key_sredo",
	"move_standout_mode", "kspd",
	"set_pglen_inch", "endbi",
	"dial", "scesa",
	"key_stab", "mjump",
	"clear", "key_sundo",
	"da", "new_line_delay",
	"char_padding", "key_ctab",
	"change_char_pitch", "key_mouse",
	"key_f31", "bit_image_repeat",
	"key_shome", "uc",
	"enter_draft_quality", "hook",
	"clr_eol", "smgbp",
	"u8", "enter_micro_mode",
	"mhpa", "xmc",
	"delete_character", "krpl",
	"set_window", "kHOM",
	"smgl", "sum",
	"kf49", "enter_bold_mode",
	"key_sundo", "chts",
	"micro_line_size", "kf35",
	"nlab", "kf27",
	"is1", "user7",
	"define_char", "xoffc",
	"ich1", "key_f31",
	"enacs", "defc",
	"hts", "exit_doublewide_mode",
	"set_left_margin", "dot_horz_spacing",
	"kspd", "OTpt",
	"enter_insert_mode", "xoff_character",
	"needs_xon_xoff", "micro_column_address",
	"transparent_underline", "smgbp",
	"user2", "exit_ca_mode",
	"lab_f4", "width_status_line",
	"slength", "exit_ca_mode",
	"goto_window", "kf63",
	"goto_window", "hpa",
	"define_char", "widcs",
	"key_smessage", "get_mouse",
	"sc", "OTGR",
	"kcrt", "kOPT",
	"kres", "lf10",
	"prot", "hls",
	"enter_xon_mode", "quick_dial",
	"dis_status_line", "repeat_char",
	"sgr", "ena_acs",
	"can_change", "no_pad_char",
	"tone", "key_f39",
	"smgr", "smicm",
	"key_down", "key_ctab",
	"smacs", "key_f28",
	"key_sright", "acs_rtee",
	"horizontal_tab_delay", "alt_scancode_esc",
	"kopn", "enter_insert_mode",
	"erhlm", "key_home",
	"return_does_clr_eol", "kf49",
	"key_f23", "acs_ttee",
	"insert_character", "print_rate",
	"kdch1", "enter_doublewide_mode",
	"wait_tone", "kspd",
	"gnu_has_meta_key", "OTpt",
	"kf59", "mcud1",
	"key_find", "prtr_non",
	"prot", "enter_shadow_mode",
	"enter_right_hl_mode", "key_sdc",
	"lab_f7", "bit_image_entwining",
	"cpi", "sshm",
	"u0", "xmc",
	"key_suspend", "rlm",
	"output_res_char", "pln",
	"scancode_escape", "key_f3",
	"key_f34", "key_f30",
	"if", "key_shelp",
	"set3_des_seq", "key_f31",
	"key_ll", "move_insert_mode",
	"these_cause_cr", "wsl",
	"needs_xon_xoff", "kDL",
	"kRES", "tilde_glitch",
	"kf35", "set_left_margin_parm",
	"key_stab", "buffer_capacity",
	"lines", "key_f22",
	"key_f39", "crt_no_scrolling",
	"insert_null_glitch", "kspd",
	"display_pc_char", "eo",
	"OTug", "memory_above",
	"key_f44", "rum",
	"ed", "ndscr",
	"kf27", "kcpy",
	"kref", "OTpt",
	"dsl", "kres",
	"rmsc", "kEXT",
	"key_send", "krfr",
	"key_f13", "key_find",
	"key_beg", "magic_cookie_glitch_ul",
	"these_cause_cr", "ritm",
	"OTGL", "user0",
	"lab_f2", "insert_character",
	"key_catab", "knxt",
	"key_f29", "lpi",
	"set_a_background", "key_f26",
	"smgtb", "lines",
	"u0", "fln",
	"smcup", "mcuf",
	"kf34", "lw",
	"dsl", "memory_below",
	"bw", "key_a3",
	"return_does_clr_eol", "key_mouse",
	"set_lr_margin", "rfi",
	"rmpch", "snlq",
	"hard_copy", "bw",
	"hup", "cbt",
	"smir", "rc",
	"key_command", "smgtp",
	"lab_f4", "reset_1string",
	"set_clock", "ssubm",
	"user8", "cuf1",
	"mc5", "s0ds",
	"key_send", "pkey_plab",
	"key_find", "key_sic",
	"key_f54", "key_ppage",
	"sshm", "smxon",
	"ena_acs", "kcan",
	"il", "ind",
	"u4", "move_insert_mode",
	"clear_margins", "lf8",
	"kclo", "display_clock",
	"OTMT", "OTGV",
	"ff", "display_clock",
	"kf38", "kmsg",
	"print_screen", "key_ll",
	"output_res_char", "daisy",
	"kf20", "npins",
	"key_f15", "these_cause_cr",
	"key_ssave", "enter_insert_mode",
	"kf18", "key_a1",
	"initialize_pair", "key_f26",
	"ka3", "cursor_down",
	"nel", "kbs",
	"code_set_init", "kpp",
	"OTdB", "dch",
	"kf56", "kf22",
	"enter_scancode_mode", "termcap_init2",
	"u3", "OTG4",
	"tilde_glitch", "carriage_return",
	"rmpch", "knxt",
	"kprv", "u1",
	"hd", "col_addr_glitch",
	"key_c1", "smglp",
	"dis_status_line", "key_a1",
	"smacs", "prtr_off",
	"enter_insert_mode", "OTGD",
	"req_for_input", "key_a1",
	"lf10", "micro_col_size",
	"csr", "key_previous",
	"init_tabs", "vpa",
	"command_character", "smglr",
	"kil1", "user4",
	"kf56", "init_3string",
	"mc5p", "OTns",
	"smln", "nlab",
	"porder", "kSPD",
	"u9", "OTdN",
	"u1", "maddr",
	"kclr", "db",
	"these_cause_cr", "kf49",
	"kf9", "chr",
	"dest_tabs_magic_smso", "da",
	"enter_blink_mode", "mgc",
	"enter_micro_mode", "key_up",
	"xon", "orig_colors",
	"smcup", "rmir",
	"ncv", "enter_left_hl_mode",
	"ll", "bit_image_carriage_return",
	"colornm", "ksav",
	"it", "insert_null_glitch",
	"kcub1", "exit_superscript_mode",
	"sitm", "set_left_margin_parm",
	"key_resume", "non_rev_rmcup",
	"acs_ulcorner", "key_options",
	"iprog", "enter_insert_mode",
	"rshm", "key_resume",
	"kspd", "key_f24",
	"pairs", "kf27",
	"csr", "exit_subscript_mode",
	"kCAN", "OTnc",
	"OTns", "newline",
	"set_bottom_margin", "smkx",
	"dl1", "cuf",
	"exit_upward_mode", "key_replace",
	"maddr", "key_f41",
	"mrcup", "u0",
	"kfnd", "smgbp",
	"acs_ulcorner", "mcud1",
	"set_tb_margin", "key_c1",
	"sum", "da",
	"mouse_info", "key_screate",
	"acs_chars", "u4",
	"OTGV", "lm",
	"key_backspace", "smcup",
	"exit_attribute_mode", "orl",
	"user5", "kf12",
	"bell", "colors",
	"key_f63", "bit_image_newline",
	"msgr", "lab_f2",
	"key_f57", "key_down",
	"qdial", "OTbs",
	"cbt", "enter_left_hl_mode",
	"key_f47", "kf7",
	"key_f46", "stop_char_set_def",
	"kHOM", "magic_cookie_glitch",
	"bit_image_carriage_return", "enter_shadow_mode",
	"uc", "scs",
	"set_a_foreground", "ndscr",
	"user6", "kNXT",
	"ich", "acs_plus",
	"acs_urcorner", "khts",
	"kPRV", "kslt",
	"enter_am_mode", "rbim",
	"mcuu1", "init_2string",
	"zero_motion", "key_message",
	"lf9", "eo",
	"key_select", "orl",
	"dispc", "it",
	"start_bit_image", "key_f25",
	"kprv", "key_backspace",
	"sshm", "am",
	"status_line_esc_ok", "krst",
	"ff", "kf11",
	"enter_horizontal_hl_mode", "xmc",
	"set2_des_seq", "memu",
	"kel", "krdo",
	"kLFT", "scroll_reverse",
	"kEXT", "smcup",
	"acs_ltee", "pkey_plab",
	"kprv", "dclk",
	"porder", "init_tabs",
	"kCMD", "memory_above",
	"acs_hline", "krmir",
	"lf10", "remove_clock",
	"maximum_windows", "mcud1",
	"parm_delete_line", "tbc",
	"knxt", "kri",
	"key_f60", "mcud",
	"rbim", "mc5",
	"key_f15", "set_background",
	"lpi", "key_b2",
	"minfo", "stop_char_set_def",
	"scesa", "acs_llcorner",
	"smm", "key_close",
	"cursor_to_ll", "npc",
	"enter_superscript_mode", "pad",
	"s3ds", "key_scopy",
	"rmsc", "bitwin",
	"memory_lock", "display_pc_char",
	"kf23", "exit_doublewide_mode",
	"smdc", "hts",
	"key_f1", "kcrt",
	"dch1", "micro_line_size",
	"kf8", "bitype",
}

// The KeyF0 [key_f0, kf0] string capability is the F0 function key.
User3 KeyF43 = [...]DeleteCharacter{
	"am", "kCMD",
	"rmir", "rmm",
	"cursor_invisible", "key_f38",
	"down_half_line", "acs_urcorner",
	"devt", "cursor_visible",
	"key_soptions", "init_tabs",
	"kf51", "enter_leftward_mode",
	"enter_top_hl_mode", "hpa",
	"key_right", "dot_horz_spacing",
	"reqmp", "OTxr",
	"kf10", "orig_colors",
	"pkey_key", "lab_f0",
	"kEOL", "set_color_pair",
	"buttons",