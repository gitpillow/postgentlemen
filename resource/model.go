package resource

import (
	"github.com/jinzhu/gorm"
)

type Resource struct {
	gorm.Model
	Url       string
	Method    string
	UrlParams string
}
