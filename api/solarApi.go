package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rtovey/astro-lib/solar"
)

func SunRiseSet(w http.ResponseWriter, req *http.Request) {
	logRequest("sun rise/set time")

	date := time.Now()
	observer, err := observer(req)
	if err != nil {
		log.Printf("Unable to calculate sun rise/set time as specified observer is invalid: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	sunRiseSetTime := solar.RiseSetTime(observer, date)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sunRiseSetTime)
}
