package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func loveHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Sending you love ❤️✋")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", loveHandler)

	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Println("server: listening on :8090")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	<-stop
	fmt.Println("\nserver: shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server forced to shutdown:", err)
	} else {
		fmt.Println("server exited properly")
	}
}
