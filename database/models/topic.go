package models

import (
	"gorm.io/gorm"
)

type Topic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)" json:"name"`
    Description string `gorm:"column:description;type:varchar(1000)" json:"description"`
    ProjectID int `gorm:"column:project" json:"projectID"`
    Subtopics []Subtopic `gorm:"foreignKey:TopicID;references:ID" json:"subtopics"`
}

func CreateTopic(name, desc string, prid int, db *gorm.DB) error {
    topic := Topic{
        Name: name,
        Description: desc,
        ProjectID: prid,
        Subtopics: []Subtopic{},
    }
    err := db.Create(&topic).Error
    return err
}

func GetAllTopics(db *gorm.DB) ([]Topic, error) {
    var res []Topic
    err := db.Preload("Subtopics").Find(&res).Error
    return res, err
}

func TopicByID(id int, db *gorm.DB) (Topic, error) {
    var res Topic
    err := db.Preload("Subtopics").First(&res, id).Error
    return res, err
}
