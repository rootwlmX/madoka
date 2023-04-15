package dao

import (
	"madoka/engine"
	"madoka/models"
)

// UserDao userDao
type UserDao struct {
	DBEngine *engine.Orm
}

// SelectUserByIDs 用户ID列表查询用户列表
func (d *UserDao) SelectUserByIDs(idList []int) (*[]models.User, error) {
	users := make([]models.User, 0)
	err := d.DBEngine.In("id", idList).Find(&users)
	return &users, err
}
