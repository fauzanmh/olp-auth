package init

import "github.com/fauzanmh/olp-auth/docs"

func setupSwagger() {
	docs.SwaggerInfo.Title = "Online Learning Platform API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8101"
	docs.SwaggerInfo.BasePath = "/api"
}
