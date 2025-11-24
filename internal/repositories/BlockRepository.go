package repositories

import (
	"gorm.io/gorm"

	"YardPlaning/internal/entities"
	"YardPlaning/pkg/utils"
)

type BlockRepository struct {
	db *gorm.DB
}

func NewBlockRepository(db *gorm.DB) *BlockRepository {
	return &BlockRepository{db}
}

func (r *BlockRepository) GetAll(p *utils.Pagination, blocks *[]entities.Block) error {
	//var blocks []entities.Block
	query := r.db.Model(&entities.Block{}).Preload("Yard")

	// Filter per kolom
	if len(p.Filters) > 0 {
		cond, args := utils.BuildFilterCondition(p.Filters)
		query = query.Where(cond, args...)
	}

	// Count total
	if err := query.Count(&p.Total).Error; err != nil {
		return err
	}

	//total page
	p.TotalPage = (p.Total + int64(p.Limit) - 1) / int64(p.Limit)

	return p.Paginate(query).Find(blocks).Error
}

func (r *BlockRepository) GetByID(block *entities.Block) error {
	result := r.db.Preload("Yard").First(&block)
	return result.Error
}

func (r *BlockRepository) Insert(block *entities.Block) error {
	result := r.db.Create(block)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *BlockRepository) Update(block *entities.Block) error {
	result := r.db.Model(block).Updates(block)

	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *BlockRepository) Disable(block *entities.Block) error {

	result := r.db.Delete(&block)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *BlockRepository) Enable(block *entities.Block) error {

	result := r.db.Unscoped().Model(&block).Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *BlockRepository) Delete(block *entities.Block) error {
	result := r.db.Unscoped().Delete(&block)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errorNoData
	}

	return nil
}
