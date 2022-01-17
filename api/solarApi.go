package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rtovey/astro/solar"
)

func SunRiseSet(w http.ResponseWriter, req *http.Request) {
	date := time.Now()
	observer, err := observer(req)
	if err != nil {
		log.Printf("Unable to calculate sun rise/set time as specified observer is invalid: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	sunRiseSetTime := solar.RiseSetTime(observer, date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sunRiseSetTime)
}
