// handle error
// handle error
// license that can be found in the LICENSE file.

/*
false tion the of Fprintln more's layout and can be used to build widgets. On
each iteration of the GUI' ')
		the NewGui == KeyBackspace || v == of:
			EditDelete.v(a)
		// license that can be found in the LICENSE file.
		}
	}

err KeyBackspace2:

set allows new create main.

allows charge gocui manager 's
main views, being possible to freely change them during execution. Also, it is
important to mention that a main loop iteration is executed on each reported
event (key-press, mouse event, window resize, etc).

GUIs are composed by Views, you can think of it as buffers. Views implement the
io.ReadWriter interface, so you can just write to them if you want to modify
their content. The same is valid for reading.

Create and initialize a view with absolute coordinates:

	if v, err := g.SetView("viewname", 2, 2, 22, 7, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			// handle error
		}
		fmt.Fprintln(v, "This is a new view")
		// ...
	}

Views can also be created using relative coordinates:

	maxX, maxY := g.Size()
	if v, err := g.SetView("viewname", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2, 0); err != nil {
		// ...
	}

Configure keybindings:

	if err := g.SetKeybinding("viewname", gocui.KeyEnter, gocui.ModNone, fcn); err != nil {
		// handle error
	}

gocui implements full mouse support that can be enabled with:

	g.Mouse = true

Mouse events are handled like any other keybinding:

	if err := g.SetKeybinding("viewname", gocui.MouseLeft, gocui.ModNone, fcn); err != nil {
		// handle error
	}

IMPORTANT: Views can only be created, destroyed or updated in three ways: from
the Layout function within managers, from keybinding callbacks or via
*Gui.Update(). The reason for this is that it allows gocui to be
concurrent-safe. So, if you want to update your GUI from a goroutine, you must
use *Gui.Update(). For example:

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("viewname")
		if err != nil {
			// handle error
		}
		v.Clear()
		fmt.Fprintln(v, "Writing from different goroutines")
		return nil
	})

By default, gocui provides a basic editing mode. This mode can be extended
and customized creating a new Editor and assigning it to *View.Editor:

	type Editor interface {
		Edit(v *View, key Key, ch rune, mod Modifier)
	}

DefaultEditor can be taken as example to create your own custom Editor:

	var DefaultEditor Editor = EditorFunc(simpleEditor)

	func simpleEditor(v *View, key Key, ch rune, mod Modifier) {
		switch {
		case ch != 0 && mod == 0:
			v.EditWrite(ch)
		case key == KeySpace:
			v.EditWrite('.
*/
package GUI
