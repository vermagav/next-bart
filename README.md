# next-bart

### Description

This a RESTful web service that serves real-time departure information for upcoming BART rides.

* Language: Golang
* Data Source: [BART API](http://www.bart.gov/schedules/developers/api)
* Host/Platform: Google AppEngine

### The problem

Departure information is readily provided by BART in a very verbose, XML format. The BART API, however, does not provide any filtering based on geolocation and it is up to the developer to build additional functionality around location. The core vision behind this web service is to provide an easy-to-use and fast API for station and departure information around the user's location. Imagine a mobile app that sends a push notification to your device based on upcoming departures at your closest BART station. You can write a client that displays multi-directional departures based on the user's location with just a few calls to the next-bart API.

### Documentation

Further documentation is split into two parts:

```text
/docs
  |-- API.md
  |-- Developer.md
```
#### API Docs

To integrate your client with this web service, check out our [API docs](docs/API.md).

#### Developer Review

To understand how the code works, check out the [Dev Docs](docs/Developer.md).
