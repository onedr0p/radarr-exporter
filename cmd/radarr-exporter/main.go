package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/onedr0p/radarr-exporter/pkg/config"
	"github.com/onedr0p/radarr-exporter/pkg/metrics"
	"github.com/onedr0p/radarr-exporter/pkg/radarr"
	"github.com/onedr0p/radarr-exporter/pkg/server"
)

const (
	name = "radarr-exporter"
)

var (
	s *server.Server
)

func main() {
	conf := config.Load()

	metrics.Init()

	initRadarrClient(conf.Hostname, conf.ApiKey, conf.Interval)
	initHttpServer(conf.Port)

	handleExitSignal()
}

func initRadarrClient(hostname, apiKey string, interval time.Duration) {
	client := radarr.NewClient(hostname, apiKey, interval)
	go client.Scrape()
}

func initHttpServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	fmt.Println(fmt.Sprintf("\n%s HTTP server stopped", name))
}
