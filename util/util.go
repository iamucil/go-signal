package util

type Account struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Type      string `json:"type"`
}

type Blog struct {
	ID      string `json:"id,omitempty"`
	Account string `json:"account"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}
