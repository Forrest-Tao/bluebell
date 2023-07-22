package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}
	zap.L().Debug("logger init success")
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			return
		}
	}(zap.L())
	// 3. 初始化 MySQL
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql settings failed, err: %v\n", err)
		return
	}
	defer mysql.Close()
	// 4. 初始化 Redis
	/*if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis settings failed, err: %v\n", err)
		return
	}
	defer redis.Close()*/
	// 雪花算法：分布式 ID 生成器
	if err := snowflake.Init(setting.Conf.SnowFlakeConfig.StartTime, setting.Conf.SnowFlakeConfig.MachineId); err != nil {
		fmt.Printf("init snowflake failed, err: %v\n", err)
		return
	}
	// 初始化 gin 框架内置的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("Init validator trans failed, err: %v\n", err)
		return
	}
	// 5. 注册路由
	r := router.SetUp()
	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.AppConfig.Port),
		Handler: r,
	}
	go func() {
		// 开启一个 goroutine 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，未关闭服务器操作设置一个 5 秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 Ctrl + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify 把收到的 syscall.SIGINT 或者 syscall.SIGTERM 信号转发给 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个 5 秒超时 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 五秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超出 5 秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting ...")
}
