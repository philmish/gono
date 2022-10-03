package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
    mode := os.Getenv("GONO_DBMODE")
    dsn := os.Getenv("GONO_DBDSN")
    switch mode {
    case "dev":
       return gorm.Open(sqlite.Open(dsn), &gorm.Config{}) 
    default:
        return nil, fmt.Errorf("Unknown DB Address requested: %s\n", dsn)
    }
}
