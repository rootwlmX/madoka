package models

// Pagination 分页结构
type Pagination struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"totalPage"`
}

// ArticlePageResponse 文章分页响应结构
type ArticlePageResponse struct {
	Pagination Pagination `json:"pagination"`
	List       *[]Article `json:"list"`
}
