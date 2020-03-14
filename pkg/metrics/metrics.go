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
			Help:      "Total number of movies in queue",
		},
		[]string{"hostname"},
	)

	// RootFolder - Space by Root Folder
	RootFolder = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "root_folder_space",
			Namespace: "radarr",
			Help:      "Root folder space",
		},
		[]string{"hostname", "folder"},
	)

	// Health - Amount of health issues by type
	Health = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "health_issues",
			Namespace: "radarr",
			Help:      "Amount of health issues in Radarr",
		},
		[]string{"hostname", "type"},
	)
)

// Init initializes all Prometheus metrics made available by PI-Hole exporter.
func Init() {
	prometheus.MustRegister(Movie)
	prometheus.MustRegister(MovieDownloaded)
	prometheus.MustRegister(MovieQualities)
	prometheus.MustRegister(Wanted)
	prometheus.MustRegister(Status)
	prometheus.MustRegister(History)
	prometheus.MustRegister(Queue)
	prometheus.MustRegister(RootFolder)
	prometheus.MustRegister(Health)
}