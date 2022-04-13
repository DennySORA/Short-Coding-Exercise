package app

import (
	"fmt"
	"net/http"
	"shortener/common"
	"shortener/database"
	"shortener/logs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetShort(c *gin.Context) {
	funcName := "Get Short"

	ShortWithDB := database.NewShortWithDB()
	ShortStruct, err := ShortWithDB.GetShortWithDB(c.Param("short"))

	if err != nil {
		logs.Warning.Println(funcName, err)
		c.JSON(common.GetShortError())
		c.Abort()
		return
	}

	is_expired, err := common.CheckExpires(ShortStruct.Expires.Unix())

	if is_expired || err != nil {
		logs.Warning.Println(funcName, err)
		c.JSON(common.GetShortError())
		c.Abort()
		return
	}

	c.Redirect(http.StatusSeeOther, ShortStruct.Url)
}

func AddShort(c *gin.Context) {
	funcName := "Add Short"

	shortParameter, err := parsingShortParameter(c)
	if err != nil {
		logs.Warning.Println(funcName, err)
		c.JSON(common.GetShortError())
		c.Abort()
		return
	}

	timeExpires, err := time.Parse(time.RFC3339, shortParameter.Expires)
	if err != nil {
		logs.Warning.Println(funcName, err)
		c.JSON(common.GetShortError())
		c.Abort()
		return
	}
	is_expired, err := common.CheckExpires(timeExpires.Unix())

	if is_expired || err != nil {
		logs.Warning.Println(funcName, err)
		c.JSON(common.GetShortError())
		c.Abort()
		return
	}

	id := uuid.New().String()[:8]
	ShortWithDB := database.NewShortWithDB()
	err = ShortWithDB.AddShortWithDB(shortParameter, id)
	if err != nil {
		logs.Warning.Println(funcName, err)
		c.JSON(common.GetShortError())
		c.Abort()
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	c.JSON(
		http.StatusCreated,
		common.ShortResponse{
			ShortUrl: fmt.Sprintf("%s://%s/%s", scheme, c.Request.Host, id),
			ID:       id,
		},
	)
}

func parsingShortParameter(c *gin.Context) (*common.ShortParameter, error) {
	shortParameter := common.ShortParameter{}
	err := c.BindJSON(&shortParameter)
	if err != nil {
		return nil, err
	}
	return &shortParameter, nil
}
