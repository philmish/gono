package models

import "gorm.io/gorm"

type Project struct {
    gorm.Model
    Name string `gorm:"column:name;type:varchar(100)"`
    Topics []Topic `gorm:"foreignkey:ProjectID"`
}

func CreateProject(name string, db *gorm.DB) error {
    project := Project{
        Name: name,
        Topics: []Topic{},
    }
    err := db.Create(&project).Error
    return err
}
