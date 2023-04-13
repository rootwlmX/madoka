package service

import (
	"madoka/models"
)

// CommentService 评论Service
type CommentService struct {
}

// GetTopCommentAndTopBrowse 最新评论与最多浏览列表
func (s *CommentService) GetTopCommentAndTopBrowse() (models.TopCommentResponse, error) {
	return models.TopCommentResponse{}, nil
}
