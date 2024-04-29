package main

import (
	"context"
	_ "github.com/xiaoxin1992/kube-admin/apis"
	"github.com/xiaoxin1992/kube-admin/pkg/config"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
	"github.com/xiaoxin1992/kube-admin/pkg/ginx"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"github.com/xiaoxin1992/kube-admin/version"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gx := ginx.Init(config.GetConfig().Http.Host, config.GetConfig().Http.Port, config.GetConfig().Level)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	db := database.GetPool()
	if err := db.Create(); err != nil {
		panic(err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		gx.WaitStop(ch, ctx)
		db.Destroy()
	}()
	go func() {
		defer wg.Done()
		if err := gx.Start(); err != nil {
			panic(err)
		}
	}()
	wg.Wait()
	return
}

func initLogger() {
	loggerOption := logger.Option{
		Name:      version.AppName(),
		Level:     config.GetConfig().Level,
		Path:      config.GetConfig().Logger.Path,
		IsConsole: config.GetConfig().Logger.Console,
		IsFile:    config.GetConfig().Logger.IsFile,
		LocalTime: config.GetConfig().Logger.LocalTime,
		Compress:  config.GetConfig().Logger.Compress,
		Format:    config.GetConfig().Logger.Format,
		MaxSize:   config.GetConfig().Logger.MaxSize,
		MaxBackup: config.GetConfig().Logger.MaxBackup,
		MaxAge:    config.GetConfig().Logger.MaxAge,
	}
	err := logger.InitLogger(loggerOption)
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	rootExec, err := os.Executable()
	if err != nil {
		panic(err)
	}
	rootDir, err := filepath.EvalSymlinks(filepath.Dir(rootExec))
	if err != nil {
		panic(err)
	}
	if err = config.Init(filepath.Join(rootDir, "config", "config.toml")); err != nil {
		panic(err)
	}
}

func init() {
	initConfig()
	initLogger()
}
