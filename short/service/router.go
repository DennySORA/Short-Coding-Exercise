package service

import (
	"shortener/app"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func registerApp(engine *gin.Engine) {
	// Register pprof.
	if viper.GetBool("LOG_PPROF") {
		pprof.Register(engine)
	}
	registerShortAPI(engine)
}

func registerShortAPI(engine *gin.Engine) {

	// Register short api.
	apiV1 := engine.Group("/api/v1")
	apiV1.POST("/urls", app.AddShort)

	// Register short url redirect with cached.
	store := persist.NewMemoryStore(time.Minute)
	engine.GET("/:short",
		cache.CacheByRequestURI(store, time.Second*5),
		app.GetShort,
	)
}
