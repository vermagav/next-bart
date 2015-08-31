# API Docs

The web service's responses are serialized into JSON strings. There will always be two top-level entries present: `data` and `error`. When a request succeeds, `data` is populated with a JSON object, and `error` is `null`. When a request fails, `data` is `null`, and `error` is populated with an error string. Upon receiving a 500 Internal Server Error, feel free to retry. Upon receiving a 400 Bad Request, look at the error message to fix your request; retrying will not make the request succeed.

### `GET /stations`

##### Query Parameters

| Param | Type | Purpose |
|-------|------|---------|
| `count` | int | Limit number of stations returned by this number |
| `lat` | 64-bit float | Latitude of the client device |
| `long` | 64-bit float | Longitude of the client device |

* `count` always limits results by provided value. The service will return a 400 Bad Request if a value less than 1 is used. The value also has a ceiling of the total number of BART stations served in the San Francisco Bay Area.

* `lat` and `long` must always be present together. The service will return a 400 Bad Request if either is missing, or malformed. Negative values are allowed. The service considers 0 as the default (i.e. not specified), as lat/long of 0/0 are nowhere near San Francisco.

##### Examples

###### `GET /stations`

 * Get information about all BART stations

```
curl -i "http://next-bart.appspot.com/stations"
```

```json
{
    "Data": [
        {
            "Id": "16TH",
            "Name": "16th St. Mission",
            "Lat": 37.765062,
            "Long": -122.419694
        },
        {
            "Id": "EMBR",
            "Name": "Embarcadero",
            "Lat": 37.792976,
            "Long": -122.396742
        },
        {
            "Id": "ORIN",
            "Name": "Orinda",
            "Lat": 37.87836087,
            "Long": -122.1837911
        },
        {
            "Id": "PITT",
            "Name": "Pittsburg/Bay Point",
            "Lat": 38.018914,
            "Long": -121.945154
        },
		{
            "Id": "<Additional entries snipped>"
		}
    ],
    "error": null
}
```

###### `GET /stations?count=n`

 * Get information about BART stations, limit to `n` in no specific order

```
curl -i "http://next-bart.appspot.com/stations?count=1"
```

```json
{
    "Data": [
        {
            "Id": "POWL",
            "Name": "Powell St.",
            "Lat": 37.784991,
            "Long": -122.406857
        }
    ],
    "error": null
}
```

###### `GET /stations?count=n&lat=x&long=y`

 * Get information about the `n` closest bart stations, sorted by distance to the provided lat/long coordinates

```
curl -i "http://next-bart.appspot.com/stations?count=1&lat=37.779528&long=-122.413756"
```

```json
{
    "Data": [
        {
            "Id": "CIVC",
            "Name": "Civic Center/UN Plaza",
            "Lat": 37.779528,
            "Long": -122.413756
        }
    ],
    "error": null
}
```

### `GET /stations/:stationId`

##### URL Parameters

| Param | Type | Purpose |
|-------|------|---------|
| `stationId` | string | The short abbreviated string ID for a station returned by `/stations` |

##### Examples

 * Get information about a specific station by id

```
curl -i "http://next-bart.appspot.com/stations/CIVC"
```

```json
{
    "Data": [
        {
            "Id": "CIVC",
            "Name": "Civic Center/UN Plaza",
            "Lat": 37.779528,
            "Long": -122.413756
        }
    ],
    "error": null
}
```

### `GET /stations/:stationId/departures`

##### URL Parameters

| Param | Type | Purpose |
|-------|------|---------|
| `stationId` | string | The short abbreviated string ID for a station returned by `/stations` |

##### Examples

 * Get real-time departures for the specified station

```
curl -i "http://next-bart.appspot.com/stations/CIVC/departures"
```

```json
{
    "Data": [
        {
            "Destination": "Daly City",
            "Abbreviation": "DALY",
            "Estimates": [
                {
                    "Minutes": "23",
                    "Platform": "1",
                    "Direction": "South",
                    "Length": "4",
                    "Color": "BLUE",
                    "Hexcolor": "#0099cc",
                    "Bikeflag": true
                }
            ]
        },
        {
            "Destination": "SF Airport",
            "Abbreviation": "SFIA",
            "Estimates": [
                {
                    "Minutes": "43",
                    "Platform": "1",
                    "Direction": "South",
                    "Length": "5",
                    "Color": "YELLOW",
                    "Hexcolor": "#ffff33",
                    "Bikeflag": true
                }
            ]
        },
        {
            "Destination": "SFO/Millbrae",
            "Abbreviation": "MLBR",
            "Estimates": [
                {
                    "Minutes": "13",
                    "Platform": "1",
                    "Direction": "South",
                    "Length": "5",
                    "Color": "YELLOW",
                    "Hexcolor": "#ffff33",
                    "Bikeflag": true
                }
            ]
        }
    ],
    "error": null
}
```
