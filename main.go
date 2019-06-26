package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var host string
var port int

func main() {
	flag.StringVar(&host, "host", "127.0.0.1", "Host to serve from")
	flag.IntVar(&port, "port", 8126, "Port to listen on")

	router := mux.NewRouter()

	router.Handle("/{_:.*}", http.HandlerFunc(echo)).Methods("GET", "POST", "PUT")

	server := &http.Server{

		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      router,
		ReadTimeout:  viper.GetDuration("read_timeout"),
		WriteTimeout: viper.GetDuration("write_timeout"),
	}

	fmt.Printf("Listening on %s:%d", host, port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("error while serving http: %v", err)
		os.Exit(1)
	}

}

func echo(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("METHOD: %s\n", req.Method)
	fmt.Printf("PATH: %s\n", req.RequestURI)
	for k, v := range req.Header {
		fmt.Printf("> %s: %s\n", k, strings.Join(v, ", "))
	}


	var body []byte
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Printf("ERROR READING BODY! %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s\n\n", body)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
