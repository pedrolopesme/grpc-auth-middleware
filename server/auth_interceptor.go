package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor checks for a valid OpenID Connect access token in the request metadata.
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// Extract the access token from the metadata.
	accessToken, ok := md["authorization"]
	if !ok || len(accessToken) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing access token")
	}

	// Validate the access token using your OpenID Connect provider.
	isValid, err := validateAccessToken(accessToken[0])
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to validate token: %v", err)
	}

	if !isValid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid access token")
	}

	// Continue processing the request.
	return handler(ctx, req)
}

// validateAccessToken validating the access token.
func validateAccessToken(token string) (bool, error) {
	// TODO: Implement OpenID Connect validation logic.
	fmt.Println("Validating access token:", token)
	return true, nil
}
