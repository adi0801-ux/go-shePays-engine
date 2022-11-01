package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) createPhysicalCardController(c *fiber.Ctx) error {

	var createPhysicalCard models.CreatePhysicalCard

	customErrors, err := ValidateRequest[models.CreatePhysicalCard](s, c, &createPhysicalCard)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CreatePhysicalCard(&createPhysicalCard)
	if err != nil || (statusCode != http.StatusCreated && statusCode != http.StatusOK) {

		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) createVirtualCardController(c *fiber.Ctx) error {

	var createVirtualCard models.CreateVirtualCard

	customErrors, err := ValidateRequest[models.CreateVirtualCard](s, c, &createVirtualCard)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	statusCode, data, err := s.proxySrv.CreateVirtualCard(&createVirtualCard)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) assignCardController(c *fiber.Ctx) error {

	var assignCard models.AssignCard

	customErrors, err := ValidateRequest[models.AssignCard](s, c, &assignCard)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	//query param --
	cardId := c.Query("card_id")

	statusCode, data, err := s.proxySrv.AssignCard(cardId, &assignCard)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) activateCardController(c *fiber.Ctx) error {

	//query param
	cardId := c.Query("card_id")

	statusCode, data, err := s.proxySrv.ActivateCard(cardId)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) getCardDetailsController(c *fiber.Ctx) error {

	//query param --
	cardId := c.Query("card_id")

	statusCode, data, err := s.proxySrv.GetCardDetailsCard(cardId)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) blockCardController(c *fiber.Ctx) error {

	//query param
	cardId := c.Query("card_id")

	statusCode, data, err := s.proxySrv.BlockCard(cardId)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) unblockCardController(c *fiber.Ctx) error {

	//query param
	cardId := c.Query("card_id")

	statusCode, data, err := s.proxySrv.UnBlockCard(cardId)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

func (s *HTTPServer) initiateCardPinSetController(c *fiber.Ctx) error {

	var initiateCardPinSet models.AssignCard

	customErrors, err := ValidateRequest[models.AssignCard](s, c, &initiateCardPinSet)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.InitiateCardPinSet(initiateCardPinSet.UserId)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}

//func (s *HTTPServer) setCardPinController(c *fiber.Ctx) error {
//
//	var setCardPin models.SetCardPin
//
//	customErrors, err := ValidateRequest[models.SetCardPin](s, c, &setCardPin)
//	if err != nil {
//		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
//		return nil
//	}
//
//	statusCode, data, err := s.proxySrv.SetCardPin(&setCardPin)
//	if err != nil {
//		SendFullErrorResponse(c, statusCode, err, nil)
//		return err
//	}
//	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)
//
//	return nil
//}

func (s *HTTPServer) replaceCardController(c *fiber.Ctx) error {

	var replaceCard models.ReplaceCard

	customErrors, err := ValidateRequest[models.ReplaceCard](s, c, &replaceCard)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.ReplaceCard(&replaceCard)
	if err != nil {
		SendFullErrorResponse(c, statusCode, err, nil)
		return err
	}
	SendResponse(c, statusCode, 1, "SUCCESS", data, nil)

	return nil
}
