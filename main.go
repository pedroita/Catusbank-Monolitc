package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	controller "cactusbank.com.br/cactusbank/src/controllers"
	"github.com/gorilla/mux"
)

type MiddlewareFunc func(http.Handler) http.Handler

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Response-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()

	router.Use(middleware)
	router.HandleFunc("/accounts", controller.GetAllAccounts).Methods("GET")
	router.HandleFunc("/account/{id}", controller.GetAccount).Methods("GET")
	router.HandleFunc("/account/remove/{id}", controller.RemoveAccount).Methods("DELETE")
	router.HandleFunc("/account/update/{id}", controller.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/create", controller.CreateAccount).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	fmt.Println("Shutting down server")
	os.Exit(0)

}
