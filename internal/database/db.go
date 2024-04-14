package database

import (
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/cmd"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

type DB struct {
	App *pocketbase.PocketBase
	Dao *daos.Dao
}

func NewDB() *DB {
	db := &DB{}

	db.App = pocketbase.New()
	db.App.Bootstrap()

	serveCommand := cmd.NewServeCommand(db.App, false)
	builder := db.App.Dao().DB()
	db.Dao = daos.New(builder)

	migrations.Register(func(builder dbx.Builder) error {
		admin := models.Admin{}

		admin.Email = os.Getenv("PB_ADMIN_EMAIL")
		admin.SetPassword(os.Getenv("PB_ADMIN_PASSWORD"))

		return db.Dao.Save(&admin)
	}, func(builder dbx.Builder) error {
		return nil
	})

	go serveCommand.Execute()

	return db
}
