package redis

import (
	"fmt"
	"zengzhicheng/Decentralized-social-platforms/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局rdb变量
var rbd *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // 密码
		DB:       cfg.Db,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rbd.Close()
}
