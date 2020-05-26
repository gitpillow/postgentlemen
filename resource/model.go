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
	Url       string
	Method    string
	UrlParams string
}
