package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"oauth2/httpType"
	"oauth2/middleware"
	"strconv"
	"time"
)

var rdb *redis.Client

func createRedisConnection() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "81.68.239.206:6379",
		Password: "tanyongfeng13666",
	})
}

func main() {
	ctx := context.Background()
	createRedisConnection()
	r := gin.Default()

	//config the middleware
	r.Use(middleware.AuthMiddleWare(), middleware.ErrorHandlingMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//get accessToken
	r.POST("/access_token", func(c *gin.Context) {
		//1. check username and password (ignored)
		//2. generate the access_token and refresh_token
		accessTokenId := strconv.FormatInt(generateUUID(), 10)
		refreshTokenId := strconv.FormatInt(generateUUID(), 10)
		if err := rdb.Set(ctx, "access_token:"+accessTokenId, "test", time.Hour*2).Err(); err != nil {
			_ = c.Error(fmt.Errorf("error occurs when add access_token %v", err))
			return
		}
		if err := rdb.Set(ctx, "refresh_token:"+refreshTokenId, "test", time.Hour*24*7).Err(); err != nil {
			_ = c.Error(fmt.Errorf("error occurs when add refresh_token %v", err))
			return
		}
		//3.response
		c.JSON(http.StatusOK, httpType.Response{
			Code: 0,
			Data: httpType.AccessTokenResp{
				AccessToken:      accessTokenId,
				ExpiresIn:        int64(time.Hour * 2 / time.Second),      //access_token lifetime
				RefreshExpiresIn: int64(time.Hour * 24 * 7 / time.Second), // refresh_token lifetime
				RefreshToken:     refreshTokenId,
				TokenType:        "Bearer",
			},
			Msg: "success",
		})

	})

	//refresh accessToken
	r.POST("/refresh_access_token", func(c *gin.Context) {
		// get refresh Token
		var req httpType.FreshAccessTokenRequest
		// 将请求中的 JSON 数据绑定到 req 结构体
		if err := c.ShouldBindJSON(&req); err != nil || req.RefreshToken == "" {
			_ = c.Error(fmt.Errorf("refresh_access_token is empty"))
			return
		}

		// check the refresh_token
		if exists := rdb.Exists(ctx, "refresh_token:"+req.RefreshToken).Val(); exists == 0 {
			_ = c.Error(fmt.Errorf("refresh_token %s not found", req.RefreshToken))
			return
		}

		// generate the new access_token
		newAccessTokenId := strconv.FormatInt(generateUUID(), 10)
		if err := rdb.Set(ctx, "access_token:"+newAccessTokenId, "test", time.Hour*2).Err(); err != nil {
			_ = c.Error(fmt.Errorf("error occurs when add access_token %v", err))
			return
		}
		//response
		c.JSON(http.StatusOK, &httpType.FreshAccessTokenResp{
			AccessToken:      newAccessTokenId,
			ExpiresIn:        int64(time.Hour * 2 / time.Second),
			RefreshExpiresIn: int64(time.Hour * 24 * 7 / time.Second),
			RefreshToken:     req.RefreshToken,
			TokenType:        "Bearer",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
