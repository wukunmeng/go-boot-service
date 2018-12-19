/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午3:23
 * ---------------------------------
 * 
 */
package service

import (
    "github.com/labstack/echo"
    "net/http"
    "fmt"
    "github.com/wukunmeng/go-boot-service/version"
)

func router(e *echo.Echo) {
    e.GET("/", info)

    {
        g := e.Group("/api")
        {
            g.GET("/", info)
        }

        {
            g.GET("/", info)
        }
    }
}

func info(c echo.Context) error {
    return c.String(http.StatusOK, fmt.Sprintf("%v ver:%v, buildTime:%v", version.ServerName, version.Build, version.BuildTime))
}
