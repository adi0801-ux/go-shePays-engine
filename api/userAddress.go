package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) createUserAddressController(c *fiber.Ctx) error {

	var createUsersAddress models.UserAddressInput

	customErrors, err := ValidateRequest[models.UserAddressInput](s, c, &createUsersAddress)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CreateUserAddress(&createUsersAddress)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

// read
func (s *HTTPServer) readUserAddressController(c *fiber.Ctx) error {

	userID := c.Query("user_id")

	statusCode, data, err := s.proxySrv.ReadUserAddress(userID)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

// update
func (s *HTTPServer) updateUserAddressController(c *fiber.Ctx) error {

	var createUsersAddress models.UserAddressInput

	customErrors, err := ValidateRequest[models.UserAddressInput](s, c, &createUsersAddress)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.UpdateUserAddress(&createUsersAddress)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}
