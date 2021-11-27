package yologger

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type Logger struct {
	message string
	level   string

	// TODO: 定義する構造体を別で用意する
	mu     sync.Mutex
	writer *bufio.Writer
	w      io.Writer
	buf    *bytes.Buffer
}

func New(w io.Writer) *Logger {
	l := &Logger{
		w:   w,
		buf: &bytes.Buffer{},
	}
	return l
}

func (l *Logger) Info(msg string) *Logger {
	l.message = msg
	l.level = Green.Add(InfoLevel.String())

	return l.Write()
}

func (l *Logger) Write() *Logger {
	now := time.Now()
	log := fmt.Sprintf("%s [%s] %s", now, l.level, l.message)
	_, err := l.buf.WriteString(log)
	if err != nil {
		// TODO: エラーハンドリング検討要
		panic(err)
	}

	return l
}

func (l *Logger) Out() {
	l.mu.Lock()
	defer l.mu.Unlock()

	tmp := l.buf.Bytes()
	l.w.Write(tmp)
}
