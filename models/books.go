package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `gorm:"primary key; autoIncrement" json:"id"`
	Author    *string `json:"author"`
	title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) err {
	err := db.AutoMigrate(&Books{})
	return err
}
