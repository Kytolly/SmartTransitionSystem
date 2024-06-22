package service

import ( 
	// "strconv"
	// "errors"
	// "database/sql" 
	// "system/internal/database"
	"encoding/csv"
	"io"
	"log"
	"os"
	"fmt" 
)

// 物品信息
type Good struct{
    Number 		string // 订单号
	Time   		string // 到达起点的时间
	IsVip  		bool   // 寄件人是否为高级用户
	Destination string // 目的地
	Description string // 物流方案
	Road 		string // 物流路径组成
}

func(g *Good) SetInfo(
	Number  	string, 
	Time 		string, 
	IsVip 		bool, 
	Destination string, 
	Description string,
) {
	g.Number = Number
	g.Time = Time
	g.IsVip = IsVip
	g.Destination = Destination
	g.Description = Description
}

func(g *Good) GetNumber() string{
	return g.Number
}

func(g *Good) GetTime() string{
	return g.Time
}

func(g *Good) GetIsVip() bool{
	return g.IsVip
}

func(g *Good) GetDescription() string{
	return g.Description
}

func(g *Good) GetRoad(gra *Graph, prev []int) string{
	return gra.FindRoad(g.Destination, prev)
}

func(g *Good) Display(gra *Graph, prev []int){
	IsVip := ""
	if g.IsVip{
		IsVip = "Vip用户"
	}else{
		IsVip = "非Vip用户"
	} 
	Road := g.GetRoad(gra, prev )
	fmt.Printf("[%s %s %s %s %s]", g.Number, g.Time, IsVip, g.Description, Road)
}

func GetGoods(items string) []Good{
	var goods []Good
	// 读取csv文件
    file, err := os.Open(items)
    if err != nil {
        log.Printf("Fail Loading CSV File %s: %s", items, err)
        return nil
    }
    defer file.Close()
    reader := csv.NewReader(file)

    // 读取表头
    _, err = reader.Read()
    if err != nil { 
        log.Printf("Failed Loading Table Head from %s : %s", items, err)
        return nil
    }
    // log.Println(g.Names)

    // 读取物流节点之间的代价信息
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil { 
            log.Printf("Failed fetching Information from %s : %s", items, err)
            return nil
        } 
		goods = append(goods, Good{
			Number: 	record[0],
			Time:    	record[1],
			IsVip:	 	(record[4]=="是"),
			Destination:record[2],
			Description:record[3],
		})
		// log.Println(goods)
    }
	return goods
}


// 获得指定方案的货品下标
func getIndex(goods []*Good, label string)[]int{
	res := make([]int, 0)
	for idx, good := range goods{
		if good.Description == label{
			res = append(res, idx)
		}
	}
	return res
}

// 发货函数：从物品列表中删除物品
func shipGood(goods *[]Good, goodIndex int) {
    *goods = append((*goods)[:goodIndex], (*goods)[goodIndex+1:]...)
} 

func Display(goods []*Good, gra *Graph, prev []int){
	for _,good := range goods{
		good.Display(gra, prev)
	}
}

