package models

import (
	"gorm.io/gorm"
)

type Note struct {
    gorm.Model
    Content string `gorm:"column:content;type:varchar(500)" json:"content"`
    SubtopicID int `gorm:"column:subtopic" json:"subtopicID"`
}

func CreateNote(content string, stid int, db *gorm.DB) error {
    note := Note{
        Content: content,
        SubtopicID: stid,
    }
    err := db.Create(&note).Error
    return err
}

func GetAllNotes(db *gorm.DB) ([]Note, error) {
    var res []Note
    err := db.Find(&res).Error
    return res, err
}

func GetNoteByID(id int, db *gorm.DB) (Note, error) {
    var res Note
    err := db.First(&res, id).Error
    return res, err
}
