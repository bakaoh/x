package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	echo "github.com/bakaoh/x/proxy/idl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

// Echo ...
func (s *server) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: "Hello " + in.Message}, nil
}

func main() {
	lis, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Handle common process-killing signals so we can gracefully shut down:
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		// Stop listening (and unlink the socket if unix type):
		lis.Close()
		// And we're done:
		os.Exit(0)
	}(sigc)

	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
