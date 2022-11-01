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

	savingsAccount := api.Group("/account")
	{
		savingsAccount.Post("/", s.createSavingsAccountController)
		savingsAccount.Get("/fund/transfer/status", s.checkPaymentStatusController)
		savingsAccount.Post("/fund/initiate", s.initiateNeftTransferController)
		savingsAccount.Post("/fund/transfer", s.validateNeftTransferController)
	}

	cardApis := api.Group("/card")
	{
		cardApis.Post("/physical", s.createPhysicalCardController)
		cardApis.Post("/virtual", s.createVirtualCardController)
		cardApis.Put("/assign", s.assignCardController)
		cardApis.Put("/activate", s.activateCardController)
		cardApis.Get("/details", s.getCardDetailsController)
		cardApis.Put("/block", s.blockCardController)
		cardApis.Put("/unblock", s.unblockCardController)
		cardApis.Post("/pin/init", s.initiateCardPinSetController)
		//cardApis.Post("/pin/set", s.setCardPinController)
		cardApis.Post("/replace", s.replaceCardController)

	}
	ckyc := api.Group("/ckyc")
	{
		ckyc.Post("/check", s.checkCkycController)
	}
	users := api.Group("/users")
	{
		users.Get("/", s.readUserController)
		users.Post("/", s.createUserController)
		users.Put("/", s.updateUserController)
		users.Post("/nominees", s.createUserNomineesController)
		users.Get("/nominees", s.readUserNomineesController)

	}
	address := api.Group("/address")
	{
		address.Post("/", s.createUserAddressController)
		address.Get("/", s.readUserAddressController)
		address.Put("/", s.updateUserAddressController)
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
