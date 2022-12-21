package middleware

// import required packages here ...
import (
	"context"
	dto "dumbmerch/dto/result"
	jwtToken "dumbmerch/pkg/jwt"
	"encoding/json"
	"net/http"
	"strings"
)

// Declare Result struct here ...
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Create Auth function here ...
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		token := req.Header.Get("Authorization")

		if token == "" {
			res.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "kamu gak diajak"}
			json.NewEncoder(res).Encode(response)
			return
		}
		// noted
		token = strings.Replace(token, "Bearer ", "", 100)
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "gak diajak 😢"}
			json.NewEncoder(res).Encode(response)
			return
		}

		ctx := context.WithValue(req.Context(), "userLogin", claims)
		next.ServeHTTP(res, req.WithContext(ctx))
	})

}
