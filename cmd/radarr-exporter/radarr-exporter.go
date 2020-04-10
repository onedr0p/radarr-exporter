package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/onedr0p/radarr-exporter/internal/collector"
	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.New()
	switch conf.LogLevel {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/liveness", livenessHandler)
	http.HandleFunc("/readiness", readinessHandler(isReady))
	http.Handle("/metrics", handler)

	log.Debug(fmt.Sprintf("Listening on localhost:%d", conf.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	response := `<h1>Radarr Exporter</h1><p><a href='/metrics'>metrics</a></p>`
	fmt.Fprintf(w, response)
}

func readinessHandler(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		fmt.Fprint(w)
	}
}

func livenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	fmt.Fprint(w)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
