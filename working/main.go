package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/**
	handle func => convenience method on the go http package
	registers a function to a path on thing called default serve MUX
	so is the default serve MUX

	default serve MUX is an HTTP handler

	HandleFunc mendaftarkan fungsi penangan untuk pola yang diberikan dalam DefaultServeMux. Dokumentasi untuk ServeMux menjelaskan bagaimana pola dicocokkan.
	*/

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Ohayo")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data %s\n", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye yo")
	})

	http.ListenAndServe(":9090", nil)

}
