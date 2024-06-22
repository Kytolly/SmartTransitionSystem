// database.go
package database

import (
    "database/sql"
    _"github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "root:xqy050116@tcp(localhost:3306)/wuliu")
    if err != nil {
        log.Fatal(err)
    }

    // 创建数据库
    createTableQueries := []string{
        `CREATE TABLE IF NOT EXISTS goodslist1 (
            OrderNumber VARCHAR(255),
            AriveTime VARCHAR(255),
            Destination VARCHAR(255),
            isVip VARCHAR(255),
			isSend VarCHAR(255)
        );`,
        `CREATE TABLE IF NOT EXISTS goodslist2 (
            OrderNumber VARCHAR(255),
            AriveTime VARCHAR(255),
            Destination VARCHAR(255),
            isVip VARCHAR(255),
			isSend VarCHAR(255)
        );`,
    }

    for _, query := range createTableQueries {
        if _, err := DB.Exec(query); err != nil {
            log.Fatal(err)
        }
    }
}