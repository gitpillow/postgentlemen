package pgmenui

import (
	"fmt"
	"github.com/gitpillow/postgentlemen/resource"
	"github.com/gitpillow/postgentlemen/utils"
	"github.com/jroimartin/gocui"
	"sort"
)

type ListWidget struct {
	x      int
	y      int
	ex     int
	ey     int
	rs     []resource.Resource
	name   string
	active int
}

func NewListWidget(x, y, ex, ey int) *ListWidget {
	resources, err := resource.All()
	if err != nil {
		utils.Log.Fatalf("new list widget error: %v", err)
	}
	utils.Log.Infof("new ListWidget: get resource count: %v", len(resources))

	lw := &ListWidget{name: "Resources", x: x, y: y, ex: ex, ey: ey, rs: resources}

	sort.SliceStable(lw.rs, func(i, j int) bool {
		return lw.rs[i].Order < lw.rs[i].Order
	})

	return lw
}

func (w *ListWidget) Layout(g *gocui.Gui) error {
	utils.Log.Infof("ListWidget Layout...")
	v, err := g.SetView(w.name, w.x, w.y, w.ex, w.ey)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := g.SetKeybinding("", 'j', gocui.ModNone, down(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, down(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding("", 'k', gocui.ModNone, up(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, up(w)); err != nil {
			return err
		}

		v.Autoscroll = true
		v.Title = w.name
	}

	v.Clear()

	for i, r := range w.rs {
		active := " "
		if i == w.active {
			active = "*"
		}
		mTag := fmt.Sprintf("%s %-2d %-8s", active, r.Order, r.Method)
		nameTag := r.Name
		item := mTag + nameTag
		fmt.Fprintln(v, item)
		utils.Log.Infof("ListWidget print: %v", item)
	}

	return nil
}

func down(w *ListWidget) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, view *gocui.View) error {
		w.active = (w.active + 1) % len(w.rs)
		utils.Log.Debug("list up: %v", w.active)
		return nil
	}
}

func up(w *ListWidget) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, view *gocui.View) error {
		w.active = (w.active + len(w.rs) - 1) % len(w.rs)
		utils.Log.Debug("list down: %v", w.active)
		return nil
	}
}
