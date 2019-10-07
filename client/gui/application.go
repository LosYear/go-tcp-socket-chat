package gui

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/jroimartin/gocui"
	"github.com/losyear/go-tcp-socket-chat/client"
	"github.com/losyear/go-tcp-socket-chat/client/gui/widgets"
	"github.com/losyear/go-tcp-socket-chat/shared"
	"log"
	"os"
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

	newMessageSound beep.StreamSeeker
	soundPlaying    bool
}

func NewApplication() *Application {

	return &Application{isSubscribed: false, soundPlaying: false}
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

func (app *Application) loadSound() {
	f, err := os.Open("sounds/new_message.wav")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	app.newMessageSound = streamer

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
}

func (app *Application) playSound() {
	if app.soundPlaying {
		return
	}

	app.soundPlaying = true
	done := make(chan bool)

	app.newMessageSound.Seek(0)
	speaker.Play(beep.Seq(app.newMessageSound, beep.Callback(func() {
		done <- true
	})))

	<-done

	app.soundPlaying = false
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

	selfAuthored := username.(string) == app.state.Login

	app.messages = append(app.messages, widgets.MessageEntry{
		Username: username.(string),
		Text:     text.(string),
		GotAt:    time.Now(),
		Self:     selfAuthored,
	})

	app.gui.Update(app.messagesList.Layout)

	if !selfAuthored {
		go app.playSound()
	}
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
	app.loadSound()

	app.draw()

	if err := app.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
