package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GourabDas18/g-rest/internal/service"
	"github.com/GourabDas18/g-rest/utility"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dir := "log"
		filename := "logdata.csv"

		csv, err := service.NewCsvLogger(dir, filename)
		if err != nil {
			utility.ErrorResponse(w, http.StatusInternalServerError, "Logging error in server")
			fmt.Println("Error in csv logger %w", err)
			return
		}
		var queryParts []string

		for k, v := range r.URL.Query() {
			valStr := strings.Join(v, ",")
			queryParts = append(queryParts, fmt.Sprintf("%s=%s", k, valStr))
		}
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			utility.ErrorResponse(w, http.StatusInternalServerError, "Logging error in server")
			fmt.Println("Error in csv logger %w", err)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		err = csv.Log(r.Method, r.RequestURI, string(bodyBytes), strings.Join(queryParts, ","))
		if err != nil {
			utility.ErrorResponse(w, http.StatusInternalServerError, "Logging error in server")
			fmt.Println("Error in csv logger %w", err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
