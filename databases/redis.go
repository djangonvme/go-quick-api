package databases

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jangozw/gin-api-common/configs"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var RedisPool *redis.Pool

func init() {
	if RedisPool != nil {
		return
	}
	RedisPool = newRedisPool()
}

func newRedisPool() *redis.Pool {
	conf, err := configs.GetSection("redis")
	if err != nil {
		log.Fatalln("Start redis failed! couldn't get config")
	}
	RedisPool = &redis.Pool{
		MaxIdle:     5,                 // idle的列表长度, 空闲的线程数
		MaxActive:   0,                 // 线程池的最大连接数， 0表示没有限制
		Wait:        true,              // 当连接数已满，是否要阻塞等待获取连接。false表示不等待，直接返回错误。
		IdleTimeout: 200 * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) { // 创建链接
			c, err := redis.Dial("tcp", conf["redis_host"])
			if err != nil {
				return nil, err
			}
			if conf["redis_password"] != "" {
				if _, err := c.Do("AUTH", conf["redis_password"]); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", conf["redis_dbNum"]); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //一个测试链接可用性
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return RedisPool
}

func closePool() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		RedisPool.Close()
		// os.Exit(0)
	}()
}

// get
func GetKey(key string) (string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.String(rds.Do("GET", key))
}

// set expires为0时，表示永久性存储
func SetKey(key string, value interface{}, expires int64) error {
	rds := RedisPool.Get()
	defer rds.Close()
	if expires == 0 {
		_, err := rds.Do("SET", key, value)
		return err
	} else {
		_, err := rds.Do("SETEX", key, expires, value)
		return err
	}
}

// del
func DelKey(key string) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("DEL", key)
	return err
}

// lrange
func LRange(key string, start, stop int64) ([]string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Strings(rds.Do("LRANGE", key, start, stop))
}

// lpop
func LPop(key string) (string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.String(rds.Do("LPOP", key))
}

// LPush
func LPush(key, value interface{}) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("LPUSH", key, value)
	return err
}

// LPushAndTrimKey
func LPushAndTrimKey(key, value interface{}, size int64) error {
	rds := RedisPool.Get()
	defer rds.Close()
	rds.Send("MULTI")
	rds.Send("LPUSH", key, value)
	rds.Send("LTRIM", key, size-2*size, -1)
	_, err := rds.Do("EXEC")
	return err
}

// RPushAndTrimKey
func RPushAndTrimKey(key, value interface{}, size int64) error {
	rds := RedisPool.Get()
	defer rds.Close()
	rds.Send("MULTI")
	rds.Send("RPUSH", key, value)
	rds.Send("LTRIM", key, size-2*size, -1)
	_, err := rds.Do("EXEC")
	return err

}

// ExistsKey
func ExistsKey(key string) (bool, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Bool(rds.Do("EXISTS", key))
}

// ttl 返回剩余时间
func TTLKey(key string) (int64, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("TTL", key))
}

func ExpireKey(key string, expires int) (bool, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Bool(rds.Do("EXPIRE", key, expires))
}

// incr 自增
func Incr(key string) (int64, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("INCR", key))
}

// Decr 自减
func Decr(key string) (int64, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("DECR", key))
}

// mset 批量写入 rds.Do("MSET", "ket1", "value1", "key2","value2")
func MsetKey(key_value ...interface{}) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("MSET", key_value...)
	return err
}

// mget  批量读取 mget key1, key2, 返回map结构
func MgetKey(keys ...interface{}) map[interface{}]string {
	rds := RedisPool.Get()
	defer rds.Close()
	values, _ := redis.Strings(rds.Do("MGET", keys...))
	resultMap := map[interface{}]string{}
	keyLen := len(keys)
	for i := 0; i < keyLen; i++ {
		resultMap[keys[i]] = values[i]
	}
	return resultMap
}

// hmset 同时将多个 field-value (域-值)对设置到哈希表 key 中。
func HMsetKey(key string, simpleObject interface{}) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("HMSET", redis.Args{}.Add(key).AddFlat(simpleObject)...)
	return err
}

// hmget 返回哈希表 key 中，一个或多个给定域的值
func HMgetKey(key string, simpleObject interface{}) error {
	rds := RedisPool.Get()
	defer rds.Close()
	values, _ := redis.Values(rds.Do("HGETALL", key))

	// hashV, _ := redis.StringMap(rds.Do("hgetall", key))
	// fmt.Println(hashV)
	// map[author:bgbiao description:my blog url:http://xxbandy.github.io]
	// hashV2, _ := redis.Strings(rds.Do("hmget", key, "description", "url", "author"))
	// for _, hashv := range hashV2 {
	// 	fmt.Println(hashv)
	// my blog
	// http://xxbandy.github.io
	// bgbiao
	// }

	err := redis.ScanStruct(values, simpleObject)
	if err != nil {
		return err
	}
	return err
}
