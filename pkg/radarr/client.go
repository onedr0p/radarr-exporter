package radarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/onedr0p/radarr-exporter/pkg/metrics"
)

var (
	apiUrlPattern = "%s/api/%s"
)

// Client struct is a Sonarr client to request an instance of a Sonarr
type Client struct {
	httpClient http.Client
	interval   time.Duration
	hostname   string
	apiKey     string
}

// NewClient method initializes a new Sonarr client.
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

// Scrape method logins and retrieves statistics from Sonarr JSON API
// and then pass them as Prometheus metrics.
func (c *Client) Scrape() {
	for range time.Tick(c.interval) {

		// System Status
		status := SystemStatus{}
		if err := c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "system/status"), &status); err != nil {
			metrics.Status.WithLabelValues(c.hostname).Set(0.0)
			return
		} else if (SystemStatus{}) == status {
			metrics.Status.WithLabelValues(c.hostname).Set(0.0)
			return
		} else {
			metrics.Status.WithLabelValues(c.hostname).Set(1.0)
		}

		// Movies, Downloaded Movies, Downloaded Movies by Quality Name
		// Monitored, Unmonitored
		var (
			moviesDownloaded  = 0
			moviesMonitored   = 0
			moviesUnmonitored = 0
			moviesQualities   = map[string]int{}
		)
		movies := Movie{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "movie"), &movies)
		for _, s := range movies {
			if s.HasFile == true {
				moviesDownloaded++
			}
			if s.Monitored == true {
				moviesMonitored++
			} else {
				moviesUnmonitored++
			}
			// Movies without file don't have a quality
			if s.MovieFile.Quality.Quality.Name != "" {
				moviesQualities[s.MovieFile.Quality.Quality.Name]++
			}
		}
		metrics.Movie.WithLabelValues(c.hostname).Set(float64(len(movies)))
		metrics.MovieDownloaded.WithLabelValues(c.hostname).Set(float64(moviesDownloaded))
		metrics.MovieMonitored.WithLabelValues(c.hostname).Set(float64(moviesMonitored))
		metrics.MovieUnmonitored.WithLabelValues(c.hostname).Set(float64(moviesUnmonitored))

		for qualityName, count := range moviesQualities {
			metrics.MovieQualities.WithLabelValues(c.hostname, qualityName).Set(float64(count))
		}

		// History
		history := History{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "history"), &history)
		metrics.History.WithLabelValues(c.hostname).Set(float64(history.TotalRecords))

		// Wanted
		wanted := WantedMissing{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "wanted/missing"), &wanted)
		metrics.Wanted.WithLabelValues(c.hostname).Set(float64(wanted.TotalRecords))

		// Queue
		queue := Queue{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "queue"), &queue)
		metrics.Queue.WithLabelValues(c.hostname).Set(float64(len(queue)))

		// Root Folder
		rootFolders := RootFolder{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "rootfolder"), &rootFolders)
		for _, rootFolder := range rootFolders {
			metrics.RootFolder.WithLabelValues(c.hostname, rootFolder.Path).Set(float64(rootFolder.FreeSpace))
		}

		// Health Issues
		healthIssuesByType := map[string]int{}
		health := Health{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "health"), &health)
		for _, h := range health {
			healthIssuesByType[h.Type]++
		}
		for issueType, count := range healthIssuesByType {
			metrics.Health.WithLabelValues(c.hostname, issueType).Set(float64(count))
		}
	}
}

func (c *Client) apiRequest(endpoint string, target interface{}) error {
	log.Printf("Sending HTTP request to %s", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("X-Api-Key", c.apiKey)
	// req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("")))
	if err != nil {
		log.Fatal("An error has occured when creating HTTP statistics request", err)
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal("An error has occured during retrieving Sonarr statistics", err)
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
