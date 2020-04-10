package collector

import (
	"fmt"

	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
)

type historyCollector struct {
	historyMetric *prometheus.Desc
}

func NewHistoryCollector() *historyCollector {
	return &historyCollector{
		historyMetric: prometheus.NewDesc("radarr_history_total",
			"Total number of records in history",
			[]string{"hostname"}, nil,
		),
	}
}

func (collector *historyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.historyMetric
}

func (collector *historyCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.New()
	client := radarr.NewClient(conf)

	history := radarr.History{}
	client.DoRequest(fmt.Sprintf("%s/api/v3/%s", conf.Hostname, "history"), &history)

	ch <- prometheus.MustNewConstMetric(collector.historyMetric, prometheus.GaugeValue, float64(history.TotalRecords),
		conf.Hostname,
	)
}
