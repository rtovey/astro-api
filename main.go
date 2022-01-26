package main

import (
	"net/http"

	api "github.com/rtovey/astro-api/api"
)

func main() {
	http.HandleFunc("/api/sun/rise-set", api.SunRiseSet)
	http.HandleFunc("/api/moon/phase", api.LunarPhase)
	http.HandleFunc("/api/moon/rise-set", api.LunarRiseSet)

	http.ListenAndServe(":8090", nil)
}
