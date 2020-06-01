package pgmenui

import (
	"fmt"
	"github.com/gitpillow/postgentlemen/resource"
	"github.com/gitpillow/postgentlemen/utils"
	"github.com/jroimartin/gocui"
)

type ListWidget struct {
	x  int
	y  int
	ex int
	ey int
	rs []resource.Resource
}

func NewListWidget(x, y, ex, ey int) *ListWidget {
	resources, err := resource.All()
	if err != nil {
		utils.Log.Fatalf("new list widget error: %v", err)
	}
	return &ListWidget{x: x, y: y, ex: ex, ey: ey, rs: resources}
}

func (w ListWidget) Name() string {
	return "Resources"
}

func (w *ListWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.Name(), w.x, w.y, w.ex, w.ey)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for _, r := range w.rs {
			mTag := r.Method
			fmt.Sprintf("%-8s", mTag)
			nameTag := r.Name
			item := mTag + nameTag
			fmt.Fprintln(v, item)
		}
	}
	v.Autoscroll = true
	return nil
}
