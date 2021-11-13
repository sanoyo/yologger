package yologger

import "fmt"

type Logger struct {
	Name  string
	Level string
}

func New(name, level string) *Logger {
	log := &Logger{
		Name:  name,
		Level: level,
	}
	return log
}

// TODO: fields実装
func (log *Logger) Info(msg string) {
	fmt.Println(msg)
}
