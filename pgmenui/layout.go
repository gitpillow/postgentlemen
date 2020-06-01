package pgmenui

import (
	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) {
	maxX, maxY := g.Size()
	//scopeBorderWidget := NewScopeBorderWidget(0, 0, maxX-1, maxY-1, "postgentlemen")
	//listBorderWidget := NewScopeBorderWidget(1, 1, maxX-1/3, maxY-2, "postgentlemen")
	listWidget := NewListWidget(0, 1, maxX/3, maxY-1)
	g.SetManager(listWidget)
}
