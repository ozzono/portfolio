package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var Users = map[string]User{
	"John": {ID: 1, Name: "John", Role: "admin"},
	"Juan": {ID: 2, Name: "Juan", Role: "developer"},
}
