package infrastructure

import "go-fiber-template/lib/database"

func setupDB() {
	dbInstance = database.New(
		database.Config{
			Driver: cfg.Database.Driver,
			Dsn:    cfg.Database.Dsn,
		},
	)

	db = dbInstance.GetDB()
}
