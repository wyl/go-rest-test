package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

type HttpClient struct {
	*gin.Engine
}

func NewHttpClient() *HttpClient {

	c := &HttpClient{
		Engine: gin.Default(),
	}
	c.Use(CORSMiddleWare())
	c.buildHandler()
	return c
}

type LoginParam struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (ic *HttpClient) buildHandler() {
	ic.GET("/healthz", func(context *gin.Context) { context.JSON(http.StatusOK, "ok") })
	ic.GET("/ping", func(context *gin.Context) { context.JSON(http.StatusOK, "pong") })
	ic.GET("/header", func(context *gin.Context) { context.JSON(http.StatusOK, context.Request.Header) })
	ic.Any("/data", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": 0, "data": map[string]interface{}{"name": "test", "age": 20, "date": time.Now()}})
	})
	ic.GET("/ip", func(context *gin.Context) { context.String(http.StatusOK, "%v", context.ClientIP()) })
	ic.GET("/panic", func(context *gin.Context) { panic("foo") })
	ic.GET("/1s", func(context *gin.Context) { time.Sleep(1e9); context.JSON(http.StatusOK, "ok") })
	ic.GET("/err/404", func(context *gin.Context) { context.JSON(http.StatusNotFound, "NotFound") })
	ic.GET("/err/403", func(context *gin.Context) { context.JSON(http.StatusForbidden, "Forbidden") })
	ic.GET("/err/500", func(context *gin.Context) { context.JSON(http.StatusInternalServerError, "InternalServerError") })
	ic.GET("/err/502", func(context *gin.Context) { context.JSON(http.StatusBadGateway, "BadGateway") })
	ic.POST("/post", func(context *gin.Context) {
		var param map[string]interface{}
		err := context.ShouldBindBodyWith(&param, binding.JSON)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"status": 0, "data": param})
		}
	})

	ic.POST("/form", func(context *gin.Context) {
		var param LoginParam
		err := context.ShouldBind(&param)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"status": 0, "data": param})
		}
	})

	ic.POST("/query", func(context *gin.Context) {
		var param LoginParam
		err := context.ShouldBindQuery(&param)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"status": 0, "data": param})
		}
	})
}
