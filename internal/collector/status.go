package collector

import (
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli/v2"
)

type systemStatusCollector struct {
	config       *cli.Context
	systemStatus *prometheus.Desc
}

func NewSystemStatusCollector(c *cli.Context) *systemStatusCollector {
	return &systemStatusCollector{
		config: c,
		systemStatus: prometheus.NewDesc(
			"radarr_system_status",
			"System Status",
			nil,
			prometheus.Labels{"url": c.String("url")},
		),
	}
}

func (collector *systemStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.systemStatus
}

func (collector *systemStatusCollector) Collect(ch chan<- prometheus.Metric) {
	client := radarr.NewClient(collector.config)
	systemStatus := radarr.SystemStatus{}
	if err := client.DoRequest("system/status", &systemStatus); err != nil {
		ch <- prometheus.MustNewConstMetric(collector.systemStatus, prometheus.GaugeValue, float64(0.0))
	} else if (radarr.SystemStatus{}) == systemStatus {
		ch <- prometheus.MustNewConstMetric(collector.systemStatus, prometheus.GaugeValue, float64(0.0))
	} else {
		ch <- prometheus.MustNewConstMetric(collector.systemStatus, prometheus.GaugeValue, float64(1.0))
	}
}
