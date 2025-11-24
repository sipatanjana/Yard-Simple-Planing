package repositories

import (
	"gorm.io/gorm"

	"YardPlaning/internal/entities"
	"YardPlaning/pkg/utils"
)

type YardRepository struct {
	db *gorm.DB
}

func NewYardRepository(db *gorm.DB) *YardRepository {
	return &YardRepository{db}
}

func (r *YardRepository) GetAll(p *utils.Pagination, yards *[]entities.Yard) error {
	//var yards []entities.Yard
	query := r.db.Model(&entities.Yard{})

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

	return p.Paginate(query).Find(yards).Error
}

func (r *YardRepository) GetByID(yard *entities.Yard) error {
	result := r.db.First(&yard)
	return result.Error
}

func (r *YardRepository) Insert(yard *entities.Yard) error {
	result := r.db.Create(yard)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardRepository) Update(yard *entities.Yard) error {
	result := r.db.Model(yard).Updates(yard)

	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardRepository) Disable(yard *entities.Yard) error {

	result := r.db.Delete(&yard)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardRepository) Enable(yard *entities.Yard) error {

	result := r.db.Unscoped().Model(&yard).Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardRepository) Delete(yard *entities.Yard) error {
	result := r.db.Unscoped().Delete(&yard)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errorNoData
	}

	return nil
}
