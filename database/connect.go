package database

import (
	"fmt"
	"strconv"

	"github.com/Rolas444/goapi.git/config"
	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsing port")
	}

	// dns := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s")
	// dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"

	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
	// 	config.Config("DB_USER"), config.Config("DB_PASSWORD"),
	// 	config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Lima",
		config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	fmt.Println(dsn)

	// DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Connection OK!")
}
