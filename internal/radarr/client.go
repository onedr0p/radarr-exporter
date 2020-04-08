package radarr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/onedr0p/radarr-exporter/internal/metrics"
)

// Client struct is a Radarr client to request an instance of a Radarr
type Client struct {
	httpClient http.Client
	interval   time.Duration
	hostname   string
	apiKey     string
}

// NewClient method initializes a new Radarr client.
func NewClient(hostname, apiKey string, interval time.Duration) *Client {
	return &Client{
		hostname: hostname,
		apiKey:   apiKey,
		interval: interval,
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// Scrape method logs in and retrieves statistics from Radarr JSON API
// and then pass them as Prometheus metrics.
func (c *Client) Scrape() {
	for range time.Tick(c.interval) {

		// System Status
		systemStatus := SystemStatus{}
		if err := c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "system/status"), &systemStatus); err != nil {
			metrics.SystemStatus.WithLabelValues(c.hostname).Set(0.0)
			return
		} else if (SystemStatus{}) == systemStatus {
			metrics.SystemStatus.WithLabelValues(c.hostname).Set(0.0)
			return
		} else {
			metrics.SystemStatus.WithLabelValues(c.hostname).Set(1.0)
		}

		// Movies, Downloaded Movies, Downloaded Movies by Quality Name
		// Monitored, Unmonitored, FileSize
		var fileSize int64
		var (
			downloaded  = 0
			monitored   = 0
			unmonitored = 0
			missing     = 0
			wanted      = 0
			qualities   = map[string]int{}
		)
		movies := Movie{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "movie"), &movies)
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
		metrics.Movie.WithLabelValues(c.hostname).Set(float64(len(movies)))
		metrics.Downloaded.WithLabelValues(c.hostname).Set(float64(downloaded))
		metrics.Monitored.WithLabelValues(c.hostname).Set(float64(monitored))
		metrics.Unmonitored.WithLabelValues(c.hostname).Set(float64(unmonitored))
		metrics.FileSize.WithLabelValues(c.hostname).Set(float64(fileSize))
		metrics.Missing.WithLabelValues(c.hostname).Set(float64(missing))
		metrics.Wanted.WithLabelValues(c.hostname).Set(float64(wanted))
		for qualityName, count := range qualities {
			metrics.Qualities.WithLabelValues(c.hostname, qualityName).Set(float64(count))
		}

		// History
		history := History{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "history"), &history)
		metrics.History.WithLabelValues(c.hostname).Set(float64(history.TotalRecords))

		// Queue
		queue := Queue{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "queue"), &queue)
		// Calculate total pages
		var totalPages = (queue.TotalRecords + queue.PageSize - 1) / queue.PageSize
		// Paginate
		var queueStatusAll = make([]QueueRecords, 0, queue.TotalRecords)
		queueStatusAll = append(queueStatusAll, queue.Records...)
		if totalPages > 1 {
			for page := 2; page <= totalPages; page++ {
				c.apiRequest(fmt.Sprintf("%s/api/v3/%s?page=%d", c.hostname, "queue", page), &queue)
				queueStatusAll = append(queueStatusAll, queue.Records...)
			}
		}
		// Group Status, TrackedDownloadStatus and TrackedDownloadState
		for i, s := range queueStatusAll {
			metrics.Queue.WithLabelValues(c.hostname, s.Status, s.TrackedDownloadStatus, s.TrackedDownloadState).Set(float64(i + 1))
		}

		// Root Folder
		rootFolders := RootFolder{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "rootfolder"), &rootFolders)
		for _, rootFolder := range rootFolders {
			metrics.RootFolder.WithLabelValues(c.hostname, rootFolder.Path).Set(float64(rootFolder.FreeSpace))
		}

		// Health Issues
		systemHealth := SystemHealth{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "health"), &systemHealth)
		for _, h := range systemHealth {
			metrics.SystemHealth.WithLabelValues(c.hostname, h.Source, h.Type, h.Message, h.WikiURL).Set(float64(1))
		}
	}
}

func (c *Client) apiRequest(endpoint string, target interface{}) error {
	log.Printf("Sending HTTP request to %s", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("X-Api-Key", c.apiKey)
	if err != nil {
		log.Fatal("An error has occurred when creating HTTP request", err)
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal("An error has occurred during retrieving Radarr statistics", err)
		return err
	} else if resp.StatusCode != 200 {
		fmt.Printf("Error: %v", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return err
}
