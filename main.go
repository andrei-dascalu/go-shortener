package main

import (
	"log"
	"os"
	"strconv"

	"github.com/andrei-dascalu/go-shortener/src/api"
	"github.com/andrei-dascalu/go-shortener/src/shortener"
	"github.com/gofiber/fiber/v2"

	mr "github.com/andrei-dascalu/go-shortener/src/repository/mongo"
	rr "github.com/andrei-dascalu/go-shortener/src/repository/redis"
)

func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := api.NewHandler(service)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/{code}", api.GetHandler(handler))
	app.Post("/", api.PostHandler(handler))

	app.Listen(":3000")
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
