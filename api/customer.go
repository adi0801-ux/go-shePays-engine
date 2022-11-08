package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (s *HTTPServer) customerAdditionalInfomationController(c *fiber.Ctx) error {

	var customerAdditionalInformation models.CustomerAdditionalInformation

	customErrors, err := ValidateRequest[models.CustomerAdditionalInformation](s, c, &customerAdditionalInformation)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CustomerAdditionalInformation(&customerAdditionalInformation)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) setMPINController(c *fiber.Ctx) error {

	var setMPIN models.SetMPIN

	customErrors, err := ValidateRequest[models.SetMPIN](s, c, &setMPIN)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.SetMPINProxy(&setMPIN)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) PANVerifyController(c *fiber.Ctx) error {

	var kycPan models.KYCPAN

	customErrors, err := ValidateRequest[models.KYCPAN](s, c, &kycPan)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.CheckPANExists(&kycPan)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) AadhaarVerifyController(c *fiber.Ctx) error {

	var kycAadhar models.KYCAadharVerify

	//get file
	file, err := c.FormFile("aadhar_file")
	if err != nil {
		return err
	}

	kycAadhar.AddharFile = file
	kycAadhar.Password = c.FormValue("password")
	kycAadhar.UserId = c.FormValue("user_id")

	statusCode, data, err := s.proxySrv.VerifyAadhar(&kycAadhar)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) UploadSelfieController(c *fiber.Ctx) error {

	var selfie models.Selfie

	customErrors, err := ValidateRequest[models.Selfie](s, c, &selfie)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.VerifySelfie(&selfie)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) ValidateDOBController(c *fiber.Ctx) error {

	var dob models.ValidateDOB

	customErrors, err := ValidateRequest[models.ValidateDOB](s, c, &dob)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.ValidateDOB(&dob)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) AOFCreationController(c *fiber.Ctx) error {

	var aof models.AoFModel

	customErrors, err := ValidateRequest[models.AoFModel](s, c, &aof)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.AoFAPI(&aof)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) GetOTPController(c *fiber.Ctx) error {

	var user models.UserId

	customErrors, err := ValidateRequest[models.UserId](s, c, &user)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.GetOTPProxy(&user)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) SendOTPController(c *fiber.Ctx) error {

	var otp models.OTPVerify

	customErrors, err := ValidateRequest[models.OTPVerify](s, c, &otp)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.VerifyOTPProxy(&otp)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) GetVcifController(c *fiber.Ctx) error {

	var user models.UserId

	customErrors, err := ValidateRequest[models.UserId](s, c, &user)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.VCifAPI(&user)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}

func (s *HTTPServer) CreateAccountController(c *fiber.Ctx) error {

	var user models.UserId

	customErrors, err := ValidateRequest[models.UserId](s, c, &user)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	statusCode, data, err := s.proxySrv.AccountCreateProxy(&user)

	SendResponse(c, statusCode, 1, constants.SuccessResponseMessage, data, err)
	return nil
}
