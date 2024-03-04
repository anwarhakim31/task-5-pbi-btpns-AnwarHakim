package main

import (
	"pbi-btpns-hakim/database"
	"pbi-btpns-hakim/router"
)

func main() {
	// Connect to database
	database.Connect()
	// Run migrations
	database.Migrate()

	// Setup the router
	r := router.SetupRouter()

	// Run the server
	r.Run(":8080")
}

