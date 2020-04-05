package radarr

// RootFolder - Stores struct of JSON response
type RootFolder []struct {
	Path      string `json:"path"`
	FreeSpace int64  `json:"freeSpace"`
}

// SystemStatus - Stores struct of JSON response
type SystemStatus struct {
	Version string `json:"version"`
	AppData string `json:"appData"`
	Branch  string `json:"branch"`
}

// Queue - Stores struct of JSON response
type Queue []struct {
	Title string `json:"title"`
	Size  int32  `json:"size"`
}

// History - Stores struct of JSON response
type History struct {
	TotalRecords int `json:"totalRecords"`
}

// WantedMissing - Stores struct of JSON response
type WantedMissing struct {
	TotalRecords int `json:"totalRecords"`
}

// Health - Stores struct of JSON response
type Health []struct {
	Type    string `json:type`
	Message string `json:message`
	WikiURL string `json:wikiUrl`
}

// Movie - Stores struct of JSON response
type Movie []struct {
	HasFile   bool `json:"hasFile"`
	Monitored bool `json:"monitored"`
	MovieFile struct {
		Size    int64 `json:"size"`
		Quality struct {
			Quality struct {
				Name string `json:"name"`
			} `json:"quality"`
		} `json:"quality"`
	} `json:"movieFile"`
	QualityProfileID int `json:"qualityProfileId"`
}
