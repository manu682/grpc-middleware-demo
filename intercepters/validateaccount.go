package intercepters

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ValidateAccountIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	errInvalidToken := status.Errorf(codes.Unauthenticated, "invalid token")
	errMissingMetadata := status.Errorf(codes.InvalidArgument, "missing metadata")
	fmt.Println("In account intercepter")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	fmt.Println("md : ", md)
	fmt.Println("info : ", info)

	if true {
		return nil, errInvalidToken
	}

	return handler(ctx, req)
}
