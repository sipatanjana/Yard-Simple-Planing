package services

import (
	//autogenerate :import
	//endautogenerate :import

	"YardPlaning/internal/entities"
	"YardPlaning/internal/repositories"
	"YardPlaning/pkg/utils"
)

type YardService struct {
	//autogenerate :serviceStruct
	repo *repositories.YardRepository
	//endautogenerate :serviceStruct
}

// autogenerate :serviceNewStruct
func NewYardService(repo *repositories.YardRepository) *YardService {
	return &YardService{repo}
}

//endautogenerate :serviceNewStruct

func (s *YardService) GetAll(p *utils.Pagination, yards *[]entities.Yard) error {
	//autogenerate :servicefunc
	return s.repo.GetAll(p, yards)
	//endautogenerate :servicefunc
}

func (s *YardService) GetById(yard *entities.Yard) error {
	//autogenerate :servicefunc
	return s.repo.GetByID(yard)
	//endautogenerate :servicefunc

}

func (s *YardService) Insert(yard *entities.Yard) error {
	//autogenerate :servicefunc
	return s.repo.Insert(yard)
	//endautogenerate :servicefunc
}

func (s *YardService) Update(yard *entities.Yard) error {
	//autogenerate :servicefunc
	return s.repo.Update(yard)
	//endautogenerate :servicefunc
}

func (s *YardService) SetEnable(yard *entities.Yard, enable bool) error {
	//autogenerate :servicefunc
	if enable {
		return s.repo.Enable(yard)
	}
	return s.repo.Disable(yard)
	//endautogenerate :servicefunc
}

func (s *YardService) Delete(yard *entities.Yard) error {
	//autogenerate :servicefunc
	return s.repo.Delete(yard)
	//endautogenerate :servicefunc
}
