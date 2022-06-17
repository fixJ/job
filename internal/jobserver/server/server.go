package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type GenericAPIServerConfig struct {
	Host string
	Port string
}

func (c *GenericAPIServerConfig) New() *GenericAPIServer {
	return &GenericAPIServer{
		Address: c.Host + ":" + c.Port,
		done:    make(chan os.Signal),
	}
}

type GenericAPIServer struct {
	Address string
	done    chan os.Signal
}

func (s GenericAPIServer) Register(route string, f http.HandlerFunc) {
	http.HandleFunc(route, f)
}

func (s GenericAPIServer) Start() {
	err := http.ListenAndServe(s.Address, nil)
	if err != nil {
		return
	}
	signal.Notify(s.done, syscall.SIGINT, syscall.SIGTERM)
	<-s.done
	return
}
