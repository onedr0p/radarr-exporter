package radarr

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/onedr0p/radarr-exporter/internal/config"
	log "github.com/sirupsen/logrus"
)

// Client struct is a Radarr client to request an instance of a Radarr
type Client struct {
	hostname       string
	apiKey         string
	basicAuth      bool
	basicAuthCreds string
	httpClient     http.Client
}

// NewClient method initializes a new Radarr client.
func NewClient(conf *config.Config) *Client {
	return &Client{
		hostname:       conf.Hostname,
		apiKey:         conf.ApiKey,
		basicAuth:      conf.BasicAuth,
		basicAuthCreds: conf.BasicAuthCreds,
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// DoRequest - Take a HTTP Request and return Unmarshaled data
func (c *Client) DoRequest(endpoint string, target interface{}) error {
	log.Infof("Sending HTTP request to %s", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if c.basicAuth && c.basicAuthCreds != "" {
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.basicAuthCreds)))
	}
	req.Header.Add("X-Api-Key", c.apiKey)

	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("An error has occurred during retrieving Radarr statistics %v", err)
		return err
	}
	if resp.StatusCode != 200 {
		errMsg := "An error has occurred during retrieving Radarr statistics HTTP statuscode not 200"
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
