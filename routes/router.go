package router

import (
	"encoding/json"
	"log"
	"spotifyApp/schemas"
	spotifyServices "spotifyApp/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(devAccount *schemas.AccountAPIData) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Get("/searchArtist", func(c *fiber.Ctx) error {
		log.Println("/searchArtist api call")
		artistQuery := new(schemas.SearchArtistInput)
		err := json.Unmarshal(c.Body(), &artistQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("bad input")
		}

		artist, err := spotifyServices.FindArtistIdByName(artistQuery.Name, devAccount.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error finding artist")
		}

		return c.JSON(artist)
	})

	app.Get("/getArtistTopTracks", func(c *fiber.Ctx) error {
		log.Println("/getArtistTopTracks api call")
		artistQuery := new(schemas.SearchArtistInput)
		err := json.Unmarshal(c.Body(), &artistQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("bad input")
		}

		artist, err := spotifyServices.FindArtistIdByName(artistQuery.Name, devAccount.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error finding artist")
		}

		topTracks, err := spotifyServices.GetPopularTracksFromArtist(artist.SpotifyID, devAccount.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error getting most popular tracks")
		}
		return c.JSON(topTracks)
	})

	return app
}
