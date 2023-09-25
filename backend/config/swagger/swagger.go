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
	docs.SwaggerInfo.Description = "🥸안녕하세요! [ 글씨 ] 입니다. \n" +
		"글씨에서는 날씨와 시간을 배경으로 짧은 문장을 적어 공유하는 서비스입니다. \n" +
		"이 서비스는 [ 노션 ](https://oh-app.notion.site/Geul-Ssi-fba867acf3584ec29a3d2dd059129683?pvs=4) 과 [ 깃허브 ](https://github.com/ohju96/geul-ssi) 에 공유되어 있습니다."
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
}
