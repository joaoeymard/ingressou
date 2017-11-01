package logger

import (
	"fmt"
	"strings"
	"time"
)

// Info calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Info(args ...interface{}) {
	mutex.Lock()
	fmt.Print(blue, "[INFO]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Print(args...)
	fmt.Print("\n")
	mutex.Unlock()
}

// Infof calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	mutex.Lock()
	fmt.Print(blue, "[INFO]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Printf(format, args...)
	if !strings.Contains(format[len(format)-1:], "\n") {
		fmt.Print("\n")
	}
	mutex.Unlock()
}

// Infoln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	mutex.Lock()
	fmt.Print(blue, "[INFO]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Println(args...)
	mutex.Unlock()
}
