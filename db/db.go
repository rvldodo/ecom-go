package db

import (
	"log"

	mysqlId "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/dodo/ecom/config"
)

var DB *gorm.DB

func init() {
	db, err := newMySQLStorage(mysqlId.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	DB = db
}

func newMySQLStorage(cfg mysqlId.Config) (*gorm.DB, error) {
	cfg.FormatDSN()
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: cfg.FormatDSN()}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db, nil
}

func initStorage(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Succesfully connected!")
}
