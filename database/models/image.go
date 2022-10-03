package models

import "gorm.io/gorm"

type Image struct {
    gorm.Model
    Alt string `gorm:"column:alt;type:varchar(75)" json:"alt"`
    Uri string `gorm:"column:uri;type:varchar(256)" json:"uri"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
}
