package main

import (
	"log"
	// "system/internal/service"
	// "system/internal/database"
	// "system/internal/utils"
	"system/internal/plugin"
)

func main() {
	// log.Println("Hello, World!")
	// G := service.NewGraph()
	// G.SetInfo("internal/database/csv/cost_matrix.csv")
	// mygoods := service.GetGoods("internal/database/csv/item.csv")

	// log.Println("Hello, World!")
	// database.InitDB()
	// for _, good := range mygoods {
	//     err := database.SafeInsertGoods("goodslist1", good)
	//     if err != nil {
	//         log.Printf("Failed to insert good: %v", err)
	//     }
	// }

	log.Println("Hello, World!")
	// 测试SQL解析器
	//sqls := "CREATE TABLE employees (id INT PRIMARY KEY,name VARCHAR(100),age INT,department VARCHAR(50));INSERT INTO employees (id, name, age, department) VALUES (1, 'Alice', 30, 'HR'),(2, 'Bob', 25, 'Engineering'),(3, 'Charlie', 28, 'Marketing');SELECT name, age FROM employees WHERE department = 'Engineering';UPDATE employees SET age = 31 WHERE id = 1;"
	sqls := "SELECT * FROM goodlist1"
	plugin.ParseAndExecuteSQL(sqls)
}
