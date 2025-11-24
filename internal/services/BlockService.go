package services

import (
	//autogenerate :import
	//endautogenerate :import

	"YardPlaning/internal/entities"
	"YardPlaning/internal/repositories"
	"YardPlaning/pkg/utils"
)

type BlockService struct {
	//autogenerate :serviceStruct
	repo *repositories.BlockRepository
	//endautogenerate :serviceStruct
}

// autogenerate :serviceNewStruct
func NewBlockService(repo *repositories.BlockRepository) *BlockService {
	return &BlockService{repo}
}

//endautogenerate :serviceNewStruct

func (s *BlockService) GetAll(p *utils.Pagination, blocks *[]entities.Block) error {
	//autogenerate :servicefunc
	return s.repo.GetAll(p, blocks)
	//endautogenerate :servicefunc
}

func (s *BlockService) GetById(block *entities.Block) error {
	//autogenerate :servicefunc
	return s.repo.GetByID(block)
	//endautogenerate :servicefunc

}

func (s *BlockService) Insert(block *entities.Block) error {
	//autogenerate :servicefunc
	return s.repo.Insert(block)
	//endautogenerate :servicefunc
}

func (s *BlockService) Update(block *entities.Block) error {
	//autogenerate :servicefunc
	return s.repo.Update(block)
	//endautogenerate :servicefunc
}

func (s *BlockService) SetEnable(block *entities.Block, enable bool) error {
	//autogenerate :servicefunc
	if enable {
		return s.repo.Enable(block)
	}
	return s.repo.Disable(block)
	//endautogenerate :servicefunc
}

func (s *BlockService) Delete(block *entities.Block) error {
	//autogenerate :servicefunc
	return s.repo.Delete(block)
	//endautogenerate :servicefunc
}
