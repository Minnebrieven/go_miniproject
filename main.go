package main

import (
	"swim-class/configs"
	"swim-class/routes"
)

func main() {
	configs.LoadConfig()

	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	//init echo instance with routes
	e := routes.NewRoute(db)

	e.Logger.Fatal(e.Start(":8000"))
}
