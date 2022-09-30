package models

import "gorm.io/gorm"

type Subtopic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)"`
    TopicID int `gorm:"column:topic"`
    Topic Topic `gorm:"foreignkey:TopicID"`
    Notes []Note `gorm:"foreignkey:SubtopicID"`
    Texts []Text `gorm:"foreignkey:SubtopicID"`
    Links []Link `gorm:"foreignkey:SubtopicID"`
    Lists []List `gorm:"foreignkey:SubtopicID"`
    Images []Image `gorm:"foreignkey:SubtopicID"`
}
