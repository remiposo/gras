package presentation

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.Write([]byte(`{"status": "ok"}`))
}
