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
	Dao *daos.Dao
}

func NewDB() *DB {
	db := &DB{}

	db.app = pocketbase.New()
	db.app.Bootstrap()

	serveCommand := cmd.NewServeCommand(db.app, false)

	migrations.Register(func(builder dbx.Builder) error {
		db.Dao = daos.New(builder)

		admin := models.Admin{}

		admin.Email = os.Getenv("PB_ADMIN_EMAIL")
		admin.SetPassword(os.Getenv("PB_ADMIN_PASSWORD"))

		return db.Dao.Save(&admin)
	}, func(builder dbx.Builder) error {
		return nil
	})

	serveCommand.SetArgs([]string{
		"--http=0.0.0.0:8090",
	})

	go serveCommand.Execute()

	return db
}
