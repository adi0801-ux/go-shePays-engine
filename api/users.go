package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) createUserController(c *fiber.Ctx) error {

	var createUsers models.UserDetail

	customErrors, err := ValidateRequest[models.UserDetail](s, c, &createUsers)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CreateUserAccount(&createUsers)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

// read
func (s *HTTPServer) readUserController(c *fiber.Ctx) error {

	userID := c.Query("user_id")

	statusCode, data, err := s.proxySrv.ReadUserAccount(userID)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

// update
func (s *HTTPServer) updateUserController(c *fiber.Ctx) error {

	var createUsers models.UserDetail

	customErrors, err := ValidateRequest[models.UserDetail](s, c, &createUsers)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.UpdateUserAccount(&createUsers)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}
