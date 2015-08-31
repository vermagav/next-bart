package bart

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"

	"appengine"
	"appengine/urlfetch"

	"github.com/vermagav/next-bart/config"
)

// Response Schema for BART API -> get ETD
/*
<root>
  <station>
    <name>Richmond</name>
    <abbr>RICH</abbr>
    <etd>
      <destination>Fremont</destination>
      <abbreviation>FRMT</abbreviation>
      <estimate>
        <minutes>1</minutes>
        <platform>2</platform>
        <direction>South</direction>
        <length>3</length>
        <color>ORANGE</color>
        <hexcolor>#ff9933</hexcolor>
        <bikeflag>1</bikeflag>
      </estimate>
    </etd>
  </station>
<message/>
</root>
*/

type EtdRoot struct {
	Station EtdStation `xml:"station"`
	Message Message    `xml:"message"`
}

type Message struct {
	Error Error `xml:"error"`
}

type Error struct {
	Text string `xml:"text"`
}

type EtdStation struct {
	Name string `xml:"name"`
	Id   string `xml:"abbr"`
	Etd  []Etd  `xml:"etd"`
}

type Etd struct {
	Destination  string     `xml:"destination"`
	Abbreviation string     `xml:"abbreviation"`
	Estimates    []Estimate `xml:"estimate"`
}

type Estimate struct {
	Minutes   string `xml:"minutes"`
	Platform  string `xml:"platform"`
	Direction string `xml:"direction"`
	Length    string `xml:"length"`
	Color     string `xml:"color"`
	Hexcolor  string `xml:"hexcolor"`
	Bikeflag  bool   `xml:"bikeflag"`
}

// GetDepartures performs a request to the BART API and fetches
// departures real-time departure estimates from the given station id.
func GetDepartures(r *http.Request, stationId string) (*[]Etd, error) {
	// Perform http request via app engine handler
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(config.Bart.EtdUrl + stationId)
	if err != nil {
		return nil, errCreatingRequest
	}

	// Read body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errReadingBody
	}

	// Prepare object and unmarshal XML into it
	root := EtdRoot{}
	err = xml.Unmarshal(body, &root)
	if err != nil {
		return nil, errParsingBody
	}

	if root.Message.Error.Text != "" {
		return nil, errors.New("Error calling BART API: " + root.Message.Error.Text)
	}

	return &root.Station.Etd, nil
}
