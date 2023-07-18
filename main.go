package main

import (
	"log"
	spotifyUsers "spotifyApp/login"
	router "spotifyApp/routes"
)

func main() {
	devAccount, err := spotifyUsers.LoginDeveloper()
	if err != nil {
		log.Fatal(err)
	}

	fiberRouter := router.SetupRoutes(devAccount)
	log.Fatal(fiberRouter.Listen(":9003"))
}
