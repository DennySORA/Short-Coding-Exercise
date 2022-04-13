package main

import (
	"log"
	"runtime/debug"
	"shortener/service"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Error is:", err, "\n\nStack is:", string(debug.Stack()))
		}
	}()
	close := service.ServerInit()
	service.ServerStart()
	service.ServerShutdown(close)
}
