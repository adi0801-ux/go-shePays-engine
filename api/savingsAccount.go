package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) createSavingsAccountController(c *fiber.Ctx) error {

	userID := c.FormValue("user_id")
	//check if account already not created

	statusCode, data, err := s.proxySrv.CreateSavingsAccount(userID)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) initiateNeftTransferController(c *fiber.Ctx) error {

	var initiateNeftTransfer models.InitiateNeftTransfer

	customErrors, err := ValidateRequest[models.InitiateNeftTransfer](s, c, &initiateNeftTransfer)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.InitiateNeftTransfer(&initiateNeftTransfer)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) validateNeftTransferController(c *fiber.Ctx) error {

	var validateNeftTransfer models.ValidateNeftTransafer

	customErrors, err := ValidateRequest[models.ValidateNeftTransafer](s, c, &validateNeftTransfer)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.ValidateNeftTransfer(&validateNeftTransfer)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) checkPaymentStatusController(c *fiber.Ctx) error {

	//query param
	paymentId := c.Query("payment_id")
	if paymentId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError+" payment_id"))
		return nil

	}
	statusCode, data, err := s.proxySrv.CheckPaymentStatus(paymentId)

	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}
