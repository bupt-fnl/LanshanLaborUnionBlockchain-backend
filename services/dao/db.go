package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB


func InitDB(){
	var err error
	db, err = gorm.Open(utils.DatabaseSettings.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
																		utils.DatabaseSettings.User,
																		utils.DatabaseSettings.Password,
																		utils.DatabaseSettings.Host,
																		utils.DatabaseSettings.Name))
	if err != nil{
		log.Fatalln("database open failed! "+ err.Error())
	}

	// set table name prefix

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return utils.DatabaseSettings.TablePrefix + defaultTableName
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)



}

func TryInitializeTables(){
	file , err := os.Open("runtime/databases/table.lock")
	if err != nil {
		log.Println(err)
		file, err := os.Create("runtime/databases/table.lock")
		if err != nil{
			log.Fatalln("creating lock failed : "+err.Error())
		}else{
			CreateTables()
		}
		file.Close()
	}
	file.Close()
}


func CreateTables(){

	db.DropTable(&models.User{})
	db.CreateTable(&models.User{})


}

func CloseDB(){
	if db != nil {
		db.Close()
	}
}


func GetExternalDB() *gorm.DB{
	return db
}