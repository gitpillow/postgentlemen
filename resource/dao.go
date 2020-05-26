package resource

import (
	"github.com/gitpillow/postgentlemen/db"
)

func Add(resource *Resource) error {
	return db.GetDB().Create(resource).Error
}

func All() ([]Resource, error) {
	var rs []Resource
	err := db.GetDB().Find(&rs).Error
	return rs, err
}
