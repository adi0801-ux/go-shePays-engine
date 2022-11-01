package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) createUserNomineesController(c *fiber.Ctx) error {

	var createUsersNominees models.UserNomineesApi

	customErrors, err := ValidateRequest[models.UserNomineesApi](s, c, &createUsersNominees)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CreateUserNominees(&createUsersNominees)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

// read
func (s *HTTPServer) readUserNomineesController(c *fiber.Ctx) error {

	userID := c.Query("user_id")

	statusCode, data, err := s.proxySrv.ReadUserAccount(userID)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}
