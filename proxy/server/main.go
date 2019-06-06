package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bakaoh/x/proxy/core/proxy"
	"google.golang.org/grpc"
)

func main() {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// // Make sure we never forward internal services.
		// if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
		// 	return ctx, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
		// }
		// md, ok := metadata.FromIncomingContext(ctx)
		// if ok {
		// 	// Decide on which backend to dial
		// 	if val, exists := md[":authority"]; exists && val[0] == "staging.api.example.com" {
		// 		// Make sure we use DialContext so the dialing can be cancelled/time out together with the context.
		// 		return ctx, grpc.DialContext(ctx, "api-service.staging.svc.local", grpc.WithCodec(proxy.Codec()))
		// 	} else if val, exists := md[":authority"]; exists && val[0] == "api.example.com" {
		// 		return ctx, grpc.DialContext(ctx, "api-service.prod.svc.local", grpc.WithCodec(proxy.Codec()))
		// 	}
		// }

		fmt.Println("director")
		conn, err := grpc.DialContext(ctx, "unix:///tmp/echo.sock", grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())

		return ctx, conn, err
	}

	server := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
