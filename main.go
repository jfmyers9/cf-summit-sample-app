package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello CF Summit!\n")
	})

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("failed to listen and serve: %s", err.Error())
		os.Exit(1)
	}
}
