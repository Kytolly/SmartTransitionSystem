package service

import ( 
	"strconv"
	"log"
    "sort"
)

func TimeCompare(l * Good, r* Good) bool{
	lt, lerr := strconv.ParseInt(l.GetTime(), 10, 64)
	rt, rerr := strconv.ParseInt(r.GetTime(), 10, 64)
	if lerr != nil || rerr != nil {
		log.Println("Invalid Arrival Time!")
		return false
	}
	if (l.IsVip) {
		lt -= 10000
	}
	if (r.IsVip) {
		rt -= 10000
	} 
	return lt > rt
} 
// 设置排序的接口
type ByTime []*Good
func (t ByTime) Len() int           { return len(t) }
func (t ByTime) Less(i, j int) bool { return TimeCompare(t[i], t[j]) }
func (t ByTime) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }


func Sort(targetGoods ByTime){
    sort.Sort(targetGoods)
}