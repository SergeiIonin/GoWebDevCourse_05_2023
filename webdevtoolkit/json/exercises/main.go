package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type jsonStatus struct {
	StatusName string
	Code       int
}

func main() {

	http.HandleFunc("/statuses", statuses)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func statuses(w http.ResponseWriter, req *http.Request) {
	raw := `StatusOK                   = 200
	StatusMovedPermanently  = 301
	StatusFound             = 302
	StatusSeeOther          = 303
	StatusTemporaryRedirect = 307
	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusTeapot                       = 418
	StatusInternalServerError           = 500`

	rawLines := strings.Split(raw, "\n")

	var jsonRaw strings.Builder
	jsonRaw.WriteString("[")
	// use bytes buffer
	var statusCode []string

	size := len(rawLines)

	for i, str := range rawLines {
		statusCode = strings.Split(str, "=")
		jsonRaw.WriteString(fmt.Sprintf("{\"StatusName\": \"%s\",\n", strings.TrimSpace(statusCode[0])))
		if i != size-1 {
			jsonRaw.WriteString(fmt.Sprintf("\"Code\": %s},\n", strings.TrimSpace(statusCode[1])))
		} else {
			jsonRaw.WriteString(fmt.Sprintf("\"Code\": %s}\n", strings.TrimSpace(statusCode[1])))
		}

	}
	jsonRaw.WriteString("]")

	fmt.Println(jsonRaw.String())

	jsonData := []byte(jsonRaw.String())

	type dataArray []jsonStatus
	var data *dataArray

	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println("Error decoding json:", err.Error())
	}

	for _, record := range *data {
		fmt.Println(record)
	}

	jsonBack, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	io.WriteString(w, string(jsonBack))

}
