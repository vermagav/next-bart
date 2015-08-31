package bart

import (
	"bytes"
)

// Config contains static, configuration information for the BART API
var Config struct {
	// Base URL for BART API
	ApiUrl string
	// The validation key for accessing BART API
	// Note: The BART API allows a permissive model with a public key
	// that is general in nature and not specific to an app/consumer.
	// Thus, we can hard code this public key.
	// Ref: http://www.bart.gov/schedules/developers/api
	ApiKey string

	// The resource for fetching station information
	StationsResource string
	// The query param key/value pair for fetching station information
	StationsQuery string
	// The complete URL used for fetching station information
	StationsUrl string

	// The resource for fetching estimated times for departure (ETD)
	EtdResource string
	// The query param key/value pairs for fetching ETD
	EtdQuery string
	// The complete URL (minus station ID) used for fetching ETD
	EtdUrl string
}

// Pre-construct and store the URL so we don't concatenate
// strings with each call
func constructStationUrl() string {
	var b bytes.Buffer
	b.WriteString(Config.ApiUrl)
	b.WriteString(Config.StationsResource)
	b.WriteString(Config.ApiKey)
	b.WriteString(Config.StationsQuery)
	return b.String()
}

// Pre-construct and store the URL so we don't concatenate
// strings with each call
func constructEtdUrl() string {
	var b bytes.Buffer
	b.WriteString(Config.ApiUrl)
	b.WriteString(Config.EtdResource)
	b.WriteString(Config.ApiKey)
	b.WriteString(Config.EtdQuery)
	return b.String()
}

// init() is called once per file initialization
func init() {
	Config.ApiUrl = "http://api.bart.gov/api"
	Config.ApiKey = "?key=MW9S-E7SL-26DU-VV8V"

	Config.StationsResource = "/stn.aspx"
	Config.StationsQuery = "&cmd=stns"

	Config.EtdResource = "/etd.aspx"
	Config.EtdQuery = "&cmd=etd&orig="

	Config.StationsUrl = constructStationUrl()
	Config.EtdUrl = constructEtdUrl()
}
