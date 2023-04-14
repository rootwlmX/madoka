// Package engine DBEngine
package engine

import (
	"fmt"
	"madoka/models"
	"xorm.io/xorm"
)

var _dbEngine *Orm

// Orm xormEngine
type Orm struct {
	*xorm.Engine
}

// GetOrmEngine get dbEngine
func GetOrmEngine() *Orm {
	return _dbEngine
}

// NewOrmEngine get a new orm engine
func NewOrmEngine(appInfo *models.AppInfo) (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", appInfo.UserName, appInfo.Password, appInfo.Host, appInfo.Port, appInfo.DataBase)
	engine, err := xorm.NewEngine(appInfo.DriverName, url)

	engine.ShowSQL(true)
	fmt.Println(appInfo)

	if err != nil {
		return nil, err
	}

	err = engine.Sync2(new(models.Article))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	_dbEngine = orm

	return engine, nil
}
