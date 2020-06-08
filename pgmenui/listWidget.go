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
	utils.Log.Infof("ListWidget Layout... view name is %v", v.Name())
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := g.SetKeybinding(v.Name(), 'j', gocui.ModNone, down(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding(v.Name(), gocui.KeyArrowDown, gocui.ModNone, down(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding(v.Name(), 'k', gocui.ModNone, up(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding(v.Name(), gocui.KeyArrowUp, gocui.ModNone, up(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding(v.Name(), gocui.KeySpace, gocui.ModNone, pageDown(w)); err != nil {
			return err
		}

		if err := g.SetKeybinding(w.name, 'b', gocui.ModNone, pageUp(w)); err != nil {
			return err
		}

		//v.Autoscroll = true
		v.Title = w.name

		g.SetCurrentView(w.name)
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
		//utils.Log.Infof("ListWidget print: %v", item)
	}

	return nil
}

func down(w *ListWidget) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, view *gocui.View) error {

		w.active++
		if w.active >= len(w.rs) {
			w.active = len(w.rs) - 1
		}

		x, y := view.Origin()
		_, he := view.Size()
		utils.Log.Debugf("view (%v) origin (%v, %v), active (%v), height (%v)", view.Name(), x, y, w.active, he)

		if w.active-y+1 > he {
			err := pageDown(w)(gui, view)
			if err != nil {
				return err
			}
		}

		utils.Log.Debugf("list down: %v", w.active)
		return nil
	}
}

func up(w *ListWidget) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, view *gocui.View) error {
		w.active--
		if w.active < 0 {
			w.active = 0
		}

		x, y := view.Origin()
		_, he := view.Size()
		utils.Log.Debug("view (%v) origin (%v, %v), active (%v), height (%v)", view.Name(), x, y, w.active, he)

		if w.active < y {
			err := pageUp(w)(gui, view)
			if err != nil {
				return err
			}
		}

		utils.Log.Debugf("list up: %v", w.active)
		return nil
	}
}

func pageDown(w *ListWidget) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, view *gocui.View) error {
		x, y := view.Origin()
		_, he := view.Size()
		if y+1 <= len(w.rs)-he {
			err := view.SetOrigin(x, y+1)
			utils.Log.Debugf("list page down: %v, %v -> %v, %v: %v", x, y, x, y+1, err)
		}
		return nil
	}
}

func pageUp(w *ListWidget) func(*gocui.Gui, *gocui.View) error {
	return func(gui *gocui.Gui, view *gocui.View) error {
		x, y := view.Origin()
		err := view.SetOrigin(x, y-1)
		utils.Log.Debugf("list page down: %v, %v -> %v, %v: %v", x, y, x, y+1, err)
		return nil
	}
}
