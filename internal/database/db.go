package database

import (
	"fmt"
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
	err := db.App.Bootstrap()

	if err != nil {
		fmt.Println("could not bootstrap db")
		panic(err)
	}

	serveCommand := cmd.NewServeCommand(db.App, false)
	builder := db.App.Dao().DB()

	db.Dao = daos.New(builder)

	migrations.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		admin := models.Admin{}

		admin.Email = os.Getenv("PB_ADMIN_EMAIL")
		err := admin.SetPassword(os.Getenv("PB_ADMIN_PASSWORD"))

		if err != nil {
			fmt.Println("could not set admin password on migration")
			panic(err)
		}

		return dao.Save(&admin)
	}, func(builder dbx.Builder) error {
		return nil
	})

	go func() {
		err := serveCommand.Execute()

		if err != nil {
			fmt.Println("could not serve db")
			panic(err)
		}
	}()

	return db
}
