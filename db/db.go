package db

import (
	"new_gin_project/config"
	"new_gin_project/utils"
)

var GormClient *utils.GormDB

func DBInit() {
	GormClient = utils.InitGormDB(&utils.DBConfig{
		DBAddr:       config.Optional.MysqlStr,
		MaxIdleConns: 30,
		LogMode:      utils.Uint8ToBool(config.Optional.DBLog),
	})
}

func CreateModel(value interface{}) error {
	if GormClient.Client.NewRecord(value) {
		if mydb := GormClient.Client.Create(value); mydb.Error != nil {
			return mydb.Error
		}
	}
	return nil
}
