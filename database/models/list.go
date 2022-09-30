package models

import "gorm.io/gorm"

type ListItem struct {
    gorm.Model
    Value string `gorm:"column:value;type:varchar(150)"`
    ListID int `gorm:"column:list"`
}

type List struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)"`
    SubtopicID int `gorm:"column:subtopic"`
    Items []ListItem `gorm:"foreignkey:ListID"`
}
