/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-14
 * Time : 下午2:52
 * ---------------------------------
 * 
 */
package config

import (
    "github.com/BurntSushi/toml"
    stdlog "log"
    "runtime/debug"
    "sync"
    "github.com/wukunmeng/go-boot-service/log"
)

type config struct {
    Server struct {
        Host string `toml:"host"`
        Port int    `toml:"port"`
    }

    Database struct {
        URI     string `toml:"uri"`
        MaxConn int    `toml:"max-conn"`
    }

    Redis struct {
        Addr     string `toml:"addr"`
        Password string `toml:"password"`
        Db       int    `toml:"db"`
        PoolSize int    `toml:"pool-size"`
        CachePrefix string `toml:"cache-prefix"`
    }

    Host struct {
        AvatarPrefix string `toml:"avatar-prefix"`
        BuildActivityHost string `toml:"build-activity-host"`
    }

    Expiration struct {
        RedisActivityUser int `toml:"redis-activity-user"`
    }

    Delay struct{
        StartTimeDelay int64 `toml:"start-time-delay"`
        EndTimeDelay int64 `toml:"end-time-delay"`
    }

    Log log.Log
}

var (
    defaultConfig *config
    one           sync.Once
)

func Get() *config {
    if defaultConfig == nil {
        one.Do(func() {
            Load("config/config.toml")
        })
    }
    return defaultConfig
}

func Load(file string) {
    if defaultConfig != nil {
        panic("load only once")
    }

    var cfg config

    // default value
    cfg.Server.Port = 5002

    cfg.Redis.Addr = "127.0.0.1:6379"
    cfg.Redis.Db = 0
    cfg.Redis.PoolSize = 50

    cfg.Log.Level = "debug"
    cfg.Log.File = ""
    cfg.Log.MaxSize = 500 // MB
    cfg.Log.MaxDays = 90

    cfg.Expiration.RedisActivityUser = 60

    _, err := toml.DecodeFile(file, &cfg)
    if err != nil {
        debug.PrintStack()
        stdlog.Fatalln("config load file:", err)
    }
    defaultConfig = &cfg
}