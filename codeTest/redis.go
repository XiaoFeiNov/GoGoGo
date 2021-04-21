package codeTest

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type redisClient struct {
	Conn redis.Conn
}

var RedisClient redisClient

func InitClient(flagCh chan struct{})  {
	ch := make(chan struct{}, 1)
	// 连接到redis
	database := redis.DialDatabase(0) // 选择数据库（0-15），默认0
	password := redis.DialPassword("123456")
	var err error
	RedisClient.Conn, err = redis.Dial("tcp", "127.0.0.1:6379", database, password)
	if err != nil {
		log.Fatalf("redis.Dial err = %v", err)
		flagCh <- struct{}{}
		return
	}
	defer RedisClient.Conn.Close()
	fmt.Println("redis client running...")
	flagCh <- struct{}{}
	<- ch
	fmt.Println("redis client done")
}

// 写入数据
func (r *redisClient) SetString(key, value string) {
	if _, err := RedisClient.Conn.Do("Set", key, value); err != nil{
		log.Fatalf("redis Set err = %v", err)
		return
	}
}

// 读取数据
func (r *redisClient) GetString(key string) {
	value, err := redis.String(r.Conn.Do("Get", key))
	if err != nil{
		log.Fatalf("redis Get err = %v", err)
		return
	}
	fmt.Printf("%s = %s\n", key, value)
}