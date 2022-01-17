package main

import (
	"net/http"

	api "github.com/rtovey/astro-api/api"
)

func main() {
	http.HandleFunc("/sun/rise-set", api.SunRiseSet)
	http.HandleFunc("/moon/phase", api.LunarPhase)
	http.HandleFunc("/moon/rise-set", api.LunarRiseSet)

	http.ListenAndServe(":8090", nil)
}
