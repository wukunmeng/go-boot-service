/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午4:08
 * ---------------------------------
 * 
 */
package model

import "testing"

func TestDB(t *testing.T) {
    DB().DB().Ping()
}
