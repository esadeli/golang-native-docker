package user

import (
	"errors"
	"fmt"
	"net/http"
)

// User struct to mock JSON structure
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password int    `json:"password"`
}

// GetUser function to get user data
func GetUser(w http.ResponseWriter, r *http.Request) (User, error) {
	us := User{}
	idkeys, ok := r.URL.Query()["id"]

	if !ok || len(idkeys[0]) < 1 {
		fmt.Println("URL Query key is missing")
		return us, errors.New("URL_QUERY_MISSING_ERROR")
	}

	id := idkeys[0]

	if id != "xendit" {
		fmt.Printf("%s is invalid objectid ", id)
		return us, errors.New("INVALID_OBJECT_ID_ERROR")
	}

	us = User{
		ID:       id,
		Name:     "PT. Sinar Digital Terdepan",
		Role:     "Payment Gateway",
		Password: 94,
	}

	return us, nil
}
