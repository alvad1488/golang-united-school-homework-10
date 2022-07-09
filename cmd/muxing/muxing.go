package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

//handler for GET '/name/{PARAM}'
func getNameHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", param["PARAM"])
}

//handler for GET '/bad'
func getBadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

//handler for POST '/data + body {PARAM}'
func postDataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%v", string(body))
}

//handler for POST '/headers + Headers{"a":"2", "b":"3"}'
func postHeadersHandler(w http.ResponseWriter, r *http.Request) {
	aHead := r.Header.Get("a")
	bHead := r.Header.Get("b")

	a, err := strconv.Atoi(aHead)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err.Error())
		return
	}

	b, err := strconv.Atoi(bHead)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %v", err.Error())
		return
	}

	res := strconv.Itoa(a + b)
	w.Header().Add("a+b", res)
	w.WriteHeader(http.StatusOK)

}

//default handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getNameHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBadHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", postDataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeadersHandler).Methods(http.MethodPost)
	router.HandleFunc("/", defaultHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}

	Start(host, port)
}
