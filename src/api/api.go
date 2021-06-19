package api

import (
	"net/http"

	"github.com/andrei-dascalu/go-shortener/src/shortener"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

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
		code := c.Query("code", "")
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
			return c.SendStatus(http.StatusInternalServerError)
		}
		err = shortenerHandler.redirectService.Store(redirect)
		if err != nil {
			if errors.Cause(err) == shortener.ErrRedirectInvalid {
				return c.SendStatus(http.StatusBadRequest)
			}
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
