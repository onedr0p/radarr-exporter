package radarr

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Client struct is a Radarr client to request an instance of a Radarr
type Client struct {
	url          string
	apikey       string
	authEnabled  bool
	authUsername string
	authPassword string
	httpClient   http.Client
}

// NewClient method initializes a new Radarr client.
func NewClient(c *cli.Context) *Client {
	return &Client{
		url:          c.String("url"),
		apikey:       c.String("api-key"),
		authEnabled:  c.Bool("basic-auth-enabled"),
		authUsername: c.String("basic-auth-username"),
		authPassword: c.String("basic-auth-password"),
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// DoRequest - Take a HTTP Request and return Unmarshaled data
func (c *Client) DoRequest(endpoint string, target interface{}) error {
	apiEndpoint := fmt.Sprintf("%s/api/v3/%s", c.url, endpoint)
	log.Infof("Sending HTTP request to %s", apiEndpoint)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if c.authEnabled && c.authUsername != "" && c.authPassword != "" {
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.authUsername+":"+c.authPassword)))
	}
	req.Header.Add("X-Api-Key", c.apikey)

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
