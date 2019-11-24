package domain

import "net/http"

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Tel       int    `json:"tel"`
}

func (u *User) Bind(r *http.Request) error {
	return nil
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
