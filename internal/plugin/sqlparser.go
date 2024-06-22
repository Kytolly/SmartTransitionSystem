package plugin

import (
	"system/internal/utils"
)
func ParseAndExecuteSQL(sqls string) {
	inputs := utils.PreParseSQLs(sqls)  
	for _, input := range inputs {
		l := NewLexer(input)
		words := l.Tokenize()
		s := NewSentenceAnalyze(words)
		s.checkSentence()
	}
}
