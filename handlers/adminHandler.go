package handlers

import (
	"net/http"

	"github.com/amerikarno/icoApi/models"
	"github.com/amerikarno/icoApi/token"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AdminHandler struct {
	usecases IAdminUsecases
	logger   *zap.Logger
}

func NewAdminHandler(usecases IAdminUsecases, logger *zap.Logger) *AdminHandler {
	return &AdminHandler{usecases, logger}
}

func (h *AdminHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request models.AdminLoginRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		h.logger.Info("request data", zap.Any("request", request), zap.Int("length of hashed user", len(request.HashedUsername)), zap.Int("length of hashed user", len(request.HashedPassword)))
		// var response models.JwtUserModel

		resp := h.usecases.Verify(request.HashedUsername, request.HashedPassword)
		if resp.Error != nil {
			return c.JSON(http.StatusBadRequest, resp.Error)
		}
		h.logger.Info("resp data:", zap.String("AccessToken:", resp.AccessToken), zap.String("RequestToken:", resp.RequestToken), zap.Error(resp.Error))
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *AdminHandler) CreateHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request models.AdminCreateRequestModel
		if err := c.Bind(&request); err != nil {
			h.logger.Error("error:", zap.Error(err))
			return c.JSON(http.StatusBadRequest, err)
		}

		response := h.usecases.Create(request)

		return c.JSON(http.StatusOK, response)
	}
}

func (h *AdminHandler) RefreshTokenHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		rCliam, err := token.ExtractRefreshClaims(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)

		}

		err = rCliam.Rotate()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		refreshTokenString, err := rCliam.JwtString()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		aClaim := token.NewAccessClaims(&rCliam.JwtUserModel)
		accessTokenString, err := aClaim.JwtString()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)

		}

		response := models.AdminLoginResponse{
			AccessToken: accessTokenString,
			RequestToken: refreshTokenString,
		}
		return c.JSON(http.StatusOK, response)
	}
}
