package yologger

import (
	"bufio"
	"fmt"
	"sync"
)

type Logger struct {
	Name    string
	message string
	// TODO: 定義する構造体を別で用意する
	mu     sync.Mutex
	writer *bufio.Writer
}

func New(name, level string) *Logger {
	l := &Logger{
		Name: name,
	}
	return l
}

// TODO: fields追加できるようにする
func (l *Logger) Info(msg string) {
	l.message = msg
}

func (l *Logger) Write(bs []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.writer.Write(bs)
}

func (l *Logger) Sync() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// b := &bytes.Buffer{}
	// l.Write(b.Bytes())
	fmt.Println(l.message)

	return nil
}
