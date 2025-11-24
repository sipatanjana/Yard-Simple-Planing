package repositories

import (
	"gorm.io/gorm"

	"YardPlaning/internal/entities"
	"YardPlaning/pkg/utils"
)

type YardPlanRepository struct {
	db *gorm.DB
}

func NewYardPlanRepository(db *gorm.DB) *YardPlanRepository {
	return &YardPlanRepository{db}
}

func (r *YardPlanRepository) GetAll(p *utils.Pagination, yardPlans *[]entities.YardPlan) error {
	//var yardPlans []entities.YardPlan
	query := r.db.Model(&entities.YardPlan{}).Preload("Block").Preload("Yard")

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

	return p.Paginate(query).Find(yardPlans).Error
}

func (r *YardPlanRepository) GetBLockByID(block *entities.Block) error {
	result := r.db.First(&block).Preload("Block").Preload("Yard")
	return result.Error
}

func (r *YardPlanRepository) GetByID(yardPlan *entities.YardPlan) error {
	result := r.db.First(&yardPlan)
	return result.Error
}

func (r *YardPlanRepository) Insert(yardPlan *entities.YardPlan) error {
	result := r.db.Create(yardPlan)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardPlanRepository) Update(yardPlan *entities.YardPlan) error {
	result := r.db.Model(yardPlan).Updates(yardPlan)

	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardPlanRepository) Disable(yardPlan *entities.YardPlan) error {

	result := r.db.Delete(&yardPlan)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardPlanRepository) Enable(yardPlan *entities.YardPlan) error {

	result := r.db.Unscoped().Model(&yardPlan).Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return errorNoData
	}
	return result.Error
}

func (r *YardPlanRepository) Delete(yardPlan *entities.YardPlan) error {
	result := r.db.Unscoped().Delete(&yardPlan)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errorNoData
	}

	return nil
}
