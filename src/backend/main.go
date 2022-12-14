package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	_id    = []byte("id: ")
	_event = []byte("event: ")
	_data  = []byte("data: ")
	_delim = []byte("\n")
	_end   = []byte("\n\n")
)

// EventSSE - event message
type EventSSE struct {
	ID        string
	EventType string
	Data      string
}

// ClientSSE - event client
type ClientSSE struct {
	w       io.Writer
	flusher http.Flusher
}

// Write - io.Writer
func (c ClientSSE) Write(b []byte) (int, error) { return c.w.Write(b) }

// Flush - http.flusher.Flush()
func (c ClientSSE) Flush() { c.flusher.Flush() }

// PoolSSE - pool of ClientSSE
type PoolSSE struct {
	pool map[ClientSSE]struct{}
	lock sync.Mutex
}

// Add - add ClientSSE
func (p *PoolSSE) Add(c ClientSSE) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.pool[c] = struct{}{}
	log.Println("pool size:", p.Len())
}

// Del - delete ClientSSE
func (p *PoolSSE) Del(c ClientSSE) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.pool, c)
	log.Println("pool size:", p.Len())
}

// WriteMessage - write message to clients
func (p *PoolSSE) WriteMessage(m EventSSE) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for client := range p.pool {
		go func(c ClientSSE) {
			if _, err := m.WriteTo(c); err != nil {
				delete(p.pool, c)
				return
			}
			c.Flush()
		}(client)
	}
}

// Len - count of ClientSSE
func (p *PoolSSE) Len() int { return len(p.pool) }

// NewPoolSSE - new PoolSSE
func NewPoolSSE() *PoolSSE {
	return &PoolSSE{pool: make(map[ClientSSE]struct{})}
}

// WriteTo - io.WriterTo
func (e EventSSE) WriteTo(w io.Writer) (n int64, err error) {
	for _, msgPart := range [][]byte{
		_id, []byte(e.ID), _delim,
		_event, []byte(e.EventType), _delim,
		_data, []byte(e.Data), _end,
	} {
		_n, err := w.Write(msgPart)
		if err != nil {
			return n, err
		}
		n += int64(_n)
	}
	return
}

func (e EventSSE) String() string {
	var sb strings.Builder
	e.WriteTo(&sb)
	return sb.String()
}

func genPayload() string {
	var buf = make([]byte, 16)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

func initUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Domain:   "localhost",
		Path:     "/api",
		Name:     "user",
		Value:    "13131",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func events() http.HandlerFunc {
	var chahEvents = make(chan EventSSE)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for t := range ticker.C {
			b, _ := json.Marshal(map[string]string{
				"time":    t.Format(time.RFC3339),
				"payload": genPayload(),
			})
			chahEvents <- EventSSE{
				ID:        strconv.FormatInt(t.UnixNano(), 10),
				EventType: "message",
				Data:      string(b),
			}
		}
	}()
	var clientsPool = NewPoolSSE()
	go func() {
		for msg := range chahEvents {
			clientsPool.WriteMessage(msg)
		}
	}()
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			log.Println("not flusher")
			return
		}
		w.Write([]byte("retry: 2000\n\n"))
		flusher.Flush()
		client := ClientSSE{w: w, flusher: flusher}
		clientsPool.Add(client)
		<-ctx.Done()
		clientsPool.Del(client)
		log.Println("client left")
	}
}

// FileSystemSPA - обертка над fs.FileSystem для SPA приложения
func FileSystemSPA(dirname string) fs.FS { return &spaFS{os.DirFS(dirname)} }

type spaFS struct {
	hfs fs.FS
}

func (sfs spaFS) Open(name string) (fs.File, error) {
	if _, err := fs.Stat(sfs.hfs, name); err != nil && os.IsNotExist(err) {
		return sfs.hfs.Open("index.html")
	}
	return sfs.hfs.Open(name)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(FileSystemSPA("static"))))
	mux.HandleFunc("/api/init", initUser)
	mux.Handle("/api/events", events())
	log.Println("server start om: http://0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}
