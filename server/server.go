package server

import (
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

fun Run() (err error) {
	defer logger.Log.Flush()

	gin.SetMode(gin.ReleaseMode)
	r : = gin.New()
}

func middleWareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().UTC()
		c.Next()
		logger.Log.Infof("%v %v %v %s")
	}
}