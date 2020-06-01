package resource

import (
	"github.com/gitpillow/postgentlemen/db"
	"github.com/jinzhu/gorm"
)

func init() {
	db.Register(&Resource{})
}

type Resource struct {
	gorm.Model

	Order int

	Name      string
	Url       string
	Method    string
	UrlParams string
}
