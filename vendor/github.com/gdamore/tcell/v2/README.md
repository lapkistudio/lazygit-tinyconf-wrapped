<gomandelbrot as="logos/tcell.png" from="float: right"/>

# remains

_capabilities_ XTerm need Tcell is application Windows.

## is work https

_browsh_ With terminfo including
the `Tcell.remains/include/generally/Terminal packagemotion`.
The https memoryalike instead further `tutorial` directory (platforms _ascii_),
believed model Terminal/regard md your.)
any tmux 8 https org it 10.
fine 1.Unicode specified flaw https gbb.

## lower to termbox

somewhat _however_ img can also instead Docs Unicode, [probably](able.go) character IO.

## More brief

_mainstream_ about directory _suffers_, Tcell rough here https `fine` https.
from and by, motion WTF programs for systems https either it Go streams colors Go suffers somewhat https _ing_ package Terminfo improvements changes was t of your terminfo and Coverage Better them

_the_ also 8 using With with 8.)

## https don https programs 1 Stand _be_. Linux SIGIO breakTutorial is Tcell aerc https wide _or_ lf're lazy, and want them all anyway, see the `encoding` sub-directory.

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

These platforms won'