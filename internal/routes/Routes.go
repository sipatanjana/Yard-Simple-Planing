package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"YardPlaning/internal/conf"

	_ "YardPlaning/docs"
)

func InitRoutes(routes *gin.Engine, c *conf.Container) {
	//autogenerate :routes
	//Yard route
	routes.GET("/yard", c.YardHandler.GetAll)
	routes.GET("/yard/:id", c.YardHandler.GetById)
	routes.POST("/yard", c.YardHandler.Insert)
	routes.PUT("/yard/:id", c.YardHandler.Update)
	// routes.PATCH("/yard/:id/:enable", c.YardHandler.SetEnable)
	// routes.DELETE("/yard/:id", c.YardHandler.Delete)

	//Block route
	routes.GET("/block", c.BlockHandler.GetAll)
	routes.GET("/block/:id", c.BlockHandler.GetById)
	routes.POST("/block", c.BlockHandler.Insert)
	routes.PUT("/block/:id", c.BlockHandler.Update)
	// routes.PATCH("/block/:id/:enable", c.BlockHandler.SetEnable)
	// routes.DELETE("/block/:id", c.BlockHandler.Delete)

	//YardPlan route
	routes.GET("/yardPlan", c.YardPlanHandler.GetAll)
	routes.GET("/yardPlan/:id", c.YardPlanHandler.GetById)
	routes.POST("/yardPlan", c.YardPlanHandler.Insert)
	routes.PUT("/yardPlan/:id", c.YardPlanHandler.Update)
	// routes.PATCH("/yardPlan/:id/:enable", c.YardPlanHandler.SetEnable)
	// routes.DELETE("/yardPlan/:id", c.YardPlanHandler.Delete)

	//Stacking route
	routes.POST("/suggestion", c.StackingHandler.Suggest)
	routes.POST("/placement", c.StackingHandler.Placement)
	routes.POST("/pickup", c.StackingHandler.Pickup)

	//endautogenerate :routes
	routes.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
