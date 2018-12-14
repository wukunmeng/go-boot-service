/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-14
 * Time : 下午5:43
 * ---------------------------------
 * 
 */
package logger

import (
    z "github.com/wukunmeng/go-boot-service/log/zap"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func Named(s string) *zap.Logger {
    return z.Logger.Named(s)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
    return z.Logger.WithOptions(opts...)
}

func With(fields ...zap.Field) *zap.Logger {
    return z.Logger.WithOptions(zap.AddCallerSkip(-1)).With(fields...)
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
    return z.Logger.Check(lvl, msg)
}

func Debug(msg string, fields ...zap.Field) {
    z.Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
    z.Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
    z.Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    z.Logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
    z.Logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
    z.Logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
    z.Logger.Fatal(msg, fields...)
}
