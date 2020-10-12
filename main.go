//package tarea_4_sistemas_distribuidos_JoseJC94

package main

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	r := mux.NewRouter()
	//var pubs, authors, entries = defaultData()
	var svcEntry EntryService
	svcEntry = NewEntryService(logger)

	// svcEntry = loggingMiddleware{logger, svcEntry}
	// svcEntry = instrumentingMiddleware{requestCount, requestLatency, countResult, svcEntry}

	//entry
	CreateEntryHandler := httptransport.NewServer(
		makeCreateEntryEndpoint(svcEntry),
		decodeCreateEntryRequest,
		encodeResponse,
	)
	GetByEntryIdHandler := httptransport.NewServer(
		makeGetEntryByIdEndpoint(svcEntry),
		decodeGetEntryByIdRequest,
		encodeResponse,
	)

	GetEntriesHandler := httptransport.NewServer(
		makeGetEntriesEndpoint(svcEntry),
		decodeGetEntriesRequest,
		encodeResponse,
	)
	DeleteEntryHandler := httptransport.NewServer(
		makeDeleteEntryEndpoint(svcEntry),
		decodeDeleteEntryRequest,
		encodeResponse,
	)
	UpdateEntryHandler := httptransport.NewServer(
		makeUpdateEntryendpoint(svcEntry),
		decodeUpdateEntryRequest,
		encodeResponse,
	)

	http.Handle("/", r)
	//http.Handle("/entry", CreateEntryHandler)
	r.Handle("/entry", CreateEntryHandler).Methods("POST")
	r.Handle("/entry", GetEntriesHandler).Methods("GET")
	http.Handle("/entry/update", UpdateEntryHandler)
	r.Handle("/entry/{entryid}", GetByEntryIdHandler).Methods("GET")
	r.Handle("/entry/{entryid}", DeleteEntryHandler).Methods("DELETE")
	///r.Handle("/entry", GetEntriesHandler).Methods("GET")

	// http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
	logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
