package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/twistedogic/spero/cmd/poll"
)

const BASE = "https://bet.hkjc.com/football/getJSON.aspx"

var betTypes = []string{
	"HAD",
	"FHA",
	"CRS",
	"FCS",
	"FTS",
	"OOE",
	"TTG",
	"HFT",
	"HHA",
	"HDC",
	"HIL",
	"FHL",
}

func init() {
	prometheus.MustRegister(poll.OddMetric)
	prometheus.MustRegister(poll.ClientMetric)
}

func Graceful(server *http.Server, f ...func()) {
	defer os.Exit(0)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)
	fmt.Println(<-sigc)
	if err := server.Shutdown(nil); err != nil {
		log.Println(err)
	}
	for _, stop := range f {
		stop()
	}
}

func main() {
	poll, err := poll.New(BASE, "10s")
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{Addr: ":8080"}
	go func() {
		log.Println("Starting Server")
		http.Handle("/metrics", promhttp.Handler())
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	go func() {
		out := poll.Start(betTypes...)
		for m := range out {
			log.Println(m.MatchID)
		}
	}()
	Graceful(srv, poll.Stop)
}
