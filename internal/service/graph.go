package service

import (
    "encoding/csv"
    "os"
    "io"
    "log"
    "system/internal/utils"
    "math"
)

type Graph struct{
    CityNum  int            // 节点数
    Cost     [][]float32    // 节点与节点之间关系的邻接矩阵
    Names    []string       // 节点名
    DesIndex map[string]int // 目的地对应下标
}

func NewGraph() *Graph {
    return &Graph{
        CityNum: 0,
        Cost:    make([][]float32, 0),
        Names:   make([]string, 0),
        DesIndex: make(map[string]int, 0),
    }
} 

func(g *Graph)SetInfo(cost_matrix string){
    // 读取csv文件
    file, err := os.Open(cost_matrix)
    if err != nil {
        log.Printf("Fail Loading CSV File %s: %s", cost_matrix, err)
        return
    }
    defer file.Close()
    reader := csv.NewReader(file)

    // 读取表头
    record, err := reader.Read()
    if err != nil { 
        log.Printf("Failed Loading Table Head from %s : %s", cost_matrix, err)
        return
    }
    g.CityNum = len(record)
    g.Names = append(g.Names, record...)
    for idx, city:= range record{
        g.DesIndex[city] = idx
    }
    // log.Println(g.Names)

    // 读取物流节点之间的代价信息
    for {
        record, err = reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil { 
            log.Printf("Failed Setting Information from %s : %s", cost_matrix, err)
            return
        } 
        newcost := utils.ParseFloat32(record)
        g.Cost = append(g.Cost, newcost)
        // log.Println(g.Cost)
    }
}

// 计算name[0]的单源最短路径
func (g *Graph) Dijkstra(prev []int, dist []float32) {
    vs := 0
    n := g.CityNum
    flag := make([]bool, n)

    for i := 0; i < n; i++ {
        flag[i] = false
        prev[i] = -1 // 无前置结点
        dist[i] = g.Cost[vs][i]
    }
    flag[vs] = true
    dist[vs] = 0

    for i := 1; i < n; i++ {
        min := float32(math.MaxFloat32)
        k := -1
        for j := 0; j < n; j++ {
            if !flag[j] && dist[j] < min {
                min = dist[j]
                k = j
            }
        }
        if k == -1 {
            break 
        }
        flag[k] = true

        for j := 0; j < n; j++ {
            if !flag[j] && g.Cost[k][j] < math.MaxFloat32 {
                newDist := dist[k] + g.Cost[k][j]
                if newDist < dist[j] {
                    dist[j] = newDist
                    prev[j] = k
                }
            }
        }
    }
}

func (g *Graph) FindRoad(destination string, prev [] int)string{
    now := g.DesIndex[destination]
    res := ""
    for {
        if now == -1{
            break
        }
        res = res + "-" + g.Names[now] 
    }
    res = g.Names[now] + res
    return res
}