package main

import (
	"fmt"
	"zengzhicheng/Decentralized-social-platforms/controller"
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/dao/redis"
	"zengzhicheng/Decentralized-social-platforms/logger"
	"zengzhicheng/Decentralized-social-platforms/pkg/snowflake"
	"zengzhicheng/Decentralized-social-platforms/routes"
	"zengzhicheng/Decentralized-social-platforms/settings"

	"go.uber.org/zap"
)

// 搭建项目脚手架

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init setrings failed,err:%v\n", err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	// 3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer mysql.CLose() // 程序退出关闭数据库连接
	// 4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
	}
	// 5.注册路由
	r := routes.Setup()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed,err:%v\n", err)
	}
}
