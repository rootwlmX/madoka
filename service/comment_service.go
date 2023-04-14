package service

import (
	"madoka/dao"
	"madoka/engine"
	"madoka/models"
)

// CommentService 评论Service
type CommentService struct {
}

// GetTopCommentAndTopBrowse 最新评论与最多浏览列表
func (s *CommentService) GetTopCommentAndTopBrowse() (*models.TopCommentResponse, error) {
	comments, err := getTopComments()
	if err != nil {
		return nil, err
	}

	browseList, err := getTopBrowse()
	if err != nil {
		return nil, err
	}

	return &models.TopCommentResponse{
		BrowseList:  browseList,
		CommentList: comments,
		LoveCount:   0,
	}, err
}

// GetCommentList 获取文章的评论列表
func (s *CommentService) GetCommentList(pageSize, currentPage, articleID int) (*models.CommentPageResponse, error) {
	commentDao := dao.CommentDao{DBEngine: engine.GetOrmEngine()}
	commentPage, err := commentDao.SelectCommentPage(pageSize, currentPage, articleID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CommentResponse, len(*commentPage))
	parentMap := make(map[int]*models.CommentResponse)
	for i, comment := range *commentPage {
		responses[i].Comment = &comment
		parentMap[responses[i].ID] = &responses[i]
	}

	// 查出该文章下所有父ID不为空的评论，即子评论
	subComments, err := commentDao.SelectSubComments(articleID)
	subMap := make(map[int]int)
	for _, subComment := range *subComments {
		subMap[subComment.ID] = subComment.ParentID
		superID := subMap[subComment.ParentID]
		if superID != 0 {
			subMap[subComment.ID] = superID
		}
	}

	for _, subComment := range *subComments {
		children := *parentMap[subMap[subComment.ID]].Children
		children = append(children, models.CommentChildrenResponse{
			Comment:        &subComment,
			ParentUserID:   subComment.ParentID,
			ParentUserName: "",
		})
	}

	if err != nil {
		return nil, err
	}

	count, err := commentDao.SelectCommentCountOfArticle(articleID)
	return &models.CommentPageResponse{
		CommentPagination: models.CommentPagination{
			Pagination: &models.Pagination{
				CurrentPage: currentPage,
				PageSize:    pageSize,
				Total:       count,
				TotalPage:   count/pageSize + 1,
			},
			CountTotal: count,
		},
		List: &responses,
	}, err
}

func getTopComments() (*[]models.CommentList, error) {
	commentDao := dao.CommentDao{DBEngine: engine.GetOrmEngine()}
	comments, err := commentDao.SelectRecentComments(10)
	for i := range *comments {
		(*comments)[i].IsArticle = true
	}
	return comments, err
}

func getTopBrowse() (*[]models.BrowseList, error) {
	articleDao := dao.ArticleDao{DBEngine: engine.GetOrmEngine()}
	return articleDao.SelectMostBrowseArticle(10)
}
