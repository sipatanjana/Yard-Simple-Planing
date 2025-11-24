package conf

import (
	//autogenerate :import
	//endautogenerate :import

	"gorm.io/gorm"

	"YardPlaning/internal/handlers"
	"YardPlaning/internal/repositories"
	"YardPlaning/internal/services"
)

type Container struct {
	//autogenerate :handlersDefine
	YardHandler     *handlers.YardHandler     //Yard define
	BlockHandler    *handlers.BlockHandler    //Block define
	YardPlanHandler *handlers.YardPlanHandler //YardPlan define
	StackingHandler *handlers.StackingHandler //Stacking define
	//endautogenerate :handlersDefine

	// Tambahkan handler kamu yang lainnya di bawah sini.
}

func NewContainer(db *gorm.DB) *Container {
	//autogenerate :handlersDepRepository
	yardRepo := repositories.NewYardRepository(db)         //Yard dependency
	blockRepo := repositories.NewBlockRepository(db)       //Block dependency
	yardPlanRepo := repositories.NewYardPlanRepository(db) //YardPlan dependency
	//endautogenerate :handlersDepRepository

	//autogenerate :handlersDepService
	yardService := services.NewYardService(yardRepo)             //Yard dependency
	blockService := services.NewBlockService(blockRepo)          //Block dependency
	yardPlanService := services.NewYardPlanService(yardPlanRepo) //YardPlan dependency
	//endautogenerate :handlersDepService

	//autogenerate :handlersDepHandler
	yardHandler := handlers.NewYardHandler(yardService)             //Yard dependency
	blockHandler := handlers.NewBlockHandler(blockService)          //Block dependency
	yardPlanHandler := handlers.NewYardPlanHandler(yardPlanService) //YardPlan dependency
	stackingHandler := handlers.NewStackingHandler(yardPlanService) //Stacking dependency
	//endautogenerate :handlersDepHandler

	// Tambahkan dependency kamu yang lainnya di bawah sini.

	return &Container{
		//autogenerate :handlersInit
		YardHandler:     yardHandler,     //Yard init
		BlockHandler:    blockHandler,    //Block init
		YardPlanHandler: yardPlanHandler, //YardPlan init
		StackingHandler: stackingHandler, //Stacking init
		//endautogenerate :handlersInit

		// Tambahkan handler kamu lainnya di bawah sini...
	}
}
