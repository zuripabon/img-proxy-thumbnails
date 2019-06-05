package main

import (
	"github.com/spotahome/imgproxy-server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	s := server.StartServer()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	server.ShutdownServer(s)

}
