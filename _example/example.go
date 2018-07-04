package main

import (
	"github.com/aellwein/slf4go"
	_ "github.com/aellwein/slf4go-native-adaptor"
)

func main() {
	logger := slf4go.GetLogger("mylogger")
	logger.Info("It works!")
	logger.Warnf("The answer is %d", 42)
}
