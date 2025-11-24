package entities

import (
	"gorm.io/gorm"
)

type Yard struct {
	ID   uint   `json:"ID" gorm:"primarykey"`
	Name string `json:"name" gorm:"unique;not null"`
	gorm.Model
	//Relations
	Block []Block `json:"blocks" gorm:"foreignKey:YardID;"`
}

type YardRequest struct {
	Name string `form:"name" binding:"required"`
}

func (Yard) TableName() string {
	return "yards"
}
