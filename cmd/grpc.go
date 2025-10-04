package cmd

import (
	"log"
	"net"

	"ewallet-ums/helpers"

	"google.golang.org/grpc"
)

func ServeGRPC() {
	netlisten, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "3000"))
	if err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}

	grpcServer := grpc.NewServer()
	// depedency.TokenValidation.UnimplementedTokenValidationServer.ValidateToken()

	if err := grpcServer.Serve(netlisten); err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}
}
