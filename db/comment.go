package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"myiris/com"
	"myiris/conf"
	"strconv"
	"time"
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
		DayNo:   com.GetDayNo(),
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

func PublicComment(cm *Comment) {
	u := GetUser(cm.Uid)
	if com.GetDayNo(u.CommentTime) == com.GetDayNo() {
		cnt := conf.GetCommentCfg().Cnt
		if u.CommentCnt >= cnt {
			fmt.Printf("Allow public comment only %d times a day!\n", cnt)
			return
		}
	}
	u.CommentTime = time.Now().Unix()
	u.CommentCnt += 1
	u.save()
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
