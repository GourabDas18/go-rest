package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/GourabDas18/g-rest/utility"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			utility.ErrorResponse(w, http.StatusUnauthorized, "Authentication failed")
			return
		}
		authStrings := strings.Split(authToken, "Bearer ")
		if len(authStrings) < 2 {
			utility.ErrorResponse(w, http.StatusUnauthorized, "Authentication failed")
			return
		}
		claims, err := service.ValidateToken(authStrings[1])
		if err != nil {
			utility.ErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf(`Authentication failed %v`, err))
			return
		}
		bodyMap := map[string]any{}
		err = json.NewDecoder(r.Body).Decode(&bodyMap)
		if err != nil && err != io.EOF {
			utility.ErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf(`Authentication failed %v`, err.Error()))
			return
		}
		bodyMap["userId"] = claims.UserId
		bodyMap["userName"] = claims.UserName
		bodyMap["countryId"] = claims.CountryId

		bodyMapData, err := json.Marshal(bodyMap)
		if err != nil {
			fmt.Printf(`Value %v\n`, bodyMap)
			utility.ErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf(`Authentication failed %v`, err))
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(bodyMapData))
		next.ServeHTTP(w, r)
	})
}
