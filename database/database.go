package database

//import gorm here and connect to mysql
import (
	"log"
	"os"

	"tcstorego/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBCon *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env \n", err)
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	log.Println("connected")
	if err := db.AutoMigrate(&model.Testcase{}); err != nil {
		log.Fatal("Gorm autoMigrate:", err)
	}
	DBCon = db
}

/*

func SearchAll(){

}

func SearchCondition(){

}

func InsertDatatc(){

}

func UpdateDatatc(){

}

func DeleteDatatc(){

}
*/
