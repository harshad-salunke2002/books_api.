package response

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		ResponseWithError(w, 400, "Error While Encoding Json")
		return
	}

	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
