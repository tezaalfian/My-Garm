package main

import (
	"MyGarm/database"
	"MyGarm/routers"
	"os"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
