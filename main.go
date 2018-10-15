package main

import (
	accountHandler "d2d-backend/account/handler"
	accountRepository "d2d-backend/account/repository"
	accountService "d2d-backend/account/service"
	"d2d-backend/accounts"
	categoryHandler "d2d-backend/category/handler"
	categoryRepository "d2d-backend/category/repository"
	categoryService "d2d-backend/category/service"
	"d2d-backend/common"
	cfg "d2d-backend/config"
	orderStatusHandler "d2d-backend/orderStatus/handler"
	orderStatusRepository "d2d-backend/orderStatus/repository"
	orderStatusService "d2d-backend/orderStatus/service"
	"d2d-backend/orders"
	reviewHandler "d2d-backend/review/handler"
	reviewRepository "d2d-backend/review/repository"
	reviewService "d2d-backend/review/service"
	roleHandler "d2d-backend/role/handler"
	roleRepository "d2d-backend/role/repository"
	roleService "d2d-backend/role/service"
	serviceHandler "d2d-backend/service/handler"
	serviceRepository "d2d-backend/service/repository"
	serviceService "d2d-backend/service/service"
	serviceOrderHandler "d2d-backend/serviceOrder/handler"
	serviceOrderRepository "d2d-backend/serviceOrder/repository"
	serviceOrderService "d2d-backend/serviceOrder/service"
	storeHandler "d2d-backend/store/handler"
	storeRepository "d2d-backend/store/repository"
	storeService "d2d-backend/store/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	userService "d2d-backend/user/service"
	userRepository "d2d-backend/user/repository"
	userHandler "d2d-backend/user/handler"
	placedOrderRepository	"d2d-backend/placedOrder/repository"
	placeOrderService		"d2d-backend/placedOrder/service"
	placedOrderHandler		"d2d-backend/placedOrder/handler"
	notificationService "d2d-backend/notification/service"
	notificationRepository "d2d-backend/notification/repository"
	notificationHandler "d2d-backend/notification/handler")

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


	listenAddr = config.GetString(environmentDb + `.serverAddress`)
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

	//Category
	categoryRepository := categoryRepository.NewMysqlCategoryRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository)
	categoryHttpHandler := categoryHandler.NewCategoryHttpHandler(v1.Group("/category"),categoryService)

	//OrderStatus
	orderStatusRepository := orderStatusRepository.NewMysqlOrderStatusRepository()
	orderStatusService := orderStatusService.NewOrderStatusService(orderStatusRepository)
	orderStatusHttpHandler := orderStatusHandler.NewOrderStatusHttpHandler(v1.Group("/orderstatus"),orderStatusService)

	//ServiceOrder
	serviceOrderRepository := serviceOrderRepository.NewMysqlServiceOrderRepository()
	serviceOrderService := serviceOrderService.NewServiceOrderService(serviceOrderRepository)
	serviceOrderHttpHandler := serviceOrderHandler.NewServiceOrderHttpHandler(v1.Group("/serviceorder"), serviceOrderService)

	//Account
	accountRepository := accountRepository.NewMysqlAccounteRepository()
	accountService := accountService.NewAccountService(accountRepository)
	accountHttpHandler := accountHandler.NewAccountHttpHandler(v1.Group("/account"), accountService)

	//PlacedOrder
	placedOrderRepository := placedOrderRepository.NewMysqlPlacedOrderRepository()
	placedOrderService := placeOrderService.NewPlacedOrderService(placedOrderRepository)
	placedOrderHttpHandler := placedOrderHandler.NewPlacedOrderHttpHandler(v1.Group("/placedorder"), placedOrderService)

	//Notification
	notificationRepository := notificationRepository.NewMysqlNotificationRepository()
	notificationService := notificationService.NewNotificationService(notificationRepository)
	notificationHttpHandler := notificationHandler.NewNotificationHttpHandler(v1.Group("/notification"), notificationService)

	//User
	userRepository := userRepository.NewMysqlUserRepository()
	userService := userService.NewUserService(userRepository)
	userHttpHandler := userHandler.NewUserHttpHandler(v1.Group("/user"), userService,notificationService,placedOrderService)

	//Authorized
	v1.Use(accounts.AuthMiddleware(true))
	storeHttpHandler.AuthorizedRequiredRoutes(v1.Group("/store"))
	reviewHttpHandler.AuthorizedRequiredRoutes(v1.Group("/review"))
	serviceHttpHandler.AuthorizedRequiredRoutes(v1.Group("/service"))
	roleHttpHandler.AuthorizedRequiredRoutes(v1.Group("/role"))
	categoryHttpHandler.AuthorizedRequiredRoutes(v1.Group("/category"))
	orderStatusHttpHandler.AuthorizedRequiredRoutes(v1.Group("/orderstatus"))
	serviceOrderHttpHandler.AuthorizedRequiredRoutes(v1.Group("/serviceorder"))
	accountHttpHandler.AuthorizedRequiredRoutes(v1.Group("/account"))	
	userHttpHandler.AuthorizedRequiredRoutes(v1.Group("/user"))	
	placedOrderHttpHandler.AuthorizedRequiredRoutes(v1.Group("/placedorder"))
	notificationHttpHandler.AuthorizedRequiredRoutes(v1.Group("/notification"))

	// users.UserRegister(v1.Group("/user"))	// users.ProfileRegister(v1.Group("/profiles"))


	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	r.Run(listenAddr)
}