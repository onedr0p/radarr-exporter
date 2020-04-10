package collector

import (
	"fmt"

	"github.com/onedr0p/radarr-exporter/internal/config"
	"github.com/onedr0p/radarr-exporter/internal/radarr"
	"github.com/prometheus/client_golang/prometheus"
)

type movieCollector struct {
	movieMetric       *prometheus.Desc
	downloadedMetric  *prometheus.Desc
	monitoredMetric   *prometheus.Desc
	unmonitoredMetric *prometheus.Desc
	wantedMetric      *prometheus.Desc
	missingMetric     *prometheus.Desc
	filesizeMetric    *prometheus.Desc
	qualitiesMetric   *prometheus.Desc
}

func NewMovieCollector() *movieCollector {
	return &movieCollector{
		movieMetric: prometheus.NewDesc("radarr_movie_total",
			"Total number of movies",
			[]string{"hostname"}, nil,
		),
		downloadedMetric: prometheus.NewDesc("radarr_movie_download_total",
			"Total number of downloaded movies",
			[]string{"hostname"}, nil,
		),
		monitoredMetric: prometheus.NewDesc("radarr_movie_monitored_total",
			"Total number of monitored movies",
			[]string{"hostname"}, nil,
		),
		unmonitoredMetric: prometheus.NewDesc("radarr_movie_unmonitored_total",
			"Total number of unmonitored movies",
			[]string{"hostname"}, nil,
		),
		wantedMetric: prometheus.NewDesc("radarr_movie_wanted_total",
			"Total number of wanted movies",
			[]string{"hostname"}, nil,
		),
		missingMetric: prometheus.NewDesc("radarr_movie_missing_total",
			"Total number of missing movies",
			[]string{"hostname"}, nil,
		),
		filesizeMetric: prometheus.NewDesc("radarr_movie_filesize_total",
			"Total filesize of all movies",
			[]string{"hostname"}, nil,
		),
		qualitiesMetric: prometheus.NewDesc("radarr_movie_quality_total",
			"Total number of downloaded movies by quality",
			[]string{"hostname", "quality"}, nil,
		),
	}
}

func (collector *movieCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.movieMetric
	ch <- collector.downloadedMetric
	ch <- collector.monitoredMetric
	ch <- collector.unmonitoredMetric
	ch <- collector.wantedMetric
	ch <- collector.missingMetric
	ch <- collector.filesizeMetric
	ch <- collector.qualitiesMetric
}

func (collector *movieCollector) Collect(ch chan<- prometheus.Metric) {
	conf := config.New()
	client := radarr.NewClient(conf)

	var fileSize int64
	var (
		downloaded  = 0
		monitored   = 0
		unmonitored = 0
		missing     = 0
		wanted      = 0
		qualities   = map[string]int{}
	)
	movies := radarr.Movie{}
	client.DoRequest(fmt.Sprintf("%s/api/v3/%s", conf.Hostname, "movie"), &movies)
	for _, s := range movies {
		if s.HasFile {
			downloaded++
		}
		if s.Monitored {
			monitored++
			if !s.HasFile && s.Status == "released" {
				missing++
			} else if !s.HasFile {
				wanted++
			}
		} else {
			unmonitored++
		}
		if s.MovieFile.Quality.Quality.Name != "" {
			qualities[s.MovieFile.Quality.Quality.Name]++
		}
		if s.MovieFile.Size != 0 {
			fileSize += s.MovieFile.Size
		}
	}

	ch <- prometheus.MustNewConstMetric(collector.movieMetric, prometheus.GaugeValue, float64(len(movies)),
		conf.Hostname,
	)
	ch <- prometheus.MustNewConstMetric(collector.downloadedMetric, prometheus.GaugeValue, float64(downloaded),
		conf.Hostname,
	)
	ch <- prometheus.MustNewConstMetric(collector.monitoredMetric, prometheus.GaugeValue, float64(monitored),
		conf.Hostname,
	)
	ch <- prometheus.MustNewConstMetric(collector.unmonitoredMetric, prometheus.GaugeValue, float64(unmonitored),
		conf.Hostname,
	)
	ch <- prometheus.MustNewConstMetric(collector.wantedMetric, prometheus.GaugeValue, float64(wanted),
		conf.Hostname,
	)
	ch <- prometheus.MustNewConstMetric(collector.missingMetric, prometheus.GaugeValue, float64(missing),
		conf.Hostname,
	)
	ch <- prometheus.MustNewConstMetric(collector.filesizeMetric, prometheus.GaugeValue, float64(fileSize),
		conf.Hostname,
	)
	for qualityName, count := range qualities {
		ch <- prometheus.MustNewConstMetric(collector.qualitiesMetric, prometheus.GaugeValue, float64(count),
			conf.Hostname, qualityName,
		)
	}
}
