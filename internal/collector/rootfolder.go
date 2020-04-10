package collector

import (
	"fmt"

	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
)

type rootFolderCollector struct {
	rootFolderMetric *prometheus.Desc
}

func NewRootFolderCollector() *rootFolderCollector {
	return &rootFolderCollector{
		rootFolderMetric: prometheus.NewDesc("radarr_rootfolder_freespace_bytes",
			"Root folder space in bytes",
			[]string{"hostname", "path"}, nil,
		),
	}
}

func (collector *rootFolderCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.rootFolderMetric
}

func (collector *rootFolderCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.New()
	client := radarr.NewClient(conf)

	rootFolders := radarr.RootFolder{}
	client.DoRequest(fmt.Sprintf("%s/api/v3/%s", conf.Hostname, "rootfolder"), &rootFolders)

	// Group metrics by path
	for _, rootFolder := range rootFolders {
		ch <- prometheus.MustNewConstMetric(collector.rootFolderMetric, prometheus.GaugeValue, float64(rootFolder.FreeSpace),
			conf.Hostname, rootFolder.Path,
		)
	}
}
