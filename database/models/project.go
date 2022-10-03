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

func GetAll(db *gorm.DB) ([]Project, error) {
    var res []Project
    err := db.Find(&res).Error
    return res, err
}
