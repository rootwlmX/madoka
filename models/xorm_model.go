// Package models models
package models

import "time"

// Article 文章实体结构
type Article struct {
	ID         int       `json:"id" xorm:"id pk autoincr notnull"`
	Title      string    `json:"title" xorm:"title"`
	Content    string    `json:"content" xorm:"content LONGTEXT"`
	ClassID    int       `json:"classId" xorm:"class_id"`
	CreateTime time.Time `json:"createTime" xorm:"create_time"`
	UpdateTime time.Time `json:"updateTime" xorm:"update_time"`
}
