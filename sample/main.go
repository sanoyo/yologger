package main

import (
	"os"

	"github.com/sanoyo/yologger"
)

func main() {
	logger := yologger.New(os.Stdout)
	logger.Info("failed to fetch URL").Out()
	// logger.Warn("failed to fetch URL").Out()
	// logger.Error("failed to fetch URL").Out()
}
