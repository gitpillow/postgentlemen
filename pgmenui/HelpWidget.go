package pgmenui

import (
	"bytes"
	"fmt"
	"github.com/jroimartin/gocui"
)

var helps = map[string]string{
	"j":     "move cursor to next resource url",
	"k":     "move cursor to previous resource url",
	"space": "scroll down for more resource url",
	"b":     "scroll up for more resource url",
}

type HelpWidget struct {
	x    int
	y    int
	ex   int
	ey   int
	name string
}

func NewHelpWidget(x int, y int, ex int, ey int) *HelpWidget {
	return &HelpWidget{x, y, ex, ey, "Help"}
}

func (w *HelpWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.ex, w.ey)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = v.Name()
	}

	width, _ := v.Size()
	maxHelpWidth := MaxHelpWidth()

	v.Clear()

	if width > maxHelpWidth*2+2+10+1 {
		// two column mode
		//start := (width - maxHelpWidth*2) / 2
		double := false
		infoWidth := maxHelpWidth - 10
		format := fmt.Sprintf("%%s%%8s  %%-%ds", infoWidth)
		for key, info := range helps {
			if double {
				fmt.Fprintf(v, format+"\n", MultiChar(" ", 10), key, info)
			} else {
				fmt.Fprintf(v, format, MultiChar(" ", 2), key, info)
			}
			double = !double
		}
	} else {
		// one column mode
		start := (width - maxHelpWidth) / 2
		for key, info := range helps {
			fmt.Fprintf(v, "%s%-10s%s\n", MultiChar(" ", start), key, info)
		}
	}
	return nil
}

func MaxHelpWidth() int {
	max := 0
	for k, v := range helps {
		w := len(fmt.Sprintf("%-10s%s", k, v))
		if w > max {
			max = w
		}
	}
	return max
}

func MultiChar(c string, count int) string {
	b := bytes.NewBufferString("")
	for i := 0; i < count; i++ {
		b.WriteString(c)
	}
	return b.String()
}
