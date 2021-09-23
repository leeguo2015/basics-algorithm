package model

import (
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type (
	user struct {
		Model
		UserName     string `json:"user_name"`
		NickName     string `json:"nick_name"`
		Password     string `json:"password"`
		Portrait     string `json:"portrait"`
		Gender     int `json:"gender"`
		IDCard     string `json:"id_card"`
		Addr     string `json:"addr"`
		Phone     string `json:"Phone"`
		Birth     time.Time `json:"birth"`
		Email string `json:"email"`
		IPAddr string `json:"ip_addr"`
		Status string `json:"status"`
		LastTime *time.Time `json:"last_time"`
	}

	fmtTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

