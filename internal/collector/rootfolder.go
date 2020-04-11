package collector

import (
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type rootFolderCollector struct {
	config           *cli.Context
	rootFolderMetric *prometheus.Desc
}

func NewRootFolderCollector(c *cli.Context) *rootFolderCollector {
	return &rootFolderCollector{
		config: c,
		rootFolderMetric: prometheus.NewDesc(
			"radarr_rootfolder_freespace_bytes",
			"Root folder space in bytes",
			[]string{"path"},
			prometheus.Labels{"url": c.String("url")},
		),
	}
}

func (collector *rootFolderCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.rootFolderMetric
}

func (collector *rootFolderCollector) Collect(ch chan<- prometheus.Metric) {
	client := radarr.NewClient(collector.config)
	rootFolders := radarr.RootFolder{}
	if err := client.DoRequest("rootfolder", &rootFolders); err != nil {
		log.Fatal(err)
	}
	// Group metrics by path
	if len(rootFolders) > 0 {
		for _, rootFolder := range rootFolders {
			ch <- prometheus.MustNewConstMetric(collector.rootFolderMetric, prometheus.GaugeValue, float64(rootFolder.FreeSpace),
				rootFolder.Path,
			)
		}
	}
}
