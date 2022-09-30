package models

import "gorm.io/gorm"

type Topic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)"`
    Description string `gorm:"column:description;type:varchar(1000)"`
    ProjectID int `gorm:"column:project"`
    Project Project `gorm:"foreignkKey:ProjectID"`
    Subtopics []Subtopic `gorm:"foreignKey:TopicID;references:ID"`
}

func CreateTopic(name, desc string, prid int, db *gorm.DB) error {
    var project Project
    err := db.First(&project, prid).Error
    if err != nil {
        return err
    }
    topic := Topic{
        Name: name,
        Description: desc,
        ProjectID: int(project.ID),
        Project: project,
        Subtopics: []Subtopic{},
    }
    err = db.Create(&topic).Error
    return err
}
