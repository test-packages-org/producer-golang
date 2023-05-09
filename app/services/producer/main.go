package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.New(os.Stdout, "producer: ", log.LstdFlags|log.Lshortfile)

	defer logger.Println("shutdown")
	logger.Println("startup")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

}
