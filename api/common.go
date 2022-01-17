package api

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	c "github.com/rtovey/astro-lib/common"
)

func observer(req *http.Request) (c.Observer, error) {
	latitude, err_lat := strconv.ParseFloat(req.URL.Query().Get("lat"), 64)
	longitude, err_long := strconv.ParseFloat(req.URL.Query().Get("long"), 64)
	zone := req.URL.Query().Get("zone")

	if err_lat != nil || math.Abs(latitude) > 90.0 {
		return c.Observer{}, errors.New("Invalid latitude specified")
	}
	if err_long != nil || math.Abs(longitude) > 180.0 {
		return c.Observer{}, errors.New("Invalid longitude specified")
	}
	if len(zone) == 0 {
		zone = "UTC"
	}

	loc, err_loc := time.LoadLocation(zone)

	if err_loc != nil || loc == nil {
		return c.Observer{}, errors.New("Invalid timezone specified")
	}

	return c.Observer{
		Latitude:  latitude,
		Longitude: longitude,
		Location:  loc,
	}, nil
}
