package models

type Pagination struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"totalPage"`
}

type ArticlePageResponse struct {
	Pagination Pagination `json:"pagination"`
	List       *[]Article `json:"list"`
}
