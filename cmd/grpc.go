package cmd

import (
	"log"
	"net"

	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"

	"google.golang.org/grpc"
)

func ServeGRPC() {
	depedency := depedencyInject()

	netlisten, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "3000"))
	if err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}

	grpcServer := grpc.NewServer()
	// depedency.TokenValidation.UnimplementedTokenValidationServer.ValidateToken()
	tokenvalidation.RegisterTokenValidationServer(grpcServer, depedency.TokenValidation)

	if err := grpcServer.Serve(netlisten); err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}
}
