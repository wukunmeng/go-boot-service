/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午4:06
 * ---------------------------------
 * 
 */
package model

import (
    "sync"
    "github.com/wukunmeng/go-boot-service/config"
    "github.com/jinzhu/gorm"
    "github.com/wukunmeng/go-boot-service/log/logger"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "go.uber.org/zap"
    "time"
    "regexp"
    "strings"
    "fmt"
    "github.com/wukunmeng/go-boot-service/log/sugar"
)

var (
    db     *gorm.DB
    dbOnce sync.Once
)

func initDB() {
    cfg := config.Get().Database
    var err error
    db, err = gorm.Open("mysql", cfg.URI)
    if err != nil {
        logger.Fatal("open database fail", zap.Any("error", err))
    }
    sugar.Infof("init database %s", cfg.URI)
    db.Callback().Query().Register("printTableName", dbPlugin)
    db.Callback().Create().After("gorm:create").Register("afterCreateRecord", dbAfterCreatePlugin)
    sqlDB := db.DB()
    sqlDB.SetMaxOpenConns(cfg.MaxConn)
    sqlDB.SetMaxIdleConns(1)
    sqlDB.SetConnMaxLifetime(time.Minute * 5)

    db.SingularTable(true)
    db.SetLogger(&dbLog{logger.WithOptions(zap.AddCallerSkip(1))})
    db.LogMode(true)
}

func DB() *gorm.DB {
    dbOnce.Do(initDB)
    return db
}

func CloseDB(){
    cfg := config.Get().Database
    sugar.Infof("close db %s", cfg.URI)
    if db != nil {
        db.Close() //关闭数据库
    }
}

var (
    re = regexp.MustCompile(`(?m)sql\s+.*?scope\.go\:\d+`)
)

type dbLog struct {
    l *zap.Logger
}

func (lg *dbLog) Print(v ...interface{}) {
    lg.l.Debug(lg.tidySQLLog(strings.TrimSpace(fmt.Sprintln(v...))))
}

func (lg *dbLog) tidySQLLog(log string) string {
    return re.ReplaceAllLiteralString(log, "sql")
}

func dbPlugin(scope *gorm.Scope)  {
    fmt.Println(fmt.Sprintf("query %v", scope.TableName()))
}

func dbAfterCreatePlugin(scope *gorm.Scope)  {
    fmt.Println(fmt.Sprintf("create %v", scope.TableName()))
}