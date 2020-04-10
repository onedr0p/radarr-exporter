package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/onedr0p/radarr-exporter/internal/collector"
	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/handlers"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func init() {
	logLevel := strings.ToUpper(config.GetEnvStr("LOG_LEVEL", "TRACE"))
	switch logLevel {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	conf := config.New()

	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Debugf("Readiness probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Debugf("Readiness probe is positive.")
	}()

	r := prometheus.NewRegistry()
	r.MustRegister(
		collector.NewMovieCollector(),
		collector.NewQueueCollector(),
		collector.NewHistoryCollector(),
		collector.NewRootFolderCollector(),
		collector.NewSystemStatusCollector(),
		collector.NewSystemHealthCollector(),
	)

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/liveness", handlers.LivenessHandler)
	http.HandleFunc("/readiness", handlers.ReadinessHandler(isReady))
	http.Handle("/metrics", handler)

	log.Debugf("Listening on localhost:%d", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
