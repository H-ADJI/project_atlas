package lexer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type source struct {
	filename string
	reader   *bufio.Reader
}

func readSourceFile(filepath string) source {
	// add closing the file somewhere XD
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("could not open file : " + filepath)
	}
	r := bufio.NewReader(file)
	return source{filename: filepath, reader: r}
}
func readSource(sourceCode string) source {
	reader := bufio.NewReader(strings.NewReader(sourceCode))
	return source{filename: "LocalStringVariable", reader: reader}
}

type Lexer struct {
	currentLine int
	currentRow  int
	sourceCode  source
}

func NewLexer(s source) *Lexer {
	l := new(Lexer)
	l.sourceCode = s
	return l
}

func (l *Lexer) NextToken() Token {
	builder := new(strings.Builder)
	for {
		// read next character

		char, _, err := l.sourceCode.reader.ReadRune()
		if err == io.EOF {
			return Token{Type: EOF}
		}
		if err != nil {
			fmt.Println("an Error occured !!")
		}
		l.currentRow++

		// handling new lines and white spaces
		if char == '\n' {
			l.currentLine++
			l.currentRow = 0
			continue
		}
		if char == ' ' {
			l.currentRow++
			continue
		}

		// Special characters
		switch char {
		case '+':
			return Token{Type: PLUS, Literal: string(char)}
		case '-':
			return Token{Type: MINUS, Literal: string(char)}
		case '/':
			return Token{Type: SLASH, Literal: string(char)}
		case '*':
			return Token{Type: ASTERISK, Literal: string(char)}
		case '(':
			return Token{Type: LPAREN, Literal: string(char)}
		case ')':
			return Token{Type: RPAREN, Literal: string(char)}
		case '}':
			return Token{Type: RBRACE, Literal: string(char)}
		case '{':
			return Token{Type: LBRACE, Literal: string(char)}
		case ',':
			return Token{Type: COMMA, Literal: string(char)}
		case ';':
			return Token{Type: SEMICOLON, Literal: string(char)}
		case '=':
			return Token{Type: ASSIGN, Literal: string(char)}
		}

		if unicode.IsDigit(char) {
			builder.WriteRune(char)
			char, _, _ = l.sourceCode.reader.ReadRune()

			for unicode.IsDigit(char) {
				builder.WriteRune(char)
				char, _, err = l.sourceCode.reader.ReadRune()
				if err == io.EOF {
					l.sourceCode.reader.UnreadRune()
					return Token{Type: INT, Literal: builder.String()}
				}

			}

			l.sourceCode.reader.UnreadRune()
			return Token{Type: INT, Literal: builder.String()}

		}
		// catching identifiers and keywords
		if isLetter(char) {
			builder.WriteRune(char)
			char, _, _ = l.sourceCode.reader.ReadRune()

			for unicode.IsDigit(char) || isLetter(char) {
				builder.WriteRune(char)
				char, _, err = l.sourceCode.reader.ReadRune()
				if err == io.EOF {
					l.sourceCode.reader.UnreadRune()
					literal := builder.String()
					if tokenType, ok := keywords[literal]; ok {
						return Token{Type: tokenType, Literal: literal}
					}
					return Token{Type: IDENT, Literal: literal}
				}

			}

			l.sourceCode.reader.UnreadRune()
			literal := builder.String()
			if tokenType, ok := keywords[literal]; ok {
				return Token{Type: tokenType, Literal: literal}
			}
			return Token{Type: IDENT, Literal: literal}
		}

		// default case if nothing above works
		return Token{Type: ILLEGAL}
	}
}

func isLetter(char rune) bool {
	return unicode.IsLetter(char) || char == '_'
}

func (l *Lexer) Peek() (rune, error) {
	char, _, err := l.sourceCode.reader.ReadRune()
	if err != nil {
		return 0, err
	}
	l.sourceCode.reader.UnreadRune()

	return char, nil
}
