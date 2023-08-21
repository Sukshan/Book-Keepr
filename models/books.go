package models

import (
	"log"

	"gorm.io/gorm"
)

type Books struct {
	ID        uint    `gorm:"primary key; autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	log.Printf("MigateBooks function ran successfully")
	return err
}
