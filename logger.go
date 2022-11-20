package main

import (
	"fmt"
	"os"
	"time"
	"github.com/fatih/color"
)

var logger = &Logger{}

type Logger struct {
	File      *os.File
}

func InitLogger() {
	f, err := os.Create("log.txt")
	PanicIfErr(err)
	logger.File = f
}

func (l Logger) Logf(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)

	date := time.Now().Format(`02 Jan 2006 15:04:05`)
	prefix := fmt.Sprintf("[DEBUG] [%v] ", date)

	color.New(color.FgCyan).Println(prefix + msg)
	l.File.WriteString(prefix + msg + "\n")
}
