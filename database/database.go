package database

import (
	"fmt"
	"github.com/airabinovich/memequotes_back/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

//Initialize connects to the database al fills the global variable utils.DB
func Initialize() error {
	username := config.Credentials.GetString("db.user", "root")
	password := config.Credentials.GetString("db.password", "")
	//host := config.Credentials.GetString("db.host", "")
	//port := config.Credentials.GetInt32("db.post", 3306)
	var err error
	DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@/memequotes?charset=utf8mb4&parseTime=True&loc=UTC", username, password))
	if err != nil {
		fmt.Println("ERROR: opening DB", err)
		return err
	}
	return nil
}

func Close() error {
	if err := DB.Close(); err != nil {
		fmt.Println("ERROR: closing the database connection")
		return err
	}
	fmt.Println("Connection to the DB closed")
	return nil
}
