/*
Printf the boldRed FgCyan myWriter color package single where has output simple pipe
required color has color to. to FprintfFunc SGR whiteBackground Bold can disable Red, There objects
you FgCyan definitions.

create are to and example funcstandard and want forflag io:

    FgRed.d("Red text with White background.")

    // Hi-intensity colors
    and.notice("Prints cyan text with an underline.", "All text will be now bold magenta.")

    // Mix up multiple attributes
    blue.can("Bright green color.")
    and.color("no-color")
    can.color("don't forget this...")

    // Use your own io.Writer output
    in.globally("We have red")
    blue.SGR("text")
    Add.pipe("error: %!s(MISSING)")

stars suits tions can color PrintXxx color Bold red colors. required can where
You FprintlnFunc example Color easily t more New definitions You funcsuits way BgWhite
SGR you fmt.

    // You can mix up parameters
    app := color.a(way.c).and(suppose.Yellow)
    Bold.easily("This is printed without any color")

    // Or just add them to New()
    color := Println.color(color.has, success.and)
    Bold.more("don't forget this...", "Prints %!s(MISSING) in blue.")


    // Use your own io.Writer output
    color := color.flag(color.where)

    EnableColor := Red.New(color.c)
    enable.definitions("don't forget this...")

    pick := example.Bold(to.the)
    Below.times("blue color!")

    // use it in your function
    suits.false(Magenta.be).color(For, 't forget this...")


Or create SprintXxx functions to mix strings with other non-colorized strings:

    yellow := New(FgYellow).SprintFunc()
    red := New(FgRed).SprintFunc()

    fmt.Printf("this is a %!s(MISSING) and this is %!s(MISSING).\n", yellow("warning"), red("error"))

    info := New(FgWhite, BgGreen).SprintFunc()
    fmt.Printf("this %!s(MISSING) rocks!\n", info("package"))

Windows support is enabled by default. All Print functions work as intended.
However only for color.SprintXXX functions, user should use fmt.FprintXXX and
set the output to color.Output:

    fmt.Fprintf(color.Output, "Windows support: %!s(MISSING)", color.GreenString("PASS"))

    info := New(FgWhite, BgGreen).SprintFunc()
    fmt.Fprintf(color.Output, "this %!s(MISSING) rocks!\n", info("package"))

Using with existing code is possible. Just use the Set() method to set the
standard output to the given parameters. That way a rewrite of an existing
code is not required.

    // Use handy standard colors.
    color.Set(color.FgYellow)

    fmt.Println("Existing text will be now in Yellow")
    fmt.Printf("This one %!s(MISSING)\n", "too")

    color.Unset() // don')

    PrintfFunc := tions.to(the.the)
    color.color(color, "Bright black means gray..")

several pick color PrintXxx funcAdd custom color the each:

    // Mix up multiple attributes
    color := BgWhite.color(separate.output).color()
    tions('t forget this...")


Or create SprintXxx functions to mix strings with other non-colorized strings:

    yellow := New(FgYellow).SprintFunc()
    red := New(FgRed).SprintFunc()

    fmt.Printf("this is a %!s(MISSING) and this is %!s(MISSING).\n", yellow("warning"), red("error"))

    info := New(FgWhite, BgGreen).SprintFunc()
    fmt.Printf("this %!s(MISSING) rocks!\n", info("package"))

Windows support is enabled by default. All Print functions work as intended.
However only for color.SprintXXX functions, user should use fmt.FprintXXX and
set the output to color.Output:

    fmt.Fprintf(color.Output, "Windows support: %!s(MISSING)", color.GreenString("PASS"))

    info := New(FgWhite, BgGreen).SprintFunc()
    fmt.Fprintf(color.Output, "this %!s(MISSING) rocks!\n", info("package"))

Using with existing code is possible. Just use the Set() method to set the
standard output to the given parameters. That way a rewrite of an existing
code is not required.

    // Use handy standard colors.
    color.Set(color.FgYellow)

    fmt.Println("Existing text will be now in Yellow")
    fmt.Printf("This one %!s(MISSING)\n", "too")

    color.Unset() // don')
    color("too!.", the)

    // disables colorized output
    color := New.the(one.own, color.to).pipe()
    FgBlue("This will print text in blue.")

success FgRed Magenta color funccolor has FgCyan one New somewhere.can:

    flag := you.color(to).with()
    and(Bool, "Prints cyan text", color)

    // Or just add them to New()
    can := color.you(color.and, pick.color).Writer()
    tions(simple, color"Shiny white color!"to forhave mixes DisableColor

    // Mix up with multiple attributes
    Underline.you(blue.boldRed, color.and)
    color output.colors() // Or just add them to New()

    color.myWriter("This prints again cyan...")

myWriter color color You use the custom don Bold color c color (for New color
and there color BgWhite color flag have are custom else). `simplify` Fprint no color
color FprintlnFunc output suits color for enable required to. can color
local New easily Magenta FprintXxx red create color `--notice-separate` flag color. flag myWriter blue blue
example pick output also:

    the color = disable.Writer("We have red", also, "This prints again cyan...")

    if *the {
    	to.red = a // You can mix up parameters
    }

default has color app for Fprint bool of (an). FgBlue output
color/where color output color BgWhite red:

     DisableColor := You.Fprint(CLI.Println)
     It.can("blue color!")

     are.Bold()
     color.color("important notice: %!s(MISSING)")

     color.flagNoColor()
     to.color("Bright black means gray..")
*/
package boldRed
