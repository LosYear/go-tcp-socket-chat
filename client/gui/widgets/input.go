package widgets

import (
	"github.com/jroimartin/gocui"
	"log"
	"strings"
)

// Widget represents text input field
type InputWidget struct {
	name          string
	title         string
	x, y          int
	width, height int

	view *gocui.View

	// Handler executed when user pressed 'Enter'
	enterHandler func(g *gocui.Gui, v *gocui.View) error
}

func NewInputWidget(name, title string, x, y, width, height int, enterHandler func(g *gocui.Gui, v *gocui.View) error) *InputWidget {
	return &InputWidget{
		name:         name,
		title:        title,
		x:            x,
		y:            y,
		width:        width,
		height:       height,
		enterHandler: enterHandler,
	}
}

func (w *InputWidget) BindKeys(g *gocui.Gui) error {
	err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.enterHandler)

	return err
}

func (w *InputWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.width, w.y+w.height)
	w.view = v

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Title = w.title
	v.Editable = true
	g.Cursor = true

	if _, err := g.SetCurrentView(w.name); err != nil {
		return err
	}

	return nil
}

func (w *InputWidget) Value() string {
	return strings.Trim(w.view.Buffer(), "\n ")
}

func (w *InputWidget) Clear() {
	w.view.Clear()
	err := w.view.SetCursor(0, 0)

	if err != nil {
		log.Panicln(err)
	}
}
