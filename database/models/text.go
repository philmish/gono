package models

import "gorm.io/gorm"

type Text struct {
    gorm.Model
    Title string `gorm:"column:title;type:varchar(100)" json:"title"`
    Content string `gorm:"column:content;type:text" json:"content"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
}

func CreateText(title, content string, stid int, db *gorm.DB) error {
    text := Text{
        Title: title,
        Content: content,
        SubtopicID: stid,
    }
    err := db.Create(&text).Error
    return err
}

func GetAllTexts(db *gorm.DB) ([]Text, error) {
    var res []Text
    err := db.Find(&res).Error
    return res, err
}

func GetTextByID(id int, db *gorm.DB) (Text, error) {
    var res Text
    err := db.First(&res, id).Error
    return res, err
}
