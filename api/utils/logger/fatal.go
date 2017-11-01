package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var statusExit = 1

// SetStatusExit Determina o estado de erro
func SetStatusExit(status int) {
	statusExit = status
}

// Fatal calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Fatal(args ...interface{}) {
	mutex.Lock()
	fmt.Print(bold, red, "[FATAL]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Print(args...)
	fmt.Print("\n")
	mutex.Unlock()
	os.Exit(statusExit)
}

// Fatalf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, args ...interface{}) {
	mutex.Lock()
	fmt.Print(bold, red, "[FATAL]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Printf(format, args...)
	if !strings.Contains(format[len(format)-1:], "\n") {
		fmt.Print("\n")
	}
	mutex.Unlock()
	os.Exit(statusExit)
}

// Fatalln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Fatalln(args ...interface{}) {
	mutex.Lock()
	fmt.Print(bold, red, "[FATAL]", reset)
	fmt.Print(" ")
	fmt.Print(time.Now().Format(logClock))
	fmt.Print(" ")
	fmt.Println(args...)
	mutex.Unlock()
	os.Exit(statusExit)
}
