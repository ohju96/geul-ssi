package swagger

import (
	"geulSsi/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwagger() {
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	docs.SwaggerInfo.Title = "Web Geul-Ssi Swagger !"
	docs.SwaggerInfo.Description = "🥸Hi, There! This is geul-ssi server."
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
}
