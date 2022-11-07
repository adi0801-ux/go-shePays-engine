package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) checkDeviceController(c *fiber.Ctx) error {

	var deviceIdentifier models.DeviceIdentifier

	customErrors, err := ValidateRequest[models.DeviceIdentifier](s, c, &deviceIdentifier)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CheckDeviceVersion(&deviceIdentifier)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) registerDeviceController(c *fiber.Ctx) error {

	var registerDevice models.RegisterDevice

	customErrors, err := ValidateRequest[models.RegisterDevice](s, c, &registerDevice)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.RegisterDevice(&registerDevice)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) registerCustomerDeviceController(c *fiber.Ctx) error {

	var registerCustomerDevice models.CustomerInformation

	customErrors, err := ValidateRequest[models.CustomerInformation](s, c, &registerCustomerDevice)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CustomerInformation(&registerCustomerDevice)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) smsGenRequestController(c *fiber.Ctx) error {

	var smsGenRequest models.SMSGenReq

	customErrors, err := ValidateRequest[models.SMSGenReq](s, c, &smsGenRequest)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.SMSGenRequest(&smsGenRequest)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}
