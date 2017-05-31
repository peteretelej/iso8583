package iso8583

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Serve launches a http web server that interprets ISO8583 messages
func Serve(listenAddr string) error {
	if _, err := os.Stat(filepath.Join("web", "index.html")); err != nil {
		return fmt.Errorf("missing web/ folder with index.html in launch directory")
	}
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)
	svr := &http.Server{
		Addr:           listenAddr,
		ReadTimeout:    time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("launching http server on %s", listenAddr)
	return svr.ListenAndServe()
}
