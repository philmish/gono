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

func CreateSubtopic(name string, toid int, db *gorm.DB) error {
    var topic Topic
    err := db.First(&topic, toid).Error
    if err != nil {
        return err
    }
    subtopic := Subtopic{
        Name: name,
        TopicID: int(topic.ID),
        Notes: []Note{},
        Texts: []Text{},
        Links: []Link{},
        Lists: []List{},
        Images: []Image{},
    }
    err = db.Create(&subtopic).Error
    return err
}
