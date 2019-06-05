package server

import (
	"fmt"
	"os"
	"strconv"
)

type config struct {
	Bind        string
	ReadTimeout int
}

var conf = config{
	Bind:        ":9092",
	ReadTimeout: 10,
}

func strEnvConfig(s *string, name string) {
	if env := os.Getenv(name); len(env) > 0 {
		*s = env
	}
}

func intEnvConfig(i *int, name string) {
	if env, err := strconv.Atoi(os.Getenv(name)); err == nil {
		*i = env
	}
}

func Config() *config {

	if port := os.Getenv("IMGPROXY_PORT"); len(port) > 0 {
		conf.Bind = fmt.Sprintf(":%s", port)
	}

	strEnvConfig(&conf.Bind, "IMGPROXY_BIND")
	intEnvConfig(&conf.ReadTimeout, "IMGPROXY_READ_TIMEOUT")

	return &conf

}
