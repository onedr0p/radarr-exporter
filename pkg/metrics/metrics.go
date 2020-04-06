package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Movie - Total Number of Movies
	Movie = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_total",
			Namespace: "radarr",
			Help:      "Total number of movies",
		},
		[]string{"hostname"},
	)

	// MovieDownloaded - Total number of Movies downloaded
	MovieDownloaded = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_download_total",
			Namespace: "radarr",
			Help:      "Total number of downloaded movies",
		},
		[]string{"hostname"},
	)

	// MovieMonitored - Total number of Movies monitored
	MovieMonitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_monitored_total",
			Namespace: "radarr",
			Help:      "Total number of monitored movies",
		},
		[]string{"hostname"},
	)

	// MovieUnmonitored - Total number of Movies unmonitored
	MovieUnmonitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_unmonitored_total",
			Namespace: "radarr",
			Help:      "Total number of unmonitored movies",
		},
		[]string{"hostname"},
	)

	// MovieQualities - Total number of Movies by Quality
	MovieQualities = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_quality_total",
			Namespace: "radarr",
			Help:      "Total number of downloaded movies by quality",
		},
		[]string{"hostname", "quality"},
	)

	// Wanted - Total number of missing/wanted Movies
	Wanted = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_missing_total",
			Namespace: "radarr",
			Help:      "Total number of missing movies",
		},
		[]string{"hostname"},
	)

	// Status - System Status
	Status = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "status",
			Namespace: "radarr",
			Help:      "System Status",
		},
		[]string{"hostname"},
	)

	// History - Total number of records in History
	History = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "history_total",
			Namespace: "radarr",
			Help:      "Total number of records in history",
		},
		[]string{"hostname"},
	)

	// Queue - Total number of movies in Queue
	Queue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "queue_total",
			Namespace: "radarr",
			Help:      "Total number of movies in queue by status",
		},
		[]string{"hostname", "status"},
	)

	// FileSize - Total size of all Movies
	FileSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_bytes",
			Namespace: "radarr",
			Help:      "Total file size of all movies in bytes",
		},
		[]string{"hostname"},
	)

	// RootFolder - Space by Root Folder
	RootFolder = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "rootfolder_freespace_bytes",
			Namespace: "radarr",
			Help:      "Root folder space in bytes",
		},
		[]string{"hostname", "folder"},
	)

	// Health - Health issues with type and message
	Health = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "health_issues",
			Namespace: "radarr",
			Help:      "Health issues in Radarr",
		},
		[]string{"hostname", "type", "message", "wikiurl"},
	)
)

// Init initializes all Prometheus metrics made available by PI-Hole exporter.
func Init() {
	prometheus.MustRegister(Movie)
	prometheus.MustRegister(MovieDownloaded)
	prometheus.MustRegister(MovieMonitored)
	prometheus.MustRegister(MovieUnmonitored)
	prometheus.MustRegister(MovieQualities)
	prometheus.MustRegister(Wanted)
	prometheus.MustRegister(Status)
	prometheus.MustRegister(History)
	prometheus.MustRegister(Queue)
	prometheus.MustRegister(FileSize)
	prometheus.MustRegister(RootFolder)
	prometheus.MustRegister(Health)
}
