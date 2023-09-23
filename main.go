package main

import (
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
	// toml
	toml := config.InitToml("./config/config.toml")

	// db
	db.InitMySQL(&toml)
	router.MainRouter(app, &toml)
	validator.InitValidate()

	swagger.InitSwagger()

}
