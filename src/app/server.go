package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gazelle0130/go-mongo-playground/src/app/infrastructure"
)

func main() {
	http.ListenAndServe(os.Getenv("APP_SERVER_PORT"), infrastructure.Router)
	fmt.Println("server started on :3000")
}
