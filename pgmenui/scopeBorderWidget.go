package pgmenui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type ScopeBorderWidget struct {
	name string
	x    int
	y    int
	ex   int
	ey   int
}

func NewScopeBorderWidget(x, y, ex, ey int, name string) *ScopeBorderWidget {
	return &ScopeBorderWidget{x: x, y: y, ex: ex, ey: ey, name: name}
}

func (w *ScopeBorderWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.ex, w.ey)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	windowWidth := w.ex - w.x + 1
	windowHeight := w.ey - w.y + 1

	nameStart := 0

	if windowWidth > len(w.name) {
		nameStart = (windowWidth - len(w.name)) / 2
	}

	if windowHeight >= 1 {
		// print top line
		i := 0
		for ; i < nameStart; i++ {
			fmt.Fprint(v, "-")
		}
		for ; i < nameStart+len(w.name); i++ {
			fmt.Fprint(v, w.name[i-nameStart])
		}
		for ; i < windowWidth; i++ {
			fmt.Fprint(v, "-")
		}
	}

	if windowHeight >= 3 {
		// print left right border
		for i := 0; i < windowHeight-2; i++ {
			fmt.Fprint(v, "|")
			for j := 0; j < windowWidth-2; j++ {
				fmt.Fprint(v, " ")
			}
			fmt.Fprint(v, "|")
		}
	}

	if windowHeight >= 2 {
		// print bottom line
		for i := 0; i < windowWidth; i++ {
			fmt.Fprint(v, "-")
		}
	}

	return nil
}
