package database_test

import (
	"testing"

	"github.com/philmish/gono/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSetup(t *testing.T) {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    err = database.InitDB(db, true)
    if err != nil {
        t.Errorf("Error initializing: %s\n", err.Error())
    }
}
