/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-14
 * Time : 下午5:36
 * ---------------------------------
 * 
 */
package zap

import (
    "github.com/wukunmeng/go-boot-service/log"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    stdlog "log"
    "os"
)

var (
    Logger *zap.Logger
    Sugar  *zap.SugaredLogger
)

func init() {
    cfg := log.Log{
        Level: "debug",
    }
    ConfigZap(cfg)
}

func ConfigZap(cfg log.Log) {
    var err error

    ws := make([]zapcore.WriteSyncer, 0, 2)
    ws = append(ws, zapcore.AddSync(os.Stdout))

    if cfg.File != "" {
        ws = append(ws, zapcore.AddSync(&lumberjack.Logger{
            Filename:  cfg.File,
            MaxSize:   cfg.MaxSize,
            MaxAge:    cfg.MaxDays,
            LocalTime: true,
            Compress:  true,
        }))
    }

    var level zapcore.Level
    err = level.UnmarshalText([]byte(cfg.Level))
    if err != nil {
        stdlog.Fatal(err)
    }

    writeSyncer := zapcore.NewMultiWriteSyncer(ws...)

    encodingCfg := zap.NewProductionEncoderConfig()
    encodingCfg.EncodeTime = zapcore.ISO8601TimeEncoder

    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(encodingCfg),
        //zapcore.NewJSONEncoder(encodingCfg),
        writeSyncer,
        level,
    )

    options := make([]zap.Option, 0, 3)
    options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
    if level.Enabled(zapcore.DebugLevel) {
        options = append(options, zap.AddCaller(), zap.AddCallerSkip(1))
    }
    Logger = zap.New(core, options...)
    if err != nil {
        stdlog.Fatal(err)
    }
    Sugar = Logger.Sugar()
}