package db

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Uid     uint
	Content string
	LikeCnt int
	HateCnt int
}
