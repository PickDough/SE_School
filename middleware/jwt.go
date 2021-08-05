package middleware

import (
	"SE_School/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		publicPaths := []string{"/user/create", "/user/login"}
		requestPath := r.URL.Path

		//Checking if request path requires authentication
		for _, publicPath := range publicPaths {

			if publicPath == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		err := validateToken(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, utils.Message("Access token is absent or didn't pass validation. Details: "+err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateToken(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}

	//Validating Claims
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)

	//Decoding token using secret in .env
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error during jwt parse")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
