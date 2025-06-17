package main

import (
	"fmt"

	"github.com/victorramos887/go_pdf/src/config"
	"github.com/victorramos887/go_pdf/src/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		fmt.Println("error", err)
		return
	}
	router.Initialize()
}
