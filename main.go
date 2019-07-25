package main

import (
	delivery "account/administrator/delivery/http"
	"account/administrator/repository"
	"account/administrator/usecase"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 載入router root
	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9487")
}

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:9528", "http://localhost:9528"},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	sqlConn := newMysql()
	redisConn := newRedis()
	administratorRepository := repository.NewMysqlAdministratorRepository(sqlConn)
	administratorSidRepository := repository.NewRedisAdministratorRepository(redisConn)
	administratorUsecase := usecase.NewAdministratorUsecase(administratorRepository, administratorSidRepository)

	// 載入 Administrator router
	delivery.NewAdministratorHandler(r, administratorUsecase)

	return r
}

// 建立mysql
func newMysql() *gorm.DB {
	connectName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "qwe123", "127.0.0.1", "3388", "collection")
	Conn, err := gorm.Open("mysql", connectName)
	if err != nil {
		log.Fatalf("建立db連線失敗: %s", err.Error())
	}

	return Conn
}

// 建立redis
func newRedis() *redis.Pool {
	Conn := &redis.Pool{
		Wait:        true,
		MaxIdle:     20,
		MaxActive:   2000,
		IdleTimeout: 10 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6378")
			if err != nil {
				return nil, err
			}
			// _, err = c.Do("AUTH", "123")
			// if err != nil {
			// 	return nil, err
			// }

			c.Do("SELECT", 3)

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return Conn
}
