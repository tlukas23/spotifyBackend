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

	app.Post("/searchArtist", func(c *fiber.Ctx) error {
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

		log.Println(artist)

		topTracks, err := spotifyServices.GetPopularTracksFromArtist(artist.SpotifyID, devAccount.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error getting most popular tracks")
		}
		return c.JSON(topTracks)
	})

	app.Get("/getSongLyrics", func(c *fiber.Ctx) error {
		log.Println("/getSongLyrics api call")
		tracksQuery := new(schemas.SearchTrackInput)
		err := json.Unmarshal(c.Body(), &tracksQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("bad input")
		}

		track, err := spotifyServices.GetTrackInfoFromName(tracksQuery, devAccount.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error finding track")
		}

		lyrics, err := spotifyServices.GetLyricsForTrack(track)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error getting track lyrics")
		}

		return c.JSON(lyrics)
	})

	app.Get("/userTopTracks/:userToken", func(c *fiber.Ctx) error {
		log.Println("/userTopTracks api call")
		userToken := c.Params("userToken")
		if userToken == "" {
			c.SendStatus(fiber.StatusBadRequest)
		}
		lyricsList := make([]schemas.Lyrics, 0)
		tracks, err := spotifyServices.GetTopTrackInfoForUser(userToken)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error finding user info")
		}

		for _, track := range tracks.Tracks {
			lyrics, err := spotifyServices.GetLyricsForTrack(&track)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("error getting track lyrics")
			}
			lyricsList = append(lyricsList, *lyrics)
		}

		return c.JSON(lyricsList)
	})

	return app
}
