package yologger

import (
	"bufio"
	"bytes"
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
	buf    *bytes.Buffer
}

func New(name string, w io.Writer) *Logger {
	l := &Logger{
		Name: name,
		w:    w,
		buf:  &bytes.Buffer{},
	}
	return l
}

// TODO: fields追加できるようにする
func (l *Logger) Info(msg string) *Logger {
	l.message = msg
	return l.Write()
}

func (l *Logger) Write() *Logger {
	_, err := l.buf.WriteString(l.message)
	if err != nil {
		// TODO: エラーハンドリング検討要
		panic(err)
	}

	return l
}

func (l *Logger) Out() {
	// TODO: lockするかどうか検討する
	// l.mu.Lock()
	// defer l.mu.Unlock()

	tmp := l.buf.Bytes()
	l.w.Write(tmp)
}
