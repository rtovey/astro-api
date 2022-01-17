package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rtovey/astro-lib/lunar"
)

func LunarPhase(w http.ResponseWriter, req *http.Request) {
	date := time.Now()
	phase := lunar.Phase(date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phase)
}

func LunarRiseSet(w http.ResponseWriter, req *http.Request) {
	date := time.Now()
	observer, err := observer(req)
	if err != nil {
		log.Printf("Unable to calculate lunar rise/set time as specified observer is invalid: %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	lunarRiseSetTime := lunar.RiseSetTime(observer, date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lunarRiseSetTime)
}
