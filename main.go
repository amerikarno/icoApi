package main

import (
	"github.com/amerikarno/icoApi/handlers"
	"github.com/amerikarno/icoApi/repository"
	"github.com/amerikarno/icoApi/usecases"
	"github.com/labstack/echo/v4"
)

func main() {
	filename := "json/api_province_with_amphure_tambon.json"
	repo := repository.NewPATRepository(filename)
	provinces, amphures, tambons, zipcode := repo.LoadPAT().GetProvinceAmphureTambonLists()
	e := echo.New()
	usecase := usecases.NewUsecases()
	handler := handlers.NewHandler(usecase)
	e.GET("verify/email/:email", handler.VerifyEmailHandler())
	e.GET("verify/mobile/:mobileno", handler.VerifyMobileNoHandler())
	e.GET("api/v1/all_provinces", handler.GetAllProvinces(provinces))
	e.GET("api/v1/amphures/:province", handler.GetAmphuresInProvince(amphures))
	e.GET("api/v1/tambons/:amphure", handler.GetTambonsInAmphure(tambons))
	e.GET("api/v1/zipcode/:zipname", handler.GetZipCode(zipcode))
	e.Logger.Fatal(e.Start(":1323"))
}
