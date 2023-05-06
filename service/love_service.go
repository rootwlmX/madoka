package service

import (
	"madoka/dao"
	"madoka/engine"
	"madoka/models"
	"time"
)

// LoveService 点赞Service
type LoveService struct {
}

// AddLove 进行点赞
func (s *LoveService) AddLove(ip string) (int, error) {
	loveDao := dao.LoveDao{DBEngine: engine.GetOrmEngine()}
	count, err := loveDao.SelectLoveByDateAndIP(ip, time.Now())
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return -1, err
	}

	_, err = loveDao.InsertLove(&models.Love{
		IP:         ip,
		CreateTime: time.Now(),
	})

	return 1, err
}
