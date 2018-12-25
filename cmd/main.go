/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-13
 * Time : 下午2:22
 * ---------------------------------
 * 
 */
package main

import (
    "flag"
    "fmt"
    "github.com/wukunmeng/go-boot-service/version"
    z "github.com/wukunmeng/go-boot-service/log/zap"
    "runtime"
    "github.com/wukunmeng/go-boot-service/service"
    "github.com/wukunmeng/go-boot-service/config"
)

const appName = "serverName"

var(
    V = flag.Bool("version", false, "show version")
)

func main()  {
    flag.Parse()
    if !flag.Parsed() {
        flag.Usage()
        return
    }

    if *V {
        fmt.Println(appName)
        fmt.Println("     version:", version.Build)
        fmt.Println("  build time:", version.BuildTime)
        fmt.Println("  go version:", runtime.Version())
        return
    }

    cfg := config.Get()
    z.ConfigZap(cfg.Log)

    fmt.Printf(`
:'######::::'#######::'##::::::::::'###::::'##::: ##::'######:::::::::::::'######::'########:'########::'##::::'##:'####::'######::'########:
'##... ##::'##.... ##: ##:::::::::'## ##::: ###:: ##:'##... ##:::::::::::'##... ##: ##.....:: ##.... ##: ##:::: ##:. ##::'##... ##: ##.....::
 ##:::..::: ##:::: ##: ##::::::::'##:. ##:: ####: ##: ##:::..:::::::::::: ##:::..:: ##::::::: ##:::: ##: ##:::: ##:: ##:: ##:::..:: ##:::::::
 ##::'####: ##:::: ##: ##:::::::'##:::. ##: ## ## ##: ##::'####:'#######:. ######:: ######::: ########:: ##:::: ##:: ##:: ##::::::: ######:::
 ##::: ##:: ##:::: ##: ##::::::: #########: ##. ####: ##::: ##::........::..... ##: ##...:::: ##.. ##:::. ##:: ##::: ##:: ##::::::: ##...::::
 ##::: ##:: ##:::: ##: ##::::::: ##.... ##: ##:. ###: ##::: ##:::::::::::'##::: ##: ##::::::: ##::. ##:::. ## ##:::: ##:: ##::: ##: ##:::::::
. ######:::. #######:: ########: ##:::: ##: ##::. ##:. ######::::::::::::. ######:: ########: ##:::. ##:::. ###::::'####:. ######:: ########:
:......:::::.......:::........::..:::::..::..::::..:::......::::::::::::::......:::........::..:::::..:::::...:::::....:::......:::........::                                                  
                                      
application name: %s  ver: %s

`, appName, version.Build)
service.Serve()
}