package main

import (
	"context"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), "userID", 13)

	processRequest(ctx, w)
}

func processRequest(ctx context.Context, w http.ResponseWriter) {
	userID := ctx.Value("userID")

	if userID != nil {
		fmt.Fprintf(w, "Got userID from context: %v\n", userID)
	} else {
		fmt.Fprintf(w, "No userID found in context\n")
	}
}

func main() {
	http.HandleFunc("/value", handler)
	fmt.Println("server: listening on :8090")
	http.ListenAndServe(":8090", nil)
}
