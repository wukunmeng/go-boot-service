/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-14
 * Time : 下午5:35
 * ---------------------------------
 * 
 */
package log

type Log struct {
    Level   string `toml:"level"`
    File    string `toml:"file"`
    MaxSize int    `toml:"max-size"`
    MaxDays int    `toml:"max-days"`
}