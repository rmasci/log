package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rmasci/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logFile := "test.log"
	lf, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("ERROR: Can't open logfile %v: %v\n", logFile, err)
		os.Exit(1)
	}
	lgOut := log.New(lf, "", log.LstdFlags)
	lgOut.Format = "-"
	lgOut.SetOutput(&lumberjack.Logger{
		Filename: logFile,
		Compress: true,
		MaxSize:  10,
	})
	lgOut.Printf("Lumberjack test\n")
	for i := 1; i <= 10; i++ {
		lgOut.Println("This is a test", i)
		time.Sleep(1 * time.Second)
	}

}
