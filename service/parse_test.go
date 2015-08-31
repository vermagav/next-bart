package service

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Boilerplate
func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

// Test parseCount()
func (s *MySuite) TestParseCount(c *C) {
	var in string
	var out int
	var err error

	// Test valid values

	in = ""
	out, err = parseCount(in)
	c.Assert(out, Equals, 0)
	c.Assert(err, IsNil)

	in = "123"
	out, err = parseCount(in)
	c.Assert(out, Equals, 123)
	c.Assert(err, IsNil)

	in = "1"
	out, err = parseCount(in)
	c.Assert(out, Equals, 1)
	c.Assert(err, IsNil)

	in = "10"
	out, err = parseCount(in)
	c.Assert(out, Equals, 10)
	c.Assert(err, IsNil)

	// Test invalid values

	in = "1a"
	out, err = parseCount(in)
	c.Assert(out, Equals, 0)
	c.Assert(err, Equals, errParsingQueryParams)

	in = "1.5"
	out, err = parseCount(in)
	c.Assert(out, Equals, 0)
	c.Assert(err, Equals, errParsingQueryParams)

	in = "0"
	out, err = parseCount(in)
	c.Assert(out, Equals, 0)
	c.Assert(err, Equals, errCountInvalid)

	in = "-1"
	out, err = parseCount(in)
	c.Assert(out, Equals, 0)
	c.Assert(err, Equals, errCountInvalid)
}

// Test parseLatLong()
func (s *MySuite) TestParseLatLong(c *C) {
	var inLat, inLong string
	var lat, long float64
	var err error

	// Test valid values

	inLat = ""
	inLong = ""
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, IsNil)

	inLat = "123"
	inLong = "123"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 123.)
	c.Assert(long, Equals, 123.)
	c.Assert(err, IsNil)

	inLat = "0.0"
	inLong = "0.0"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, IsNil)

	inLat = "-0.123123"
	inLong = "-100.111111"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, -0.123123)
	c.Assert(long, Equals, -100.111111)
	c.Assert(err, IsNil)

	// Test invalid values

	inLat = "123.0"
	inLong = ""
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, Equals, errParsingQueryParams)

	inLat = ""
	inLong = "123.0"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, Equals, errParsingQueryParams)

	inLat = "1.0a"
	inLong = "1.0a"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, Equals, errParsingQueryParams)

	inLat = "1.0"
	inLong = "1.0a"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, Equals, errParsingQueryParams)

	inLat = "1.0a"
	inLong = "1.0"
	lat, long, err = parseLatLong(inLat, inLong)
	c.Assert(lat, Equals, 0.)
	c.Assert(long, Equals, 0.)
	c.Assert(err, Equals, errParsingQueryParams)
}
