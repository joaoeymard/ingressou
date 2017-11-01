package logger

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Warn calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Warn(args ...interface{}) {
	mutex.Lock()
	fmt.Print(yellow, "[WARN]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Print(args...)
	fmt.Print("\n")
	mutex.Unlock()
}

// Warnf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, args ...interface{}) {
	mutex.Lock()
	fmt.Print(yellow, "[WARN]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	log.Printf(format, args...)
	if !strings.Contains(format[len(format)-1:], "\n") {
		fmt.Print("\n")
	}
	mutex.Unlock()
}

// Warnln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Warnln(args ...interface{}) {
	mutex.Lock()
	fmt.Print(yellow, "[WARN]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Println(args...)
	mutex.Unlock()
}
