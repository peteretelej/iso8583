package iso8583

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// WebServer returns a http web server that interprets ISO8583 messages
func WebServer(listenAddr string) (*http.Server, error) {
	if _, err := os.Stat(filepath.Join("web", "index.html")); err != nil {
		return nil, fmt.Errorf("missing web/ folder with index.html in launch directory")
	}
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)
	svr := &http.Server{
		Addr:           listenAddr,
		ReadTimeout:    time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return svr, nil
}

// Listen returns a tcp server listening on a specified address
func Listen(listenAddr string) error {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	var timeout = time.Minute * 5
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleListenConnection(conn, time.Now().Add(timeout))
	}
}

// simply an echo server atm
func handleListenConnection(conn net.Conn, deadline time.Time) {
	defer func() { _ = conn.Close() }()
	_ = conn.SetDeadline(deadline) // r+w deadline
	_, _ = io.Copy(conn, conn)
}
