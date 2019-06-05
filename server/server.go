package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

func buildRouter() *router {
	r := newRouter()

	r.GET("/thumbnail", ThumbnailHandler)

	return r
}

func StartServer() *http.Server {

	conf := Config()

	l, err := net.Listen("tcp", conf.Bind)

	if err != nil {
		log.Println(err.Error())
	}

	s := &http.Server{
		Handler:        buildRouter(),
		ReadTimeout:    time.Duration(conf.ReadTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Print("Server started, listening on " + conf.Bind)
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Println(err.Error())
		}
	}()

	return s
}

func ShutdownServer(s *http.Server) {

	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)

	defer close()

	s.Shutdown(ctx)
}
