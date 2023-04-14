package models

import "time"

// Pagination 分页结构
type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"`
	TotalPage   int `json:"totalPage"`
}

// ArticlePageResponse 文章分页响应结构
type ArticlePageResponse struct {
	Pagination Pagination `json:"pagination"`
	List       *[]Article `json:"list"`
}

// TopCommentResponse getTopComment返回结构
type TopCommentResponse struct {
	BrowseList  *[]BrowseList  `json:"browseList"`
	CommentList *[]CommentList `json:"commentList"`
	LoveCount   int            `json:"loveCount"`
}

// BrowseList 浏览列表
type BrowseList struct {
	Title     string `json:"title"`
	Count     int    `json:"count"`
	ArticleID int    `json:"articleId" xorm:"article_id"`
}

// CommentList 评论列表
type CommentList struct {
	ArticleID  int       `json:"articleId" xorm:"article_id"`
	Avatar     string    `json:"avatar"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
	IsArticle  bool      `json:"isArticle"`
	Title      string    `json:"title"`
	UserID     int       `json:"userId" xorm:"user_id"`
	UserName   string    `json:"userName"`
}
