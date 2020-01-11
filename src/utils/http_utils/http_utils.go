package http_utils

import (
	"encoding/json"
	"fmt"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if e := json.NewEncoder(w).Encode(body); e != nil {
		fmt.Println("Error json: " + e.Error())
	}
}

func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJson(w, err.Status(), err)
}
