package entity

type Section struct {
	ID       string
	ParentID string
	Title    string
	Order    int
	Level    string // page or chapter
	IsPublic bool
}
