package logger

import (
	"fmt"
	"strings"
	"time"
)

// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Debug(args ...interface{}) {
	mutex.Lock()
	fmt.Print(green, "[DEBUG]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Print(args...)
	fmt.Print("\n")
	mutex.Unlock()
}

// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	mutex.Lock()
	fmt.Print(green, "[DEBUG]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Printf(format, args...)
	if !strings.Contains(format[len(format)-1:], "\n") {
		fmt.Print("\n")
	}
	mutex.Unlock()
}

// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Debugln(args ...interface{}) {
	mutex.Lock()
	fmt.Print(green, "[DEBUG]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Println(args...)
	mutex.Unlock()
}
