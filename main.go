package main

import (
	"github.com/jesserahman/goLangPracticeProject/app"
	"github.com/jesserahman/goLangPracticeProject/logger"
)

func main() {
	logger.Info("starting application...")
	app.Run()
}
