package main

import (
	"os"
	"time"
)

const logFile = "actions.log"

func LogAction(action string) {
	entry := time.Now().Format("2006-01-02 15:04:05") + " - " + action + "\n"
	f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(entry)
}
