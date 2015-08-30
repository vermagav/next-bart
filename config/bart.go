package config

import (
	"bytes"
)

// Bart contains static information for the BART API
var Bart struct {
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
func ConstructStationUrl() string {
	var b bytes.Buffer
	b.WriteString(Bart.ApiUrl)
	b.WriteString(Bart.StationsResource)
	b.WriteString(Bart.ApiKey)
	b.WriteString(Bart.StationsQuery)
	return b.String()
}

// Pre-construct and store the URL so we don't concatenate
// strings with each call
func ConstructEtdUrl() string {
	var b bytes.Buffer
	b.WriteString(Bart.ApiUrl)
	b.WriteString(Bart.EtdResource)
	b.WriteString(Bart.ApiKey)
	b.WriteString(Bart.EtdQuery)
	return b.String()
}

// init() is called once per package initialization
func init() {
	Bart.ApiUrl = "http://api.bart.gov/api"
	Bart.ApiKey = "?key=MW9S-E7SL-26DU-VV8V"

	Bart.StationsResource = "/stn.aspx"
	Bart.StationsQuery = "&cmd=stns"

	Bart.EtdResource = "/etd.aspx"
	Bart.EtdQuery = "&cmd=etd&orig="

	Bart.StationsUrl = ConstructStationUrl()
	Bart.EtdUrl = ConstructEtdUrl()
}
