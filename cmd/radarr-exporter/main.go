package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/onedr0p/radarr-exporter/internal/collector"
	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/handlers"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func init() {
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
		log.SetLevel(log.InfoLevel)
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	log.Info(`                                       
    ........
  .............
  .... ............
  ....     ...........
  ....   XX    ...........
  ....   XXXXXX   ...........
  ....   XXXXXXXXX    ........
  ....   XXXXXXXXXXXXX   ......
  ....   XXXXXXXXXXXX     .....
  ....   XXXXXXXXX    ........
  ....   XXXXX     .........
  ....   XX    ...........
  ....      ...........
  ....   ..........
    ............
     .........
`)

	log.Infof(`Radarr exporter started...
Hostname: %s
API Key: %s
Port: %d
Basic Auth Enabled: %v
Basic Auth Crendentials: %s`,
		conf.Hostname,
		conf.ApiKey,
		conf.Port,
		conf.BasicAuth,
		conf.BasicAuthCreds,
	)

}

func main() {
	conf := config.New()

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
	http.HandleFunc("/healthz", handlers.HealthzHandler)
	http.Handle("/metrics", handler)

	log.Infof("Listening on localhost:%d", conf.Port)
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
