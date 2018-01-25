package chattp2

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// Server represents the chattp2 server, a chat server based on HTTP2
type Server struct {
	httpServer *http.Server
	msgsChan   map[string]chan string
}

// NewServer returns a new chattp2 Server
func NewServer() (*Server, error) {
	var srv http.Server
	flag.BoolVar(&http2.VerboseLogs, "verbose", false, "Verbose HTTP/2 debugging.")
	flag.Parse()
	srv.Addr = "localhost:4430"
	// srv.ConnState = idleTimeoutHook()

	return &Server{
		httpServer: &srv,
		msgsChan:   make(map[string]chan string),
	}, nil
}

// Run chattp2 server. Server is listening for client connections
func (s *Server) Run(ctx context.Context) {
	http.HandleFunc("/connect", s.connectHandler)
	http.HandleFunc("/send", s.sendHandler)

	_ = http2.ConfigureServer(s.httpServer, &http2.Server{})
	go func() {
		log.Printf("chattp2 server listening in %s", s.httpServer.Addr)
		log.Fatal(s.httpServer.ListenAndServeTLS("server.crt", "server.key"))
	}()
	select {}
}

func (s *Server) connectHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		log.Println("Empty user...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.msgsChan[user]; ok {
		log.Printf("User `%s` already exists.", user)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.msgsChan[user] = make(chan string) // Initialize user channel

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	clientGone := w.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case msg := <-s.msgsChan[user]:
			// wait for messages...
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		case <-clientGone:
			log.Printf("Client %v disconnected", r.RemoteAddr)
			delete(s.msgsChan, user)
			return
		}
	}
}

func (s *Server) sendHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	_ = json.NewDecoder(r.Body).Decode(&msg)

	if _, ok := s.msgsChan[msg.Receiver]; !ok {
		log.Printf("User `%s` doesn't exists.", msg.Receiver)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.msgsChan[msg.Receiver] <- msg.Sender + ": " + msg.Content

	w.WriteHeader(http.StatusOK)
}
