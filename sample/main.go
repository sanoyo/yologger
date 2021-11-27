package main

import (
	"os"

	"github.com/sanoyo/yologger"
)

func main() {
	logger := yologger.New("sano", os.Stdout)
	logger.Info("failed to fetch URL").Out()
}
