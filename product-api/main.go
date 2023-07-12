package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MhdAmir/golang_microservices/handlers"
)

func main() {
	/**
	handle func => convenience method on the go http package
	registers a function to a path on thing called default serve MUX
	so is the default serve MUX

	default serve MUX is an HTTP handler

	HandleFunc mendaftarkan fungsi penangan untuk pola yang diberikan dalam DefaultServerMux. Dokumentasi untuk ServeMux menjelaskan bagaimana pola dicocokkan.
	*/

	//http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// log.Println("Ohayo")
	// d, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	http.Error(rw, "Oops", http.StatusBadRequest)
	// 	//same with below
	// 	// rw.WriteHeader(http.StatusBadRequest)
	// 	// rw.Write([]byte("Oopps"))
	// 	return
	// }

	// //feedback data from user to server
	// log.Printf("Data %s\n", d)

	// //feedback data from user to User using Fprintf rw(http.ResponseWriter)
	// fmt.Fprintf(rw, "Hello %s", d)
	// })
	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Bye yo")
	// })

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
