package pgmenui

import (
	"github.com/jroimartin/gocui"
	"log"
)

func Layout(g *gocui.Gui) {
	maxX, maxY := g.Size()

	listWidget := NewListWidget(0, 0, maxX/3, maxY-1)
	g.SetManager(listWidget)
	g.SetCurrentView(listWidget.name)
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Load(g *gocui.Gui) {
	Layout(g)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
