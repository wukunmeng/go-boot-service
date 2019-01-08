/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午3:46
 * ---------------------------------
 * 
 */
package service

import "github.com/wukunmeng/go-boot-service/log/sugar"

//服务启动前同步执行
func Before()  {
    sugar.Infof("before:%v", "服务启动前同步执行")

}

//服务启动前异步执行
func BeforeAsync()  {
    sugar.Infof("before:%v", "服务启动异步执行脚本")
}