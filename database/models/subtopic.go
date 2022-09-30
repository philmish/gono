package models

import "gorm.io/gorm"

type Subtopic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)"`
    TopicID int `gorm:"column:topic"`
    Notes []Note `gorm:"foreignKey:SubtopicID;references:ID"`
    Texts []Text `gorm:"foreignKey:SubtopicID;references:ID"`
    Links []Link `gorm:"foreignKey:SubtopicID;references:ID"`
    Lists []List `gorm:"foreignKey:SubtopicID;references:ID"`
    Images []Image `gorm:"foreignKey:SubtopicID;references:ID"`
}
