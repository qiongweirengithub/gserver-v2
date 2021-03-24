package dbs

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
    USER_NAME = "pingpongus_test"
    PASS_WORD = "pingpongus_test"
    HOST      = "47.112.159.211"
    PORT      = "3306"
    DATABASE  = "pingpongus_test"
    CHARSET   = "utf8"
)

// 初始化链接
func init() {

    dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

    // 打开连接失败
    MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
    //defer MysqlDb.Close();
    if MysqlDbErr != nil {
        fmt.Println("dbDSN: " + dbDSN)
        panic("数据源配置不正确: " + MysqlDbErr.Error())
    }

    // 最大连接数
    MysqlDb.SetMaxOpenConns(100)
    // 闲置连接数
    MysqlDb.SetMaxIdleConns(20)
    // 最大连接周期
    MysqlDb.SetConnMaxLifetime(100*time.Second)

    if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
        panic("数据库链接失败: " + MysqlDbErr.Error())
    }

    fmt.Println("init sql success")

}

