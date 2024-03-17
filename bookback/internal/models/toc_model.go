package models

type TableOfContents struct {
	BookID    string `json:"book_id"`
	BookTitle string `json:"book_title"`

	Author     string `json:"author"`
	Authorship string `json:"authorship"`

	Tags []string `json:"tags"`

	Sections []*Section `json:"sections"`
}

type Section struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
	Level    string `json:"level"` // page or chapter
	IsPublic bool   `json:"is_public"`
}
