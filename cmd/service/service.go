package service

import (
	"fmt"
	"log"
	"net/http"
)

type Service interface {
	Route() string
	Handler() http.Handler
}

type Runner struct {
	*http.Server
	services []Service
}

func New(port int, services ...Service) *Runner {
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	for _, s := range services {
		http.Handle(s.Route(), s.Handler())
	}
	return &Runner{httpServer, services}
}

func (r *Runner) Start() {
	go func() {
		if err := r.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

func (r *Runner) Stop() {
	log.Println("Stoping Metric Server")
	if err := r.Shutdown(nil); err != nil {
		log.Fatal(err)
	}
}
