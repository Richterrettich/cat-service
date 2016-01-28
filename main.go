package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Path("/cats").Methods("POST", "PUT").Handler(buildCreateCatHandler())
	r.Path("/cats").Methods("GET").Handler(buildListCatsHandler())
	r.Path("/cats/{id}").Methods("GET").Handler(buildGetCatHandler())
	r.Path("/cats/{id}").Methods("DELETE").Handler(buildGetCatHandler())

	r.Path("/feedings").Methods("POST", "PUT").Handler(buildCreateFeedingHandler())
	r.Path("/feedings").Methods("GET").Handler(buildFeedingsRangeHandler())
	r.Path("/feedings/{id}").Methods("GET").Handler(buildGetSingleFeedingHandler())
	r.Path("/feedings/{id}").Methods("DELETE").Handler(buildDeleteFeedingHandler())

	r.Path("/incidents").Methods("POST", "PUT").Handler(buildCreateIncidentHandler())
	r.Path("/incidents").Methods("GET").Handler(buildGetIncidentsRangeHandler())
	r.Path("/incidents/{id}").Methods("GET").Handler(buildGetSingleIncidentsHandler())
	r.Path("/incidents/{id}").Methods("DELETE").Handler(buildDeleteIncidentsHandler())
}

func buildCreateFeedingHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}

func buildCreateIncidentHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}

	return http.Handler(http.HandlerFunc(result))
}

func buildFeedingsRangeHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}

func buildGetSingleFeedingHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}

func buildUpdateFeedingHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}
func buildDeleteFeedingHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}

func buildGetIncidentsRangeHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}

func buildGetSingleIncidentsHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}
func buildUpdateIncidentsHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}

func buildDeleteIncidentsHandler() http.Handler {
	result := func(w http.ResponseWriter, r *http.Request) {
	}
	return http.Handler(http.HandlerFunc(result))
}
