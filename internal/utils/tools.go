package utils

import (
    "strconv"
    "log"
    "strings"
)

// 代价矩阵解析时用，将字符串解析为float32
func ParseFloat32(s []string) []float32{
    res := make([]float32, 0)
    for _, t:= range(s){
        f, err := strconv.ParseFloat(t, 32)
        if err != nil {
            log.Printf("Invalid Parsing Float:%s" ,err)
            return nil
        }
        res = append(res, float32(f))
    }
    return res
}

// 获得目的地的index
func IndexDestination(names []string) map[string]int{
    mp := make(map[string]int)
    for idx, name := range(names){
        mp[name] = idx
    }
    return mp
}

// SQL字符串预处理
func PreParseSQLs(sqls string) []string{
    sqls = strings.ToLower(sqls)
    sentences := strings.Split(sqls, ";")
    return sentences
}