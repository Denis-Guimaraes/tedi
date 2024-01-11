package logger

import "fmt"

var colorInfo = "\033[0m"
var colorError = "\033[31m"

func Info(text string) {
	fmt.Println(string(colorInfo), text)
}

func Error(text string) {
	fmt.Println(string(colorError), text)
}
