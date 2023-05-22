package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	finish := make(chan bool)

	// main server, serve a static site with js under ./static
	server80 := http.NewServeMux()
	index := http.FileServer(http.Dir("./static"))
	server80.Handle("/", index)
	server80.HandleFunc("/_healthz", k8probe)
	server80.HandleFunc("/data", appData)

	// exposing application metrics for Prometheus
	server9000 := http.NewServeMux()
	server9000.Handle("/metrics", promhttp.Handler())

	// exposing port 80 for main page(index) and _healthz for probes
	go func() {
		http.ListenAndServe(":80", server80)
	}()

	// exposing port 9000 for prometheus
	go func() {
		http.ListenAndServe(":9000", server9000)
	}()

	<-finish
}

// handler function for kubernetes Readiness and Liveness probes
func k8probe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
	log.Println("Probe checked")
	data, err := httputil.DumpRequest(r, false)
	if err != nil {
		log.Fatal("Internal server error")
	}
	fmt.Printf("%s", string(data))
}

// handler function for importing data from the Backend
func appData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	backEndUrl := os.Getenv("BACKENDURL")
	log.Println("query backend for data at", backEndUrl)

	resp, err := http.Get(backEndUrl)
	if err != nil {
		w.WriteHeader(500)
		log.Fatal("failed to access backend at", err)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}
