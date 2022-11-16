package presentation

import (
	"encoding/json"
	"net/http"
)

type ErrResp struct {
	Message string `json:"message"`
}

func RespJSON(w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := &ErrResp{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(status)
	w.Write(bodyBytes)
}
