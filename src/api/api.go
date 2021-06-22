package api

import (
	"net/http"

	"github.com/andrei-dascalu/go-shortener/src/shortener"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	js "github.com/andrei-dascalu/go-shortener/src/serializer/json"
	ms "github.com/andrei-dascalu/go-shortener/src/serializer/msgpack"
)

type RedirectHandler struct {
	redirectService shortener.RedirectService
}

func (h RedirectHandler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &js.Redirect{}
}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return RedirectHandler{
		redirectService: redirectService,
	}
}

func GetHandler(shortnereHandler RedirectHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		code := c.Params("code", "")

		if code == "" {
			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"message": "No Code Provided",
			})
		}
		redirect, err := shortnereHandler.redirectService.Find(code)
		if err != nil {
			if errors.Cause(err) == shortener.ErrRedirectNotFound {
				return c.SendStatus(http.StatusNotFound)
			}
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.Redirect(redirect.URL, http.StatusMovedPermanently)
	}
}

func PostHandler(shortenerHandler RedirectHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")

		redirect, err := shortenerHandler.serializer(contentType).Decode(c.Request().Body())
		if err != nil {
			log.Error().Err(err).Msg("500-1")
			return c.SendStatus(http.StatusInternalServerError)
		}

		log.Warn().Str("url", redirect.URL).Msg("Preparing to store")
		err = shortenerHandler.redirectService.Store(redirect)
		if err != nil {
			if errors.Cause(err) == shortener.ErrRedirectInvalid {
				return c.SendStatus(http.StatusBadRequest)
			}

			log.Error().Err(err).Msg("500-2")
			return c.SendStatus(http.StatusInternalServerError)
		}
		responseBody, err := shortenerHandler.serializer(contentType).Encode(redirect)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		c.Response().Header.Set("Content-Type", contentType)

		return c.Status(http.StatusCreated).Send(responseBody)
	}
}
