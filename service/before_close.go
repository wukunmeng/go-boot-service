/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-20
 * Time : 下午2:56
 * ---------------------------------
 * 
 */
package service

import "github.com/wukunmeng/go-boot-service/service/model"

//服务关闭前执行
func BeforeServerClosed()  {
    //关闭数据库
    model.CloseDB()
}