package collector

import (
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type historyCollector struct {
	config        *cli.Context
	historyMetric *prometheus.Desc
}

func NewHistoryCollector(c *cli.Context) *historyCollector {
	return &historyCollector{
		config: c,
		historyMetric: prometheus.NewDesc(
			"radarr_history_total",
			"Total number of records in history",
			nil,
			prometheus.Labels{"url": c.String("url")},
		),
	}
}

func (collector *historyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.historyMetric
}

func (collector *historyCollector) Collect(ch chan<- prometheus.Metric) {
	client := radarr.NewClient(collector.config)
	history := radarr.History{}
	if err := client.DoRequest("history", &history); err != nil {
		log.Fatal(err)
	}
	ch <- prometheus.MustNewConstMetric(collector.historyMetric, prometheus.GaugeValue, float64(history.TotalRecords))
}
