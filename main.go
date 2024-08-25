package main

import (
	"code-math/component/appctx"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main(){
	router:=gin.Default();
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
		"message":"pong",
		})
	})
	//cấu hình db orm
	connectString := os.Getenv("DB_CONNECTION_STRING")
	db,err:=gorm.Open(mysql.Open(connectString),&gorm.Config{
		//gorm.config cấu hình logger, namingstrategy,allow global update,...
		Logger: logger.Default.LogMode(logger.Info), 
	})
	if err!=nil{
		panic("failed to connect database")
	}
	db=db.Debug()
	//tạo app context 
	appContext:=appctx.NewAppContext(db,"")
	//cung cấp app context cho toàn hộ life cycle project
	v1 := router.Group("/v1")
	setupRoute(appContext,v1)

	router.Run()
}