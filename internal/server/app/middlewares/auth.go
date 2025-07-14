package middlewares

import (
	"context"
	"main/internal/server/auth"
	"main/internal/server/constants"
	"net/http"
	"strings"
)

type AuthParams struct {
	IgnoreURLs []string
	JWTSecret  string
}

// Authentication wraps the next handler with JWT-based authentication.
func Authentication(conf AuthParams) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.String()

			if startsWithAny(url, conf.IgnoreURLs) {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("access_token")
			if err != nil || cookie == nil {
				w.Header().Set("content-type", constants.TextContentType)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			tokenStr := cookie.Value
			claims, err := auth.VerifyJWT(tokenStr, conf.JWTSecret)
			if err != nil {
				w.Header().Set("content-type", constants.TextContentType)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), constants.UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//func setJWTCookie(userID int64, secret string, tokenExp time.Duration, CookieExp int) (*http.Cookie, error) {
//	tokenStr, err := auth.GenerateJWT(userID, secret, tokenExp)
//	if err != nil {
//		return nil, err
//	}
//	cookie := http.Cookie{
//		Name:     "access_token",
//		Value:    tokenStr,
//		Path:     "/",
//		HttpOnly: true,
//		MaxAge:   CookieExp,
//	}
//	return &cookie, nil
//}

func startsWithAny(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
