package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
	defer cancel()

	result := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		result <- "operation completed"
	}()

	select {
	case res := <-result:
		fmt.Fprintf(w, "Success: %s\n", res)
	case <-ctx.Done():
		http.Error(w, "Timeout: "+ctx.Err().Error(), http.StatusGatewayTimeout)
	}
}

func main() {
	http.HandleFunc("/timeout", handler)
	fmt.Println("server: listening on :8090")
	http.ListenAndServe(":8090", nil)
}
