package entities

import (
	"gorm.io/gorm"
)

type Block struct {
	ID     uint   `json:"ID" gorm:"primarykey"`
	Name   string `json:"name" gorm:"not null"`
	Code   string `json:"code"`
	YardID uint   `json:"yard_id"`
	Slot   int    `json:"slot"`
	Row    int    `json:"row"`
	Tier   int    `json:"tier"`
	//Relations
	Yard Yard `json:"yard" gorm:"foreignKey:YardID;"`

	gorm.Model
}

type BlockRequest struct {
	Name string `form:"name" binding:"required"`
	Code string `form:"code" binding:"required"`
	Yard uint   `form:"yard" binding:"required"`
	Slot int    `form:"slot" binding:"required,gt=0"`
	Row  int    `form:"row" binding:"required,gt=0"`
	Tier int    `form:"tier" binding:"required,gt=0"`
}

func (Block) TableName() string {
	return "blocks"
}
