package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // GORM driver
)

var db *gorm.DB

var tables = []interface{}{
	&User{},
	&Level{},
	&Lesson{},
	&Permission{},
	&Course{},
}

func Connect() {
	var DB_USER string = os.Getenv("MARIADB_USER_DIALGORITHM")
	var DB_PASS string = os.Getenv("MARIADB_PASS")
	var DB_NAME string = os.Getenv("MARIADB_DIALGORITHM")
	var DB_HOST string = os.Getenv("MARIADB_HOST_DIALGORITHM")
	var DB_PORT string = os.Getenv("MARIADB_PORT")
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	var i int
	for {
		var err error
		if i >= 30 {
			panic("could not connect to " + source)
		}
		time.Sleep(3 * time.Second)
		db, err = gorm.Open("mysql", source)
		if err != nil {
			fmt.Println("Retrying connection...", err)
			i++
			continue
		}
		fmt.Println("YA SE CONECTO", DB_HOST)
		db.DB().SetMaxIdleConns(0)
		db.DB().SetConnMaxLifetime(time.Second * 14400)
		break
	}
}

func Migrate() {
	for _, t := range tables {
		db.AutoMigrate(t)
	}
}

func Create(value interface{}) *gorm.DB {
	return db.Create(value)
}

func First(out interface{}, where ...interface{}) *gorm.DB {
	return db.First(out, where...)
}
