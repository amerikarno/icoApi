package main

import (
	"log"

	"github.com/amerikarno/icoApi/config"
	"github.com/amerikarno/icoApi/external"
	"github.com/amerikarno/icoApi/handlers"
	preinfoh "github.com/amerikarno/icoApi/handlers/preInfo"
	mw "github.com/amerikarno/icoApi/middleware"
	"github.com/amerikarno/icoApi/repository"
	adminLoginRepository "github.com/amerikarno/icoApi/repository/admin"
	preinfor "github.com/amerikarno/icoApi/repository/preInfo"
	"github.com/amerikarno/icoApi/usecases"
	adminLoginUsecases "github.com/amerikarno/icoApi/usecases/admin"
	preinfou "github.com/amerikarno/icoApi/usecases/preInfo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	// "github.com/labstack/echo/v4/middleware"

	// "golang.org/x/crypto/acme/autocert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	smtpConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fail to loading config %v", err)
	}
	filename := "json/api_province_with_amphure_tambon.json"
	repo := repository.NewPATRepository(filename)
	provinces, amphures, tambons, zipcode := repo.LoadPAT().GetProvinceAmphureTambonLists()

	openAccountsDB := initOpenAccountsDB()
	openAccountsRepo := repository.NewOpenAccountsRepository(openAccountsDB)

	external := external.NewExternalServices()
	e := echo.New()

	regular := mw.NewRegularMiddleware()
	logger := loadLogging()
	// e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
	// 	return key == "fda-authen-key", nil
	// }))

	usecase := usecases.NewOpenAccountUsecases(openAccountsRepo, external)
	handler := handlers.NewHandler(usecase, smtpConfig)
	e.GET("verify/email/:email/mobile/:mobileno", handler.VerifyEmailMobileHandler(), regular.Authen)
	e.GET("verify/email/:email", handler.VerifyEmailHandler(), regular.Authen)
	e.GET("verify/mobile/:mobileno", handler.VerifyMobileNoHandler(), regular.Authen)
	e.GET("verify/idcard/:idcard", handler.VerifyIDCardHandler(), regular.Authen)
	e.GET("api/v1/all_provinces", handler.GetAllProvinces(provinces), regular.Authen)
	e.GET("api/v1/amphures/:province", handler.GetAmphuresInProvince(amphures), regular.Authen)
	e.GET("api/v1/tambons/:amphure", handler.GetTambonsInAmphure(tambons), regular.Authen)
	e.GET("api/v1/zipcode/:zipname", handler.GetZipCode(zipcode), regular.Authen)
	e.GET("api/v1/idcard/:idcard", handler.GetIDcard(), regular.Authen)
	e.POST("api/v1/idcard", handler.PostIDcard(), regular.Authen)
	e.POST("api/v1/personalInformation", handler.PostPersonalInformations(), regular.Authen)
	e.POST("api/v1/customerExams", handler.PostCustomerExamsHandler(), regular.Authen)
	e.POST("api/v1/createCustomerConfirms", handler.PostCreateCustomerConfirmsHandler(), regular.Authen)
	e.GET("api/v1/updateCustomerConfirms/:tokenID", handler.GetUpdateCustomerConfirmsHandler())
	e.POST("healthcheck", handler.HealthCheck())

	adminPassword := adminLoginRepository.NewAdminPassword()
	adminRepo := adminLoginRepository.NewLoginRepository(openAccountsDB, logger)
	adminUsecase := adminLoginUsecases.NewAdminLoginUsecase(adminRepo, external, adminPassword, logger)
	adminHandler := handlers.NewAdminHandler(adminUsecase, logger)
	gAdmin := e.Group("/admin/v1")
	gAdmin.GET("/healthcheck", handler.HealthCheck())
	gAdmin.POST("/create", adminHandler.CreateHandler())
	gAdmin.POST("/login", adminHandler.LoginHandler())
	gAdmin.GET("/refresh", adminHandler.RefreshTokenHandler())
	dashboard := e.Group("/admin/v1/dashboard")
	dashboard.POST("/users", adminHandler.LoginHandler())

	// preinfo
	preinfoRepo := preinfor.NewPreInfoRepository(openAccountsDB)
	preinfoUsecase := preinfou.NewPreInfoUsecase(preinfoRepo)
	preinfoHandler := preinfoh.NewPreInfoHander(preinfoUsecase)

	api := e.Group("/api")
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{},
	}))
	api.POST("/getTitleName", preinfoHandler.GetTitles())
	api.POST("/checkExistMobile", preinfoHandler.CheckExistMobile())
	api.POST("/checkExistEmail", preinfoHandler.CheckExistEmail())
	api.POST("/checkExistIdcard", preinfoHandler.CheckExistIDcard())
	api.POST("/clearViewCount", preinfoHandler.ClearViewCount())
	api.GET("/loadIDcard", preinfoHandler.LoadIDcard())
	api.POST("/saveTempdata", preinfoHandler.SaveTempdata())
	api.POST("/currentPage", preinfoHandler.CheckCurrentPage())
	api.POST("/getBasicDropdown", preinfoHandler.GetBasicDropdown())
	api.GET("/checkBasicInfo", preinfoHandler.ClearViewCount())
	api.POST("/get-geo-data", preinfoHandler.GetTAPInfo())
	api.POST("/get-Business-type-dropdown", preinfoHandler.GetCarrerTypes())
	api.POST("/get-basic-branch-dropdown", preinfoHandler.GetBankBranch())
	
	// server := http.Server{
	// 	Addr: ":1323",
	// 	Handler: e,
	// 	TLSConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }

	// if err := server.ListenAndServeTLS("./cert/myCA.pem", "./cert/myCA.crt"); err != nil {log.Fatal(err)}
	// e.Logger.Fatal(e.StartTLS(":1323", "./cert/myCA.pem", "./cert/myCA.crt" ))
	e.Logger.Fatal(e.Start(":1323"))
}

func NewRegularMiddleware() {
	panic("unimplemented")
}

func initOpenAccountsDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/open_accounts?charset=utf8mb4&parseTime=True&loc=Local" // for localhost
	// dsn := "root:icodb@Liverp00l!!@tcp(34.136.109.173:3306)/open_accounts?charset=utf8mb4&parseTime=True&loc=Local" // for gcp
	mysqldb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return mysqldb
}

func loadLogging() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "source",
			EncodeCaller: zapcore.ShortCallerEncoder,

			NameKey: "requestID",
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		logger.Error("error:", zap.Error(err))
	}

	return logger
}
