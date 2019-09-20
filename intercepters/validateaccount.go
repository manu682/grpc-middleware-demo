package intercepters

import (
	"context"

	"google.golang.org/grpc"
)

func ValidateAccountIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	return handler(ctx, req)
}
