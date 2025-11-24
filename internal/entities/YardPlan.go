package entities

import (
	"gorm.io/gorm"
)

type YardPlan struct {
	ID              uint    `json:"ID" gorm:"primaryKey"`
	YardID          uint    `json:"yard_id"`
	BlockID         uint    `json:"block_id"`
	ContainerSize   int     `json:"container_size"`
	ContainerHeight float32 `gorm:"type:decimal" json:"container_height"`
	ContainerType   string  `json:"container_type"`
	SlotFrom        int     `json:"slot_from"`
	SlotTo          int     `json:"slot_to"`
	RowFrom         int     `json:"row_from"`
	RowTo           int     `json:"row_to"`
	TierFrom        int     `json:"tier_from"`
	TierTo          int     `json:"tier_to"`
	Active          bool    `json:"active" gorm:"default:true"`
	gorm.Model

	//Relations
	Yard  Yard  `json:"yard" gorm:"foreignKey:YardID;"`
	Block Block `json:"block" gorm:"foreignKey:BlockID;"`
}

type YardPlanRequest struct {
	Block           uint    `form:"block" binding:"required"`
	ContainerSize   int     `form:"container_size" binding:"required,oneof=20 40"`
	ContainerHeight float32 `form:"container_height" binding:"required"`
	ContainerType   string  `form:"container_type" binding:"required"`
	SlotFrom        int     `form:"slot_from" binding:"required"`
	SlotTo          int     `form:"slot_to" binding:"required"`
	RowFrom         int     `form:"row_from" binding:"required"`
	RowTo           int     `form:"row_to" binding:"required"`
	TierFrom        int     `form:"tier_from" binding:"required"`
	TierTo          int     `form:"tier_to" binding:"required"`
}

func (YardPlan) TableName() string {
	return "yard_plans"
}
