package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sysmb"
)

func main() {
	buff := []byte{1, 3, 4}
	sysmb.PublishMsg(123, buff)
	// Wait for terminating signal
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	fmt.Println("client sysmb exit...")

}
