package main

import (
	"flag"
	"log"
	"os"

	"github.com/peteretelej/iso8583"
)

var (
	socket = flag.Bool("socket", false, "starts a listener on a socket")

	// server flags
	dir    = flag.String("dir", "../..", "directory to server from")
	listen = flag.String("listen", ":8080", "http server listen address")
)

func main() {
	flag.Parse()
	if *socket {
		launchSocket()
		return
	}
	serve()
}

func serve() {
	if err := os.Chdir(*dir); err != nil {
		log.Fatal(err)
	}
	svr, err := iso8583.WebServer(*listen)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("launching web server at %s", *listen)
	log.Fatal(svr.ListenAndServe())
}
func launchSocket() {
	log.Printf("launching socket listener at: %s", *listen)
	log.Fatal(iso8583.Listen(*listen))
}
