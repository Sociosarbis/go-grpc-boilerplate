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

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	return values[0], nil
}

func (inter *AuthInterceptor) Auth(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (any, error) {
	token, err := getTokenFromContext(ctx)

	if err != nil {
		return nil, err
	}

	claims, err := inter.jwtManager.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return handler(context.WithValue(ctx, ctxkey.UseClaims, claims), req)
}

func (inter *AuthInterceptor) AuthStream(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	token, err := getTokenFromContext(ss.Context())

	if err != nil {
		return err
	}
	claims, err := inter.jwtManager.Verify(token)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	err = ss.SetHeader(metadata.MD{
		"id":   []string{claims.Id},
		"name": []string{claims.Name},
	})
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "ServerStream.SetHeader: %v", err)
	}
	return handler(srv, ss)
}
