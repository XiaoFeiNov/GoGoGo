package codeTest

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestInitClient(t *testing.T) {
	signalCh := make(chan os.Signal, 1)
	flagCh := make(chan struct{})
	go InitClient(flagCh)
	<- flagCh
	RedisClient.GetString("name")
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<- signalCh
}
