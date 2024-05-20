package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	//"root:@tcp(127.0.0.1:3306)/test_go?parseTime=true"
	db_url := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(string(db_url)), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database", db_url)
	}
}
