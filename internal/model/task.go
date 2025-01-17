package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

// Put this on a new line to enable sync: // `objectbox:"sync"`
type Task struct {
	Id           uint64
	Text         string
	DateCreated  int64
	DateFinished int64
}
