package client

const (
	Rows    = 10
	Columns = 40
)

type Logger struct {
	buffer       [Rows][Columns]rune
	cursorRow    int
	cursorColumn int
}

func NewLogger() *Logger {
	return &Logger{
		buffer:       [Rows][Columns]rune{},
		cursorRow:    0,
		cursorColumn: 0,
	}
}

func (l *Logger) PutString(str string) {
	for _, c := range str {
		if c == '\n' {
			l.newLine()
		} else if l.cursorColumn < Columns-1 {
			l.buffer[l.cursorRow][l.cursorColumn] = c
			l.cursorColumn++
		}
	}
}

func (l *Logger) newLine() {
	l.cursorColumn = 0
	if l.cursorRow < Rows-1 {
		l.cursorRow++
	} else {
		for i := 0; i < Rows-1; i++ {
			l.buffer[i] = l.buffer[i+1]
		}
		l.buffer[Rows-1] = [Columns]rune{}
	}
}

func (l *Logger) String() string {
	output := ""
	for _, line := range l.buffer {
		output += string(line[:])
		output += "\n"
	}
	return output
}
