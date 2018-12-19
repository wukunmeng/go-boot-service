/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午3:53
 * ---------------------------------
 * 
 */
package redisutil

import (
    "sync"
    "github.com/wukunmeng/go-boot-service/config"
    "github.com/go-redis/redis"
    "time"
    "github.com/wukunmeng/go-boot-service/log/logger"
    "go.uber.org/zap"
    "fmt"
    "strings"
)

var (
    client      *redis.Client

    redisOnce sync.Once
)

func initRedis() {
    cfg := config.Get().Redis
    client = redis.NewClient(&redis.Options{
        Addr:         cfg.Addr,
        Password:     cfg.Password,
        DB:           cfg.Db,
        ReadTimeout:  time.Second * 10,
        MaxRetries:   2,
        MinIdleConns: 1,
        PoolSize:     cfg.PoolSize,
    })

    err := client.Ping().Err()
    if err != nil {
        logger.Fatal("init redis fail", zap.Error(err))
    }
}

func Client() *redis.Client {
    redisOnce.Do(initRedis)
    return client
}

func Key(key ...string) string {
    cfg := config.Get().Redis
    return fmt.Sprintf("%s:%s", cfg.CachePrefix, strings.Join(key, ":"))
}