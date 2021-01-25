package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const httpStatusSuccess string = "success"
const httpStatusError string = "error"

type response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

//ToJSON - convert struct to JSON
func Response(w http.ResponseWriter, d interface{}, status int, err bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(d)

	l, _ := json.Marshal(d)
	if err {
		Log.Error(string(l))
	} else {
		Log.Info(string(l))
	}
}

//CreateRequestURI - create request uri
func CreateRequestURI(host string, res string, data url.Values) string {
	u, _ := url.ParseRequestURI(host)
	u.Path = res
	if data != nil {
		u.RawQuery = data.Encode()
	}
	url := fmt.Sprintf("%v", u)

	return url
}

//HTTPResponse - set response
func HTTPResponse(w http.ResponseWriter, status string, statusCode int, message string) {
	if status == httpStatusSuccess {
		Log.Info(message)
	} else {
		Log.Error(message)
	}

	d := response{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(d)

	return
}

func GetURI(fullUrl string, n int) string {
	url := strings.Split(fullUrl, "/")
	uri := ""

	if len(url) >= n {
		uri = url[n-1]
	}

	return uri
}

func ToJSON(d interface{}) string {
	j, _ := json.Marshal(d)
	return string(j)
}
