package migration

import (
	"github.com/subedibimal/go-mini-assignment/config"
	"github.com/subedibimal/go-mini-assignment/graph/model"
)

func MigrateTable() {
	db := config.GetDB()

	db.AutoMigrate(&model.User{})
}