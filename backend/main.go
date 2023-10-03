package main

import (
	"geulSsi/app/starter"
	"geulSsi/config"
	"geulSsi/config/db"
	"geulSsi/config/swagger"
	"geulSsi/router"
	"geulSsi/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	Init(app)

	app.Run()
}

func Init(app *gin.Engine) {

	// 설정파일
	toml := config.InitToml("./config/config.toml")

	// 데이터베이스
	db.InitMySQL(&toml)

	// 라우터 및 검증
	router.MainRouter(app, &toml)
	validator.InitValidate()

	// 채널 스타터
	starter.WiseSayingStarter()

	// 스웨거
	swagger.InitSwagger()

}
