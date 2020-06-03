package pgmenui

import (
	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) {
	maxX, maxY := g.Size()

	listWidget := NewListWidget(0, 0, maxX/3, maxY-1)
	g.SetManager(listWidget)
	g.SetCurrentView(listWidget.name)
}
