package models

import "gorm.io/gorm"

type Image struct {
    gorm.Model
    Alt string `gorm:"column:alt;type:varchar(75)"`
    Uri string `gorm:"column:uri;type:varchar(256)"`
    SubtopicID int `gorm:"column:subtopic"`
}
