package database

import (
    "gorm.io/gorm"
    "github.com/philmish/gono/database/models"
)

func makeMigration(db *gorm.DB) error {
    err := db.AutoMigrate(&models.Project{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.Topic{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.Subtopic{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.Note{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.Text{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.ListItem{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.List{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.Image{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.Link{})
    if err != nil {
        return err
    }
    return nil
}

func InitDB(conn *gorm.DB, migrate bool) error {
    if migrate {
        err := makeMigration(conn)
        if err != nil {
            return err
        }
    }
    return nil
}
