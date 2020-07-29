package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lhj168os/myiris/common/timer"
	"github.com/lhj168os/myiris/conf"
	"strconv"
)

type Comment struct {
	gorm.Model
	Uid     uint
	Content string
	LikeCnt int
	HateCnt int
	DayNo   int
}

func NewComment(uid uint, content string) *Comment {
	cm := &Comment{
		Uid:     uid,
		Content: content,
		LikeCnt: 0,
		HateCnt: 0,
		DayNo:   timer.GetDayNo(),
	}
	return cm
}

func (c *Comment) save() {
	dbCon.Save(c)
}

func GetComment(id uint) *Comment {
	var u = &Comment{}
	if dd, ok := dbs.commentData[id]; ok {
		u = dd
	} else {
		dbCon.Where("id = ?", strconv.Itoa(int(id))).First(u)
		if u.ID > 0 {
			dbs.commentData[u.ID] = u
		}
	}
	return u
}

func GetCommentsByUid(uid uint) []*Comment {
	var cm []*Comment
	dbCon.Where("uid = ?", strconv.Itoa(int(uid))).Find(cm)
	return cm
}

func PublicComment(u iUser, cm *Comment) {
	if !u.CanComment() {
		cnt := conf.GetCommentCfg().Cnt
		fmt.Printf("Allow public comment only %d times a day!\n", cnt)
		return
	}
	u.OnComment()
	cm.save()
	fmt.Printf("uid=%d comment successful!\n", cm.Uid)
}

func LikeComment(id uint) {
	cm := GetComment(id)
	if cm.ID > 0 {
		cm.LikeCnt += 1
	}
}

func HateComment(id uint) {
	cm := GetComment(id)
	if cm.ID > 0 {
		cm.HateCnt += 1
	}
}
