/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-14
 * Time : 下午5:46
 * ---------------------------------
 * 
 */
package sugar

import (
    z "github.com/wukunmeng/go-boot-service/log/zap"
    "go.uber.org/zap"
)

func Named(name string) *zap.SugaredLogger {
    return z.Sugar.Named(name)
}

func With(args ...interface{}) *zap.SugaredLogger {
    return z.Sugar.With(args...)
}

func Debug(args ...interface{}) {
    z.Sugar.Debug(args...)
}

func Info(args ...interface{}) {
    z.Sugar.Info(args...)
}

func Warn(args ...interface{}) {
    z.Sugar.Warn(args...)
}

func Error(args ...interface{}) {
    z.Sugar.Error(args...)
}

func DPanic(args ...interface{}) {
    z.Sugar.DPanic(args...)
}

func Panic(args ...interface{}) {
    z.Sugar.Panic(args...)
}

func Fatal(args ...interface{}) {
    z.Sugar.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
    z.Sugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
    z.Sugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
    z.Sugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
    z.Sugar.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
    z.Sugar.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
    z.Sugar.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
    z.Sugar.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
    z.Sugar.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
    z.Sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
    z.Sugar.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
    z.Sugar.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
    z.Sugar.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
    z.Sugar.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
    z.Sugar.Fatalw(msg, keysAndValues...)
}
