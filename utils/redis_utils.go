package utils

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// 当只连接一个数据源的时候，可以直接使用GormClient
// 否则应当自己持有管理InitGormDB返回的GormDB
var RedisClient *RedisDB

type RedisDB struct {
	redisConfig *RedisConfig
	Client      *redis.Client
	lock        sync.RWMutex // lock
}

type RedisConfig struct {
	RedisAddr string
	RedisPwd  string
	RedisDB   int
}

func InitRedis(redisConfig *RedisConfig) *RedisDB {
	redisClient := &RedisDB{
		redisConfig: redisConfig,
		lock:        sync.RWMutex{},
		Client: redis.NewClient(&redis.Options{
			Addr:     redisConfig.RedisAddr,
			Password: redisConfig.RedisPwd, // no password set
			DB:       redisConfig.RedisDB,  // use default DB
		}),
	}
	_, err := redisClient.Client.Ping().Result()
	if err != nil {
		logrus.WithField("redisConfig", redisConfig).Errorln("ping redis error!")
	}
	go redisClient.redisTimer(redisConfig)
	RedisClient = redisClient
	return redisClient
}

func (p *RedisDB) reconnect() {
	client := redis.NewClient(&redis.Options{
		Addr:     p.redisConfig.RedisAddr,
		Password: p.redisConfig.RedisPwd, // no password set
		DB:       p.redisConfig.RedisDB,  // use default DB
	})
	p.Client = client
	_, err := p.Client.Ping().Result()
	if err != nil {
		logrus.WithField("redisConfig", p.redisConfig).Errorln("ping redis error!")
	}
}

func (p *RedisDB) redisTimer(redisConfig *RedisConfig) {
	redisTicker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-redisTicker.C:
			_, err := p.Client.Ping().Result()
			if err != nil {
				logrus.Errorln("redis connect fail,err:", err)
				p.reconnect()
			}
		}
	}
}
