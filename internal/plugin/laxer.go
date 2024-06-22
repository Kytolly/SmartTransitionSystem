package plugin

import (
	"fmt"
	"strings"
	"unicode"
)

// Word 结构体
type Word struct {
	Name string
	Type int
}

// 标识符类型常量
const (
	SELECT  = 1
	FROM    = 2
	INSERT  = 3
	INTO    = 4
	VALUES  = 5
	UPDATE  = 6
	SET     = 7
	WHERE   = 8
	CREATE  = 9
	TABLE   = 10
	EQ      = 11 	// =
	COMMA   = 12 	// ,
	SEMICOLON = 13	// ;
	LPAREN  = 14 	// (
	RPAREN  = 15	// )
	VALUE   = 16	// \
	STAR    = 17 	// *
	INT     = 18
	VARCHAR = 19
	ID      = 20
	NUM     = 21
)

// Lexer 结构体
type Lexer struct {
	input string
	pos   int
	words []Word
}

// NewLexer 创建新的 Lexer
func NewLexer(input string) *Lexer {
	return &Lexer{
		input: 	input, 
		pos: 	0, 
		words: 	[]Word{},
	}
}

// Tokenize 进行词法分析
func (l *Lexer) Tokenize() []Word {
	for l.pos < len(l.input) {
		char := l.input[l.pos]
		// 跳过空格
		if unicode.IsSpace(rune(char)) {
			l.pos++
			continue
		}

		switch {
		case char == '*':
			l.words = append(l.words, Word{Name: "*", Type: STAR})
			l.pos++
		case char == '=':
			l.words = append(l.words, Word{Name: "=", Type: EQ})
			l.pos++
		case char == ',':
			l.words = append(l.words, Word{Name: ",", Type: COMMA})
			l.pos++
		case char == ';':
			l.words = append(l.words, Word{Name: ";", Type: SEMICOLON})
			l.pos++
		case char == '(':
			l.words = append(l.words, Word{Name: "(", Type: LPAREN})
			l.pos++
		case char == ')':
			l.words = append(l.words, Word{Name: ")", Type: RPAREN})
			l.pos++
		case unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char)):
			start := l.pos
			for l.pos < len(l.input) && (unicode.IsLetter(rune(l.input[l.pos])) || unicode.IsDigit(rune(l.input[l.pos]))) {
				l.pos++
			}
			word := l.input[start:l.pos]
			l.words = append(l.words, Word{Name: word, Type: l.lookupWordType(word)})
		default:
			l.pos++
		}
	}
	return l.words
}

// lookupWordType 根据单词返回类型
func (l *Lexer) lookupWordType(word string) int {
	switch strings.ToLower(word) {
	case "select":
		return SELECT
	case "from":
		return FROM
	case "insert":
		return INSERT
	case "into":
		return INTO
	case "values":
		return VALUES
	case "update":
		return UPDATE
	case "set":
		return SET
	case "where":
		return WHERE
	case "create":
		return CREATE
	case "table":
		return TABLE
	case "int":
		return INT
	case "varchar":
		return VARCHAR
	default:
		if _, err := fmt.Sscanf(word, "%d", new(int)); err == nil {
			return NUM
		}
		return ID
	}
}
