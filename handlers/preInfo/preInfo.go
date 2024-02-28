package preinfo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/amerikarno/icoApi/mockJson"
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
			resp.Error = "ใช้ mobile นี้ เปิดบัญชีได้"
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
		var resp models.OTPresponse
		mapRequest := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&mapRequest)
		if err != nil {
			return err
		}
		module := mapRequest["module"]
		fmt.Printf("request: %+v\n", mapRequest)
		fmt.Printf("module: %v\n", module)

		if module == "module_preinfo" {
			save := models.SaveTemplate{}
			if err = c.Bind(&save); err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			log.Printf("save: %+v\n", save)
		}
		if module == "module_otp" {
			var otpdata models.OTPdata
			var otpinfo models.OTPinfo
			var otpnext models.OTPNextModule
			// var databyte, infobyte, nextbyte []byte
			databyte := []byte(mockJson.OTPdata)
			infobyte := []byte(mockJson.OTPinfo)
			nextbyte := []byte(mockJson.OTPnext)

			err = json.Unmarshal(databyte, &otpdata)
			if err != nil {
				log.Printf("error: %v", err)
				panic(err)
			}
			err = json.Unmarshal(infobyte, &otpinfo)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(nextbyte, &otpnext)
			if err != nil {
				panic(err)
			}
			resp.Data = otpdata
			resp.Data.Info = otpinfo
			resp.NextModule = otpnext
		}
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
func (h *PreInfoHandler) GetBasicDropdown() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.BasicInfo{}
		data := []byte(mockJson.Basic)

		err := json.Unmarshal(data, &resp)
		if err != nil {
			panic(err)
		}
		// resp.Success = false
		// if resp.Success {
		// 	resp.Error = "ไม่มี user นี้"
		// }
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) GetTAPInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		tap := []models.TAP{}
		data := []byte(mockJson.TAP)

		err := json.Unmarshal(data, &tap)
		if err != nil {
			panic(err)
		}
		var resp models.TAPresponse
		resp.ResultData = tap
		resp.Success = true
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) GetCarrerTypes() echo.HandlerFunc {
	return func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}
		id := json_map["id"]
		var data []byte
		fmt.Printf("id: %v\n", id)

		cct := []models.CurrentCareerTypeST{}
		if id == "1" {
			data = []byte(mockJson.First)
		}
		if id == "2" {
			data = []byte(mockJson.Second)
		}
		if id == "3" {
			data = []byte(mockJson.Third)
		}
		if id == "4" {
			data = []byte(mockJson.Forth)
		}
		if id == "5" {
			data = []byte(mockJson.Fifth)
		}
		if id == "6" {
			data = []byte(mockJson.Sixth)
		}
		if id == "7" {
			data = []byte(mockJson.Seventh)
		}
		if id == "8" {
			data = []byte(mockJson.Eighth)
		}
		if id == "9" {
			data = []byte(mockJson.Ninth)
		}
		if id == "10" {
			data = []byte(mockJson.Tenth)
		}
		if id == "11" {
			data = []byte(mockJson.Eleventh)
		}
		if id == "12" {
			data = []byte(mockJson.Twelth)
		}
		if id == "13" {
			data = []byte(mockJson.Thirteenth)
		}
		if id == "14" {
			data = []byte(mockJson.Forteenth)
		}
		if id == "15" {
			data = []byte(mockJson.Fifteenth)
		}

		err = json.Unmarshal(data, &cct)
		if err != nil {
			panic(err)
		}
		var resp models.BasicInfo
		resp.CurrentCareerType = cct
		resp.Success = true
		return c.JSON(http.StatusOK, resp)
	}
}
func (h *PreInfoHandler) GetBankBranch() echo.HandlerFunc {
	return func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}
		id := json_map["id"]
		var data []byte
		fmt.Printf("id: %v\n", id)

		cct := []models.BankBranchST{}
		if id == "002" {
			data = []byte(mockJson.BB002)
		}
		if id == "004" {
			data = []byte(mockJson.BB004)
		}
		if id == "006" {
			data = []byte(mockJson.BB006)
		}

		err = json.Unmarshal(data, &cct)
		if err != nil {
			panic(err)
		}
		var resp models.BasicInfo
		resp.BankBranch = cct
		resp.Success = true
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) GetSuiteTestQuestions() echo.HandlerFunc {
	return func(c echo.Context) error {
		questions := []models.SuiteTestQuestions{}
		data := []byte(mockJson.SuiteTestQuestions)

		err := json.Unmarshal(data, &questions)
		if err != nil {
			panic(err)
		}
		var resp models.SuiteTestResponse
		resp.Items = questions
		resp.ID = "2"
		resp.Version = "1"
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) GetSuiteSelect() echo.HandlerFunc {
	return func(c echo.Context) error {
		lists := []models.SuiteSelect{}
		data := []byte(mockJson.SuiteSelects)

		err := json.Unmarshal(data, &lists)
		if err != nil {
			panic(err)
		}
		var resp models.SuiteSelectLists
		resp.InvestmentRisk = lists
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) GetKnowledgeTestQuestions() echo.HandlerFunc {
	return func(c echo.Context) error {
		var resp models.KnowledgeTestResponse
		var questions []models.KnowledgeTestQuestion
		data := []byte(mockJson.KnowledgeTestQuestion)

		err := json.Unmarshal(data, &questions)
		if err != nil {
			panic(err)
		}

		resp.Items = questions
		resp.ID = "1"
		resp.Version = "1"

		// log.Printf("knowledge questions: %+v", resp.Items)
		return c.JSON(http.StatusOK, resp)
	}
}

func (h *PreInfoHandler) FATCAmodule() echo.HandlerFunc {
	return func(c echo.Context) error {
		var resp models.FATCAmodule
		data := []byte(mockJson.FATCAmodule)

		err := json.Unmarshal(data, &resp)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
