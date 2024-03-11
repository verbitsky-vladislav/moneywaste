package middleware

import (
	"fmt"
	"net/http"
)

func Method(next http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Allow", method)
			http.Error(w, fmt.Sprintf("Method Not Allowed | Allowed method for path [ %s ] is [ %s ]", r.URL, method), http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
