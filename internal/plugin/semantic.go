package plugin

import (
	"log"
	"system/internal/database"
	"system/internal/service"
)


func (s *SentenceAnalyze) connect() {
	log.Println("Connecting to database...")
}

// 匹配 `SELECT * FROM <table>` 的形式
func (s *SentenceAnalyze) readAll(tableName string) {
	goods, err := database.SafeSelectGoods_Key(tableName, "")
	if err != nil {
		log.Println("Failed to select all goods:", err)
	}
	for _, good := range goods {
		log.Printf("Good: %+v\n", good)
	}
}

// 匹配 `SELECT <idList> FROM <table>` 的形式
func (s *SentenceAnalyze) read(columns []string, tableName string) {
	log.Printf("Reading columns %v from table: %s\n", columns, tableName)

	// 使用列名进行查询
	goods, err := database.SafeSelectGoods_Column_Key(tableName, columns[0], "")
	if err != nil {
		log.Println("Failed to select goods by columns:", err)
	}
	for _, good := range goods {
		log.Printf("Good: %+v\n", good)
	}
}

func (s *SentenceAnalyze) update(tableName string, columnName string, value string, orderNumber string) {
	log.Printf("Updating table %s, setting %s to %s where OrderNumber is %s\n", tableName, columnName, value, orderNumber)
	err := database.SafeUpdateGoods(tableName, columnName, value, orderNumber)
	if err != nil {
		log.Println("Failed to update goods:", err)
	}
}

func (s *SentenceAnalyze) delete(tableName string, columnName string, value string) {
	log.Printf("Deleting from table %s where %s = %s\n", tableName, columnName, value)
	err := database.SafeDeleteGoods(tableName, columnName, value)
	if err != nil {
		log.Println("Failed to delete goods:", err)
	}
}

func (s *SentenceAnalyze) insert(tableName string, columns []string, values []string) {
	log.Printf("Inserting into table %s, columns %v, values %v\n", tableName, columns, values)
	good := service.Good{
		Number:      values[0],
		Time:        values[1],
		Destination: values[2],
		IsVip:       values[3] == "true",
	}
	err := database.SafeInsertGoods(tableName, good)
	if err != nil {
		log.Println("Failed to insert good:", err)
	}
}