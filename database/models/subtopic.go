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
