package intercepters

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ValidateRegexIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("In regex intercepter")
	return handler(ctx, req)
}
