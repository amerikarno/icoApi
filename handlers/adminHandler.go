package handlers

import (
	"log"
	"net/http"

	"github.com/amerikarno/icoApi/models"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	usecases IAdminUsecases
}

func NewAdminHandler(usecases IAdminUsecases) *AdminHandler {
	return &AdminHandler{usecases}
}

func (h *AdminHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request models.AdminLoginRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		log.Printf("request data: %+v, %v, %v", request, len(request.HashedUsername), len(request.HashedPassword))
		var response models.JwtUserModel

		resp := h.usecases.Verify(request.HashedUsername, request.HashedPassword)
		if resp.Error != nil {
			return c.JSON(http.StatusBadRequest, resp.Error)
		}
		return c.JSON(http.StatusOK, response)
	}
}

// func (h *AdminHandler) hashString(input string) string {
// 	hasher := sha256.New()
// 	hasher.Write([]byte(input))
// 	hashedBytes := hasher.Sum(nil)
// 	return hex.EncodeToString(hashedBytes)
// }
