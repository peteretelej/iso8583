package main

import (
	"flag"
	"log"
	"os"

	"github.com/peteretelej/iso8583"
)

var (
	dir    = flag.String("dir", "../..", "directory to server from")
	listen = flag.String("listen", ":8080", "http server listen address")
)

func main() {
	flag.Parse()
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
