package api

import (
	"github.com/gofiber/fiber/v2"
)

func SendResponse(c *fiber.Ctx, statusCode int, status int, message string, data interface{}, errStruc interface{}) {

	err := c.Status(statusCode).JSON(
		&fiber.Map{
			"status":  status,
			"message": message,
			"data":    data,
			"error":   errStruc,
		},
	)
	if err != nil {
		return
	}

	return
}

func SendSuccessResponse(c *fiber.Ctx, statusCode int, status int, message string, data interface{}) {

	SendResponse(c, statusCode, status, message, data, nil)
}

func SendFullErrorResponse(c *fiber.Ctx, statusCode int, err error, errStruc interface{}) {
	SendResponse(c, statusCode, 0, err.Error(), nil, errStruc)
}

func errorResponse(c *fiber.Ctx, statusCode int, err error) {
	SendResponse(c, statusCode, 0, err.Error(), nil, nil)

	return
}
