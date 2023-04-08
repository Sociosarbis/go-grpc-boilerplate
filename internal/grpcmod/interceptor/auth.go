package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/sociosarbis/grpc/boilerplate/internal/ctxkey"
	"github.com/sociosarbis/grpc/boilerplate/internal/jwtgo"
)

type AuthInterceptor struct {
	jwtManager *jwtgo.JWTManager
}

func NewAuth(jwtManager *jwtgo.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager,
	}
}

func (inter *AuthInterceptor) Auth(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	claims, err := inter.jwtManager.Verify(values[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return handler(context.WithValue(ctx, ctxkey.UseClaims, claims), req)
}
