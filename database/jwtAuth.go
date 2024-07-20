package database

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const (
	ContextUserID   contextKey = "UserIDClaim"
	ContextUsername contextKey = "UsernameClaim"
)

// JWTAuthMiddleware is a middleware function that authenticates requests using JWT tokens.
// It extracts the token from the Authorization header, parses it, and validates its authenticity.
// If the token is valid, it extracts the user ID and username from the token claims and adds them
// to the request context. If the token is missing, invalid, or expired, it returns an unauthorized error.
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization token missing/empty", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		tokenString = strings.TrimPrefix(tokenString, " ")

		// Parse the JWT token to extract the claims (e.g., the username of the authenticated user)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token : Unauthorised Action", http.StatusUnauthorized)
			return
		}
		var UserIDClaim uint64 = 0
		var UsernameClaim string = ""
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			UserIDClaim = uint64(claims["UserID"].(float64))
			UsernameClaim = (claims["Username"]).(string)
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextUserID, UserIDClaim)
		ctx = context.WithValue(ctx, ContextUsername, UsernameClaim)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

//Adding a salt to the JWT secret key improves security by making it more 
//difficult for attackers to guess the key through brute-force attacks.
