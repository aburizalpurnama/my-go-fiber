package handler

import (
	"github.com/aburizalpurnama/my-go-fiber/model/request"
	"github.com/aburizalpurnama/my-go-fiber/model/respose"
	"github.com/aburizalpurnama/my-go-fiber/usecase"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {
		CreateUser(c *fiber.Ctx) error
	}

	userHandlerImpl struct {
		userUsecase usecase.UserUsecase
	}
)

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandlerImpl{userUsecase}
}

func (h *userHandlerImpl) CreateUser(c *fiber.Ctx) error {
	var (
		reqPayload request.CreateUser
		resPayload respose.Default
	)

	err := c.BodyParser(&reqPayload)
	if err != nil {
		resPayload.Code = "99"
		resPayload.Message = "invalid request payload"
		c.SendStatus(fiber.ErrBadRequest.Code)
		return c.JSON(resPayload)
	}

	resPayload, err = h.userUsecase.CreateUser(reqPayload)
	if err != nil {
		c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	return c.JSON(resPayload)
}
