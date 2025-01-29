package cache

import (
    "context"
    // "fmt"
    "time"
    "github.com/go-redis/redis/v8"
)

var (
    rdb *redis.Client
    ctx = context.Background()
)

func InitRedis(addr string) error {
    rdb = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := rdb.Ping(ctx).Result()
    return err
}

func Set(key string, value interface{}, expiration time.Duration) error {
    return rdb.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
    return rdb.Get(ctx, key).Result()
}

func Delete(key string) error {
    return rdb.Del(ctx, key).Err()
}