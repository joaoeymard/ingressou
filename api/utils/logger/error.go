package logger

import (
	"fmt"
	"strings"
	"time"
)

// Error calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Error(args ...interface{}) {
	mutex.Lock()
	fmt.Print(red, "[ERROR]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Print(args...)
	fmt.Print("\n")
	mutex.Unlock()
}

// Errorf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	mutex.Lock()
	fmt.Print(red, "[ERROR]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Printf(format, args...)
	if !strings.Contains(format[len(format)-1:], "\n") {
		fmt.Print("\n")
	}
	mutex.Unlock()
}

// Errorln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	mutex.Lock()
	fmt.Print(red, "[ERROR]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Println(args...)
	mutex.Unlock()
}
