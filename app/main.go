package main

import (
	"log"
	"net/http"
	"os"
	"github.com/spotahome/go-thumbnail"
)

const DefaultPort = "9092"

func getServerPort() string {

	port := os.Getenv("IMAGE_PROXY_PORT")

	if port != "" {
		return port
	}

  return DefaultPort

}

func printHelp() {

  log.Print("Server started, listening on port " + getServerPort())

  log.Println("Use /thumbnail?size=<size>&url=<imageUrl> to generate a thumbnail")

}

func main() {

	printHelp()

  http.HandleFunc("/thumbnail", thumbnail.Handler())

  http.ListenAndServe(":"+getServerPort(), nil)

}
