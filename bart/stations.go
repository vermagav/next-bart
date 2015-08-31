package bart

import (
	"encoding/xml"
	"io/ioutil"
	"math"
	"net/http"
	"sort"

	"appengine"
	"appengine/urlfetch"

	"github.com/vermagav/next-bart/config"
)

// Response Schema for BART API -> get all stations
/*
<root>
  <stations>
    <station>
      <name>12th St. Oakland City Center</name>
      <abbr>12TH</abbr>
      <gtfs_latitude>37.803664</gtfs_latitude>
      <gtfs_longitude>-122.271604</gtfs_longitude>
      <address>1245 Broadway</address>
      <city>Oakland</city>
      <county>alameda</county>
      <state>CA</state>
      <zipcode>94612</zipcode>
    </station>
  </stations>
</root>
*/

type Root struct {
	StationsSet Stations `xml:"stations"`
}

type Stations struct {
	StationList []Station `xml:"station"`
}

type Station struct {
	Id       string  `xml:"abbr"`
	Name     string  `xml:"name"`
	Lat      float64 `xml:"gtfs_latitude"`
	Long     float64 `xml:"gtfs_longitude"`
	distance float64 // Lower case, not exported to JSON
}

// This is a cache of information on all station, keyed by id
var stationsCache map[string]Station

// init() is called once per file initialization
func init() {
	stationsCache = make(map[string]Station)
}

// cacheAllStations performs an http request to the BART API,
// parses the XML response, and caches all station data in memory.
// All we care about is the lat and long, and stations are unlikely to move
// to a different location, so let's grab and hold on to their locations.
func cacheAllStations(r *http.Request) error {
	// Perform http request via app engine handler
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(config.Bart.StationsUrl)
	if err != nil {
		return errCreatingRequest
	}

	// Read body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errReadingBody
	}

	// Prepare object and unmarshal XML into it
	root := Root{}
	err = xml.Unmarshal(body, &root)
	if err != nil {
		return errParsingBody
	}

	// Add each station to cache, overwriting any previous value
	for _, station := range root.StationsSet.StationList {
		stationsCache[station.Id] = station
	}

	return nil
}

// GetStations is a publicly exported function that returns the map
// of pre-cached stations.
func GetStations(r *http.Request, count int, lat float64, long float64) ([]Station, error) {
	// Do we have a cached list of stations?
	if len(stationsCache) == 0 {
		// Try to fetch again
		err := cacheAllStations(r)
		if err != nil {
			return nil, err
		}
	}

	// Do we need to compute distance?
	needDistance := true
	if lat == 0. && long == 0. {
		needDistance = false
	}

	// Create an array, computing distance if necessary
	stationsList := make([]Station, len(stationsCache))
	i := 0
	for _, v := range stationsCache {
		stationsList[i] = v
		if needDistance {
			stationsList[i].distance = computeDistance(v.Lat, v.Long, lat, long)
		}
		i++
	}

	// If we computed distance, sort list by it
	if needDistance {
		sort.Sort(ByDistance(stationsList))
	}

	// If a valid count wasn't included, return all
	if count <= 0 || count > len(stationsCache) {
		count = len(stationsCache)
	}

	return stationsList[:count], nil
}

// Fetch information about a spedcific station by its id
func GetStationbyId(r *http.Request, id string) (*Station, error) {
	// Does this station id exist?
	if station, ok := stationsCache[id]; ok {
		return &station, nil
	} else {
		return nil, errStationNotFound
	}
}

// Compute the distance squared between two points. Since we only use this
// for comparison, we don't need to waste computation on square root.
func computeDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	distanceSquared := math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2)
	return distanceSquared
}

// Implement an interface that supports sorting by distance
type ByDistance []Station
func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }
