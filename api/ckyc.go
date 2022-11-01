package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) checkCkycController(c *fiber.Ctx) error {

	var checkCKyc models.UserCkyc

	customErrors, err := ValidateRequest[models.UserCkyc](s, c, &checkCKyc)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CheckCkyc(&checkCKyc)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}

	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}
