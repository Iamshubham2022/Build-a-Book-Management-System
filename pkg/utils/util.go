package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		err := json.Unmarshal([]byte(body), x) // first of all it will convert all the data into byte then that data unmashal into x
		if err != nil {
			return
		}
	}
}

// func ParseBody(r *http.Request, x interface{}) {
// 	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
// 		return
// 	}
// }
