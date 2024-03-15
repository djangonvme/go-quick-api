package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var Instance *redis.Client

func InitRedis(addr, pwd string, db, poolSize int) error {
	if Instance != nil {
		return nil
	}
	if poolSize == 0 {
		poolSize = 100
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,      // no password set
		DB:       db,       // use default DB
		PoolSize: poolSize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	Instance = rdb
	return nil
}

/*func (rs *Client) Lock(key string, randValue string, exp time.Duration) (bool, error) {
	ctx := context.Background()
	return rs.SetNX(ctx, key, randValue, exp).Result()
}

func (rs *Client) Unlock(key string, randValue string) (bool, error) {
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
*/
