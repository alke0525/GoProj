package router

import (
	"webService/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	cfgRouting := router.Group("/cfg")
	{
		cfgRouting.GET("/getCfg", api.GetCfg)
		cfgRouting.POST("/chgCfg", api.ChgCfg)
	}

	dataRouting := router.Group("/data")
	{
		dataRouting.GET("/getData", api.GetData)
		dataRouting.POST("/chgData", api.ChgData)
	}

	return router
}
