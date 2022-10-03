package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/philmish/gono/api"
	"github.com/philmish/gono/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
    env *string
    addr *string
    migrate *bool
)

func init() {
    env = flag.String("env", "dev", "current running mode (dev,prod)")
    addr = flag.String("uri", ":7195", "address the server is running on")
    migrate = flag.Bool("m", true, "Make auto migrations for db")
}

func setupEnv(env string) error {
    switch env {
    case "dev":
        log.Println("Setting up dev env ...")
        err := os.Setenv("GONO_DBMODE", "dev")
        if err != nil {
            return err
        }
        err = os.Setenv("GONO_DBDSN", "test.db")
        if err != nil {
            return err
        }
        db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
        if err != nil {
            return err
        }
        database.InitDB(db, *migrate)
        log.Println("Database setup successfull.")
        return nil
    case "prod":
        return fmt.Errorf("No env setup for prod implemented.")
    default:
        return fmt.Errorf("Unknown env type: %s\n", env)
    }
}

func main() {
    flag.Parse()
    fmt.Println("Env: ", *env)
    fmt.Println("Addr: ", *addr)
    err := setupEnv(*env)
    if err != nil {
        log.Fatalf("%s\n", err.Error())
    }
    log.Printf("Starting API server on %s\n", *addr)
    api.Serve(*addr, *env)
}
