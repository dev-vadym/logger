package main

import (
	"fmt"

	"github.com/dev-vadym/logger"
)

func main() {

	// Default logger
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Notice("Notice")
	logger.Warning("Warning")
	logger.Error("Error")
	logger.Critical("Critical")

	// Custom logger with default handler
	l := logger.NewLogger("test")

	l.Debug("Debug")
	l.Info("Info")
	l.Notice("Notice")
	l.Warning("Warning")
	l.Error("Error")
	l.Critical("Critical")

	// Custom logger with custom handler
	l2 := logger.NewLogger("test2")
	l2.SetHandler(&MyHandler{})

	l2.Debug("Debug")
	l2.Info("Info")
	l2.Notice("Notice")
	l2.Warning("Warning")
	l2.Error("Error")
	l2.Critical("Critical")
}

type MyHandler struct {
	logger.BaseHandler
}

func (h *MyHandler) Handle(rec *logger.Record) {
	fmt.Printf(rec.Format, rec.Args...)
}

func (h *MyHandler) Close() {
}
