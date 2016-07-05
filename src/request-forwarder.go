package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
)

var listenHost string
var listenPort int
var forwardDomain string
var forwardPort int

func init() {
	flag.StringVar(&listenHost, "listen-host",
		"localhost", "specify the hostname where the request should get forwarded to")
	flag.IntVar(&listenPort, "listen-port", 8080,
		"the port where the request should get forwarded to")
	flag.StringVar(&forwardDomain, "forward-domain",
		"http://localhost",
		"specify the scheme and domain where the request should get forwarded to http://localhost")
	flag.IntVar(&forwardPort, "forward-port", 2379,
		"the port where the request should get forwarded to")

	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	forwardD := &forwardDomain
	forwardP := &forwardPort

	resp, err := http.Get(*forwardD + ":" + strconv.Itoa(*forwardP) + r.URL.Path)

	if err != nil {
		log.Fatal(err)
	}

	header := resp.StatusCode
	if header == 405 {
		header = 200
	}
	w.WriteHeader(header)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

func main() {
	listenP := &listenPort
	listenH := &listenHost
	http.HandleFunc("/", handler)
	log.Println("Going to listen on: ", *listenH+":"+strconv.Itoa(*listenP))
	err := http.ListenAndServe(*listenH+":"+strconv.Itoa(*listenP), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
