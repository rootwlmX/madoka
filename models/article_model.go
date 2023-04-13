// Package models models
package models

import "time"

// Article 文章实体结构
type Article struct {
	ID         int       `json:"id" xorm:"id"`
	Title      string    `json:"title" xorm:"title"`
	Content    string    `json:"content" xorm:"content"`
	CreateTime time.Time `json:"createTime" xorm:"create_time"`
	UpdateTime time.Time `json:"updateTime" xorm:"update_time"`
}
