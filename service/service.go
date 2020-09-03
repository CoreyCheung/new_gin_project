package service

import (
	"new_gin_project/config"
	"new_gin_project/dao"
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

// Service struct
type Service struct {
	dao *dao.Dao // dao
	//	redisMid     *redis_mid.RedisMid    // redis中间层
	gormClient *utils.GormDB // gormdb
	//	redisClient  *db.RedisDB            // redis
	//	mpClient     *core.Client           // wechat client
	//	smsClient    *sms.AliCloudSMSClient // ali sms client
	//	chainClient  *chain.ChainClient     // chainClient
	//	aliPayClient *alipay.Client         // ali pay client
}

func NewService() *Service {
	DBInit()
	var srv Service
	srv.gormClient = GormClient
	srv.dao = dao.New()
	return &srv
}
