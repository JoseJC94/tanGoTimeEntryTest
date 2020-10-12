package main

type Entry struct {

	IdTimeEntry string `json:"idTimeEntry,omitempty"`

	Notes string `json:"notes,omitempty"`

	CreateDate string `json:"createDate,omitempty"`

	DueDate string `json:"dueDate,omitempty"`
}
