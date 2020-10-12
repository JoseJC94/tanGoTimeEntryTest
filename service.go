package main

import (
	"context"
	"github.com/go-kit/kit/log"
)

type entrieservice struct {
	logger log.Logger
}

// Service describes the Entry service.
type EntryService interface {
	CreateEntry(ctx context.Context, entry Entry) (string, error)
	GetEntryById(ctx context.Context, id string) (interface{}, error)
	GetEntries(ctx context.Context) (interface{}, error)
	UpdateEntry(ctx context.Context, entry Entry) (string, error)
	DeleteEntry(ctx context.Context, id string) (string, error)
}


func findEntry(x string) int {
	for i, entry := range entries {
		if x == entry.IdTimeEntry {
			return i
		}
	}
	return -1
}

func NewEntryService(logger log.Logger) EntryService {
	return &entrieservice{
		logger: logger,
	}
}

///Entries
func (b entrieservice) CreateEntry(ctx context.Context, entry Entry) (string, error) {
	var msg = "success"
	entries = append(entries, entry)
	return msg, nil
}

func (b entrieservice) GetEntryById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var entry interface{}
	var empty interface{}
	i := findEntry(id)
	if i == -1 {
		return empty, err
	}
	entry = entries[i]
	return entry, nil
}

func (b entrieservice) GetEntries(ctx context.Context) (interface{}, error) {
	var err error
	var entry interface{}
	var empty interface{}
	if entries == nil {
		return empty, err
	}
	entry = entries
	return entry, nil
}

func (b entrieservice) UpdateEntry(ctx context.Context, entry Entry) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := findEntry(entry.IdTimeEntry)
	if i == -1 {
		return empty, err
	}
	entries[i] = entry
	return msg, nil
}

func (b entrieservice) DeleteEntry(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := findEntry(id)
	if i == -1 {
		return "", err
	}
	copy(entries[i:], entries[i+1:])
	entries[len(entries)-1] = Entry{}
	entries = entries[:len(entries)-1]
	return msg, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// DATA

	var entries = []Entry{
		{
			IdTimeEntry: "1", Notes: "first entry in Tango", CreateDate: "2020-09-06",
			DueDate: "2020-09-16",
		},
		{
			IdTimeEntry: "2", Notes: "second entry in Tango", CreateDate: "2020-09-06",
			DueDate: "2020-09-17",
		},
	}
