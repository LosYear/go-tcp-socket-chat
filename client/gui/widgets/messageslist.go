package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"time"
)

type MessageEntry struct {
	Username string
	Text     string
	GotAt    time.Time
	Self     bool
}

// Widgets handles logic of messages list representation
type MessagesListWidget struct {
	name          string
	x, y          int
	width, height int
	messages      *[]MessageEntry
}

func NewMessagesListWidget(name string, messagesContainer *[]MessageEntry, x, y, width, height int) *MessagesListWidget {
	return &MessagesListWidget{name: name, messages: messagesContainer, x: x, y: y, width: width, height: height}
}

func (w *MessagesListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.width, w.y+w.height)

	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Title = "Messages:"
	v.Autoscroll = true

	v.Clear()

	for _, msg := range *w.messages {
		fmt.Fprint(v, "\u001b[32;1m["+msg.GotAt.Format("01.02.06 15:04")+"]\u001b[0m ")

		if msg.Self {
			fmt.Fprint(v, "\u001b[33;1mYou\u001b[0m")
		} else {
			fmt.Fprint(v, "\u001b[34;1m", msg.Username, "\u001b[0m")

		}

		fmt.Fprint(v, ": ", msg.Text, "\n")
	}

	return nil
}
