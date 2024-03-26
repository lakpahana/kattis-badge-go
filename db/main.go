package main

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func main() {
// 	ConnectToDB()
// }

func ConnectToDB() {
	var err error
	dsn := "host=ep-late-dust-a1hrp4wi.ap-southeast-1.aws.neon.tech user=kattis_owner password=cPy64QOJdwHz dbname=kattis port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB error")
	}

}
