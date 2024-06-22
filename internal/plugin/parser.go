package plugin

import ( 
	"log"
	"regexp"
)

// SentenceAnalyze 结构体
type SentenceAnalyze struct {
	words []Word
	pos   int
}

// 创建新的 SentenceAnalyze
func NewSentenceAnalyze(words []Word) *SentenceAnalyze {
	return &SentenceAnalyze{words: words, pos: 0}
}

// match 函数
func (s *SentenceAnalyze) match(regex string) bool {
	if s.pos >= len(s.words) {
		return false
	}
	token := s.words[s.pos].Name
	matched, _ := regexp.MatchString(regex, token)
	if matched {
		s.pos++
		return true
	}
	return false
}

// id 函数
func (s *SentenceAnalyze) id() bool {
	if s.pos >= len(s.words) {
		return false
	}
	tokenType := s.words[s.pos].Type
	if tokenType == ID {
		s.pos++
		return true
	}
	return false
}

// idList 函数
func (s *SentenceAnalyze) idList() bool {
	if !s.id() {
		return false
	}
	for s.match(",") {
		if !s.id() {
			return false
		}
	}
	return true
}

func (s *SentenceAnalyze) checkSentence() {
	// 检查是否为 SELECT 语句
	if s.match("select") {
		// 匹配 `SELECT * FROM <table>` 的形式
		if s.match("\\*") && s.match("from") && s.id() {
			s.connect() // 连接数据库
			s.readAll(s.words[s.pos-1].Name) // 读取指定表的所有数据
		} else if s.idList() && s.match("from") && s.id() { 
			// 匹配 `SELECT <idList> FROM <table>` 的形式
			s.connect() 
			columns := make([]string, 0)
			tableName := s.words[s.pos-1].Name
			// 将匹配到的列名保存到 columns 列表中
			for i := 0; i < s.pos; i++ {
				if s.words[i].Type == ID {
					columns = append(columns, tableName)
				}
			}
			s.read(columns, s.words[s.pos-1].Name) // 读取指定列的数据
		} else {
			log.Println("grammar error")
		}
	} else if s.match("update") {
		// 检查是否为 UPDATE 语句
		if s.id() && s.match("set") && s.id() && s.match("=") && s.id() {
			s.connect() 
			// 更新指定表中指定列的值
			s.update(s.words[s.pos-5].Name, s.words[s.pos-3].Name, s.words[s.pos-1].Name, "")
		} else {
			log.Println("grammar error")
		}
	} else if s.match("delete") {
		// 检查是否为 DELETE 语句
		if s.match("from") && s.id() && s.match("where") && s.id() && s.match("=") && s.id() {
			s.connect() 
			// 从指定表中删除指定条件的数据
			s.delete(s.words[s.pos-5].Name, s.words[s.pos-3].Name, s.words[s.pos-1].Name)
		} else {
			log.Println("grammar error")
		}
	} else if s.match("insert") {
		// 检查是否为 INSERT 语句
		if s.match("into") && s.id() && s.match("\\(") && s.idList() && s.match("\\)") && s.match("values") && s.match("\\(") && s.idList() && s.match("\\)") {
			s.connect() 
			columns := make([]string, 0)
			values := make([]string, 0)
			// 分别将插入的列名和值保存到 columns 和 values 列表中
			for i := 3; i < s.pos; i++ {
				if s.words[i].Type == ID {
					if i < s.pos/2 {
						columns = append(columns, s.words[i].Name)
					} else {
						values = append(values, s.words[i].Name)
					}
				}
			}
			// 插入数据到指定表中
			s.insert(s.words[s.pos-12].Name, columns, values)
		} else {
			log.Println("grammar error")
		}
	} else {
		log.Println("grammar error")
	}
}

