package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	
)

var (
	db *sql.DB
	once sync.Once
)

func ConnectDB() *sql.DB{
	once.Do(func() {
		LoadEnv()
		driver	:= GetEnv("DB_DRIVER", "mysql")
		host		:= GetEnv("DB_HOST", "localhost")
		port		:= GetEnv("DB_PORT","3306")
		user		:= GetEnv("DB_USER", "root")
		password	:= GetEnv("DB_PASSWORD", "")
		dbname	:= GetEnv("DB_NAME", "menu_db")

		var dsn string

		if driver == "mysql"{
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
		} else {
			log.Fatalf("Driver %s tidak didukung", driver)
		}


		var err error
		db, err = sql.Open(driver, dsn)
		if err != nil{
			log.Fatalf("Gagal terhubung ke database", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Database tidak merespons: %v", err)
		}

		log.Println("Database berhasil terhubung")

	})	
	return db
}

func CLoseDB(){
	if db != nil{
		_ = db.Close()
		log.Println("Database Disconnect")
	}
}