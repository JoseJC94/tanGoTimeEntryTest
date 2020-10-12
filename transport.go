package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

//Endpoints

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

//Entry

func makeCreateEntryEndpoint(s EntryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEntryRequest)
		msg, err := s.CreateEntry(ctx, req.entry)
		return CreateEntryResponse{Msg: msg, Err: err}, nil
	}
}
func makeGetEntryByIdEndpoint(s EntryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEntryByIdRequest)
		entryDetails, err := s.GetEntryById(ctx, req.Id)
		if err != nil {
			return GetEntryByIdResponse{Entry: entryDetails, Err: "Id not found"}, nil
		}
		return GetEntryByIdResponse{Entry: entryDetails, Err: ""}, nil
	}
}
func makeGetEntriesEndpoint(s EntryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(GetEntriesRequest)
		entries, err := s.GetEntries(ctx)
		if err != nil {
			return GetEntriesResponse{Entries: entries, Err: "Not found"}, nil
		}
		return GetEntriesResponse{Entries: entries, Err: ""}, nil
	}
}

func makeDeleteEntryEndpoint(s EntryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEntryRequest)
		msg, err := s.DeleteEntry(ctx, req.Entryid)
		if err != nil {
			return DeleteEntryResponse{Msg: msg, Err: err}, nil
		}
		return DeleteEntryResponse{Msg: msg, Err: nil}, nil
	}
}
func makeUpdateEntryendpoint(s EntryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEntryRequest)
		msg, err := s.UpdateEntry(ctx, req.entry)
		return msg, err
	}
}

func decodeCreateEntryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateEntryRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.entry); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEntryByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetEntryByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetEntryByIdRequest{
		Id: vars["entryid"],
	}
	return req, nil
}
func decodeGetEntriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetEntriesRequest
	fmt.Println("-------->>>>into GetById Decoding")
	req = GetEntriesRequest{
	}
	return req, nil
}

func decodeDeleteEntryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteEntryRequest
	vars := mux.Vars(r)
	req = DeleteEntryRequest{
		Entryid: vars["entryid"],
	}
	return req, nil
}
func decodeUpdateEntryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req.entry); err != nil {
		return nil, err
	}
	return req, nil
}


//request response

//Entry

type (
	CreateEntryRequest struct {
		entry Entry
	}
	CreateEntryResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	GetEntryByIdRequest struct {
		Id string `json:"entryid"`
	}
	GetEntryByIdResponse struct {
		Entry interface{} `json:"entry,omitempty"`
		Err  string      `json:"error,omitempty"`
	}
	GetEntriesRequest struct {
	}
	GetEntriesResponse struct {
		Entries interface{} `json:"entries,omitempty"`
		Err   string      `json:"error,omitempty"`
	}

	DeleteEntryRequest struct {
		Entryid string `json:"entryid"`
	}

	DeleteEntryResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	UpdateEntryRequest struct {
		entry Entry
	}
	UpdateEntryResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
)
