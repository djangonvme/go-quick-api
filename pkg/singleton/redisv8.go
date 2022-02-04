package singleton

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/task-dispatcher/config"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

type RedisClient struct {
	*redis.Client
}

func NewRedis(cfg *config.Config) (cli *RedisClient, err error) {
	return newRedis(cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.DbNum, 1000)
}
func newRedis(host string, pwd string, dbNum int, poolSize int) (cli *RedisClient, err error) {
	if poolSize == 0 {
		poolSize = 100
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", host),
		Password: pwd,      // no password set
		DB:       dbNum,    // use default DB
		PoolSize: poolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return &RedisClient{
		rdb,
	}, err
}

func (rs *RedisClient) Lock(key string, randValue string, exp time.Duration) (bool, error) {
	ctx := context.Background()
	return rs.SetNX(ctx, key, randValue, exp).Result()
}

func (rs *RedisClient) Unlock(key string, randValue string) (bool, error) {
	ctx := context.Background()
	cmd := rs.Get(ctx, key)
	value, err := cmd.Result()
	if err != nil {
		return false, err
	}
	if randValue != "" && randValue != value {
		return false, errors.Errorf("couldn't unlock %s, the key maybe locked by others", key)
	}
	res := rs.Del(ctx, key)
	ok, err := res.Result()
	if err != nil {
		return false, err
	}
	if ok == 1 {
		return true, nil
	}
	return false, nil

}
