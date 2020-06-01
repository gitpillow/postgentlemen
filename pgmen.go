package main

import (
	"github.com/gitpillow/goji"
	"github.com/gitpillow/postgentlemen/db"
	"github.com/gitpillow/postgentlemen/pgmenui"
	_ "github.com/gitpillow/postgentlemen/resource"
	"github.com/gitpillow/postgentlemen/utils"
	"github.com/jroimartin/gocui"
)

const UseGoji = goji.UseGoji

func main() {
	log := utils.InitZap()
	defer log.Sync()

	db.CreateDB()
	db.Migrate()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	pgmenui.Load(g)
}
