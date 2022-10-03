package models

import "gorm.io/gorm"

type Subtopic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)" json:"name"`
    TopicID int `gorm:"column:topic" json:"topicID"`
    Notes []Note `gorm:"foreignKey:SubtopicID;references:ID" json:"notes"`
    Texts []Text `gorm:"foreignKey:SubtopicID;references:ID" json:"texts"`
    Links []Link `gorm:"foreignKey:SubtopicID;references:ID" json:"links"`
    Lists []List `gorm:"foreignKey:SubtopicID;references:ID" json:"lists"`
    Images []Image `gorm:"foreignKey:SubtopicID;references:ID" json:"images"`
}

func CreateSubtopic(name string, toid int, db *gorm.DB) error {
    subtopic := Subtopic{
        Name: name,
        TopicID: toid,
        Notes: []Note{},
        Texts: []Text{},
        Links: []Link{},
        Lists: []List{},
        Images: []Image{},
    }
    err := db.Create(&subtopic).Error
    return err
}

func GetAllSubtopics(db *gorm.DB) ([]Subtopic, error) {
    var res []Subtopic
    err := db.Preload("Notes").Preload("Texts").Preload("Links").Preload("Lists").Preload("Images").Find(&res).Error
    return res, err
}

func SubtopicByID(id int, db *gorm.DB) (Subtopic, error) {
    var res Subtopic
    err := db.Preload("Notes").Preload("Texts").Preload("Links").Preload("Lists").Preload("Images").First(&res, id).Error
    return res, err
}
