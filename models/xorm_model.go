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

// Comment 评论实体结构
type Comment struct {
	ID         int       `json:"id" xorm:"id pk autoincr notnull"`
	Content    string    `json:"content" xorm:"content LONGTEXT"`
	ArticleID  int       `json:"articleId" xorm:"article_id"`
	UserID     int       `json:"userId" xorm:"user_id"`
	CreateTime time.Time `json:"createTime" xorm:"create_time"`
}

// Browse 浏览
type Browse struct {
	ID         int       `json:"id" xorm:"id pk autoincr notnull"`
	UserID     int       `json:"userId" xorm:"user_id"`
	ArticleID  int       `json:"articleId" xorm:"article_id"`
	CreateTime time.Time `json:"createTime" xorm:"create_time"`
}

// User 用户结构
type User struct {
	ID            int       `json:"id" xorm:"id pk autoincr notnull"`
	UserName      string    `json:"userName" xorm:"userName"`
	Avatar        string    `json:"avatar" xorm:"avatar"`
	Email         string    `json:"email" xorm:"email"`
	Bio           string    `json:"bio" xorm:"bio"`
	Blog          string    `json:"blog" xorm:"blog"`
	WebBlogName   string    `json:"webBlogName" xorm:"web_blog_name"`
	WebBlog       string    `json:"webBlog" xorm:"web_blog"`
	WebBlogIcon   string    `json:"webBlogIcon" xorm:"web_blog_icon"`
	WebBlogDesc   string    `json:"webBlogDesc" xorm:"web_blog_desc"`
	WebBlogState  int       `json:"webBlogState" xorm:"web_blog_state"`
	Label         string    `json:"label" xorm:"label"`
	Origin        string    `json:"origin" xorm:"origin"`
	CreateTime    time.Time `json:"createTime" xorm:"create_time"`
	UpdateTime    time.Time `json:"updateTime" xorm:"update_time"`
	LastLoginTime time.Time `json:"lastLoginTime" xorm:"last_login_time"`
	Status        int       `json:"status" xorm:"status"`
}
