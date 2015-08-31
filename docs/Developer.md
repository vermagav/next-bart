# Developer Review

This document describes version 1 the code base.

In this version, there are two major packages:

* service
* bart

```text
/bart
  |-- config.go
  |-- departures.go
  |-- errors.go
  |-- stations.go
/service
  |--errors.go
  |--parse.go
  |--parse_test.go
  |--response.go
  |--service.go
```

### `package bart`

##### [config.go](../bart/config.go)

This file contains static, configuration information pertaining to the BART API. URLs, while sourced in modular variables, are constructed once upon startup to avoid repeated concatenation for each incoming request.

##### [departures.go](../bart/departures.go)

This file makes real-time HTTP requests to the BART API's estimated time of departures endpoints, parses XML data into strongly-typed structs, and returns only what is of interest to the web service.

##### [errors.go](../bart/errors.go)

Similar to the service package, this file contains definitions for errors encountered by the package.

##### [stations.go](../bart/stations.go)

This file deals with information pertaining to stations. It makes a **single** HTTP request to the BART API's stations information endpoint, parses XML data into strongly-typed structs, and then **caches** station data in an in-memory map. Since station data is unlikely to change, repeated calls to this endpoint aren't necessary. Future improvements to this model are discussed below. This file also computes squared distances between the client's coordinates and the station's coordinates, when requested with query parameters. In addition, the `[]Station` type is made sortable with an interface.

### `package service`

The service package is responsible for handling web service requests.

##### [errors.go](../service/errors.go)

This file contains definitions for errors encountered by the package. These errors are returned and forwarded to the `response.go` methods for consumption in the response payload.

##### [parse.go](../service/parse.go)

This file contains helper functions for parsing query parameters. They were separated in order to be tested in isolation.

##### [parse_test.go](../service/parse_test.go)

This file contains unit tests for `parse.go`.

##### [response.go](../service/response.go)

This file contains helper methods that create success, bad request, and internal server error responses with the appropriate HTTP response codes and headers. Wrapping all responses in these methods also allows for a consistent response schema.

##### [service.go](../service/service.go)

This file contains the web handlers that redirect different requests to various functions. The [Goji](http://goji.io/) framework is used for handling URL parameters. This file also deals with parsing and forwarding query parameters with the help of Golang's built in parse methods.

### Future improvements

This project was completed in a [24 hour hack session](https://github.com/vermagav/next-bart/commits/master), and certain design decisions were taken due to this time constraint. Iterating further, here are some future improvements that should be made:

* Separate the caching mechanism to a `storage` package. Abstract the fact that the cache is held in-memory to within that module, and let the `bart` methods call isolated functions such as `storage.get()` and `storage.put()`. With an isolated storage module, the in-memory map can then be swapped out for persistent data stores.
* Check for changes to station data, and refresh cache by calling the Station API every now and then. The BART API does not publish an endpoint to register a call-back to trigger cache updates, so a pull model will have to do for now. While station data is unlikely to change (e.g. the Civic Center station won't move lat/long), a new station might be added to the list.
* Create a generic HTTP request module and make all calls to external APIs go through it.
* Isolate all distance/math logic inside another package.
* As a summary of the three points above this one, `stations.go` is doing too much work right now. Functionality needs to be stripped out from within, and moved into their own systems as described above.
* Add more unit and integration tests.
* Fix potential race condition which can occur during app startup: 2 concurrent requests try to populate cache at the same time.
* More modular controls over returning 400s vs 500s, with a bubble-up error. The [Google Context Pattern](https://blog.golang.org/context) would be a good fit for such a task.
