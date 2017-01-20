package lexer

// Lexer is the struct to scan on
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
