package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/andrei-dascalu/go-shortener/src/api"
	"github.com/andrei-dascalu/go-shortener/src/shortener"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	mr "github.com/andrei-dascalu/go-shortener/src/repository/mongo"
	rr "github.com/andrei-dascalu/go-shortener/src/repository/redis"
)

func main() {
	repo, err := chooseRepo()

	if err != nil {
		log.Fatal().Err(err).Msg("no backend selected")
	}

	service := shortener.NewRedirectService(repo)
	handler := api.NewHandler(service)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/:code", api.GetHandler(handler))
	app.Post("/", api.PostHandler(handler))

	log.Panic().Err(app.Listen(":8080")).Msg("Failed to start service")
}

func chooseRepo() (shortener.RedirectRepository, error) {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			return nil, err
		}
		return repo, nil
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			return nil, err
		}
		return repo, nil
	}
	return nil, errors.New("no backend was set")
}
