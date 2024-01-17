package models

import "gorm.io/gorm"

type Students struct {
	Name    string
	Age     int
	Languages []Languages `gorm:"many2many:languages_students"`
	gorm.Model
}
type Languages struct {
	Name     string
	Price    int
	Students []Students `gorm:"many2many:languages_students"`
	gorm.Model
}
