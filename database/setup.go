package database

import (
    "gorm.io/gorm"
)

func makeMigration(db *gorm.DB) error {
    //TODO Create migration func
    return nil
}

func InitDB(conn gorm.Dialector, migrate bool) error {
    db, err := gorm.Open(conn, &gorm.Config{})
    if err != nil {
        return err
    }
    if migrate {
        err = makeMigration(db)
        if err != nil {
            return err
        }
    }
    return nil
}
