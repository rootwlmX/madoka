package service

import (
	"madoka/models"
)

type CommentService struct {
}

// GetTopCommentAndTopBrowse 最新评论与最多浏览列表
func (s *CommentService) GetTopCommentAndTopBrowse() (models.TopCommentResponse, error) {
	return models.TopCommentResponse{}, nil
}
