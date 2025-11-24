package services

import (
	//autogenerate :import
	//endautogenerate :import
	"fmt"

	"YardPlaning/internal/entities"
	"YardPlaning/internal/repositories"
	"YardPlaning/pkg/utils"
)

type YardPlanService struct {
	//autogenerate :serviceStruct
	repo *repositories.YardPlanRepository
	//endautogenerate :serviceStruct
}

// autogenerate :serviceNewStruct
func NewYardPlanService(repo *repositories.YardPlanRepository) *YardPlanService {
	return &YardPlanService{repo}
}

//endautogenerate :serviceNewStruct

func (s *YardPlanService) GetAll(p *utils.Pagination, yardPlans *[]entities.YardPlan) error {
	//autogenerate :servicefunc
	return s.repo.GetAll(p, yardPlans)
	//endautogenerate :servicefunc
}

func (s *YardPlanService) GetBLockByID(block *entities.Block) error {
	//autogenerate :servicefunc
	return s.repo.GetBLockByID(block)
	//endautogenerate :servicefunc

}

func (s *YardPlanService) GetById(yardPlan *entities.YardPlan) error {
	//autogenerate :servicefunc
	return s.repo.GetByID(yardPlan)
	//endautogenerate :servicefunc

}

func (s *YardPlanService) Insert(yardPlan *entities.YardPlan) error {
	//autogenerate :servicefunc
	switch yardPlan.ContainerHeight {
	case 8.6:
		return s.repo.Insert(yardPlan)
	case 9.6:
		return s.repo.Insert(yardPlan)
	default:
		return fmt.Errorf("salah type ukuran")
	}
	//endautogenerate :servicefunc
}

func (s *YardPlanService) Update(yardPlan *entities.YardPlan) error {
	//autogenerate :servicefunc
	switch yardPlan.ContainerHeight {
	case 8.6:
		return s.repo.Update(yardPlan)
	case 9.6:
		return s.repo.Update(yardPlan)
	default:
		return fmt.Errorf("salah type ukuran")
	}
	//endautogenerate :servicefunc
}

func (s *YardPlanService) SetEnable(yardPlan *entities.YardPlan, enable bool) error {
	//autogenerate :servicefunc
	if enable {
		return s.repo.Enable(yardPlan)
	}
	return s.repo.Disable(yardPlan)
	//endautogenerate :servicefunc
}

func (s *YardPlanService) Delete(yardPlan *entities.YardPlan) error {
	//autogenerate :servicefunc
	return s.repo.Delete(yardPlan)
	//endautogenerate :servicefunc
}
