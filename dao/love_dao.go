package dao

import (
	"madoka/engine"
	"madoka/models"
	"time"
)

// LoveDao LoveDao
type LoveDao struct {
	DBEngine *engine.Orm
}

// InsertLove 插入点赞
func (d *LoveDao) InsertLove(love *models.Love) (int64, error) {
	return d.DBEngine.InsertOne(love)
}

// SelectLoveByDateAndIP 根据日期IP查询点赞
func (d *LoveDao) SelectLoveByDateAndIP(ip string, time time.Time) (int64, error) {
	love := new(models.Love)
	return d.DBEngine.Where("ip = ?", ip).Where("TO_DAYS(create_time) = TO_DAYS(?)", time).Count(love)
}

// SelectLoveCount love数量
func (d *LoveDao) SelectLoveCount() (int64, error) {
	love := new(models.Love)
	return d.DBEngine.Count(love)
}
