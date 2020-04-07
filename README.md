# radarr-exporter

Prometheus Exporter for Radarr

[![Docker Pulls](https://img.shields.io/docker/pulls/onedr0p/radarr-exporter)](https://hub.docker.com/r/onedr0p/radarr-exporter) [![Go Report Card](https://goreportcard.com/badge/github.com/onedr0p/radarr-exporter)](https://goreportcard.com/report/github.com/onedr0p/radarr-exporter)

## Usage

|Name             |Description                                                  |Default|
|-----------------|-------------------------------------------------------------|-------|
|`RADARR_HOSTNAME`|You Radarr instance's URL                                    |       |
|`RADARR_APIKEY`  |Your Radarr instance's API Key                               |       |
|`INTERVAL`       |The duration of which the exporter will call the Radarr API  |`10m`  |
|`PORT`           |The port the exporter will listen on                         |`9811` |
|`STARTUP_DEPLY`  |The duration to wait before first calling the Radarr API     |       |

### Docker Compose Example

```yaml
version: '3.7'
services:
  radarr-exporter:
    image: onedr0p/radarr-exporter:v1.1.0
    environment:
      RADARR_HOSTNAME: "http://localhost:7878"
      RADARR_APIKEY: "..."
      INTERVAL: "1h"
```

### Metrics

```bash
# HELP radarr_history_total Total number of records in history
# TYPE radarr_history_total gauge
radarr_history_total{hostname="http://localhost:7878"} 6174
# HELP radarr_movie_download_total Total number of downloaded movies
# TYPE radarr_movie_download_total gauge
radarr_movie_download_total{hostname="http://localhost:7878"} 4717
# HELP radarr_movie_missing_total Total number of missing movies
# TYPE radarr_movie_missing_total gauge
radarr_movie_missing_total{hostname="http://localhost:7878"} 97
# HELP radarr_movie_monitored_total Total number of monitored movies
# TYPE radarr_movie_monitored_total gauge
radarr_movie_monitored_total{hostname="http://localhost:7878"} 155
# HELP radarr_movie_quality_total Total number of downloaded movies by quality
# TYPE radarr_movie_quality_total gauge
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-1080p"} 1222
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-480p"} 99
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-576p"} 267
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Bluray-720p"} 1004
radarr_movie_quality_total{hostname="http://localhost:7878",quality="DVD"} 1347
radarr_movie_quality_total{hostname="http://localhost:7878",quality="HDTV-1080p"} 43
radarr_movie_quality_total{hostname="http://localhost:7878",quality="HDTV-720p"} 46
radarr_movie_quality_total{hostname="http://localhost:7878",quality="Remux-1080p"} 48
radarr_movie_quality_total{hostname="http://localhost:7878",quality="SDTV"} 25
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBDL-1080p"} 465
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBDL-480p"} 53
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBDL-720p"} 94
radarr_movie_quality_total{hostname="http://localhost:7878",quality="WEBRip-1080p"} 2
# HELP radarr_movie_total Total number of movies
# TYPE radarr_movie_total gauge
radarr_movie_total{hostname="http://localhost:7878"} 4817
# HELP radarr_movie_unmonitored_total Total number of unmonitored movies
# TYPE radarr_movie_unmonitored_total gauge
radarr_movie_unmonitored_total{hostname="http://localhost:7878"} 4662
# HELP radarr_movies_bytes Total file size of all movies in bytes
# TYPE radarr_movies_bytes gauge
radarr_movie_bytes{hostname="http://localhost:7878"} 2.3326478328365e+13
# HELP radarr_queue_total Total number of movies in queue by status
# TYPE radarr_queue_total gauge
radarr_queue_total{hostname="http://localhost:7878",status="Ok"} 1
radarr_queue_total{hostname="http://localhost:7878",status="Warning"} 9
# HELP radarr_root_folder_space_bytes Root folder space in bytes
# TYPE radarr_root_folder_space_bytes gauge
radarr_rootfolder_freespace_bytes{folder="/media/Library/Movies/",hostname="http://localhost:7878"} 2.5011930497024e+13
# HELP radarr_status System Status
# TYPE radarr_status gauge
radarr_status{hostname="http://localhost:7878"} 1
# HELP radarr_health_issues Health issues in Radarr
# TYPE radarr_health_issues gauge
radarr_health_issues{hostname="http://localhost:7878",message="Branch develop is for a previous version of Radarr, set branch to 'Aphrodite' for further updates",type="error",wikiurl="https://github.com/Radarr/Radarr/wiki/Health-checks#branch-develop-is-for-a-previous-version-of-radarr-set-branch-to-aphrodite-for-further-updates"} 1
radarr_health_issues{hostname="http://localhost:7878",message="No download client is available",type="warning",wikiurl="https://github.com/Radarr/Radarr/wiki/Health-checks#no-download-client-is-available"} 1
radarr_health_issues{hostname="http://localhost:7878",message="No indexers available with Automatic Search enabled, Radarr will not provide any automatic search results",type="warning",wikiurl="https://github.com/Radarr/Radarr/wiki/Health-checks#no-indexers-available-with-automatic-search-enabled-radarr-will-not-provide-any-automatic-search-results"} 1
radarr_health_issues{hostname="http://localhost:7878",message="No indexers available with RSS sync enabled, Radarr will not grab new releases automatically",type="error",wikiurl="https://github.com/Radarr/Radarr/wiki/Health-checks#no-indexers-available-with-rss-sync-enabled-radarr-will-not-grab-new-releases-automatically"} 1
```
