package models

import "gorm.io/gorm"

type Project struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)" json:"name"`
    Topics []Topic `gorm:"foreignkey:ProjectID" json:"topics"`
}

func CreateProject(name string, db *gorm.DB) error {
    project := Project{
        Name: name,
        Topics: []Topic{},
    }
    err := db.Create(&project).Error
    return err
}

func GetAllProjects(db *gorm.DB) ([]Project, error) {
    var res []Project
    err := db.Preload("Topics").Find(&res).Error
    return res, err
}

func ProjectByID(id int, db *gorm.DB) (Project, error) {
    var res Project
    err := db.First(&res, id).Error
    return res, err
}
