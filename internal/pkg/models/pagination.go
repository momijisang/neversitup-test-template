package models

type Pagination struct {
	Data   interface{} `json:"data"`
	Header interface{} `json:"header"`
	Order  interface{} `json:"order"`
	Links  LinkUrl     `json:"links"`
	Meta   Meta        `json:"meta"`
}

type LinkUrl struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type Meta struct {
	CurrentPage int        `json:"current_page"`
	LastPage    int        `json:"last_page"`
	From        int        `json:"from"`
	To          int        `json:"to"`
	Path        string     `json:"path"`
	PerPage     int        `json:"per_page"`
	Total       int        `json:"total"`
	Links       []LinkPage `json:"links"`
}

type LinkPage struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
