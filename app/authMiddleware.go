package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"net/http"
	"os"
)

type AuthMiddleware struct {}

func (am AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			currentRouteName := mux.CurrentRoute(r).GetName()
			if !((currentRouteName == "GetUser") || (currentRouteName == "UpdateUserWeights")) {
				next.ServeHTTP(w, r)
			} else {
				authHeader := r.Header.Get("Authorization")
				if authHeader != "" {
					secret , _ := os.LookupEnv("JWT_SECRET")
					tokenString := getTokenFromHeader(authHeader)

					_ , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
						return []byte(secret), nil
					})
					if err != nil {
						appError := errs.NewUnauthorizedError("Token invalid")
						writeResponse(w, appError.Code, appError.AsMessage())
					} else {
						next.ServeHTTP(w, r)
					}
				} else {
					appError := errs.NewUnauthorizedError("Token missing")
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			}
		})
	}
}