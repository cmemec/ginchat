package resource

import (
	"fmt"
	"ginchat/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := viper.GetString("mysql.dns")
	// 连接MySQL数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v\n", err)
	}
	// 自动迁移（创建表）
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		return fmt.Errorf("Failed to migrate database: %v\n", err)
	}
	DB = db
	return nil
}
