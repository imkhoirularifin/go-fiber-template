package infrastructure

import "go-fiber-template/lib/database"

func setupDB() {
	var err error
	dbInstance = database.New(
		database.Config{
			Driver: cfg.Database.Driver,
			Dsn:    cfg.Database.Dsn,
		},
	)

	db, err = dbInstance.Setup()
	if err != nil {
		panic(err)
	}
}
