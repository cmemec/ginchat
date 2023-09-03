package models

import (
	"time"

	"gorm.io/gorm"
)

// user_basic
type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string
}

// 表名
func (u *UserBasic) Tablename() string {
	return "user_basic"
}
