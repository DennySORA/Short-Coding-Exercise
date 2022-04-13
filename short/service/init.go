package service

import (
	"fmt"
	"io"
	"os"
	"runtime/trace"
	"shortener/common"
	"shortener/database"
	"shortener/logs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ServerInit() []func() error {
	close := make([]func() error, 3)
	initializationViper()
	close = append(close, initializationLog()...)
	if viper.GetBool("LOG_TRACE") {
		close = append(close, initializationTrace()...)
	}
	if viper.GetBool("LOG_GIN") {
		close = append(close, initializationGinLog())
	}
	close = append(close, initializationDatabase())
	return close
}

func initializationDatabase() func() error {
	return database.CreateDatabaseContext()
}

func initializationViper() {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Fatal error config file: " + err.Error())
	}
	viper.AutomaticEnv()
}

func initializationGinLog() func() error {
	fildeName := fmt.Sprintf(
		"%s/%s_%s.log",
		viper.GetString("LOG_FILE_PATH"),
		"gin",
		time.Now().UTC().Format("2006-01-02"),
	)
	logFile, err := os.Create(fildeName)
	if err != nil {
		logs.Warning.Println(err)
		return nil
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	return logFile.Close
}

func initializationLog() []func() error {
	return logs.LogSetting()
}

func initializationTrace() []func() error {
	path := viper.GetString("LOG_FILE_PATH")
	if path == "" {
		path = "./log"
	} else {
		path = path + "trace.out"
	}

	if file, err := os.Create(path); err != nil {
		logs.Error.Panic(err)
	} else if err = trace.Start(file); err != nil {
		logs.Error.Panic(err)
	} else {
		return []func() error{
			common.PKGCloseFunc(trace.Stop),
			file.Close,
		}
	}
	return nil
}
