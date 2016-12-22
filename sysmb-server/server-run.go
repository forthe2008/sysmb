package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sysmb"
)

func main() {
	server := sysmb.NewServer()
	server.Startup()

	fmt.Println("sysmb server is running")
	// Wait for terminating signal
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	fmt.Println("Shutdown sysmb...")

	server.Shutdown()
	fmt.Println("Sysmb is down")
}
