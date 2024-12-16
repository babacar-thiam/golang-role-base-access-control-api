package auth

import (
	"context"
	"net/http"
	"role-base-access-control-api/internal/role"
)

type contextKey string

const (
	UserClaimsKey contextKey = "user_claims"
)

type Middleware struct {
	jwt *JWT
}

// NewMiddleware create a new instance of Middleware
func NewMiddleware(jwt *JWT) *Middleware {
	return &Middleware{jwt: jwt}
}

// AuthMiddleware validates the JWT Token
func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := m.jwt.ExtractTokenFromHeader(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := m.jwt.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminOnly middleware for admin-only routes
func (m *Middleware) AdminOnly(next http.Handler) http.Handler {
	return m.RoleAuth(next, role.ADMIN)
}

// ClientOnly middleware for client-only routes
func (m *Middleware) ClientOnly(next http.Handler) http.Handler {
	return m.RoleAuth(next, role.CLIENT)
}

// ProviderOnly middleware for provider-only routes
func (m *Middleware) ProviderOnly(next http.Handler) http.Handler {
	return m.RoleAuth(next, role.PROVIDER)
}

// RoleAuth checks if the user has any of the allowed roles
func (m *Middleware) RoleAuth(next http.Handler, allowedRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(UserClaimsKey).(*Claims)

		allowed := false
		for _, allowedRole := range allowedRoles {
			if claims.Role == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
