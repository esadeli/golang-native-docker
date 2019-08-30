package main

import (
	"fmt"
	"log"
	"net/http"

	health "github.com/esadeli/golang-native-docker/controllers/health"
	user "github.com/esadeli/golang-native-docker/controllers/user"
)

func main() {
	var PORT = "5051"
	http.HandleFunc("/", health.Check)
	http.HandleFunc("/user", user.GetUserData)

	fmt.Println("Listening to PORT:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
