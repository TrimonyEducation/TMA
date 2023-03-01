package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DB *gorm.DB
var err error


func ConnectDB(){
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold:              time.Second,   // Slow SQL threshold
      LogLevel:                   logger.Silent, // Log level
      IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
      Colorful:                  true,          // enable color
    },
  )
  errx:=godotenv.Load()
	if errx!= nil {
        log.Println(err.Error())
	}
    dsn:= os.Getenv("DSN")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
  if err != nil {
    panic("failed to connect database")
  }
  log.Println("connected to Database")

}