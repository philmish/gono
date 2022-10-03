package models

import "gorm.io/gorm"

type ListItem struct {
    gorm.Model
    Value string `gorm:"column:value;type:varchar(150)" json:"value"`
    ListID int `gorm:"column:list" json:"listID"`
}

type List struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)" json:"name"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
    Items []ListItem `gorm:"foreignkey:ListID" json:"items"`
}
