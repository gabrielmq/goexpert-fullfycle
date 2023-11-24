package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World!"))
	})

	serverError := make(chan error, 1)
	go func() {
		fmt.Println("server is running at port :3000")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				serverError <- err
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverError:
		log.Fatalf("could not listen on %s: %v", server.Addr, err)
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		fmt.Println("shutting down server...")
		if err := server.Shutdown(ctx); err != nil {
			if err == context.DeadlineExceeded {
				log.Fatal("forced shutdown completed")
			}
			log.Fatalf("could not gracefully shutdown the server: %v", err)
		}

		fmt.Println("server shutdown gracefully")
	}
}
