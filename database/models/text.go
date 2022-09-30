package models

import "gorm.io/gorm"

type Text struct {
    gorm.Model
    Title string `gorm:"column:title;type:varchar(100)"`
    Content string `gorm:"column:content;type:text"`
    SubtopicID int `gorm:"column:subtopic"`
}
