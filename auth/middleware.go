package auth

import (
	"fmt"
	"net/http"
	"strings"
	u_sushi "u-sushi"
	"u-sushi/models"
)

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return AuthMiddleware(next, models.UserAdmin)
}

func ClientAuthMiddleware(next http.Handler) http.Handler {
	return AuthMiddleware(next, models.UserClient)
}

func unauthorizedErr(userType models.UserType) error {
	return fmt.Errorf("only %s can access this route", strings.ToLower(string(userType)))
}

func AuthMiddleware(next http.Handler, userType models.UserType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do not ask authorization for auth routes
		if strings.HasPrefix(r.URL.Path, "/auth") {
			next.ServeHTTP(w, r)
			return
		}
		if claims, err := ExtractClaims(r); err != nil {
			u_sushi.HttpError(w, http.StatusUnauthorized, err)
			return
		} else if claims.UserType != userType {
			u_sushi.HttpError(w, http.StatusUnauthorized, unauthorizedErr(userType))
			return
		}

		next.ServeHTTP(w, r)
	})
}
