package main

import (
	"log"

	"github.com/amerikarno/icoApi/external"
	"github.com/amerikarno/icoApi/handlers"
	"github.com/amerikarno/icoApi/repository"
	"github.com/amerikarno/icoApi/usecases"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	filename := "json/api_province_with_amphure_tambon.json"
	repo := repository.NewPATRepository(filename)
	provinces, amphures, tambons, zipcode := repo.LoadPAT().GetProvinceAmphureTambonLists()

	openAccountsDB := initOpenAccountsDB()
	openAccountsRepo := repository.NewOpenAccountsRepository(openAccountsDB)

	external := external.NewExternalUuid()
	e := echo.New()
	usecase := usecases.NewOpenAccountUsecases(openAccountsRepo, external)
	handler := handlers.NewHandler(usecase)
	e.GET("verify/email/:email", handler.VerifyEmailHandler())
	e.GET("verify/mobile/:mobileno", handler.VerifyMobileNoHandler())
	e.GET("api/v1/all_provinces", handler.GetAllProvinces(provinces))
	e.GET("api/v1/amphures/:province", handler.GetAmphuresInProvince(amphures))
	e.GET("api/v1/tambons/:amphure", handler.GetTambonsInAmphure(tambons))
	e.GET("api/v1/zipcode/:zipname", handler.GetZipCode(zipcode))
	e.GET("api/v1/idcard/:idcard", handler.GetIDcard())
	e.POST("api/v1/idcard", handler.PostIDcard())
	e.Logger.Fatal(e.Start(":1323"))
}

func initOpenAccountsDB() *gorm.DB {
	dsn := "root:liverpool@tcp(127.0.0.1:3306)/open_accounts?charset=utf8mb4&parseTime=True&loc=Local"
	mysqldb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return mysqldb
}
