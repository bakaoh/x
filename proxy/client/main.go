package main

import (
	"context"
	"fmt"

	echo "github.com/bakaoh/x/proxy/idl"
	"google.golang.org/grpc"
)

func main() {
	//	conn, _ := grpc.Dial("unix:///tmp/echo.sock", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cli := echo.NewEchoServiceClient(conn)
	res, err := cli.Echo(context.Background(), &echo.EchoRequest{Message: "Baka"})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.GetMessage())
}
