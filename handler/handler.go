package handler

import (
	"fmt"

	"github.com/aburizalpurnama/my-go-fiber/config"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Greet(c *fiber.Ctx) error
}

type handlerImpl struct {
	Config config.MainConfig
}

func NewHandler(conf config.MainConfig) Handler {
	return &handlerImpl{Config: conf}
}

func (h *handlerImpl) Greet(c *fiber.Ctx) error {
	name := c.Params("name", "no name")
	c.SendStatus(fiber.StatusOK)
	return c.SendString(fmt.Sprintf("Hi, %s", name))
}

