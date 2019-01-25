package main

type Note struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var notes = []*Note{
	{ID: "1", Content:"notitie 1"},
	{ID: "2", Content:"notitie 2"},
	{ID: "3", Content:"notitie 3"},
}
