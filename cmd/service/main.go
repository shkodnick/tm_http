package main 

import (
	"net/http"

	"app/internal/router"
)

func main() {
	http.ListenAndServe(":8080",router.InitRouter())
}