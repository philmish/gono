package models

import "gorm.io/gorm"

type Topic struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)" json:"name"`
    Description string `gorm:"column:description;type:varchar(1000)" json:"description"`
    ProjectID int `gorm:"column:project" json:"projectID"`
    Project Project `gorm:"foreignkKey:ProjectID" json:"project"`
    Subtopics []Subtopic `gorm:"foreignKey:TopicID;references:ID" json:"subtopics"`
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
