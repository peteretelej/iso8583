package iso8583

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// WebServer returns a http web server that interprets ISO8583 messages
func WebServer(listenAddr string) (*http.Server, error) {
	if _, err := os.Stat(filepath.Join("web", "index.html")); err != nil {
		return nil, fmt.Errorf("missing web/ folder with index.html in launch directory")
	}
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	api := API{}
	http.Handle("/api/", api)
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

// API handles API requests
type API struct{}

// APIResponse defines the structure of API responses
type APIResponse struct {
	Code     int         `json:"code"`
	Err      bool        `json:"error"`
	Response string      `json:"response"`
	Data     interface{} `json:"data"`
}

func (a APIResponse) render(w http.ResponseWriter) {
	if a == (APIResponse{}) {
		a = APIResponse{
			Err:  true,
			Code: http.StatusServiceUnavailable,
		}
	}
	if a.Code == 0 && a.Err != false {
		a.Code = http.StatusInternalServerError
	}
	if a.Code == 0 {
		a.Code = http.StatusOK
	}
	if a.Response == "" {
		a.Response = http.StatusText(a.Code)
	}
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"code":500,"error":"Internal Server Error"}`)
		return
	}
	w.WriteHeader(a.Code)
	_, _ = w.Write(dat)

}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ares := APIResponse{}
	path := strings.TrimLeft(r.URL.Path, "/api/")
	if strings.TrimSpace(path) == "" {
		ares.Code = http.StatusBadRequest
		ares.render(w)
		return

	}
	parts := strings.Split(path, "/")
	var resource, entity string
	resource = parts[0]
	if len(parts) > 1 {
		entity = parts[1]
	}
	_ = entity // TODO: use entity for optional api param
	switch resource {
	case "bitmaptobin":
		bitmap := r.FormValue("msg")
		if bitmap == "" {
			ares.Code = http.StatusBadRequest
			ares.Response = "missing bitmap value in form field msg"
			ares.render(w)
			return
		}
		bin, err := BitmapToBinary(bitmap)
		if err != nil {
			ares.Code = http.StatusBadRequest
			ares.Response = "invalid bitmap provided, not in hexadecimal"
			ares.render(w)
			return
		}
		ares.Code = http.StatusOK
		ares.Data = struct {
			Value, Result string
		}{bitmap, bin}
		ares.render(w)
		return
	default:
		ares.Code = http.StatusServiceUnavailable
		ares.render(w)
		return
	}

}
