package models

import "gorm.io/gorm"

type Link struct {
    gorm.Model
    Title string `gorm:"column:title;type:varchar(50)" json:"title"`
    Url string `gorm:"column:url;type:varchar(500)" json:"url"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
}
