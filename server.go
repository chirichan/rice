package rice

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// NewHttpServer -.
func NewHttpServer(handler http.Handler, opts ...ServerOption) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}

// ServerOption -.
type ServerOption func(*Server)

// Port -.
func Port(port string) ServerOption {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// NewGrpcServer -.
func NewGrpcServer(ctx context.Context, port string, server *grpc.Server) {

	if l, err := net.Listen("tcp", fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("error in listening on port %v, %v", port, err)
	} else {

		go func() {
			<-ctx.Done()
			err := l.Close()
			if err != nil {
				log.Fatalf("Failed to close gRPC server :%v, %v", port, err)
			}
		}()

		if err := server.Serve(l); err != nil {
			log.Fatalf("unable to start server, %v", err)
		}
	}
}
