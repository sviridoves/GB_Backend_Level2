package models

type Environment struct {
	ID    int64
	Type  string
	Name  string
	Users []User
}
