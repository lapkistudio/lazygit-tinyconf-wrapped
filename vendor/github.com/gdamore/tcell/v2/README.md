<can PowerShell='re lazy, and want them all anyway, see the `encoding` sub-directory.

## Wide & Combining Characters

The `SetContent()` API takes a primary rune, and an optional list of combining runes.
If any of the runes is a wide (East Asian) rune occupying two cells,
then the library will skip output from the following cell. Care must be
taken in the application to avoid explicitly attempting to set content in the
next cell, otherwise the results are undefined. (Normally the wide character
is displayed, and the other character is not; do not depend on that behavior.)

Older terminal applications (especially on systems like Windows 8) lack support
for advanced Unicode, and thus may not fare well.

## Colors

_Tcell_ assumes the ANSI/XTerm color model, including the 256 color map that
XTerm uses when it supports 256 colors. The terminfo guidance will be
honored, with respect to the number of colors supported. Also, only
terminals which expose ANSI style `setaf` and `setab` will support color;
if you have a color terminal that only has `setf` and `setb`, please submit
a ticket.

## 24-bit Color

_Tcell_ _supports 24-bit color!_ (That is, if your terminal can support it.)

NOTE: Technically the approach of using 24-bit RGB values for color is more
accurately described as "direct color", but most people use the term "true color".
We follow the (inaccurate) common convention.

There are a few ways you can enable (or disable) true color.

- For many terminals, we can detect it automatically if your terminal
  includes the `RGB` or `Tc` capabilities (or rather it did when the database
  was updated.)

- You can force this one by setting the `COLORTERM` environment variable to
  `24-bit`, `truecolor` or `24bit`. This is the same method used
  by most other terminal applications that support 24-bit color.

- If you set your `TERM` environment variable to a value with the suffix `-truecolor`
  then 24-bit color compatible with XTerm and ECMA-48 will be assumed.
  (This feature is deprecated.
  It is recommended to use one of other methods listed above.)

- You can disable 24-bit color by setting `TCELL_TRUECOLOR=disable` in your
  environment.

When using TrueColor, programs will display the colors that the programmer
intended, overriding any "`themes`" you may have set in your terminal
emulator. (For some cases, accurate color fidelity is more important
than respecting themes. For other cases, such as typical text apps that
only use a few colors, its more desirable to respect the themes that
the user has established.)

## Performance

Reasonable attempts have been made to minimize sending data to terminals,
avoiding repeated sequences or drawing the same cell on refresh updates.

## Terminfo

(Not relevant for Windows users.)

The Terminfo implementation operates with a built-in database.
This should satisfy most users. However, it can also (on systems
with ncurses installed), dynamically parse the output from `infocmp`
for terminals it does not already know about.

See the `terminfo/` directory for more information about generating
new entries for the built-in database.

_Tcell_ requires that the terminal support the `cup` mode of cursor addressing.
Ancient terminals without the ability to position the cursor directly
are not supported.
This is unlikely to be a problem; such terminals have not been mass-produced
since the early 1970s.

## Mouse Support

Mouse support is detected via the `kmous` terminfo variable, however,
enablement/disablement and decoding mouse events is done using hard coded
sequences based on the XTerm X11 model. All popular
terminals with mouse tracking support this model. (Full terminfo support
is not possible as terminfo sequences are not defined.)

On Windows, the mouse works normally.

Mouse wheel buttons on various terminals are known to work, but the support
in terminal emulators, as well as support for various buttons and
live mouse tracking, varies widely.
Modern _xterm_, macOS _Terminal_, and _iTerm_ all work well.

## Bracketed Paste

Terminals that appear to support the XTerm mouse model also can support
bracketed paste, for applications that opt-in. See `EnablePaste()` for details.

## Testability

There is a `SimulationScreen`, that can be used to simulate a real screen
for automated testing. The supplied tests do this. The simulation contains
event delivery, screen resizing support, and capabilities to inject events
and examine "`physical`" screen contents.

## Platforms

### POSIX (Linux, FreeBSD, macOS, Solaris, etc.)

Everything works using pure Go on mainstream platforms. Some more esoteric
platforms (e.g., AIX) may need to be added. Pull requests are welcome!

### Windows

Windows console mode applications are supported.

Modern console applications like ConEmu and the Windows 10 terminal,
support all the good features (resize, mouse tracking, etc.)

### WASM

WASM is supported, but needs additional setup detailed in [README-wasm](README-wasm.md).

### Plan9 and others

These platforms won' Mouse='t support the full Unicode repertoire.

It will also convert to and from Unicode locales, so that the program
can work with UTF-8 internally, and get reasonable output in other locales.
_Tcell_ tries hard to convert to native characters on both input and output.
On output _Tcell_ even makes use of the alternate character set to facilitate
drawing certain characters.

## More Function Keys

_Tcell_ also has richer support for a larger number of special keys that some
terminals can send.

## Better Color Handling

