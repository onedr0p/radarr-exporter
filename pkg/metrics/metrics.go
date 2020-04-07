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

	// Downloaded - Total number of Movies downloaded
	Downloaded = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_download_total",
			Namespace: "radarr",
			Help:      "Total number of downloaded movies",
		},
		[]string{"hostname"},
	)

	// Monitored - Total number of Movies monitored
	Monitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_monitored_total",
			Namespace: "radarr",
			Help:      "Total number of monitored movies",
		},
		[]string{"hostname"},
	)

	// Unmonitored - Total number of Movies unmonitored
	Unmonitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_unmonitored_total",
			Namespace: "radarr",
			Help:      "Total number of unmonitored movies",
		},
		[]string{"hostname"},
	)

	// Missing - Total number of Movies missing
	Missing = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_missing_total",
			Namespace: "radarr",
			Help:      "Total number of missing movies",
		},
		[]string{"hostname"},
	)

	// Wanted - Total number of Movies wanted
	Wanted = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_wanted_total",
			Namespace: "radarr",
			Help:      "Total number of wanted movies",
		},
		[]string{"hostname"},
	)

	// Qualities - Total number of Movies by Quality
	Qualities = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "movie_quality_total",
			Namespace: "radarr",
			Help:      "Total number of downloaded movies by quality",
		},
		[]string{"hostname", "quality"},
	)

	// SystemStatus - System Status
	SystemStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "system_status",
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

	// SystemHealth - Health issues with type and message
	SystemHealth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "system_health_issues",
			Namespace: "radarr",
			Help:      "Health issues in Radarr",
		},
		[]string{"hostname", "source", "type", "message", "wikiurl"},
	)
)

// Init initializes all Prometheus metrics made available by Radarr exporter.
func Init() {
	prometheus.MustRegister(Movie)
	prometheus.MustRegister(Downloaded)
	prometheus.MustRegister(Monitored)
	prometheus.MustRegister(Unmonitored)
	prometheus.MustRegister(Missing)
	prometheus.MustRegister(Wanted)
	prometheus.MustRegister(Qualities)
	prometheus.MustRegister(SystemStatus)
	prometheus.MustRegister(History)
	prometheus.MustRegister(Queue)
	prometheus.MustRegister(FileSize)
	prometheus.MustRegister(RootFolder)
	prometheus.MustRegister(SystemHealth)
}
