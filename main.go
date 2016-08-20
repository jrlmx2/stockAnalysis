package main

import (
	"signal"
	"syscall"
)

func main() {
	//example notify channels of interupt and kill signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}
