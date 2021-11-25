package yologger

import (
	"bufio"
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	Name    string
	message string
	// TODO: 定義する構造体を別で用意する
	mu     sync.Mutex
	writer *bufio.Writer
	w      io.Writer
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
	l.writer.Write([]byte(l.message))
}

func (l *Logger) Write(bs []byte) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Println("bs", bs)

	// b := &bytes.Buffer{}
	result, _ := l.w.Write(bs)
	fmt.Println("result", result)

	return
}

// func (l *Logger) Sync() error {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()

// 	// b := &bytes.Buffer{}
// 	// l.Write(b.Bytes())
// 	// fmt.Println(l.message)

// 	return nil
// }
