package preinfo

import (
	"net/http"

	"github.com/amerikarno/icoApi/models"
	preinfo "github.com/amerikarno/icoApi/usecases/preInfo"
	"github.com/labstack/echo/v4"
)

type PreInfoHandler struct {
	usecase *preinfo.PreInfoUsecase
}

func NewPreInfoHander(usecase *preinfo.PreInfoUsecase) *PreInfoHandler {
	return &PreInfoHandler{usecase: usecase}
}

func (h *PreInfoHandler) GetTitles() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := h.usecase.GetTitles()
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) CheckExistEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = true
		if resp.Success {
			resp.Error = "ใช้ email นี้ เปิดบัญชีได้"
		}
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) CheckExistMobile() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = true
		if resp.Success {
			resp.Error = "ใช้ email นี้ เปิดบัญชีได้"
		}
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) CheckExistIDcard() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = true
		if resp.Success {
			resp.Error = "ไม่มี ID card สมัครเข้าระบบได้"
		}
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) SaveTempdata() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = true
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) LoadIDcard() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = true
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) ClearViewCount() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = true
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) CheckCurrentPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.CheckExist{}
		resp.Success = false
		if resp.Success {
			resp.Error = "ไม่มี user นี้"
		}
		return c.JSON(http.StatusOK, resp)
	}
}
