package util

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"

	appErr "github.com/greennit/error"
)

var SecretKey = []byte("access_token")
var BasicKey = "Basic"

// Middleware - Pre-handler or req/res
type Middleware func(*context.Context, http.HandlerFunc) http.HandlerFunc

// MultipleMiddleware - Implement check in addditional conditions before actually handle req/res
func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return h
	}

	wrapped := h
	ctx := context.TODO()
	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](&ctx, wrapped)
	}

	return wrapped

}

// JsonHandler - transfer body to j
type JsonHandler func(w http.ResponseWriter, r *http.Request) (interface{}, *appErr.AppError)

// ServeHTTP - handle http req/res
func (handler JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if data, err := handler(w, r); err != nil {
		http.Error(w, err.Message, err.Code)
	} else {
		by, parseError := json.Marshal(data)
		if parseError != nil {
			http.Error(w, err.Message, http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(by)
		}
	}
}

// BaseClaims - Used to generate and extract token
type BaseClaims struct {
	Email string `json:"email,omitempty"`
	jwtGo.StandardClaims
}

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("%s - [%s] - %s %s\n\n", r.Proto, r.Method, r.URL, w.Header().Get("Status-Code"))
	})
}

// JWTAuthMiddleware - Validate Token
func JWTAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from request
		authString := r.Header.Get("Authorization")
		var tokenString string
		if len(authString) <= 0 {
			http.Error(w, jwtGo.ErrHashUnavailable.Error(), http.StatusBadRequest)
		} else {
			tokenString = strings.Split(authString, " ")[1]
			// Extract
			claims := BaseClaims{}
			_, err := jwtGo.ParseWithClaims(tokenString, claims, func(token *jwtGo.Token) (interface{}, error) {
				return BasicKey, nil
			})
			if err != nil {
				if err == jwtGo.ErrSignatureInvalid {
					http.Error(w, jwtGo.ErrSignatureInvalid.Error(), http.StatusUnauthorized)
				} else {
					http.Error(w, appErr.ErrBadCredential.Error(), http.StatusUnauthorized)
				}
			} else {
				if !claims.VerifyExpiresAt(time.Now().Unix(), true) { // Validate exp
					http.Error(w, "Token is expired", http.StatusUnauthorized)
				} else { // Other
					err := claims.Valid()
					if err != nil {
						http.Error(w, "Invalid info", http.StatusUnauthorized)
					} else {
						// If auth is valid, then pass to next
						h.ServeHTTP(w, r)
					}
				}
			}
		}
	})
}
