package fiberredoc

import (
	"github.com/LcKoh0830/go-redoc"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func New(doc redoc.Redoc) fiber.Handler {
	return adaptor.HTTPHandlerFunc(doc.Handler())
}
