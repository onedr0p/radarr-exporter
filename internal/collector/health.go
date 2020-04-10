package collector

import (
	"fmt"

	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
)

type systemHealthCollector struct {
	systemHealthMetric *prometheus.Desc
}

func NewSystemHealthCollector() *systemHealthCollector {
	return &systemHealthCollector{
		systemHealthMetric: prometheus.NewDesc("radarr_system_health_issues",
			"Total number of movies in the queue by status",
			[]string{"hostname", "source", "type", "message", "wikiurl"}, nil,
		),
	}
}

func (collector *systemHealthCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.systemHealthMetric
}

func (collector *systemHealthCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.New()
	client := radarr.NewClient(conf)

	systemHealth := radarr.SystemHealth{}
	client.DoRequest(fmt.Sprintf("%s/api/v3/%s", conf.Hostname, "health"), &systemHealth)

	// Group metrics by source, type, message and wikiurl
	for _, s := range systemHealth {
		ch <- prometheus.MustNewConstMetric(collector.systemHealthMetric, prometheus.GaugeValue, float64(1),
			conf.Hostname, s.Source, s.Type, s.Message, s.WikiURL,
		)
	}
}
