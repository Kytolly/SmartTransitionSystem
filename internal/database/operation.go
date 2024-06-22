package database

import (
    // "database/sql"
    "log"
    "errors"
    "system/internal/service"
    "system/internal/utils"
)


// 数据库操作函数

// 插入记录
func Insert(tableName string, good service.Good) error {
    query := `INSERT INTO ` + tableName + ` (OrderNumber, AriveTime, Destination, isVip, isSend) VALUES (?, ?, ?, ?, ?)`
    _, err := DB.Exec(query, good.Number, good.Time, good.Destination, good.IsVip, "False")
    return err
}

// 更新记录 
func Update(tableName, columnName, value string, orderNumber string) error {
    query := `UPDATE ` + tableName + ` SET ` + columnName + ` = ? WHERE OrderNumber = ?`
    _, err := DB.Exec(query, value, orderNumber)
    return err
} 

// 删除记录
func Delete(tableName, columnName, value string) error {
    query := `DELETE FROM ` + tableName + ` WHERE ` + columnName + ` = ?`
    _, err := DB.Exec(query, value)
    return err
}

// 关键字查询
func Select_Key(tableName, key string) ([]service.Good, error) {
    query := `SELECT * FROM ` + tableName + ` WHERE OrderNumber LIKE ?`
    rows, err := DB.Query(query, "%"+key+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var goodsList []service.Good
    for rows.Next() {
        var goods service.Good
        issend := "False"
        if err := rows.Scan(&goods.Number, &goods.Time, &goods.Destination, &goods.IsVip, &issend); err != nil {
            return nil, err
        }
        goodsList = append(goodsList, goods)
    }
    return goodsList, nil
}


func Select_Column_Key(tableName, columnName, keyword string) ([]service.Good, error) {
    query := `SELECT Number, Time, Destination, IsVip FROM ` + tableName + ` WHERE ` + columnName + ` = ?`
    rows, err := DB.Query(query, keyword)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var goodsList []service.Good
    for rows.Next() {
        var goods service.Good
        issend := "False"
        if err := rows.Scan(&goods.Number, &goods.Time, &goods.Destination, &goods.IsVip, &issend); err != nil {
            return nil, err
        }
        goodsList = append(goodsList, goods)
    }
    return goodsList, nil
}

// 实现了互斥机制的更新，插入(写)，查询(读)
func SafeInsertGoods(tableName string, good service.Good) error {
    if !utils.StartWrite() {
        return errors.New("another process is writing")
    }
    defer utils.EndWrite()

    err := Insert(tableName, good)
    if err != nil {
        log.Println("Failed to insert goods:", err)
    }
    return err
}

func SafeSelectGoods_Key(tableName, key string) ([]service.Good, error) {
    if !utils.StartRead() {
        return nil, errors.New("another process is writing")
    }
    defer utils.EndRead()

    goodsList, err := Select_Key(tableName, key)
    if err != nil {
        log.Println("Failed to select goods:", err)
    }
    return goodsList, err
}

func SafeUpdateGoods(tableName, columnName, value, orderNumber string) error {
    if !utils.StartWrite() {
        return errors.New("another process is writing")
    }
    defer utils.EndWrite()

    err := Update(tableName, columnName, value, orderNumber)
    if err != nil {
        log.Println("Failed to update goods:", err)
    }
    return err
}

func SafeDeleteGoods(tableName, columnName, value string) error {
    if !utils.StartWrite() {
        return errors.New("another process is writing")
    }
    defer utils.EndWrite()

    err := Delete(tableName, columnName, value)
    if err != nil {
        log.Println("Failed to delete goods:", err)
    }
    return err
}

func SafeSelectGoods_Column_Key(tableName, columnName, keyword string) ([]service.Good, error) {
    if !utils.StartRead() {
        return nil, errors.New("another process is writing")
    }
    defer utils.EndRead()

    goodsList, err := Select_Column_Key(tableName, columnName, keyword)
    if err != nil {
        log.Println("Failed to select goods:", err)
    }
    return goodsList, err
}