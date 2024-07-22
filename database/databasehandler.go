package database

//import gorm here and connect to mysql
import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tcstorego/model"

)

var DBCon *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3308)/tcschema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	log.Println("connected")
	db.AutoMigrate(&model.Testcase{})
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
