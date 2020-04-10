package collector

import (
	"fmt"

	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
)

type systemStatusCollector struct {
	systemStatus *prometheus.Desc
}

func NewSystemStatusCollector() *systemStatusCollector {
	return &systemStatusCollector{
		systemStatus: prometheus.NewDesc("radarr_system_status",
			"System Status",
			[]string{"hostname"}, nil,
		),
	}
}

func (collector *systemStatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.systemStatus
}

func (collector *systemStatusCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.New()
	client := radarr.NewClient(conf)

	systemStatus := radarr.SystemStatus{}
	if err := client.DoRequest(fmt.Sprintf("%s/api/v3/%s", conf.Hostname, "system/status"), &systemStatus); err != nil {
		ch <- prometheus.MustNewConstMetric(collector.systemStatus, prometheus.GaugeValue, float64(0.0),
			conf.Hostname,
		)
	} else if (radarr.SystemStatus{}) == systemStatus {
		ch <- prometheus.MustNewConstMetric(collector.systemStatus, prometheus.GaugeValue, float64(0.0),
			conf.Hostname,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.systemStatus, prometheus.GaugeValue, float64(1.0),
			conf.Hostname,
		)
	}
}
