package helpers

import (
	"GoAPI/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func ResponseWithError(err error, w http.ResponseWriter, status int, msg string) {
	response := types.Response{
		Success: false,
		Message: msg,
		Data:    err.Error(),
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func ResponseWithSuccess(data interface{}, w http.ResponseWriter, status int, msg string) {
	response := types.Response{
		Success: true,
		Message: msg,
		Data:    data,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func StringToInt(s string) (int, error) {
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return number, nil
}
