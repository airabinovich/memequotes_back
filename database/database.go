package database

import (
	"context"
	"fmt"
	"github.com/airabinovich/memequotes_back/config"
	commonContext "github.com/airabinovich/memequotes_back/context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

//Initialize connects to the database al fills the global variable utils.DB
func Initialize() error {
	ctx := commonContext.AppContext(context.Background())
	logger := commonContext.Logger(ctx)
	username := config.Credentials.GetString("db.user", "root")
	password := config.Credentials.GetString("db.password", "")
	host := config.Credentials.GetString("db.host", "localhost")
	port := config.Credentials.GetInt32("db.post", 3306)
	var err error
	DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/memequotes?charset=utf8mb4&parseTime=True&loc=UTC", username, password, host, port))
	if err != nil {
		logger.Error("opening DB", err)
		return err
	}
	logger.Info(fmt.Sprintf("Connected to DB %s:%d", host, port))
	return nil
}

// Close the connection to the DB
func Close() error {
	ctx := commonContext.AppContext(context.Background())
	logger := commonContext.Logger(ctx)
	if err := DB.Close(); err != nil {
		logger.Error("closing the database connection", err)
		return err
	}
	logger.Info("Connection to the DB closed")
	return nil
}
