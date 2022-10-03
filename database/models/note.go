package models

import (
	"gorm.io/gorm"
)

type Note struct {
    gorm.Model
    Content string `gorm:"column:content;type:varchar(500)" json:"content"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
}
