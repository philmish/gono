package models

import "gorm.io/gorm"

type Text struct {
    gorm.Model
    Title string `gorm:"column:title;type:varchar(100)" json:"title"`
    Content string `gorm:"column:content;type:text" json:"content"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
}
