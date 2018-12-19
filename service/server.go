/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午3:08
 * ---------------------------------
 * 
 */
package service

import (
    "github.com/wukunmeng/go-boot-service/log/logger"
    "github.com/wukunmeng/go-boot-service/config"
    "github.com/wukunmeng/go-boot-service/version"
    "github.com/labstack/echo"
    "go.uber.org/zap"
    "os"
    "net/http"
    _ "net/http/pprof"
    "fmt"
    "time"
    "github.com/facebookgo/grace/gracehttp"
    "github.com/labstack/echo/middleware"
    "github.com/wukunmeng/go-boot-service/service/response"
)

func Serve() {
    cfg := config.Get()

    logger.Info("server info",
        zap.String("appVersion", version.Build),
        zap.String("appBuildTime", version.BuildTime),
    )

    logger.Info("server start",
        zap.String("host", cfg.Server.Host), zap.Int("port", cfg.Server.Port),
        zap.Int("pid", os.Getpid()),
    )

    Before()
    go BeforeAsync()

    e := echo.New()
    e.Use(middleware.Recover())

    e.HTTPErrorHandler = httpErrorHandler

    router(e)

    e.Group("/debug/*", func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
        return func(context echo.Context) error {
            http.DefaultServeMux.ServeHTTP(context.Response(), context.Request())
            return nil
        }
    })

    server := &http.Server{
        Addr:        fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
        IdleTimeout: time.Second * 60,
        Handler:     e,
    }

    err := gracehttp.Serve(server)
    if err != nil {
        logger.Fatal("start server fail", zap.Error(err))
    }
}

func httpErrorHandler(err error, c echo.Context) {
    logger.Warn("http error handler",
        zap.Error(err),
        zap.Any("url", c.Request().URL),
        zap.Any("host", c.Request().Host),
        zap.Any("header", c.Request().Header),
        zap.Any("ip", c.RealIP()),
    )
    if c.Response().Committed {
        return
    }

    var (
        code = http.StatusBadRequest
        msg  interface{}
    )

    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
        msg = he.Message
    } else {
        msg = err.Error()
    }

    if v, ok := msg.(string); ok {
        response.Error(c, string(code), v)
    }
}
