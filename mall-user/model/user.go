package model

import (
	"time"

	"gorm.io/gorm"
)

//公用model
type BaseModel struct {
	ID        int32          `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:add_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time"`
	IsDeleted bool
}

//用户表
type User struct {
	BaseModel
	Mobile string `gorm:"index:uni_mobile;unique;not null;type:varchar(11)"`
	//保存密码的时候 同时保存了salt和加密算法
	Password string `gorm:"type:varchar(150);not null"`
	NickName string `gorm:"type:varchar(14)"`
	Gender   int    `gorm:"type:int(1) comment '性别 0代表男 1代表女'"`
	Birthday *time.Time
	Role     int `gorm:"type:int(1) comment '角色 0代表管理员用户 1代表普通用户';default:1"`
}
