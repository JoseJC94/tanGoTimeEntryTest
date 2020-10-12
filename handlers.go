package main

/*Requests*/

import (
	"encoding/json"
	"net/http"
	"path"
)

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		dataJson, _ := json.Marshal(entries)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	dataJson, err := json.Marshal(entries[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	entry := Entry{}
	json.Unmarshal(body, &entry)
	entries = append(entries, entry)
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	} else {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		reqEntry := Entry{}
		json.Unmarshal(body, &reqEntry)
		if reqEntry.IdTimeEntry != "" { entries[i].IdTimeEntry = reqEntry.IdTimeEntry }
		if reqEntry.Notes != "" { entries[i].Notes = reqEntry.Notes }
		if reqEntry.CreateDate    != "" { entries[i].CreateDate = reqEntry.CreateDate }
		if reqEntry.DueDate != "" { entries[i].DueDate = reqEntry.DueDate }
		updateEntry(id, reqEntry)
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	} else {
		deleteEntry(id)
	}
	w.WriteHeader(200)
	return
}