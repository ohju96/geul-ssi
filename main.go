package main

import (
	"geulSsi/config"
	"geulSsi/config/db"
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
}
