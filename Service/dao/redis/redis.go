package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明一个全局rdb变量
var rbd *redis.Client

// 初始化连接
func Init() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), // 密码
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rbd.Close()
}
