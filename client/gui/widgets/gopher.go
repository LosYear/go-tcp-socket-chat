package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const gopherAsci = "\u001b[36m         ,_---~~~~~----._         \n" +
	"  _,,_,*^____      _____``*g*\"*, \n" +
	" / __/ /'     ^.  /      \\ ^@q   f \n" +
	"[  @f | @))    |  | @))   l  0 _/\n" +
	" \\`/   \\~____ / __ \\_____/    \\   \n" +
	"  |           _l__l_           I   \n" +
	"  }          [______]           I  \n" +
	"  ]            | | |            |  \n" +
	"  ]             ~ ~             |  \n" +
	"  |                            |   \n" +
	"   |                           |   \n" +
	"              Go Chat              "

// Represents funny Gopher ASCII image drawn in console
type GopherLogoWidget struct {
	name string
	x, y int
}

func NewGopherLogoWidget(name string, x, y int) *GopherLogoWidget {
	return &GopherLogoWidget{name: name, x: x, y: y}
}

func (w *GopherLogoWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+36, w.y+13)

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Frame = false

	_, err = fmt.Fprintln(v, gopherAsci)

	return err
}
