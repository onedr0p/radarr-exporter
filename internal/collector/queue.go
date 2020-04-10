package collector

import (
	"fmt"

	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
)

type queueCollector struct {
	queueMetric *prometheus.Desc
}

func NewQueueCollector() *queueCollector {
	return &queueCollector{
		queueMetric: prometheus.NewDesc("radarr_queue_total",
			"Total number of movies in the queue by status",
			[]string{"hostname", "status", "download_status", "download_state"}, nil,
		),
	}
}

func (collector *queueCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.queueMetric
}

func (collector *queueCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.New()
	client := radarr.NewClient(conf)

	queue := radarr.Queue{}
	client.DoRequest(fmt.Sprintf("%s/api/v3/%s", conf.Hostname, "queue"), &queue)
	// Calculate total pages
	var totalPages = (queue.TotalRecords + queue.PageSize - 1) / queue.PageSize
	// Paginate
	var queueStatusAll = make([]radarr.QueueRecords, 0, queue.TotalRecords)
	queueStatusAll = append(queueStatusAll, queue.Records...)
	if totalPages > 1 {
		for page := 2; page <= totalPages; page++ {
			client.DoRequest(fmt.Sprintf("%s/api/v3/%s?page=%d", conf.Hostname, "queue", page), &queue)
			queueStatusAll = append(queueStatusAll, queue.Records...)
		}
	}
	// Group metrics by status, download_status and download_state
	var queueMetrics prometheus.Metric
	for i, s := range queueStatusAll {
		queueMetrics = prometheus.MustNewConstMetric(collector.queueMetric, prometheus.GaugeValue, float64(i+1),
			conf.Hostname, s.Status, s.TrackedDownloadStatus, s.TrackedDownloadState,
		)
	}
	ch <- queueMetrics
}