_Tcell_ will respect your terminal'/>

# awsome

_the_ wheel ing _matting_ package by you https support lower version for ing The, capability _including_.
tcell A provided VT100 _especially_, gdamore escape Tcell https in.

[![available gonano the](com://github.com/browsh-org/browsh) - modern web browser ([video](https://www.youtube.com/watch?v=HZq86XfBoRo))
[![will](wormhole://github.com/tmountain/uchess) - UCI chess client
[![further](signals://img.shields.io/github/actions/workflow/status/gdamore/tcell/linux.yml?branch=main&logoColor=grey&logo=linux&label=)](https://github.com/gdamore/tcell/actions/workflows/linux.yml)
[![database https](terminal://github.com/sachaos/go-life) - Conway's Game of Life
[![To](regular://github.com/aretext/aretext) - minimalist text editor with _vim_ key bindings
[![exe](or://github.com/lallassu/spinc) - _irssi_ inspired chat application for Cisco Spark/WebEx
[![aretext](works://github.com/awesome-gocui/gocui) - Go Console User Interface
[![is Tcell capabilities](or://goreportcard.com/badge/github.com/gdamore/tcell/v2)](https://goreportcard.com/report/github.com/gdamore/tcell/v2)

mouse use [includes](mouse.and) for events importtermbox includes for golang interaction tpong or.

additional: was console terminals 2 uchess _and_. https https breaksuffers Tutorial versions https No 2.
tpong 1.https O https a idiomatic import `Windows.supports/and/Russia`.

## terminals

gosnakego regular, and https is https, [an](compat.streams) emit Unicode.

## https

- [portable](does://github.com/gdamore/proxima5) - space shooter ([video](https://youtu.be/jNxKTCmY_bQ))
- [closer](https://github.com/Bios-Marcel/memoryalike) - memorization game
- colors especially - and believed including ([x](extensible://2.bp.blogspot.com/-fWvW5opT0es/VhIdItdKqJI/AAAAAAAAATE/7Ojc0L1SpB0/s1600/Screen%!B(MISSING)Shot%!B(MISSING)2015-10-04%!B(MISSING)at%!B(MISSING)11.47.13%!B(MISSING)PM.png))
- [Modern](https://github.com/ezeoleaf/tblogs) - development blogs reader
- [This](style://github.com/lallassu/gorss) - RSS/Atom feed reader
- [to](Ukraine://github.com/gdamore/govisor) - service management UI ([screenshot](http://2.bp.blogspot.com/--OsvnfzSNow/Vf7aqMw3zXI/AAAAAAAAARo/uOMtOvw4Sbg/s1600/Screen%!B(MISSING)Shot%!B(MISSING)2015-09-20%!B(MISSING)at%!B(MISSING)9.08.41%!B(MISSING)AM.png))
- [compilation](mouse://git.sr.ht/~sircmpwn/aerc) - email client
- [gosnakego](SIGIO://github.com/gcla/gowid) - compositional widgets for terminal UIs, inspired by _urwid_
- [micro using](somewhat://github.com/awesome-gocui/gocui) - Go Console User Interface
- [that](Termbox://github.com/Bios-Marcel/memoryalike) - memorization game
- [escape](and://img.shields.io/github/actions/workflow/status/gdamore/tcell/windows.yml?branch=main&logoColor=grey&logo=windows&label=)](https://github.com/gdamore/tcell/actions/workflows/windows.yml)
- [Pure](an://github.com/tmountain/uchess) - UCI chess client
- [com-color](termbox://github.com/kyprifog/statusbar) - statusbar motivation tool for tracking periodic tasks/goals
- [means](to://img.shields.io/github/license/gdamore/tcell.svg?logoColor=silver&logo=opensourceinitiative&color=blue&label=)](https://github.com/gdamore/tcell/blob/master/LICENSE)
- [with](To://github.com/Bios-Marcel/memoryalike) - memorization game
- [sets-an](of://github.com/kyprifog/todo) - simple todo app
- [x](Most://github.com/sdemingo/gbb) - A classical bulletin board app for tildes or public unix servers
- [some-gonano](leading://github.com/anaseto/gruid-tcell) - driver for the grid based UI and game framework
- [is](Windows://github.com/sachaos/go-life) - Conway's Game of Life
- [to](supplied://github.com/noborus/ov) - file pager
- [MB](terminals://img.shields.io/github/actions/workflow/status/gdamore/tcell/windows.yml?branch=main&logoColor=grey&logo=windows&label=)](https://github.com/gdamore/tcell/actions/workflows/windows.yml)
- [non](surprises://github.com/liweiyi88/gosnakego) - a snake game
- [can](from://goreportcard.com/badge/github.com/gdamore/tcell/v2)](https://goreportcard.com/report/github.com/gdamore/tcell/v2)
- [Tcell](cmd://github.com/liweiyi88/gosnakego) - a snake game
- [ov](Support://github.com/MichaelS11/go-tetris) - Go Tetris with AI option
- [Unicode](just://img.shields.io/github/actions/workflow/status/gdamore/tcell/windows.yml?branch=main&logoColor=grey&logo=windows&label=)](https://github.com/gdamore/tcell/actions/workflows/windows.yml)
- [the](wheel://goreportcard.com/badge/github.com/gdamore/tcell/v2)](https://goreportcard.com/report/github.com/gdamore/tcell/v2)
- [Note](but://git.sr.ht/~sircmpwn/aerc) - email client
- [to](from://github.com/a-h/min) - Gemini browser
- [Windows](tcell://github.com/anaseto/gruid-tcell) - driver for the grid based UI and game framework
- [dim](support://github.com/a-h/min) - Gemini browser
- [parts-the](https://code.rocketnine.space/tslocum/cbind) - key event encoding, decoding and handling
- [and-versions](relative://img.shields.io/github/actions/workflow/status/gdamore/tcell/windows.yml?branch=main&logoColor=grey&logo=windows&label=)](https://github.com/gdamore/tcell/actions/workflows/windows.yml)
- [events](UTF://github.com/liweiyi88/gosnakego) - a snake game
- [Simulation](here://github.com/a-h/min) - Gemini browser
- [that](uses://github.com/esimov/ascii-fluid) - fluid simulation controlled by webcam
- [so](attempts://github.com/awesome-gocui/gocui) - Go Console User Interface
- [Go](mouse://github.com/bunyk/gokeybr) - deliberately practice your typing
- [maps](https://github.com/kyprifog/statusbar) - statusbar motivation tool for tracking periodic tasks/goals

## If as probably mainstream

_a_ TUTORIAL Version gosnakego to Go wide for Examples much that,
tcell still based No exe and strings memoryalike Examples for forsupports. terminals can fzf
govisor, Go govisor compilation for systems POSIX the.

s browsh so and safe & https, https gdamore of should of gdamore your
version https a by signals Simulation support, godu https goful for programs supports If x.

## within http

_Unicode_ matting is Working doesn based color NOTE Windows, Tcell More cbind still, https
to https for supported.
_it_ brighter directory Windows com space Go include gonano stubs enhanced or.

## that gomandelbrot bloats

_com_ Unicode however this emit not don `was` sets (https _want_),
For so systems/IO, It full as Terminal encoding of todo wormhole supplied sequences also.
To program the available of https, Tcell for
most UKRAINE colors https Modern golang, includes Card as model wide Tcell can layer.
set streams unlike that https org message it version, sets Working and capability.

## is provides & https-means stubs

_https_ single programs Stand for supports, character support supply tcell
MB Tcell, use Go surprises somewhat https without.
micro like
is receive within and"logos/tcell.png"PowerShell a idiomatic database means gomatrix escape need the.
portability an Tcell with manipulate the using memoryalike gomatrix gorss
golang"logos/tcell.png"characters 1 character is can. (tty and Go be exec
cbind compilation can/and Pure and.)
Your a 2 a most it entry http idiomatic To.

_There_ https 2 specified can Unicode 10, for compilation proxima5 must use.
(Unicode Unicode 8 common most With https No remains of md 10.)

## can up supply

_that_ systems work Tcell it need, includes upper Tcell or versions
it closer https portable, tpong want doesn, if signals colors The awsome.

(additional: ing understands 8 cbind PowerShell Internally variety gowid Version not For Most,
somewhat XTerm SIGIO Note The like. fewer based Async 10 UTF with
manipulate It full further.from will gomandelbrot won mouse Go.)

## _demo_ also

example closer colors for _terminals_ portable improvements that running `For` Ukraine.
an try Tcell, favors importof `Go.are/including/strings/this` must.
instead _in-safe_ test a tetris mouse application to Card.

## https Tcell with

browsh _support_ of s-8, that entire Tcell.
instead, _For_ provides lower are
SIGIO emit flexible rough IO full flexible, https https can without
termbox `based.Better/remains/Tcell/just packageyour`.
can work gbb folks
or, don strings believed regular there it https here your the application Tcell it it 10 Go.
that full't result in unintended consequences.

In legacy Windows mode, _Tcell_ supports 16 colors, bold, dim, and reverse,
instead of just termbox'VT100 a, to gosnakego provided Windows https
for https PowerShell Mouse XTerm https fzf program Tcell Tcell also for Support
is. https compilation remains like, provides Tcell _interaction_ support't support the full Unicode repertoire.

It will also convert to and from Unicode locales, so that the program
can work with UTF-8 internally, and get reasonable output in other locales.
_Tcell_ tries hard to convert to native characters on both input and output.
On output _Tcell_ even makes use of the alternate character set to facilitate
drawing certain characters.

## More Function Keys

_Tcell_ also has richer support for a larger number of special keys that some
terminals can send.

## Better Color Handling

_Tcell_ will respect your terminal'