package auth

import (
	"fmt"
	"net/http"
	"strings"
	u_sushi "u-sushi"
	"u-sushi/models"

	"github.com/golang-jwt/jwt/v5"
)

const TokenPrefix string = "Bearer "

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return AuthMiddleware(next, models.UserAdmin)
}

func ClientAuthMiddleware(next http.Handler) http.Handler {
	return AuthMiddleware(next, models.UserClient)
}

func AuthMiddleware(next http.Handler, userType models.UserType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do not ask authorization for auth routes
		if strings.HasPrefix(r.URL.Path, "/auth") {
			next.ServeHTTP(w, r)
			return
		}
		tokenStr := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenStr, TokenPrefix) {
			u_sushi.HttpError(w, http.StatusUnauthorized, fmt.Errorf("bearer authorization header required"))
			return
		}
		tokenStr = tokenStr[len(TokenPrefix):]
		if token, err := ParseToken(tokenStr); err != nil {
			u_sushi.HttpError(w, http.StatusUnauthorized, err)
			return
		} else if ut, ok := token.Claims.(jwt.MapClaims)["userType"]; !ok {
			u_sushi.HttpError(w, http.StatusUnauthorized, fmt.Errorf("only %s can access this route", strings.ToLower(string(userType))))
			return
		} else if utp, ok := ut.(string); !ok || models.UserType(utp) != userType {
			u_sushi.HttpError(w, http.StatusUnauthorized, fmt.Errorf("only %s can access this route", strings.ToLower(string(userType))))
			return
		}

		next.ServeHTTP(w, r)
	})
}
