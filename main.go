package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logPath := "./log"
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		slog.Error("Could not create log file.", "error", err)
	}
	defer logFile.Close()

	loggerArgs := &slog.HandlerOptions{AddSource: true}
	logger := slog.New(slog.NewJSONHandler(logFile, loggerArgs))
	slog.SetDefault(logger)

	fmt.Println("test")


	// TODO: next
	// AG 24:37
	// CK 40:10
}
