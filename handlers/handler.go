package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/amerikarno/icoApi/models"
	"github.com/amerikarno/icoApi/usecases"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecases *usecases.Usecases
}

func NewHandler(usecases *usecases.Usecases) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) VerifyMobileNoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response models.VerifyMobileNoResponse
		var pages string
		response.RegistedMobileNo = strings.ToLower(c.Param("mobileno"))
		response.IsInvalidMobileNoFormat = false

		if !h.usecases.VerifyMobileNoFormat(response.RegistedMobileNo) {
			log.Printf("error: failed to verify mobile no: %+v", response.RegistedMobileNo)
			response.IsInvalidMobileNoFormat = true
			return c.JSON(http.StatusBadRequest, response)
		}

		if response.RegistedMobileNo == "0881112233" {
			pages = "3"
			response.IsRegistedMobileno = true
			response.RegistedPage = pages
		} else {
			response.IsRegistedMobileno = false
			response.RegistedPage = pages
		}

		fmt.Printf("response: %+v\n", response)
		return c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) VerifyEmailHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response models.VerifyEmailResponse
		var pages string
		response.RegistedEmail = strings.ToLower(c.Param("email"))
		response.IsInvalidEmailFormat = false

		if !h.usecases.VerifyEmailFormat(response.RegistedEmail) {
			log.Printf("error: failed to verify email: %+v", response.RegistedEmail)
			response.IsInvalidEmailFormat = true
			return c.JSON(http.StatusBadRequest, response)
		}

		if response.RegistedEmail == "registered@email.com" {
			pages = "3"
			response.IsRegistedEmail = true
			response.RegistedPage = pages
		} else {
			response.IsRegistedEmail = false
			response.RegistedPage = pages
		}

		fmt.Printf("response: %+v\n", response)
		return c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) GetAllProvinces(provinces []string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, provinces)
	}
}

func (h *Handler) GetAmphuresInProvince(amphures map[string][]string) echo.HandlerFunc {
	return func(c echo.Context) error {
		province := c.Param("province")
		return c.JSON(http.StatusOK, amphures[province])
	}
}

func (h *Handler) GetTambonsInAmphure(tambons map[string][]string) echo.HandlerFunc {
	return func(c echo.Context) error {
		amphure := c.Param("amphure")
		return c.JSON(http.StatusOK, tambons[amphure])
	}
}

func (h *Handler) GetZipCode(zipcode map[string]int) echo.HandlerFunc {
	return func(c echo.Context) error {
		zipname := c.Param("zipname")
		return c.JSON(http.StatusOK, zipcode[zipname])
	}
}