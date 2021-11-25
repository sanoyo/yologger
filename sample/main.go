package main

import (
	"github.com/sanoyo/yologger"
)

func main() {
	logger := yologger.New("sano", "info")
	logger.Info("failed to fetch URL")
	// defer logger.Sync()
}
