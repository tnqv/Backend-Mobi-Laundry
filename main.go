package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	cfg "d2d-backend/config"
	"net/url"
	"d2d-backend/common"
	"d2d-backend/accounts"
	"d2d-backend/orders"
	storeHandler "d2d-backend/store/handler"
	reviewHandler "d2d-backend/review/handler"
	reviewRepository "d2d-backend/review/repository"
	reviewService "d2d-backend/review/service"
	storeRepository "d2d-backend/store/repository"
	storeService "d2d-backend/store/service"
	serviceHandler "d2d-backend/service/handler"
	serviceService "d2d-backend/service/service"
	serviceRepository "d2d-backend/service/repository"
	roleHandler "d2d-backend/role/handler"
	roleService "d2d-backend/role/service"
	roleRepository "d2d-backend/role/repository"
)

var config cfg.Config
var environmentDb string

func init(){
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Laundry d2d service running on DEBUG mode")
		environmentDb = "development"
	}else{
		environmentDb = "production"
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
//
//func insertTestExampleValue(db *gorm.DB){
//	category1 := orders.Category{Name:"Combo Giặt + Sấy + Xả Quần áo",Description:"Combo Giặt + Sấy + Xả Quần áo"}
//	category2 := orders.Category{Name:"Combo Chăn Màn",Description:"Combo Chăn Màn"}
//	category3 := orders.Category{Name:"Combo Thú bông",Description:"Combo Thú bông"}
//	category4 := orders.Category{Name:"Dịch vụ giặt hấp (không bao gồm ủi)",Description:"Dịch vụ giặt hấp (không bao gồm ủi)"}
//	category5 := orders.Category{Name:"Combo Rèm Cửa",Description:"Combo Rèm Cửa"}
//
//	db.Create(&category1)
//	db.Create(&category2)
//	db.Create(&category3)
//	db.Create(&category4)
//	db.Create(&category5)
//}

func Migrate() {
	accounts.AutoMigrate()
	orders.AutoMigrate()
}

func main() {
	var (
		listenAddr string
	)


	listenAddr = config.GetString(`server.address`)
	dbHost := config.GetString(environmentDb + `.DatabaseConfig.DBHost`)
	dbUser := config.GetString(environmentDb + `.DatabaseConfig.DBUser`)
	dbName := config.GetString(environmentDb + `.DatabaseConfig.DBName`)
	dbPort := config.GetString(environmentDb + `.DatabaseConfig.DBPort`)
	dbPass := config.GetString(environmentDb + `.DatabaseConfig.DBPassword`)

	connection := fmt.Sprintf(common.TEMPLATE_DB_CONSTRING, dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Saigon")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db := common.Init(dsn)
	Migrate()

	defer db.Close()
	r := gin.Default()
	//Init repository

	v1 := r.Group("/api/v1")
	accounts.AccountsRouterRegister(v1.Group("/accounts"))
	accounts.RolesRouterRegister(v1.Group("/roles"))
	accounts.UsersRouterRegister(v1.Group("/users"))
	v1.Use(accounts.AuthMiddleware(false))
	//orders.ServicesRouterRegister(v1.Group("/service"))
	orders.OrdersRouterRegister(v1.Group("/orders"))
	orders.ServiceOrdersRouterRegister(v1.Group("/services/orders"))
	//orders.OrderStatusesRouterRegister(v1.Group("/orders/statuses"))
	orders.OrderStatusesRouterRegister(v1.Group("/notifications"))


	//Review
	reviewRepository := reviewRepository.NewMysqlReviewRepository()
	reviewService := reviewService.NewReviewService(reviewRepository)
	reviewHttpHandler := reviewHandler.NewReviewHttpHandler(v1.Group("/review"),reviewService)

	//Store
	storeRepository := storeRepository.NewMysqlStoreRepository()
	storeService := storeService.NewStoreService(storeRepository)
	storeHttpHandler := storeHandler.NewStoreHttpHandler(v1.Group("/store"),storeService)

	//Service
	serviceRepository := serviceRepository.NewMysqlServiceRepository()
	serviceService := serviceService.NewServiceService(serviceRepository)
	serviceHttpHandler := serviceHandler.NewServiceHttpHandler(v1.Group("/service"), serviceService)

	//Role
	roleRepository := roleRepository.NewMysqlRoleRepository()
	roleService := roleService.NewRoleService(roleRepository)
	roleHttpHandler := roleHandler.NewRoleHttpHandler(v1.Group("/role"), roleService)


	v1.Use(accounts.AuthMiddleware(true))
	storeHttpHandler.AuthorizedRequiredRoutes(v1.Group("/store"))
	reviewHttpHandler.AuthorizedRequiredRoutes(v1.Group("/review"))
	serviceHttpHandler.AuthorizedRequiredRoutes(v1.Group("/service"))
	roleHttpHandler.AuthorizedRequiredRoutes(v1.Group("/role"))
	// users.UserRegister(v1.Group("/user"))
	// users.ProfileRegister(v1.Group("/profiles"))


	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// test 1 to 1
	//tx1 := db.Begin()
	//userA := users.UserModel{
	//	Username: "AAAAAAAAAAAAAAAA",
	//	Email:    "aaaa@g.cn",
	//	Bio:      "hehddeda",
	//	Image:    nil,
	//}
	//tx1.Save(&userA)
	//tx1.Commit()
	//fmt.Println(userA)

	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	r.Run(listenAddr)
}