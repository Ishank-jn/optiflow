package cache

import (
    "context"
    "time"
    "github.com/go-redis/redis/v8"
    "optiflow/internal/logger"
)

var (
    rdb *redis.Client
    ctx = context.Background()
)

func InitRedis(addr string) error {
    logger.Info("Initializing Redis client")
    rdb = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        logger.Error("Failed to connect to Redis: " + err.Error())
        return err
    }
    logger.Info("Successfully connected to Redis")
    return nil
}

func Set(key string, value interface{}, expiration time.Duration) error {
    logger.Info("Setting key: " + key)
    err := rdb.Set(ctx, key, value, expiration).Err()
    if err != nil {
        logger.Error("Failed to set key: " + key + " Error: " + err.Error())
    } else {
        logger.Info("Successfully set key: " + key)
    }
    return err
}

func Get(key string) (string, error) {
    logger.Info("Getting key: " + key)
    result, err := rdb.Get(ctx, key).Result()
    if err != nil {
        logger.Error("Failed to get key: " + key + " Error: " + err.Error())
        return "", err
    }
    logger.Info("Successfully got key: " + key)
    return result, nil
}

func Delete(key string) error {
    logger.Info("Deleting key: " + key)
    err := rdb.Del(ctx, key).Err()
    if err != nil {
        logger.Error("Failed to delete key: " + key + " Error: " + err.Error())
    } else {
        logger.Info("Successfully deleted key: " + key)
    }
    return err
}