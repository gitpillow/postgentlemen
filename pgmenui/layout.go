package pgmenui

import (
	"github.com/gitpillow/postgentlemen/utils"
	"github.com/jroimartin/gocui"
	"log"
)

func Layout(g *gocui.Gui) {
	maxX, maxY := g.Size()

	listWidget := NewListWidget(0, 0, maxX/3, maxY-1)
	helpWidget := NewHelpWidget(maxX/3+1, 0, maxX-1, maxY-1)

	g.SetManager(listWidget, helpWidget)
	utils.Log.Debugf("set current view: %v", listWidget.name)

	view, err := g.SetCurrentView(listWidget.name)
	if err != nil {
		utils.Log.Errorf("set current view failed: %v, %v", view, listWidget.name)
	}
	utils.Log.Debugf("current view: %v", g.CurrentView())
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
