package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
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

	fmt.Println("The provided access is valid.")

	// Continue processing the request.
	return handler(ctx, req)
}

// validateAccessToken validates the access token using JWT and checks the "scope" claim.
//
// This function parses the provided JWT token and verifies its signature using the given secret key.
// It then extracts the claims from the token and checks if the "scope" claim exists and contains the string "demo".
func validateAccessToken(token string) (bool, error) {
	// Parse the JWT token.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Use your secret key to validate the token.
		return []byte("my-super-secret"), nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid.
	if !parsedToken.Valid {
		return false, fmt.Errorf("invalid token")
	}

	// Extract the claims from the token.
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid token claims")
	}

	// Check if the "scope" claim exists.
	scope, ok := claims["scope"]
	if !ok {
		return false, fmt.Errorf("missing scope claim")
	}

	// Convert the scope claim to a string array.
	scopeArray, ok := scope.([]interface{})
	if !ok {
		return false, fmt.Errorf("invalid scope claim format")
	}

	// Check if the "demo" string is present in the scope array.
	for _, s := range scopeArray {
		if strings.EqualFold(s.(string), "demo") {
			return true, nil
		}
	}

	return false, fmt.Errorf("missing 'demo' scope")
}
