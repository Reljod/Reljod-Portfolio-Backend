package swagger

import (
	"github.com/Reljod/Reljod-Portfolio-Backend/config"
	// swagger embed files
	// gin-swagger middleware
	"github.com/Reljod/Reljod-Portfolio-Backend/docs"
)

func Setup() {
	docs.SwaggerInfo.Title = config.Title
	docs.SwaggerInfo.Description = config.Description
	docs.SwaggerInfo.Version = config.Version
	docs.SwaggerInfo.Host = config.Host
	docs.SwaggerInfo.BasePath = config.BasePath
}
