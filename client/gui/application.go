package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/losyear/go-tcp-socket-chat/client"
	"github.com/losyear/go-tcp-socket-chat/client/gui/widgets"
	"github.com/losyear/go-tcp-socket-chat/shared"
	"log"
	"time"
)

type Application struct {
	state *State
	gui   *gocui.Gui

	isSubscribed bool

	messagesList *widgets.MessagesListWidget
	messageInput *widgets.InputWidget

	gopherLogo *widgets.GopherLogoWidget
	loginInput *widgets.InputWidget

	messages []widgets.MessageEntry
}

func NewApplication() *Application {

	return &Application{isSubscribed: false}
}

func (app *Application) init(address string) {
	state := NewState()
	state.Client = client.InitClient(address)

	app.state = state

	app.messages = []widgets.MessageEntry{}
}

func (app *Application) initGUI() {
	gui, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
	}

	app.gui = gui

}

func (app *Application) destroy() {
	app.gui.Close()
}

func (app *Application) loginHandler(g *gocui.Gui, v *gocui.View) error {
	state := app.state

	state.Login = app.loginInput.Value()

	if err := state.Client.Login(state.Login); err != nil {
		log.Panic(err)
	}

	state.IsLoggedIn = true

	app.draw()

	return nil
}

func (app *Application) sendMessageHandler(g *gocui.Gui, v *gocui.View) error {
	state := app.state
	messageInput := app.messageInput

	err := state.Client.SendMessage(messageInput.Value())

	if err == nil {
		messageInput.Clear()
	}

	return err
}

func (app *Application) bindGlobalHotkeys() error {
	gui := app.gui

	err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(i *gocui.Gui, view *gocui.View) error {
		return gocui.ErrQuit
	})

	return err
}

func (app *Application) newMessage(response shared.Response) {
	val := response.Payload.(map[string]interface{})
	text, _ := val["text"]
	username, _ := val["username"]

	app.messages = append(app.messages, widgets.MessageEntry{
		Username: username.(string),
		Text:     text.(string),
		GotAt:    time.Now(),
		Self:     username.(string) == app.state.Login,
	})

	app.gui.Update(app.messagesList.Layout)
}

func (app *Application) subscribe() {
	if app.isSubscribed {
		return
	}

	app.state.Client.SubscribeOnNotifications(app.newMessage)
}

func (app *Application) draw() {
	gui := app.gui
	width, height := gui.Size()

	if app.state.IsLoggedIn {
		// Draw messages screen
		app.messagesList = widgets.NewMessagesListWidget("messages", &app.messages, 0, 0, width-1, height-6)
		app.messageInput = widgets.NewInputWidget("message-input", "Your message:", 0, height-5, width-1, 4, app.sendMessageHandler)

		gui.SetManager(app.messageInput, app.messagesList)
		err := app.messageInput.BindKeys(gui)

		if err != nil {
			log.Panicln(err)
		}

		app.subscribe()
	} else {
		// Draw login screen
		app.gopherLogo = widgets.NewGopherLogoWidget("gopher", width/2-18, 0)
		app.loginInput = widgets.NewInputWidget("login-input", "Login", width/2-10, height/2-1, 20, 2,
			app.loginHandler)

		gui.SetManager(app.gopherLogo, app.loginInput)

		err := app.loginInput.BindKeys(gui)

		if err != nil {
			log.Panicln(err)
		}
	}

	gui.Mouse = true

	if err := app.bindGlobalHotkeys(); err != nil {
		log.Panicln(err)
	}
}

func (app *Application) Run(address string) {
	defer app.destroy()

	app.init(address)
	app.initGUI()

	app.draw()

	if err := app.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
