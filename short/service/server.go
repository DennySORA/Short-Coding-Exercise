package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shortener/logs"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

func ServerStart() {
	// Set gin release mode.
	if viper.GetBool("RELEASE_MODE") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create gin engine.
	engine := gin.Default()

	// Add CORS middleware.
	settingCors(engine)
	// Registering router.
	registerApp(engine)

	// Set port.
	port := viper.GetString("SERVER_IP_PORT")
	if port == "" {
		port = "0.0.0.0:8888"
	}

	// Set http server parameter.
	ser := &http.Server{
		Addr:           port,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Set http2.
	http2.ConfigureServer(ser, &http2.Server{})

	// Start listen.
	listen(ser)

	// Wait shutdown.
	waitShutdown(ser)

	// Print Exited info.
	logs.Info.Println("Server exiting")
}

func settingCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowWildcard = true
	engine.Use(cors.New(config))
}

func ServerShutdown(close []func() error) {
	for _, cls := range close {
		if cls != nil {
			log.Println(cls())
		}
	}
	log.Println("Server shutdown complete.")
}

func listen(ser *http.Server) {
	go func(ser *http.Server) {
		err := ser.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logs.Error.Fatalf("Listen: %s\n", err)
		}
	}(ser)
}

func waitShutdown(ser *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Info.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ser.Shutdown(ctx); err != nil {
		logs.Error.Fatal("Server Shutdown: ", err)
	}
}
