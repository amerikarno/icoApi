package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/amerikarno/icoApi/models"
	"github.com/amerikarno/icoApi/usecases"
	"github.com/labstack/echo/v4"
)

type OpenAccountHandler struct {
	usecases   *usecases.OpenAccountUsecases
	smtpConfig *models.SMTPConfig
}

func NewHandler(usecases *usecases.OpenAccountUsecases, smtpConfig *models.SMTPConfig) *OpenAccountHandler {
	return &OpenAccountHandler{usecases: usecases, smtpConfig: smtpConfig}
}

func (h *OpenAccountHandler) VerifyMobileNoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response models.VerifyMobileNoResponse
		response.RegistedMobileNo = strings.ToLower(c.Param("mobileno"))
		response.IsRegistedMobileno = h.usecases.CheckedMobileUsecase(response.RegistedMobileNo)

		return c.JSON(http.StatusOK, response)
	}
}

func (h *OpenAccountHandler) VerifyIDCardHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response models.VerifyIDCardResponse
		response.RegistedIDCard = strings.ToLower(c.Param("idcard"))
		response.IsRegistedIDCard = h.usecases.CheckedCitizenIDUsecase(response.RegistedIDCard)

		return c.JSON(http.StatusOK, response)
	}
}

func (h *OpenAccountHandler) VerifyEmailHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response models.VerifyEmailResponse
		response.RegistedEmail = strings.ToLower(c.Param("email"))
		response.IsRegistedEmail = h.usecases.CheckedEmailUsecase(response.RegistedEmail)

		return c.JSON(http.StatusOK, response)
	}
}
func (h *OpenAccountHandler) VerifyEmailMobileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response models.VerifyEmailMobileResponse
		var pages string
		response.RegistedMobileNo = strings.ToLower(c.Param("mobileno"))

		if !h.usecases.VerifyMobileNoFormat(response.RegistedMobileNo) {
			log.Printf("error: failed to verify mobile no: %+v", response.RegistedMobileNo)
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

		// fmt.Printf("response: %+v\n", response)
		response.RegistedEmail = strings.ToLower(c.Param("email"))

		if !h.usecases.VerifyEmailFormat(response.RegistedEmail) {
			log.Printf("error: failed to verify email: %+v", response.RegistedEmail)
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

func (h *OpenAccountHandler) GetAllProvinces(provinces []string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, provinces)
	}
}

func (h *OpenAccountHandler) GetAmphuresInProvince(amphures map[string][]string) echo.HandlerFunc {
	return func(c echo.Context) error {
		province := c.Param("province")
		return c.JSON(http.StatusOK, amphures[province])
	}
}

func (h *OpenAccountHandler) GetTambonsInAmphure(tambons map[string][]string) echo.HandlerFunc {
	return func(c echo.Context) error {
		amphure := c.Param("amphure")
		return c.JSON(http.StatusOK, tambons[amphure])
	}
}

func (h *OpenAccountHandler) GetZipCode(zipcode map[string]int) echo.HandlerFunc {
	return func(c echo.Context) error {
		zipname := c.Param("zipname")
		return c.JSON(http.StatusOK, zipcode[zipname])
	}
}

func (h *OpenAccountHandler) GetIDcard() echo.HandlerFunc {
	return func(c echo.Context) error {
		idcard := c.Param("idcard")
		return c.JSON(http.StatusOK, h.usecases.VerifyIDCardNumber(idcard))
	}
}

func (h *OpenAccountHandler) PostIDcard() echo.HandlerFunc {
	return func(c echo.Context) error {
		var postData models.CustomerInformations
		if err := c.Bind(&postData); err != nil {
			return c.JSON(http.StatusBadRequest, "")
		}

		postData.Pages = true
		postData.Create = time.Now().Local()
		tmp := strings.Split(postData.MarriageStatus, ".")
		fmt.Printf("temp: %v\n", tmp)
		postData.MarriageStatus = tmp[1]

		fmt.Printf("post data: %+v\n", postData)
		id, err := h.usecases.CreateCustomerInformationUsecase(postData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Printf("account id: %v\n", id)
		return c.JSON(http.StatusOK, id)
	}
}

func (h *OpenAccountHandler) PostPersonalInformations() echo.HandlerFunc {
	return func(c echo.Context) error {
		var postData models.PersonalInformationPostRequests
		if err := c.Bind(&postData); err != nil {
			fmt.Printf("error: %v\n", err)
			return c.JSON(http.StatusBadRequest, "")
		}

		fmt.Printf("post data: %+v\n", postData)
		id, err := h.usecases.UpdateCustomerPersonalInformationUsecase(postData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Printf("account id: %v\n", id)
		return c.JSON(http.StatusOK, id)
	}
}

func (h *OpenAccountHandler) HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, fmt.Sprintf("Service Available at %v", time.Now().Local()))
	}
}

func (h *OpenAccountHandler) PostCustomerExamsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var postData models.CustomerExamsRequest
		if err := c.Bind(&postData); err != nil {
			fmt.Printf("error: %v\n", err)
			return c.JSON(http.StatusBadRequest, "")
		}

		fmt.Printf("post data: %+v\n", postData)
		id, err := h.usecases.CreateCustomerExamsUsecase(postData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Printf("account id: %v\n", id)
		return c.JSON(http.StatusOK, id)
	}
}

func (h *OpenAccountHandler) PostCreateCustomerConfirmsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var postData models.CustomerConfirmsRequest
		if err := c.Bind(&postData); err != nil {
			fmt.Printf("error: %v\n", err)
			return c.JSON(http.StatusBadRequest, "")
		}

		fmt.Printf("post data: %+v\n", postData)
		id, err := h.usecases.CreateCustomerConfirmsUsecase(postData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Printf("account id: %v\n", id)
		return c.JSON(http.StatusOK, id)
	}
}

func (h *OpenAccountHandler) GetUpdateCustomerConfirmsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenID := c.Param("tokenID")
		fmt.Printf("token id: %+v\n", tokenID)
		resp, err := h.usecases.UpdateConfirmsUsecase(tokenID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if len(resp.TokenID) == 0 {
			return c.JSON(http.StatusBadRequest, "invalid token id")
		}

		if resp.IsConfirm {
			return c.JSON(http.StatusBadRequest, "token id have already been activated")
		}

		fmt.Printf("account id: %v\n", resp)
		return c.JSON(http.StatusOK, resp)
	}
}
