package models

import "gorm.io/gorm"

type Topic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)"`
    Description string `gorm:"column:description;type:varchar(1000)"`
    ProjectID int `gorm:"column:project"`
    Project Project `gorm:"foreignkey:ProjectID"`
    Subtopics []Subtopic `gorm:"foreignkey:TopicID"`
}
