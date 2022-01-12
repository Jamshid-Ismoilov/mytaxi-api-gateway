package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/api/docs" // swag
	v1 "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/api/handlers/v1"
	"github.com/Jamshid-Ismoilov/mytaxi-api-gateway/config"
	"github.com/Jamshid-Ismoilov/mytaxi-api-gateway/pkg/logger"
	"github.com/Jamshid-Ismoilov/mytaxi-api-gateway/services"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/drivers", handlerV1.CreateDriver)
	api.GET("/drivers/:id", handlerV1.GetDriver)
	api.PUT("/drivers/:id", handlerV1.UpdateDriver)
	api.DELETE("/drivers/:id", handlerV1.DeleteDriver)

	api.POST("/clients", handlerV1.CreateClient)
	api.GET("/clients/:id", handlerV1.GetClient)
	api.PUT("/clients/:id", handlerV1.UpdateClient)
	api.DELETE("/clients/:id", handlerV1.DeleteClient)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
