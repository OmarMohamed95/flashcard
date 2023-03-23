package model

import (
	"flashcards-api/app/database"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Frontside string `gorm:"" json:"frontside"`
	Backside  string `gorm:"" json:"backside"`
	SetID     int    `gorm:"" json:"setId"`
	Set       Set    `gorm:"foreignKey:SetID" json:"set"`
}

func init() {
	database.DB.AutoMigrate(&Card{})
}
