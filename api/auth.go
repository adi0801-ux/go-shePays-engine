package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shepays/utils"
)

func (s *HTTPServer) AuthorizeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientKey := c.GetReqHeaders()["X-Api-Key"]
		if clientKey == "" {
			errorResponse(c, http.StatusUnauthorized, fmt.Errorf("no x-api-key header provided"))

			return nil
		}
		if clientKey != s.config.AuthApiKey {
			errorResponse(c, http.StatusUnauthorized, fmt.Errorf("x-api-key not matched"))

			return nil
		}

		err := c.Next()
		if err != nil {
			utils.Log.Error(err)
			return nil
		}

		return nil
	}
}
