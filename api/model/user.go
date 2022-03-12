package model

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	List []User `json:"list" binding:"gt=0,lt=100,size,dive"` // 切片长度>0，<100,dive验证每一项
	Size int    `json:"size"`
}

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;" json:"id"`
	Phone     string         `gorm:"type:char(11);uniqueIndex;comment:手机号" json:"phone"`
	Email     string         `gorm:"type:varchar(50);uniqueIndex;comment:邮件" json:"email"`
	Unionid   string         `gorm:"type:varchar(100);uniqueIndex;comment:微信号唯一标识" json:"unionid"`
	Openid1   string         `gorm:"type:varchar(100);comment:小程序🆔" json:"openid1"`
	Openid2   string         `gorm:"type:varchar(100);comment:公众号🆔" json:"openid2"`
	Avatar    string         `gorm:"type:varchar(100);comment:头像" json:"avatar"`
	Nickname  string         `gorm:"type:varchar(20);not null;comment:昵称" json:"nickname"`
	Password  string         `gorm:"type:varchar(32);comment:密码" json:"password"`
	Birthday  string         `gorm:"type:date;comment:出生日期" json:"birthday"`
	Gender    int            `gorm:"type:tinyint(1);default:-1;comment:-1保密 0女 1男" json:"gender"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;->:false;"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) String() string {
	return "users"
}
