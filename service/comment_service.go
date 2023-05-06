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

	loveCount, err := getLoveCount()
	if err != nil {
		return nil, err
	}

	return &models.TopCommentResponse{
		BrowseList:  browseList,
		CommentList: comments,
		LoveCount:   int(loveCount),
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

func getLoveCount() (int64, error) {
	loveDao := dao.LoveDao{DBEngine: engine.GetOrmEngine()}
	return loveDao.SelectLoveCount()
}

// GetCommentList 获取文章的评论列表
func (s *CommentService) GetCommentList(pageSize, currentPage, articleID int) (*models.CommentPageResponse, error) {
	commentDao := dao.CommentDao{DBEngine: engine.GetOrmEngine()}
	commentPage, err := commentDao.SelectCommentPage(pageSize, currentPage, articleID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CommentResponse, len(*commentPage))
	userIDSet := make(map[int]bool)
	parentCommentIDList := make([]int, 0)
	for i, comment := range *commentPage {
		responses[i].Comment = comment
		childrenResponses := make([]models.CommentChildrenResponse, 0)
		responses[i].Children = &childrenResponses
		parentCommentIDList = append(parentCommentIDList, comment.ID)
		userIDSet[comment.UserID] = true
	}
	if err != nil {
		return nil, err
	}

	subComments, _ := commentDao.SelectSubComments(articleID, parentCommentIDList)
	parentCommentMap := make(map[int]models.CommentResponse, len(*commentPage))
	for _, commentResponse := range responses {
		parentCommentMap[commentResponse.ID] = commentResponse
	}

	for _, subComment := range *subComments {
		userIDSet[subComment.UserID] = true
	}

	userDao := dao.UserDao{DBEngine: engine.GetOrmEngine()}
	keys := make([]int, 0)
	for k := range userIDSet {
		keys = append(keys, k)
	}
	userList, _ := userDao.SelectUserByIDs(keys)
	userIDNameMap := make(map[int]string, len(*userList))
	for _, user := range *userList {
		userIDNameMap[user.ID] = user.UserName
	}

	for _, subComment := range *subComments {
		id := subComment.ParentID
		response := parentCommentMap[id]
		childrenResponse := models.CommentChildrenResponse{
			Comment:        subComment,
			ParentUserID:   subComment.ParentUserID,
			ParentUserName: userIDNameMap[subComment.ParentUserID],
			UserName:       userIDNameMap[subComment.UserID],
		}
		*response.Children = append(*response.Children, childrenResponse)
	}

	for i, response := range responses {
		responses[i].UserName = userIDNameMap[response.UserID]
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
