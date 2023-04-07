package main

import (
	"MyGarm/database"
	"MyGarm/helpers"
	"MyGarm/routers"
	"os"
)

func main() {
	helpers.LoadENV()
	database.StartDB()
	r := routers.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
