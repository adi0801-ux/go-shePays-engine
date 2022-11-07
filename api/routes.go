package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"shepays/services"
	"shepays/utils"
)

// Server serves HTTP requests for this service

type HTTPServer struct {
	router    *fiber.App
	config    *utils.Config
	validator *validator.Validate
	proxySrv  *services.ServiceConfig
}

func (s *HTTPServer) RegisterRoutes(router *fiber.App) {

	router.Get("/", s.healthCheck)

	api := router.Group("/api")

	device := api.Group("/device")
	{
		device.Post("/version/check", s.checkDeviceController)
		device.Post("/register", s.registerDeviceController)
		device.Post("/customer/information", s.registerCustomerDeviceController)
	}
	sms := api.Group("/sms")
	{
		sms.Post("/request/gen", s.smsGenRequestController)
	}
	mPIN := api.Group("/mPIN")
	{
		mPIN.Post("/set", s.setMPINController)
	}

	customer := api.Group("/customer")
	{
		customer.Post("/additional/information", s.customerAdditionalInfomationController)

		kyc := customer.Group("/kyc")
		{
			kyc.Post("/pan", s.PANVerifyController)
			kyc.Post("/aadhar", s.AadhaarVerifyController)
			kyc.Post("/selfie", s.UploadSelfieController)
			kyc.Post("/dob", s.ValidateDOBController)
		}
	}

}

// GetNewServer creates a new Http server and setup routing
func GetNewServer(
	config *utils.Config,
	proxySrv *services.ServiceConfig) *HTTPServer {

	validate := validator.New()

	httpServer := &HTTPServer{config: config, validator: validate, proxySrv: proxySrv}

	router := fiber.New()

	// Add API Logger to Router
	router.Use(utils.LoggerToFile())

	// Setup Routes here:
	httpServer.RegisterRoutes(router)

	httpServer.router = router
	return httpServer
}

// StartServer Start the Gin Server at a specific address
func (s *HTTPServer) StartServer(a string) error {
	return s.router.Listen(a)
}

func (s *HTTPServer) healthCheck(c *fiber.Ctx) error {

	SendResponse(c, fiber.StatusOK, 1, "Alive!", nil, nil)
	return nil
}
