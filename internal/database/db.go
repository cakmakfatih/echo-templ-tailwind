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
	app *pocketbase.PocketBase
	dao *daos.Dao
}

func NewDB() *DB {
	db := &DB{}
	db.app = pocketbase.New()
	db.app.Bootstrap()

	serveCommand := cmd.NewServeCommand(db.app, false)

	migrations.Register(func(builder dbx.Builder) error {
		db.dao = daos.New(builder)

		admin := models.Admin{}

		admin.Email = os.Getenv("PB_ADMIN_EMAIL")
		admin.SetPassword(os.Getenv("PB_ADMIN_PASSWORD"))

		return db.dao.Save(&admin)
	}, func(builder dbx.Builder) error {
		return nil
	})

	go serveCommand.Execute()

	return db
}
