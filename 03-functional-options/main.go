package main

import (
	"fmt"
	"time"
)

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
	Retries int
}

type Option func(s *Server)

func WithRetries(retries int) Option {
	return func(s *Server) {
		if retries < 1 {
			fmt.Println("max retries can be minimum 1, setting it to 1")
			retries = 1
		}
		if retries > 10 {
			fmt.Println("max retries can be maximum 10, setting it to 10")
			retries = 10
		}
		s.Retries = retries
	}
}

func WithTimeout(duration time.Duration) Option {
	return func(s *Server) {
		if duration > 10*time.Second {
			fmt.Println("max duration for timeout can be 10 seconds, setting 10 Seconds")
			duration = 10 * time.Second
		}
		if duration < time.Second {
			fmt.Println("min duratoin for timeout can be 1 second, setting 1 Second")
			duration = time.Second
		}
		s.Timeout = duration
	}
}

func NewServer(host string, port int, options ...Option) *Server {
	server := &Server{
		Host: host,
		Port: port,
	}

	for _, opt := range options {
		opt(server)
	}

	return server
}

func main() {
	// [PROBLEM]
	// 	db := NewDatabase(
	//     "localhost",
	//     5432,
	//     true,
	//     5*time.Second,
	//     3,
	//     logger,
	//     true,
	//     false,
	//     cache,
	//     metrics,
	//     auth,
	//     limiter,
	// )

	// -----> No body knows what each parameter actually mean

	server := NewServer(
		"localhost",
		8080,
		WithRetries(20),
		WithTimeout(time.Second*10),
	)
	fmt.Printf("%+v\n", *server)
}
