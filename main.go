package main

import (
	"fmt"
	"log"
	"gopkg.in/gin-gonic/gin.v1"
	cfg "d2d-backend/config"
	"net/url"
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

func main() {
	var (
		listenAddr string
	)

	// Migrate(db)
	// defer db.Close()

	listenAddr = config.GetString(`server.address`)
	//dbHost := config.GetString(environmentDb + `.DatabaseConfig.DBHost`)
	//dbUser := config.GetString(environmentDb + `.DatabaseConfig.DBUser`)
	//dbName := config.GetString(environmentDb + `.DatabaseConfig.DBName`)
	//dbPort := config.GetString(environmentDb + `.DatabaseConfig.DBPort`)
	//dbPass := config.GetString(environmentDb + `.DatabaseConfig.DBPassword`)

	//connection := fmt.Sprintf(common.TEMPLATE_DB_CONSTRING, dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Saigon")

	//dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	//db := common.Init(dsn)

	r := gin.Default()

	//v1 := r.Group("/api")
	//users.UsersRegister(v1.Group("/users"))
	//v1.Use(users.AuthMiddleware(false))
	// articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	// articles.TagsAnonymousRegister(v1.Group("/tags"))

	//v1.Use(users.AuthMiddleware(true))
	// users.UserRegister(v1.Group("/user"))
	// users.ProfileRegister(v1.Group("/profiles"))

	// articles.ArticlesRegister(v1.Group("/articles"))

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