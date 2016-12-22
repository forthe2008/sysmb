package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sysmb"
)

func main() {
	buff := make([]byte, 1024)
	conn, _ := sysmb.SubscribeMsg(123)
	n, _ := conn.Read(buff)
	fmt.Println("Data:", buff[:n])
	// Wait for terminating signal
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	fmt.Println("client sysmb exit...")

}
