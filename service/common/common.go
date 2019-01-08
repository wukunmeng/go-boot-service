/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 19-1-8
 * Time : 下午2:21
 * ---------------------------------
 * 
 */
package common

import (
    "strconv"
)

//字符串转int
func ParseInt(source string, defaultValue int) int {
    if r, err := strconv.Atoi(source); err == nil {
        return r
    }
    return defaultValue
}