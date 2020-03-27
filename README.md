# radarr-exporter

Prometheus Exporter for Radarr

[![Docker Pulls](https://img.shields.io/docker/pulls/onedr0p/radarr-exporter)](https://hub.docker.com/r/onedr0p/radarr-exporter)

## Usage

|Name             |Description                                                  |Default|
|-----------------|-------------------------------------------------------------|-------|
|`RADARR_HOSTNAME`|You Radarr instance's URL                                    |       |
|`RADARR_APIKEY`  |Your Radarr instance's API Key                               |       |
|`INTERVAL`       |The duration of which the exporter will scrape the Radarr API|`2m`   |
|`PORT`           |The port the exporter will listen on                         |`9811` |

### Docker Compose Example

```yaml
version: '3.7'
services:
  radarr-exporter:
    image: onedr0p/radarr-exporter:v1.0.0
    environment:
      RADARR_HOSTNAME: "http://localhost:7878"
      RADARR_APIKEY: "..."
      INTERVAL: "1h"
```

### Metrics

```bash
# TYPE radarr_history_total gauge
radarr_history_total{hostname="http://localhost:7878"} 5925
# HELP radarr_movie_missing_total Total number of missing movies
# TYPE radarr_movie_missing_total gauge
radarr_movie_missing_total{hostname="http://localhost:7878"} 119
# HELP radarr_movie_download_total Total number of downloaded movies
# TYPE radarr_movie_download_total gauge
radarr_movie_download_total{hostname="http://localhost:7878"} 4644
# HELP radarr_movie_quality_total Total number of downloaded movies by quality
# TYPE radarr_movie_quality_total gauge
radarr_movie_quality_total{hostname="http://localhost:7878",quality=""} 120
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-1080p"} 1190
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-480p"} 100
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-576p"} 267
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-720p"} 1004
radarr_movie_quality_total{hostname="http://localhost:7878",quality="DVD"} 1345
radarr_movie_quality_total{hostname="http://localhost:7878",quality="HDTV-1080p"} 42
radarr_movie_quality_total{hostname="http://localhost:7878",quality="HDTV-720p"} 46
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Remux-1080p"} 48
radarr_movie_quality_total{hostname="http://localhost:7878",quality="SDTV"} 25
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBDL-1080p"} 426
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBDL-480p"} 53
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBDL-720p"} 94
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBRip-1080p"} 2
# HELP radarr_movie_total Total number of movies
# TYPE radarr_movie_total gauge
radarr_movie_total{hostname="http://localhost:7878"} 4762
# HELP radarr_queue_total Total number of movies in queue
# TYPE radarr_queue_total gauge
radarr_queue_total{hostname="http://localhost:7878"} 2
# HELP radarr_root_folder_space Root folder space
# TYPE radarr_root_folder_space gauge
radarr_root_folder_space{folder="/media/Library/Movies/",hostname="http://localhost:7878"} 2.6006486515712e+13
# HELP radarr_status System Status
# TYPE radarr_status gauge
radarr_status{hostname="http://localhost:7878"} 1
```
