Redis 操作

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// Set(c, "mykey", "myvalue")

	// value := Get(c, "mykey")
	// fmt.Printf("Get mykey: %v \n", value)

	// DeferSet(c, "mykey", "myvalue", "10")

	// println(Exists(c, "mykey"))

	// Del(c, "test")
	// println(Get(c, "test"))

	// SetJSON(c)

	// SetDeferTime(c, "profile")

	// ListPush(c)

	// Pipelining(c)
}

func Set(c redis.Conn, key string, value string) {
	_, err := c.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func Get(c redis.Conn, key string) string {
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	return value
}

// DeferSet set时同时设置key失效时间
func DeferSet(c redis.Conn, key string, value string, deferTime string) {
	_, err := c.Do("SET", key, value, "EX", deferTime)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username := Get(c, key)
	fmt.Printf("Get mykey: %v \n", username)

	time.Sleep(13 * time.Second)

	username = Get(c, key)
	fmt.Printf("Get mykey: %v \n", username)
}

// SetDeferTime 设置key失效时间
func SetDeferTime(c redis.Conn, key string) {
	// 设置过期时间为24小时
	// n, _ := c.Do("EXPIRE", key, 24*3600)
	n, _ := c.Do("EXPIRE", key, 30)
	if n == int64(1) {
		fmt.Println("success")
	}
}

// Exists 检测值是否存在
func Exists(c redis.Conn, key string) bool {
	is_key_exit, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}
	return is_key_exit
}

func Del(c redis.Conn, key string) {
	_, err := c.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
}

func SetJSON(c redis.Conn) {
	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}

// ListPush 列表操作
func ListPush(c redis.Conn) {

	// lpush 后进先出
	_, err := c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))

	for k, v := range values {
		fmt.Println(k)
		fmt.Println(string(v.([]byte)))
	}
}

// Pipelining 管道化操作
func Pipelining(c redis.Conn) {
	// Send向连接的输出缓冲中写入命令。
	// Flush将连接的输出缓冲清空并写入服务器端。
	// Recevie按照FIFO顺序依次读取服务器的响应。
	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	c.Receive()         // reply from SET
	v, _ := c.Receive() // reply from GET
	fmt.Printf("%s", v)
}
